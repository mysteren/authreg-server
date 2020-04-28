package store

import (
	"context"
	"log"

	"gitlab.devkeeper.com/authreg/server/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	COLLECTION = "tokens"
)

type TokenRepository struct {
	store *Store
}

func (r *TokenRepository) Create(u *model.Token) (*model.Token, error) {
	return nil, nil
}

func (r *TokenRepository) FindById(login string) (*model.Token, error) {
	return nil, nil
}

func (r *TokenRepository) FindAll() ([]*model.Token, error) {

	var Tokens []*model.Token

	cur, err := GetStoreDB().Collection(COLLECTION).Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Token
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		Tokens = append(Tokens, &elem)
	}

	return Tokens, err
}
