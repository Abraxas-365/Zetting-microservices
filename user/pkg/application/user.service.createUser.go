package service

import (
	"fmt"
	"user/pkg/core/models"
)

func (r *userService) CreateUser(user models.User) (*models.User, error) {
	fmt.Println("---CreateUserService ---")

	/*Crear el usuario*/
	newUser, err := r.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	/*Crear carpeta con el nombre del usuario para guardar fotos y datos*/
	// folderName := fmt.Sprintf("%v", newUser.ID)
	// if err := os.Mkdir("./static/images/"+folderName, os.ModePerm); err != nil {
	// 	return nil, err
	// }

	return newUser, nil

}
