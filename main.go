package LeaGoLegends

/*
To Do:
	Implement natural rate-limiting (e x p o n e n t i a l b a c k o f f)
*/

import (
	"fmt"

	"./config"
)

// APIInterface is the structure that holds the relevant data for your api
// There's probably more things I need to put in this struct, so I'll leave it as a struct for now
type APIInterface struct {
	apiKey string
}

// CreateInterface , given an API key, returns an interface for which a user can interact with the Riot API
func CreateInterface(apiKey string) APIInterface {
	Interface := APIInterface{}
	Interface.apiKey = apiKey

	return Interface
}

func main() {
	Token, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Token)

	Interface := CreateInterface(Token)
	fmt.Println(Interface.apiKey)
}
