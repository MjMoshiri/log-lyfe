package main

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg/parser"
	"github.com/mjmoshiri/log-lyfe/gol/storage"
)

func initServerConfig(serverConfigPath string) (*models.ServerConfig, error) {
	serverConfig := &models.ServerConfig{}
	err := parser.FromYAML(serverConfigPath, serverConfig)
	if err != nil {
		return nil, err
	}
	return serverConfig, nil
}

func initDBConfig(dbConfigPath string) (*models.DBConfig, error) {
	dbConfig := &models.DBConfig{}
	err := parser.FromYAML(dbConfigPath, dbConfig)
	if err != nil {
		return nil, err
	}
	return dbConfig, nil
}

func initDB(dbConfig *models.DBConfig) (storage.DB, error) {
	db, err := storage.New(*dbConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}
