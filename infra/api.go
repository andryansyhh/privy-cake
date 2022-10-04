package infra

import (
	"privy/handler"
	"privy/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app usecase.UsecaseInterface) {
	cakeSrv := handler.NewHttpHandler(app)
	cake := r.Group("/cake")
	{
		cake.POST("", cakeSrv.PostCake)
		cake.GET("", cakeSrv.GetAllCakes)
		cake.GET("/:id", cakeSrv.GetCakeByID)
		cake.PUT("/:id", cakeSrv.UpdateCakeByID)
		cake.DELETE("/:id", cakeSrv.DeleteCakeByID)
	}
}
