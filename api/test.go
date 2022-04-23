package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}
