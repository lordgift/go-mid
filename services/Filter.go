package services

import (
	"context"
	"github.com/Nerzal/gocloak/v7"
	"github.com/gin-gonic/gin"
	"log"
)

func (in *Incoming) filter(c *gin.Context) {

	hostname := ""
	clientID := ""
	clientSecret := ""
	realms := ""
	username := ""
	password := ""

	client := gocloak.NewClient(hostname)
	ctx := context.Background()

	token, err := client.Login(ctx, clientID, clientSecret, realms, username, password)
	if err != nil {
		panic("Login failed:" + err.Error())
	}

	log.Println(token.AccessToken)

	newtoken, err := client.RefreshToken(ctx, token.RefreshToken, clientID, clientSecret, realms)
	if err != nil {
		panic("Renew failed:" + err.Error())
	}
	log.Println(newtoken.AccessToken)

	err = client.Logout(ctx, clientID, clientSecret, realms, token.RefreshToken)
	if err != nil {
		panic("Renew failed:" + err.Error())
	}
	log.Println(newtoken.AccessToken)

	//newtoken2, err := client.RefreshToken(ctx, token.RefreshToken, clientID, clientSecret, realms)
	//if err != nil {
	//	panic("Renew failed:"+ err.Error())
	//}
	//log.Println(newtoken2.AccessToken)
}
