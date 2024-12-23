package models

type Error struct {
	Code    int    `json:"StatusCode"`
	Message string `json:"ErrorMessage"`
}
