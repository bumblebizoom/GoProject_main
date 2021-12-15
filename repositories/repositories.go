package repositories

import "GoProject_1/repositories/models"

type UserRepositoryInterFace interface {
	GetByEmail(email string) models.User
}
