package steam

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserProfiles struct {
	Response Response `json:"response"`
}
type Response struct {
	Players []Players `json:"players"`
}

type Players struct {
	Steamid                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Avatarhash               string `json:"avatarhash"`
	Lastlogoff               int    `json:"lastlogoff"`
	Personastate             int    `json:"personastate"`
	Realname                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	Timecreated              int    `json:"timecreated"`
	Personastateflags        int    `json:"personastateflags"`
	Loccountrycode           string `json:"loccountrycode"`
	Locstatecode             string `json:"locstatecode"`
}

func GetProfiles(apiKey string, steamid string) (UserProfiles, error) {
	var userProfiles UserProfiles
	client := http.DefaultClient

	url := fmt.Sprintf(BaseUrl+"ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", apiKey, steamid)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return userProfiles, err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return userProfiles, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return userProfiles, err
	}

	err = json.Unmarshal(body, &userProfiles)
	if err != nil {
		return userProfiles, err
	}

	return userProfiles, err
}
