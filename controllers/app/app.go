package app

import "github.com/gin-gonic/gin"

func App(ctx *gin.Context) {

	ctx.Redirect(307, "https://cooding.io/app/")
}
