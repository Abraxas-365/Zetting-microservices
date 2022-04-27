package events

type UserCreated struct {
	UserName    string
	PerfilImage string
}

func (e UserCreated) Name() string {
	return "event.user.created"
}
