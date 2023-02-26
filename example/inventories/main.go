package main

import (
	"context"
	"fmt"

	"github.com/koo04/fnar-go/fnar"
	"github.com/koo04/fnar-go/fnar/inventory"
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

	// get the inventories of the user. In this case, the auth key user
	inventories, err := inventory.GetAll(ctx, username, auth)
	if err != nil {
		panic(err)
	}

	for _, playerInventory := range inventories {
		if playerInventory.Type == inventory.Store {
			fmt.Printf("=== Base Inventory (%s) ===\n", playerInventory.Id)
			for _, item := range playerInventory.Items {
				fmt.Printf("\tItem:\t%s\n", item.Name)
				fmt.Printf("\tAmount:\t%d\n", item.Amount)
			}
		}
	}
}
