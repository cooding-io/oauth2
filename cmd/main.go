package main

import (
	"log"
	"oauth2/controllers/app"
	cooding "oauth2/controllers/oauth2"
	"os"

	"github.com/Ensena/core/env-global"
	"github.com/elmalba/oauth2-server"
	"go.elastic.co/apm/module/apmgin"
)

var basePath, hostName, key string

func init() {
	//basePath = env.Check("basepath", "Missing Params basepath")
	basePath = "login"
	hostName = env.Check("hostname", "Missing Params hostname")
	key = env.Check("secretKey", "Missing Params secretKey")
}

func main() {
	srv, ws := oauth2.CreateServer(hostName, basePath)
	srv.SetKey(key)
	srv.MiddleWare = cooding.AuthMiddleWare
	srv.GetUser = cooding.GetUser
	srv.ValidateClientID = cooding.GetApp
	srv.ValidateClientIDAndSecretID = cooding.GetAppAndSecret

	ws.Use(apmgin.Middleware(ws))

	ws.GET("/app/", app.App)

	ws.GET(basePath+"/QR/Generate", cooding.Login)
	ws.GET(basePath+"/QR/Check", cooding.Login)

	ws.GET("/", cooding.Login)
	ws.GET(basePath+"/", cooding.Login)
	ws.POST(basePath+"/", cooding.Login)
	ws.GET(basePath+"/logout", cooding.Logout)
	ws.GET(basePath+"/google", cooding.OauthGoogleLogin)
	ws.GET(basePath+"/google/callback", cooding.OauthGoogleCallback)
	ws.POST(basePath+"/google/callback", cooding.OauthGoogleCallback)
	ws.Static(basePath+"/assets/", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	ws.Run(":" + port)

}
