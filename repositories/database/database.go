package database

import "GoProject_1/repositories/models"

type UserDBRepository struct {
}

func (u UserDBRepository) GetByEmail(email string) models.User {

	panic("implement me")
}
