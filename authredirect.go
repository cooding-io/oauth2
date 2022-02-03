package oauth2

import (
	"fmt"
	"log"

	"github.com/elmalba/oauth2-server"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(ctx *gin.Context, s *oauth2.Session) string {
	log.Println("SS", s)
	if s.ID == 0 {
		s.Save(ctx)
		ctx.Redirect(303, basePath+"/login")
		return ""

	}

	return fmt.Sprintf("%d", s.ID)
}
