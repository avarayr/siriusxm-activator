package main

type LoginTokens struct {
	ClaimsToken struct {
		Exp   int64  `json:"exp"`
		Value string `json:"value"`
	} `json:"claims_token"`
	RefreshToken string `json:"refresh_token"`
}
