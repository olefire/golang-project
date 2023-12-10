package models

import "time"

type SignUpInput struct {
	Name               string    `json:"name" bson:"name" binding:"required"`
	Email              string    `json:"email" bson:"email" binding:"required"`
	Password           string    `json:"password" bson:"password" binding:"required,min=8"`
	Role               string    `json:"role" bson:"role"`
	VerificationCode   string    `json:"verificationCode,omitempty" bson:"verificationCode,omitempty"`
	ResetPasswordToken string    `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	ResetPasswordAt    time.Time `json:"resetPasswordAt,omitempty" bson:"resetPasswordAt,omitempty"`
	Verified           bool      `json:"verified" bson:"verified"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
}

type SignInInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
