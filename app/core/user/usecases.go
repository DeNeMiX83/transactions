package user

import (
	"api/app/infrastructure/db/postgresql/gateway/user"
	"fmt"
)

type UserUseCase interface {
	IncreaseBalance(user_id string, balanc_id string, amount int) error
	DecreaseBalance(user_id string, balanc_id string, amount int) error
}

type usecase struct {
	gateway user.Gateway
}

func (u *usecase) IncreaseBalance(user_id string, balanc_id string, amount int) error{
	err := u.gateway.IncreaseBalance(user_id, balanc_id, amount)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (u *usecase) DecreaseBalance(user_id string, balanc_id string, amount int) error{
	err := u.gateway.DecreaseBalance(user_id, balanc_id, amount)
	if err != nil {
		return fmt.Errorf("%v", "Insufficient funds")
	}
	return nil
}

func NewUserUseCase(gateway user.Gateway) UserUseCase {
	return &usecase{gateway: gateway}
}