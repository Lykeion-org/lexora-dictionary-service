package main

import (
	db "github.com/Lykeion-org/lexora-dictionary-service/internal/db"
	grpcserver "github.com/Lykeion-org/lexora-dictionary-service/internal/grpc"
	"os"
	"errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	UseCli bool `yaml:"use_cli"`
	UseGrpcServer bool `yaml:"use_grpc_server"`
	UseRestApi bool `yaml:"use_rest_api"`
	DictionaryServicePort string `yaml:"dictionary_service_port"`
	DatabaseHost string `yaml:"database_host"`
	DatabasePort int `yaml:"database_port"`
	DatabaseName string `yaml:"database_name"`
	DatabaseUserName string `yaml:"database_username"`
	DatabasePassword string `yaml:"database_password"`
}

func main() {
	//Load config
	cfg := loadConfig("/Users/kevin/Repositories/lykeion/server/services/lexora/dictionary-service/config.yaml")

	//Set init variables and create objects
	connString := db.CreateConnectionString(cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName, cfg.DatabaseUserName, cfg.DatabasePassword)
	db, _ := db.ConnectAndMigrate(connString)
	
	srv := grpcserver.NewLanguageServer(db)
	
	//Run main processes
	if(cfg.UseGrpcServer) {
		srv.StartServer(cfg.DictionaryServicePort); defer srv.StopServer()
	}

	select {}

}

func loadConfig(filePath string) (*Config){
	var cfg *Config
    if _, err := os.Stat(filePath); err == nil {
        data, err := os.ReadFile(filePath)
        if err != nil {
            panic("Could not read file: " + filePath)
        }
        if err := yaml.Unmarshal(data, &cfg); err != nil {
            panic("Error parsing YAML: " + err.Error())
        }
    } else if !errors.Is(err, os.ErrNotExist) {
        panic("Unexpected error occured when loading config: " + err.Error())
    }

	return cfg
}
