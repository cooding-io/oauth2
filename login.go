package oauth2

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elmalba/oauth2-server"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	r := ctx.Request

	if r.Method == "POST" {
		email := r.FormValue("username")
		password := r.FormValue("password")

		user, exist := GetUserByEmail(ctx, email)
		fmt.Println("USER", user, exist)
		if exist && user.Password == password {

			s := oauth2.Session{}
			s.Load(ctx)
			s.ID = user.ID
			s.Email = user.Email
			s.Save(ctx)
			base := ctx.Request.Host
			fmt.Println("BASE", base)

			ctx.Redirect(http.StatusTemporaryRedirect, basePath+"/auth")
			return
		}

	}

	if ctx.Request.Host == "docencia-eit.udp.cl" {
		outputHTML(ctx, "static/login-udp.html")
	} else {
		outputHTML(ctx, "static/login.html")
	}

}

func outputHTML(ctx *gin.Context, filename string) {
	r := ctx.Request
	w := ctx.Writer
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, r, file.Name(), fi.ModTime(), file)
}

func Logout(ctx *gin.Context) {

	s := oauth2.Session{}
	s.Save(ctx)
	ctx.Redirect(http.StatusTemporaryRedirect, basePath+"/login")

	return
}
