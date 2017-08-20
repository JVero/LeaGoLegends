package main

/*
To Do:
	Implement natural rate-limiting (e x p o n e n t i a l b a c k o f f)
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"./config"
)

// APIInterface is the structure that holds the relevant data for your api
// There's probably more things I need to put in this struct, so I'll leave it as a struct for now
type APIInterface struct {
	apiKey    string
	rateLimit time.Duration
	throttler <-chan time.Time
	// probably want rate-limiting information in the interface itself, so lets try time.Ticker
}

// ChampionMasteryUnit is the response to the API query
type ChampionMasteryUnit struct {
	ChampionLevel                int  `json:"championLevel"`
	ChestGranted                 bool `json:"chestGranted"`
	ChampionPoints               int  `json:"championPoints"`
	ChampionPointsSinceLastLevel int  `json:"championPointsSinceLastLevel"`
	PlayerID                     int  `json:"playerId"`
	ChampionPointsUntilNextLevel int  `json:"championPointsUntilNextLevel"`
	TokensEarned                 int  `json:"tokensEarned"`
	ChampionID                   int  `json:"championId"`
	LastPlayTime                 int  `json:"lastPlayTime"`
}

// ChampionMasteryResponse dklfajd
type ChampionMasteryResponse []ChampionMasteryUnit

// CreateAPIInterface , given an API key, returns an interface for which a user can interact with the Riot API
func CreateAPIInterface(apiKey string, ratelimit time.Duration) APIInterface {
	Interface := APIInterface{}
	Interface.apiKey = apiKey
	Interface.rateLimit = ratelimit // 100 requests per minute
	Interface.throttler = time.Tick(time.Second / Interface.rateLimit)

	return Interface
}

// GetChampionMasteryForID returns an array of champion masteries
func (a *APIInterface) GetChampionMasteryForID(summonerID string) *ChampionMasteryResponse {

	req := "https://na1.api.riotgames.com/lol/champion-mastery/v3/champion-masteries/by-summoner/" + summonerID + "?api_key=" + a.apiKey
	resp, err := http.Get(req)

	// TODO, handle rate limiting here or at the end

	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	champresponse := new(ChampionMasteryResponse)
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &champresponse)

	return champresponse
}

func main() {
	Token, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Token)
	Interface := CreateAPIInterface(Token, 100)

	// Implementation of throttler
	for i := 0; i < 10; i++ {
		<-Interface.throttler
		fmt.Println(time.Now())
	}

	champresponse := *Interface.GetChampionMasteryForID("59459147")
	fmt.Printf("%+v", champresponse[0])
}
