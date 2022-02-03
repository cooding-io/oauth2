package main

import (
	"github.com/Ensena/core/env-global"
	ensena "github.com/Ensena/oauth2-server"
	"github.com/elmalba/oauth2-server"
	"go.elastic.co/apm/module/apmgin"
)

var basePath, hostName, key string

func init() {
	basePath = env.Check("basepath", "Missing Params basepath")
	hostName = env.Check("hostname", "Missing Params hostname")
	key = env.Check("secretKey", "Missing Params secretKey")
}

func main() {
	srv, ws := oauth2.CreateServer(hostName, basePath)
	srv.SetKey(key)
	srv.MiddleWare = ensena.AuthMiddleWare
	srv.GetUser = ensena.GetUser
	srv.ValidateClientID = ensena.GetApp
	srv.ValidateClientIDAndSecretID = ensena.GetAppAndSecret

	ws.Use(apmgin.Middleware(ws))

	ws.GET(basePath+"/login", ensena.Login)
	ws.POST(basePath+"/login", ensena.Login)
	ws.GET(basePath+"/logout", ensena.Logout)
	ws.GET(basePath+"/login/google/login", ensena.OauthGoogleLogin)
	ws.GET(basePath+"/login/google/callback", ensena.OauthGoogleCallback)
	ws.POST(basePath+"/login/google/callback", ensena.OauthGoogleCallback)
	ws.Static(basePath+"/assets/", "./assets")
	ws.Run(":8000")

}
