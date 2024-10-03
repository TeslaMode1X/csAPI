package utils

type PlayerSummariesResponse struct {
	Response struct {
		Players []struct {
			SteamID      string `json:"steamid"`
			PersonaName  string `json:"personaname"`
			ProfileURL   string `json:"profileurl"`
			Avatar       string `json:"avatar"`
			AvatarMedium string `json:"avatarmedium"`
			AvatarFull   string `json:"avatarfull"`
			LocCountry   string `json:"loccountrycode"`
		} `json:"players"`
	} `json:"response"`
}
