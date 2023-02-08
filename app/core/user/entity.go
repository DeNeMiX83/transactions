package user

import (
	"github.com/google/uuid"
)

type user_entity struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func NewUser(username, email, hashed_password string) user_entity {
	new_user := user_entity{
		Id:             uuid.New().String(),
		Username:       username,
		Email:          email,
		HashedPassword: hashed_password,
	}
	return new_user
}

type balance struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Amount int    `json:"amount"`
}

func NewBalance(user_id string, amount int) balance {
	new_balance := balance{
		Id:     uuid.New().String(),
		UserId: user_id,
		Amount: amount,
	}
	return new_balance
}
