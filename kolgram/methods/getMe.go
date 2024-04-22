package methods

import "encoding/json"

type GetMe struct {
	Ok     bool     `json:"ok"`
	Result MeResult `json:"result"`
}

type MeResult struct {
	Id                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	UserName                string `json:"username"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business"`
}

func (Telegram *Telegram) GetMe() GetMe {
	tmpResult, err := Telegram.Get("/getMe")
	result := GetMe{}

	if err != nil {
		return result
	}

	json.Unmarshal(tmpResult.Body, &result)

	return result
}
