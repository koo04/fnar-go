package main

import (
	"context"
	"fmt"

	"github.com/koo04/fnar-go/fnar"
	"github.com/koo04/fnar-go/fnar/ship"
)

func main() {
	ctx := context.Background()

	// authenticate to Fnar.net
	auth, err := fnar.Login("ThatKooGuy", "supersecretpassword")
	if err != nil {
		panic(err)
	}

	// get this username. this is an optional step, but makes it
	// easy for applications to dynamically get the username of
	// the auth key owner
	username, err := auth.GetUsername()
	if err != nil {
		panic(err)
	}

	// get the ships of the user. in this case, the auth key user
	ships, err := ship.GetAll(ctx, username, auth)
	if err != nil {
		panic(err)
	}

	for _, ship := range ships {
		fmt.Println(ship.Registration)
	}
}
