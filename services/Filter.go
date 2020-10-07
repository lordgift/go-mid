package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nerzal/gocloak/v7"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

	_, claims, err := client.DecodeAccessToken(ctx, newtoken.AccessToken, realms, "")
	if err != nil {
		panic("Renew failed:" + err.Error())
	}

	//dereference & extract value from claims
	fmt.Printf("Issuer: %v\n", (*claims)["iss"])

}

func (in *Incoming) filter2(c *gin.Context) {

	hostname := ""
	clientID := ""
	//clientSecret := ""
	realms := ""
	username := ""
	password := ""

	payload := strings.NewReader("client_id=" + clientID + "&username=" + username + "&password=" + password + "&grant_type=password")

	req, err := http.NewRequest("POST", hostname+"auth/realms/"+realms+"/protocol/openid-connect/token", payload)
	//request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatal(err)
	}

	var client http.Client
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var m Message
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}
	print(m.AC)

	//fmt.Println(string(body))
}

type Message struct {
	AC string `json:"access_token"`
}
