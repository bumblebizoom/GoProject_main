package filesystem

import (
	"GoProject_1/repositories/models"
	"encoding/json"
	"io"
	"os"
)

type UserFileRepository struct {
}

func (ufr UserFileRepository) GetByEmail(email string) (user models.User) {

	userData := []byte{}
	file, err := os.Open("./datastore/files/user_1.json")
	if err != nil {
		return models.User{}
	}
	defer file.Close()

	for {
		chunk := []byte{}
		_, err = file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		userData = append(userData, chunk...)
	}

	err = json.Unmarshal(userData, &user)
	if err != nil {
		panic(err)
	}

	return user
}

