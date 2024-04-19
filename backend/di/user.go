package di

import (
	"backend/infrastrcuture/persistence"
	handler "backend/interface/handler/echo"
	"backend/usecase"

	"gorm.io/gorm"
)

func User(db *gorm.DB) handler.UserHandler {
	up := persistence.NewUserPersistence(db)
	uu := usecase.NewUserUsecase(up)
	uh := handler.NewUserHandler(uu)
	return uh
}
