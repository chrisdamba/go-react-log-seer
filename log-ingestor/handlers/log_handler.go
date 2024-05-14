package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

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
	pgDB := database.InitPostgres()
	log.Println("Elasticsearch and Postgres initialised")
}

func IngestLog(w http.ResponseWriter, r *http.Request) {
	var logEntry models.LogEntry
	err := json.NewDecoder(r.Body).Decode(&logEntry)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := pgDB.Create(&logEntry).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	ctx := context.Background()
	_, err = esClient.Index().
		Index("logs").
		BodyJson(logEntry).
		Do(ctx)
	if err != nil {
		http.Error(w, "Failed to save log entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
