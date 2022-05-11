package repository

import (
	"actor-service/user/core/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mongoRepository) FilterUsers(filter models.Filter, page int) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	filterQuery := make(map[string]interface{})
	switch true {

	case filter.AgeMin != 0 && filter.AgeMax != 0:
		filterQuery["age"] = bson.M{"$gt": filter.AgeMin, "$lt": filter.AgeMax + 1}

	case filter.AgeMax != 0:
		filterQuery["age"] = bson.M{"$gt": 0, "$lt": filter.AgeMax + 1}

	case len(filter.Gender) != 0:
		filterQuery["gender"] = filter.Gender

	case len(filter.Features.Body) != 0:
		filterQuery["features.body"] = filter.Features.Body

	case filter.Features.Height != 0:
		filterQuery["features.height"] = filter.Features.Height

	case len(filter.Features.EyeColor) != 0:
		filterQuery["features.eye_color"] = filter.Features

	case len(filter.Features.FacialHair) != 0:
		filterQuery["features.facial_hair"] = filter.Features.FacialHair

	case len(filter.Features.HairColor) != 0:
		filterQuery["features.hair_color"] = filter.Features.HairColor

	case len(filter.Features.HairType) != 0:
		filterQuery["features.hair_type"] = filter.Features.HairType

	case len(filter.Features.HairZise) != 0:
		filterQuery["features.hair_zise"] = filter.Features.HairZise

	case len(filter.Features.Skin) != 0:
		filterQuery["features.skin"] = filter.Features.Skin

	case len(filter.Features.Skills) != 0:
		filterQuery["features.skills"] = filter.Features.Skills
	}

	users := models.Users{}
	options := options.Find()
	options.SetLimit(20)
	options.SetSkip(20 * (int64(page) - 1))

	cur, err := collection.Find(ctx, filterQuery, options)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
