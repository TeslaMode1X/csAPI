package utils

type OwnedGamesResponse struct {
	Response struct {
		GameCount int `json:"game_count"`
		Games     []struct {
			AppID    int    `json:"appid"`
			Name     string `json:"name"`
			Playtime int    `json:"playtime_forever"`
			IconURL  string `json:"img_icon_url"`
		} `json:"games"`
	} `json:"response"`
}
