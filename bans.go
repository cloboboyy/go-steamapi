package steamapi

import (
	"net/url"
	"strconv"
	"strings"
)

type playerBansJson struct {
	Players []PlayerBan
}

type PlayerBan struct {
	SteamId         uint64 `json:"SteamId,string"`
	CommunityBanned bool
	VACBanned       bool
	EconomyBan      string
}

func GetPlayerBans(ids []uint64, apiKey string) ([]PlayerBan, error) {
	getPlayerBans := NewSteamMethod("ISteamUser", "GetPlayerBans", 1)

	strIds := make([]string, len(ids))
	for _, id := range ids {
		strIds = append(strIds, strconv.FormatUint(id, 10))
	}

	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("steamids", strings.Join(strIds, ","))

	var resp playerBansJson
	err := getPlayerBans.Request(data, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Players, nil
}
