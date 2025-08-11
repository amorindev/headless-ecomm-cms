package core

type CreatePhoneReq struct {
	Number      string `json:"number"`
	CountryCode string `json:"country_code"`
	CountryIso  string `json:"country_iso" `
}
