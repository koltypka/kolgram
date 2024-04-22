package methods

import "encoding/json"

type GetWebhookInfo struct {
	Ok     bool              `json:"ok"`
	Result WebhookInfoResult `json:"result"`
}

type WebhookInfoResult struct {
	Url                  string `json:"url"`
	HasCustomCertificate bool   `json:"has_custom_certificate"`
	PendingUpdateCount   int8   `json:"pending_update_count"`
}

func (Telegram *Telegram) GetWebhookInfo() GetWebhookInfo {
	tmpResult, err := Telegram.Get("/getWebhookInfo")
	result := GetWebhookInfo{}

	if err != nil {
		return result
	}

	json.Unmarshal(tmpResult.Body, &result)

	return result
}
