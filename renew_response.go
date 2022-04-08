package main

type RenewResponse struct {
	Opstatus int    `json:"opstatus"`
	SeqValue string `json:"seqValue"`
	Errors   []struct {
		ResultCode string `json:"resultCode"`
	} `json:"errors"`
	HTTPStatusCode int `json:"httpStatusCode"`
}
