package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"Expire/config"
	"Expire/controller"
	"Expire/model"
	"Expire/router"

	bankRepository "Expire/repository/Bank"
	externalRepository "Expire/repository/External"
	leaderRepository "Expire/repository/Leader"
	reasonRepository "Expire/repository/Reason"
	reportRepository "Expire/repository/Report"
	supervisorRepository "Expire/repository/Supervisor"
	userRepository "Expire/repository/User"

	authService "Expire/service/Authentication"
	bankService "Expire/service/Bank"
	externalService "Expire/service/External"
	leaderService "Expire/service/Leader"
	reasonService "Expire/service/Reason"
	reportService "Expire/service/Report"
	supervisorService "Expire/service/Supervisor"
	tokenService "Expire/service/Token"
	userService "Expire/service/User"

	"github.com/go-playground/validator/v10"
)

func main() {
	envConf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	// Redis
	config.ConnectRedis(&envConf)

	// Database
	db := config.DatabaseConnection(&envConf)
	validate := validator.New()

	println("Message: Migrating Table...")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Report{})
	db.AutoMigrate(&model.Supervisor{})
	db.AutoMigrate(&model.Bank{})
	db.AutoMigrate(&model.Leader{})
	db.AutoMigrate(&model.Reason{})
	db.AutoMigrate(&model.User{}, &model.Supervisor{})
	db.AutoMigrate(&model.User{}, &model.External{})
	db.AutoMigrate(&model.User{}, &model.Leader{})
	db.AutoMigrate(&model.External{}, &model.Bank{})

	// Repositories
	userRepository := userRepository.NewUserRepositoryImpl(db)
	reportRepository := reportRepository.NewReportRepositoryImpl(db)
	supervisorRepository := supervisorRepository.NewSupervisorRepositoryImpl(db)
	bankRepository := bankRepository.NewBankRepositoryImpl(db)
	leaderRepository := leaderRepository.NewLeaderRepositoryImpl(db)
	reasonRepository := reasonRepository.NewReasonRepositoryImpl(db)
	externalRepository := externalRepository.NewExternalRepositoryImpl(db)

	// Services
	userService := userService.NewUserServiceImpl(bankRepository, userRepository, supervisorRepository, externalRepository, leaderRepository, reportRepository, validate)
	tokenService := tokenService.NewTokenServiceImpl(userRepository)
	authService := authService.NewAuthServiceImpl(userRepository, supervisorRepository, leaderRepository, externalRepository, validate)
	reportService := reportService.NewReportServiceImpl(reportRepository, supervisorRepository, leaderRepository, bankRepository, reasonRepository, externalRepository, validate)
	supervisorService := supervisorService.NewSupervisorServiceImpl(supervisorRepository, validate)
	bankService := bankService.NewBankServiceImpl(bankRepository, validate)
	leaderService := leaderService.NewLeaderServiceImpl(leaderRepository, validate)
	reasonService := reasonService.NewReasonServiceImpl(reasonRepository, validate)
	externalService := externalService.NewExternalServiceImpl(externalRepository, bankRepository, validate)

	// Controllers
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)
	tokenController := controller.NewTokenController(tokenService)
	reportController := controller.NewReportController(reportService)
	supervisorController := controller.NewSupervisorController(supervisorService)
	bankController := controller.NewBankController(bankService)
	leaderController := controller.NewLeaderController(leaderService)
	reasonController := controller.NewReasonController(reasonService)
	externalController := controller.NewExternalController(externalService)

	// Initialize Router
	routes := router.NewRouter(
		db,
		userController,
		authController,
		tokenController,
		reportController,
		supervisorController,
		bankController,
		leaderController,
		reasonController,
		externalController,
	)

	// Intialize Server
	server := &http.Server{
		Addr:           ":8081",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	errProvideServer := server.ListenAndServe().Error()

	fmt.Println(errProvideServer)

	println("Message: Server Successfully Running...")
}
