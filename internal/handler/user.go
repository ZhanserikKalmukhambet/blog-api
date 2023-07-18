package handler

import (
	"github.com/ZhanserikKalmukhambet/blog-api/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var newUser model.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err := h.srvs.CreateUser(ctx, &newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}
