package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/chrisdamba/go-react-log-seer/config"
	"github.com/chrisdamba/go-react-log-seer/database"
	"github.com/chrisdamba/go-react-log-seer/models"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var esClient *elastic.Client
var pgDB *gorm.DB

func init() {
	config.LoadConfig()
	esClient = database.InitElasticsearch()
	pgDB = database.InitPostgres()
	log.Println("Elasticsearch and Postgres initialised")
}

func IngestLog(w http.ResponseWriter, r *http.Request) {
	var logEntry models.LogEntry
	if err := json.NewDecoder(r.Body).Decode(&logEntry); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Save Metadata first
	if err := pgDB.Create(&logEntry.Metadata).Error; err != nil {
		log.Printf("Error saving metadata to PostgreSQL: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save LogEntry with reference to Metadata
	logEntry.MetadataID = logEntry.Metadata.ID
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := pgDB.WithContext(ctx).Create(&logEntry).Error; err != nil {
		log.Printf("Error saving to PostgreSQL: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save to Elasticsearch
	_, err := esClient.Index().
		Index("logs").
		BodyJson(logEntry).
		Do(ctx)
	if err != nil {
		log.Printf("Error saving to Elasticsearch: %v", err)
		http.Error(w, "Failed to save log entry", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Log entry ingested successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
