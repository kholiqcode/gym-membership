// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gym/cmd/interface/handler"
	"gym/cmd/interface/handler/health"
	"gym/cmd/interface/handler/member"
	"gym/config"
	"gym/infrastructure/database"
	"gym/internal/protocol/http"
	"gym/internal/protocol/http/router"
	"gym/pkg/auth"
)

// Injectors from injector.go:

func InitHttpProtocol(mode string) (*http.HttpImpl, error) {
	configConfig := config.NewConfig()
	databaseImpl := database.NewDatabaseClient(configConfig, mode)
	memberRepositoryImpl, err := member.ProvideRepository(databaseImpl)
	if err != nil {
		return nil, err
	}
	jwtTokenImpl := auth.NewJwtToken()
	memberServiceImpl, err := member.ProvideService(memberRepositoryImpl, jwtTokenImpl)
	if err != nil {
		return nil, err
	}
	memberHandlerImpl, err := member.ProvideHandler(memberServiceImpl)
	if err != nil {
		return nil, err
	}
	healthHandlerImpl, err := health.ProvideHandler()
	if err != nil {
		return nil, err
	}
	httpHandlerImpl := handler.NewHttpHandler(memberHandlerImpl, healthHandlerImpl)
	httpRouterImpl := router.NewHttpRoute(httpHandlerImpl)
	httpImpl := http.NewHttpProtocol(httpRouterImpl, configConfig)
	return httpImpl, nil
}
