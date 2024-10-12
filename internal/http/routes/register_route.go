package routes

import (
	"BACKEND_GO/internal/database/mysql"
	"BACKEND_GO/internal/domain/profile/repositories"
	"BACKEND_GO/internal/domain/profile/services"
	handler "BACKEND_GO/internal/http/handlers/profile"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	ctx := context.Background()

	db, err := mysql.NewMySQLConnection(ctx)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	pr := repositories.NewMysqlRepository(db)
	ps := services.NewProfileUseCase(pr)
	profileHandler := handler.NewProfileHandler(*ps)
	v1.GET("/profile/:id", profileHandler.GetProfile)

}
