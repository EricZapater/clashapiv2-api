package service

import (
	"clashapiv2-api/environment"
	"clashapiv2-api/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetMembersFromClash() (model.Items, error) {
	env := environment.LoadEnvironment()
	url := fmt.Sprintf("https://api.clashroyale.com/v1/clans/%s/members", env.Clan)
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + env.Bearer

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating new request.\n", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	var items model.Items
	json.NewDecoder(resp.Body).Decode(&items)
	return items, nil
}
