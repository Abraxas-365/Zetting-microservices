package service

func (r *userService) CheckEmailExist(email string) bool {

	user, _ := r.userRepo.CheckEmailExist(email)
	if user == nil {
		return false
	}
	return true

}
