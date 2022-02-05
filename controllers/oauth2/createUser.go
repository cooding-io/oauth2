package controllers

import (
	"fmt"

	"github.com/Ensena/core/graphql-client"
	"github.com/gin-gonic/gin"
)

func createUser(ctx *gin.Context, user *User) {

	g := fmt.Sprintf(`mutation MyMutation {
		__typename
		createUser(input: {user: {lastName: "%s", name: "%s", picture: "", email: "%s", moodleUdp: false, ready: true, role: "basic"}}) {
		  clientMutationId
		}
	  }`, user.LastName, user.Name, user.Email)
	graphql.Query(ctx, g)

}
