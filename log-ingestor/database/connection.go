package database

import (
	"log"

	"github.com/chrisdamba/go-react-log-seer/config"
	"github.com/chrisdamba/go-react-log-seer/models"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitElasticsearch() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(config.AppConfig.ElasticsearchURL))
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}
	return client
}

func InitPostgres() *gorm.DB {
	dsn := "host=" + config.AppConfig.PostgresHost + 
		" user=" + config.AppConfig.PostgresUser + 
		" password=" + config.AppConfig.PostgresPassword + 
		" dbname=" + config.AppConfig.PostgresDBName + 
		" port=" + config.AppConfig.PostgresPort + 
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
	}
	db.AutoMigrate(&models.LogEntry{}, &models.Metadata{})
	return db
}
