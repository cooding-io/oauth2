package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/Ensena/core/graphql-client"
	"github.com/gin-gonic/gin"
)

type outputUser struct {
	User       string `json:"user"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
	Iss        string `json:"iss"`
	Iat        int    `json:"iat"`
}

func (o *outputUser) ConvertFromGraphql(i *graphqlUser) {
	o.Email = i.Data.UserByID.Email
	o.GivenName = i.Data.UserByID.Name
	o.FamilyName = i.Data.UserByID.LastName
	o.User = string(i.Data.UserByID.ID)
	o.Picture = i.Data.UserByID.Picture
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Picture   string `json:"picture"`
	MoodleUDP bool   `json:"moodleUdp"`
}

type graphqlUser struct {
	Data struct {
		UserByID User `json:"userById"`
	} `json:"data"`
}

type graphqlUserByEmail struct {
	Data struct {
		AllUsers struct {
			TotalCount int `json:"totalCount"`
			Edges      []struct {
				Node User `json:"node"`
			} `json:"edges"`
		} `json:"allUsers"`
	} `json:"data"`
}

func GetUser(ctx *gin.Context, userID string) []byte {

	g := fmt.Sprintf(`{
		userById(id: %s) {
		  id
		  name
		  lastName
		  email
		  picture
		  moodleUdp
		}
	  }`, userID)

	response, err := graphql.Query(ctx, g)
	if err != nil {
		return []byte("")
	}
	userInput := graphqlUser{}
	json.Unmarshal(response, &userInput)
	userOut := outputUser{}
	userOut.ConvertFromGraphql(&userInput)
	out, _ := json.Marshal(&userOut)
	return out
}

func GetUserByEmail(ctx *gin.Context, email string) (*User, bool) {

	g := fmt.Sprintf(`{
		allUsers(condition: { email: "%s" }) {
		  totalCount
		  edges {
			node {
			  id
			  name
			  lastName
			  email
			  password
			  picture
			  moodleUdp
			}
		  }
		}
	  }
	  `, email)

	response, err := graphql.Query(ctx, g)
	if err != nil {
		return &User{}, false
	}
	user := graphqlUserByEmail{}
	json.Unmarshal(response, &user)

	if user.Data.AllUsers.TotalCount == 0 {
		return &User{}, false
	}
	return &user.Data.AllUsers.Edges[0].Node, true
}
