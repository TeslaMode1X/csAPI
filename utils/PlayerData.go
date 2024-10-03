package utils

type PlayerData struct {
	PlayerInfo struct {
		SteamID     string
		PersonaName string
		ProfileURL  string
		AvatarFull  string
		LocCountry  string
	}
	GamesInfo struct {
		GameCount int
		Games     []struct {
			AppID    int
			Name     string
			Playtime int
			IconURL  string
		}
	}
}
