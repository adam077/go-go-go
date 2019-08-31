package weibo

import (
	"encoding/json"
	"go-go-go/src/utils"
)

func GetGroupsFromContacts(cookie string) []string {
	result := make([]string, 0)
	a, err := utils.QueryGet("https://api.weibo.com/webim/2/direct_messages/contacts.json?source=209678993",
		nil, map[string]string{"Cookie": cookie, "Referer": "https://api.weibo.com/chat/"})
	if err == nil {
		var temp WeiboContacts
		json.Unmarshal(a, &temp)
		for _, user := range temp.Contacts {
			if user.User.Type == 2 {
				result = append(result, user.User.Idstr)
			}
		}
	}
	return result
}

type WeiboContacts struct {
	Contacts []struct {
		User struct {
			Idstr string `json:"idstr"`
			Type  int    `json:"type"`
		} `json:"user"`
	} `json:"contacts"`
}
