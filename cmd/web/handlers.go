package main

import (
	"github.com/TeslaMode1X/csAPI/utils"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"os"
)

func playerHandler(w http.ResponseWriter, r *http.Request) {
	steamAPIkey := os.Getenv("STEAM_API_KEY")
	steamID := chi.URLParam(r, "steamID")

	playerSummary, err := getPlayerInfo(steamAPIkey, steamID)
	if err != nil {
		http.Error(w, "Failed to get player data", http.StatusInternalServerError)
		log.Printf("Error: %v\n", err)
		return
	}

	if len(playerSummary.Response.Players) == 0 {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	gamesSummary, err := getPlayerGamesInfo(steamAPIkey, steamID)
	if err != nil {
		http.Error(w, "Failed to get games data", http.StatusInternalServerError)
		log.Printf("Error: %v\n", err)
		return
	}

	player := playerSummary.Response.Players[0]

	playerData := utils.PlayerData{
		PlayerInfo: struct {
			SteamID     string
			PersonaName string
			ProfileURL  string
			AvatarFull  string
			LocCountry  string
		}{
			SteamID:     player.SteamID,
			PersonaName: player.PersonaName,
			ProfileURL:  player.ProfileURL,
			AvatarFull:  player.AvatarFull,
			LocCountry:  player.LocCountry,
		},
		GamesInfo: struct {
			GameCount int
			Games     []struct {
				AppID    int
				Name     string
				Playtime int
				IconURL  string
			}
		}{
			GameCount: gamesSummary.Response.GameCount,
			Games: []struct {
				AppID    int
				Name     string
				Playtime int
				IconURL  string
			}(gamesSummary.Response.Games),
		},
	}

	tmpl, err := template.New("player.layout.tmpl").Funcs(template.FuncMap{"lower": lower}).ParseGlob("./ui/html/player.layout.tmpl")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		log.Printf("Template parsing error: %v\n", err)
		return
	}

	// Передаём объединённую структуру данных в шаблон
	err = tmpl.Execute(w, playerData)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Printf("Template execution error: %v\n", err)
		return
	}
}
