package app

import (
	"fmt"
	"log"
	"os"
	"privy/infra"
	mysql "privy/infra/mysql"
	"privy/repository"
	"privy/usecase"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	// db *sql.DB =
	Repo := repository.NewRepo(mysql.InitMysql())
	log.Println(Repo)
	app := usecase.NewUsecase(Repo)
	infra.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
