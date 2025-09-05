package models

type UserInfo struct {
	Username string `json:"username"`
	UserID int `json:"userID"`
	Role string `json:"role"`
}