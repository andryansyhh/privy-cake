package handler

import (
	"database/sql"
	"net/http"
	"privy/models"
	"privy/usecase"
	"strconv"

	"privy/utils"

	"github.com/gin-gonic/gin"
)

type handlerHttpServer struct {
	app usecase.UsecaseInterface
}

func NewHttpHandler(app usecase.UsecaseInterface) handlerHttpServer {
	return handlerHttpServer{app: app}
}

func (b *handlerHttpServer) PostCake(c *gin.Context) {
	var req models.CakeRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	err = b.app.PostCake(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "created",
		"message": "success add new cake",
	})
}

func (b *handlerHttpServer) GetAllCakes(c *gin.Context) {
	res, err := b.app.GetAllCakes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *handlerHttpServer) GetCakeByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	res, err := b.app.GetCakeByID(idInt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "cake not found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *handlerHttpServer) UpdateCakeByID(c *gin.Context) {
	var req models.CakeRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	err = b.app.UpdateCakeByID(idInt, req)
	if err == utils.ErrIdNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (b *handlerHttpServer) DeleteCakeByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	err := b.app.DeleteCakeByID(idInt)
	if err == utils.ErrIdNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
