package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := aid

	ctx.JSON(http.StatusOK, account)
}

func (c *Controller) ListAccounts(ctx *gin.Context) {

	account := "result"

	ctx.JSON(http.StatusOK, account)
}
