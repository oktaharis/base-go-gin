package dto

type AccountLoginReq struct {
	Username string `json:"uname" binding:"required,max=16"`
	Password string `json:"paswd" binding:"required,min=8,max=255"`
}

type AccountLoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
