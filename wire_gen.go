// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"2025-Lush-and-Verdant-Backend/client"
	"2025-Lush-and-Verdant-Backend/config"
	"2025-Lush-and-Verdant-Backend/controller"
	"2025-Lush-and-Verdant-Backend/dao"
	"2025-Lush-and-Verdant-Backend/middleware"
	"2025-Lush-and-Verdant-Backend/route"
	"2025-Lush-and-Verdant-Backend/service"
	"2025-Lush-and-Verdant-Backend/tool"
)

// Injectors from wire.go:

func InitApp(ConfigPath string) (*route.App, error) {
	viperSetting := config.NewViperSetting(ConfigPath)
	databaseConfig := config.NewDatabaseConfig(viperSetting)
	db, err := dao.NewDB(databaseConfig)
	if err != nil {
		return nil, err
	}
	userDAOImpl := dao.NewUserDAO(db)
	jwtConfig := config.NewJwtConfig(viperSetting)
	jwtClient := middleware.NewJwtClient(jwtConfig)
	emailDAOImpl := dao.NewEmailDAOImpl(db)
	qqConfig := config.NewQQConfig(viperSetting)
	mail := tool.NewMail(emailDAOImpl, qqConfig)
	priConfig := config.NewPriConfig(viperSetting)
	userServiceImpl := service.NewUserServiceImpl(userDAOImpl, jwtClient, mail, priConfig)
	userController := controller.NewUserController(userServiceImpl)
	userSvc := route.NewUserSvc(userController)
	sloganDAOImpl := dao.NewSloganDAOImpl(db)
	sloganServiceImpl := service.NewSloganServiceImpl(sloganDAOImpl, userDAOImpl)
	sloganController := controller.NewSloganController(sloganServiceImpl)
	sloganSvc := route.NewSloganSvc(sloganController)
	goalDAOImpl := dao.NewGoalDAOImpl(db)
	goalServiceImpl := service.NewGoalServiceImpl(goalDAOImpl)
	chatGptConfig := config.NewChatGptConfig(viperSetting)
	chatGptClient := client.NewChatGptClient(chatGptConfig)
	goalController := controller.NewGoalController(goalServiceImpl, chatGptClient)
	goalSvc := route.NewGoalSvc(goalController)
	qiNiuYunConfig := config.NewQNYConfig(viperSetting)
	imageDAOImpl := dao.NewImageDAO(db)
	imageServiceImpl := service.NewImageServiceImpl(qiNiuYunConfig, imageDAOImpl)
	imageController := controller.NewImageController(imageServiceImpl)
	imageSvc := route.NewImageSvc(imageController)
	app := route.NewApp(userSvc, sloganSvc, goalSvc, imageSvc)
	return app, nil
}
