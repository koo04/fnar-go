package main

import (
	"context"
	"fmt"

	"github.com/koo04/fnar-go/fnar"
	wf "github.com/koo04/fnar-go/fnar/workforce"
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

	// get the worforces of the user. in this case, the auth key user
	workforces, err := wf.GetAll(ctx, username, auth)
	if err != nil {
		panic(err)
	}

	for _, workforce := range workforces {
		if workforce.Required > 0 {
			fmt.Printf("%s - %s\n", workforce.PlanetNaturalId, workforce.TypeName)
			fmt.Printf("\tRequired Population: %d\n", workforce.Required)
			for _, workforceNeeds := range workforce.Needs {
				if workforceNeeds.Essential {
					fmt.Printf("\tNeeds %.2f of %s per day\n", workforceNeeds.UnitsPerOneHundred*float32(workforce.Population)/100, workforceNeeds.MaterialTicket)
				}
			}
		}
	}
}
