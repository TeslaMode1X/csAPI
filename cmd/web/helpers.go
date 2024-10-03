package main

import (
	"encoding/json"
	"fmt"
	"github.com/TeslaMode1X/csAPI/utils"
	"net/http"
	"strings"
)

func getPlayerInfo(steamAPIkey, steamID string) (*utils.PlayerSummariesResponse, error) {
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", steamAPIkey, steamID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from Steam API: %v", err)
	}
	defer resp.Body.Close()

	var result utils.PlayerSummariesResponse

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return &result, nil
}

func getPlayerGamesInfo(steamAPIkey, steamID string) (*utils.OwnedGamesResponse, error) {
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&include_appinfo=1&format=json", steamAPIkey, steamID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching games data from Steam API: %v", err)
	}
	defer resp.Body.Close()

	var gamesResult utils.OwnedGamesResponse
	if err = json.NewDecoder(resp.Body).Decode(&gamesResult); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return &gamesResult, nil
}

func lower(s string) string {
	return strings.ToLower(s)
}
