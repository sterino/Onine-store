// Code generated by Wire. DO NOT EDIT.
//go:build !wireinject
// +build !wireinject

package di

import (
	_ "github.com/lib/pq"
	"users-service/internal/api"
	"users-service/internal/api/handler"
	"users-service/internal/config"
	"users-service/internal/db"
	"users-service/internal/repository"
	"users-service/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	url, sqlxDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	err = db.Migrate(url)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(sqlxDB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	server := http.NewServer(userHandler)
	return server, nil
}