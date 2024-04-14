package di

import (
	"backend/infrastrcuture/persistence"
	"backend/interface/handler"
	"backend/usecase"

	"gorm.io/gorm"
)

func User(db *gorm.DB) handler.UserHandler {
	up := persistence.NewUserPersistence(db)
	uu := usecase.NewUserUsecase(up)
	uh := handler.NewUserHandler(uu)
	return uh
}
