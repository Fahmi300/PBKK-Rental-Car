package models


type Authorization struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}