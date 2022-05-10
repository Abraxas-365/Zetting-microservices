package repository

import (
	"context"
	"user/user/core/events"
	"user/user/core/models"
)

func (r *mongoRepository) CreateUser(user models.User) (events.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return events.UserCreated{
		ID:          user.ID,
		PerfilImage: user.PerfilImage,
		UserName:    user.Name,
		Gender:      user.Gender,
		Age:         user.Age,
		Profession:  user.Profession,
	}, nil
}
