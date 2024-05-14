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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	go func(ctx context.Context, logEntry models.LogEntry) {
		if err := pgDB.WithContext(ctx).Create(&logEntry).Error; err != nil {
			log.Printf("Error saving to PostgreSQL: %v", err)
		}

		_, err := esClient.Index().
			Index("logs").
			BodyJson(logEntry).
			Do(ctx)
		if err != nil {
			log.Printf("Error saving to Elasticsearch: %v", err)
		}
	}(ctx, logEntry)

	response := map[string]string{"message": "Log entry ingested successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
