package account

import "exness/internal/models"

func NewCreator() *creator {
	return &creator{}
}

type creator struct {
}

func (c *creator) CreateAccount() (*models.Account, error) {

}
