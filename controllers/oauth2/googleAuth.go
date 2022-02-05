package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Ensena/core/env-global"
	oauth2Server "github.com/elmalba/oauth2-server"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAccount struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	Hd            string `json:"hd"`
}

var basePath, hostName string
var RedirectURL, ClientID, ClientSecret string

func init() {
	basePath = env.Check("basepath", "Missing Params basepath")
	hostName = env.Check("hostname", "Missing Params hostname")
	googleOauthConfig.RedirectURL = env.Check("RedirectURL", "Missing Params RedirectURL")
	googleOauthConfig.ClientID = env.Check("ClientID", "Missing Params ClientID")
	googleOauthConfig.ClientSecret = env.Check("ClientSecret", "Missing Params ClientSecret")

}

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "",
	ClientID:     "",
	ClientSecret: "",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func OauthGoogleLogin(ctx *gin.Context) {

	r := ctx.Request
	w := ctx.Writer
	oauthState := generateStateOauthCookie(w)
	uri, err := url.Parse(r.Header.Get("Referer"))
	if err != nil {
		panic(err)
	}
	g := googleOauthConfig
	g.RedirectURL = fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, "/login/google/callback")
	u := g.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func redirect(w http.ResponseWriter, r *http.Request, id int, valid bool) {
	w.Write([]byte(fmt.Sprintf(`<!DOCTYPE html>
	<html>
	Hooola
	</html>`, r.Header.Get("Origin"))))

}

func OauthGoogleCallback(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		redirect(w, r, -1, false)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		redirect(w, r, -1, false)
		return
	}

	Guser := GoogleAccount{}
	if err := json.Unmarshal(data, &Guser); err != nil {
		redirect(w, r, -1, false)
		return
	}
	if Guser.VerifiedEmail == false {
		redirect(w, r, -1, false)
		return
	}

	user, exist := GetUserByEmail(ctx, Guser.Email)

	if exist == false {
		if strings.Contains(Guser.Email, "@mail.udp.cl") || strings.Contains(Guser.Email, "@uft.edu") {
			user = &User{}
			user.Email = strings.ToLower(Guser.Email)
			user.Name = strings.Title(strings.ToLower(Guser.GivenName))
			user.LastName = strings.Title(strings.ToLower(Guser.FamilyName))
			createUser(ctx, user)
			user, _ = GetUserByEmail(ctx, Guser.Email)
		} else {
			return
		}
	}

	s := oauth2Server.Session{}
	s.Load(ctx)

	s.ID = user.ID
	s.Email = user.Email
	s.Save(ctx)
	http.Redirect(w, r, "/app/", http.StatusTemporaryRedirect)
	return

}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
