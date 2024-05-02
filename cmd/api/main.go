package main

import (
	"clashapiv2-api/service"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	//recuperar dades de la batalla
	location, err := time.LoadLocation("Europe/Madrid")
    if err != nil {
        fmt.Println("Error loading location:", err)
        return
    }
	now := time.Now().In(location)
	fmt.Println(now)
	ctx := context.Background()
	err = service.SetDbPlayersFromClash(ctx)
	if err != nil {
		log.Fatal(err)
	}
	service.GetRiverRaceResultsFromClash()
	//aixecar servidor http
}