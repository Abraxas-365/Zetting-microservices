package application

import (
	"fmt"
	"user/pkg/core/models"

	"github.com/google/uuid"
)

func (r *userApplication) UpdateUser(userDataToUpdate models.User, userId uuid.UUID) error {

	// userDataToUpdate, _ := r.userRepo.GetUserByEmail(email)
	updateQuery := make(map[string]interface{})

	if len(userDataToUpdate.Name) > 0 {
		updateQuery["name"] = userDataToUpdate.Name
	}

	if len(userDataToUpdate.Password) > 0 {
		/*funcion de crear password*/
		fmt.Println("funcion para encriptar el password")
		updateQuery["password"] = userDataToUpdate.Password
	}

	if len(userDataToUpdate.PerfilImage) > 0 {
		updateQuery["perfil_image"] = userDataToUpdate.PerfilImage
	}

	if len(userDataToUpdate.Contact.Country) > 0 {
		updateQuery["contact.country"] = userDataToUpdate.Contact.Country
	}

	if len(userDataToUpdate.Contact.Email) > 0 {
		updateQuery["contact.email"] = userDataToUpdate.Contact.Email
	}

	if len(userDataToUpdate.Contact.Phone) > 0 {
		updateQuery["contact.phone"] = userDataToUpdate.Contact.Phone
	}

	err := r.userRepo.UpdateUser(updateQuery, userId)
	if err != nil {
		return err
	}

	return nil
}
