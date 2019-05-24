package models

type STSSession struct {
	SessionId    string `json:"sessionId"`
	SessionKey   string `json:"sessionKey"`
	SessionToken string `json:"sessionToken"`
}

type TemporaryCredentials struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	LoginUrl  string `json:"login_url"`
}
