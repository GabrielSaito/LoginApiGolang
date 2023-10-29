package models

type PasswordModel struct {
	NewPassword string `json: "newPassword"`
	OldPassword string `json: "oldPassword"`
}
