package entity

type TokenData struct {
	AccessToken         string `json:"AccessToken"`
	AccessTokenExpires  string `json:"AccessTokenExpires"`
	RefreshToken        string `json:"RefreshToken"`
	RefreshTokenExpires string `json:"RefreshTokenExpires"`
}
