package oauth2

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Ensena/core/graphql-client"
	"github.com/elmalba/oauth2-server"
	"github.com/gin-gonic/gin"
)

type clientsGraphql struct {
	Data struct {
		AllApps struct {
			TotalCount int `json:"totalCount"`
			Edges      []struct {
				Node struct {
					ID       int    `json:"id"`
					ClientID string `json:"clientId"`
					Secret   string `json:"secret"`
					URL      string `json:"url"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"allApps"`
	} `json:"data"`
}

func GetApp(ctx *gin.Context, clientID string) (*oauth2.Client, bool) {

	user := oauth2.Client{}
	if clientID == "" {
		clientID = "7nywNebh7Q"
	}
	g := fmt.Sprintf(`{
		allApps(condition:{clientId:"%s"}) {
		  totalCount
		  edges {
			node {
			  id
			  clientId
			  secret
			  url
			}
		  }
		}
	  }`, clientID)

	response, err := graphql.Query(ctx, g)
	if err != nil {
		log.Println("ERROR to connect graphql")
		return &user, false
	}
	userInput := clientsGraphql{}
	json.Unmarshal(response, &userInput)
	if userInput.Data.AllApps.TotalCount == 0 {
		log.Println("Not fount", userInput)
		return &user, false
	}

	user.CallBackURL = userInput.Data.AllApps.Edges[0].Node.URL
	user.ClientID = userInput.Data.AllApps.Edges[0].Node.ClientID
	user.Secret = userInput.Data.AllApps.Edges[0].Node.Secret

	return &user, true
}

func GetAppAndSecret(ctx *gin.Context, clientID, secret string) bool {

	user, valid := GetApp(ctx, clientID)
	if !valid {
		return valid
	}
	if user.Secret != secret {
		return false
	}
	return true

}
