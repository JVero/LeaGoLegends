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
	"strings"
	"time"

	"./config"
)

// APIInterface is the structure that holds the relevant data for your api
// There's probably more things I need to put in this struct, so I'll leave it as a struct for now
type APIInterface struct {
	apiKey    string
	rateLimit time.Duration
	throttler <-chan time.Time
	// probably want rate-limiting information in the interface itself, so lets try <-chan time.Time
}

// CreateAPIInterface , given an API key, returns an interface for which a user can interact with the Riot API
func CreateAPIInterface(apiKey string, ratelimit time.Duration) APIInterface {
	return APIInterface{apiKey, ratelimit, time.Tick(time.Second / ratelimit)}
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

// ChampionMasteryResponse is the slice of ChampionMasteryUnits that gets returned by the API
type ChampionMasteryResponse []ChampionMasteryUnit

// Essentially a pretty-print for the structure
func (u *ChampionMasteryUnit) String() string {
	b, _ := json.MarshalIndent(u, "", "  ")
	return string(b)
}

// ParseRateLimitPairsFromHeaders is a function that reads a header and returns the rate limit for your request
func ParseRateLimitPairsFromHeaders(h http.Header) map[string]string {
	rateArr := strings.Split(h["X-App-Rate-Limit-Count"][0], ",")
	rates := make(map[string]string, 3)
	for _, val := range rateArr {
		d := strings.Split(val, ":")
		rates[d[1]] = d[0]
	}
	return rates

}

// GetChampionMasteryForID returns an array of champion masteries
func (a *APIInterface) GetChampionMasteryForID(summonerID string) *ChampionMasteryResponse {

	req := "https://na1.api.riotgames.com/lol/champion-mastery/v3/champion-masteries/by-summoner/" + summonerID + "?api_key=" + a.apiKey
	resp, err := http.Get(req)
	ParseRateLimitPairsFromHeaders(resp.Header)
	// TODO, handle response codes
	// TODO, handle rate limiting here or at the end

	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	champresponse := new(ChampionMasteryResponse)
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &champresponse)

	return champresponse
}

//
func main() {
	Token, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Token)
	Interface := CreateAPIInterface(Token, 10)

	// Implementation of throttler
	/*for i := 0; i < 10; i++ {
		<-Interface.throttler
		fmt.Println(time.Now())
	}*/

	// Get and print the ChampionMastery slice.  TODO:  Figure out how to print it all pretty-like
	champresponse := *Interface.GetChampionMasteryForID("59459147") // 59459147 is my user id :^)
	fmt.Println(champresponse[0].String())
}
