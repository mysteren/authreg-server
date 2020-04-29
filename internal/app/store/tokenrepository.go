package store

import (
	"context"
	"log"

	"gitlab.devkeeper.com/authreg/server/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	COLLECTION = "tokens"
)

type TokenRepository struct {
	store *Store
}

//
func (r *TokenRepository) Create(u *model.Token) (*model.Token, error) {
	u.ID = primitive.NewObjectID()
	_, err := GetStoreDB().Collection(COLLECTION).InsertOne(context.TODO(), u)

	return u, err
}

//
func (r *TokenRepository) Update(u *model.Token) (*model.Token, error) {

	filter := bson.D{primitive.E{Key: "_id", Value: u.ID}}
	update := bson.M{"$set": u}

	_, err := GetStoreDB().Collection(COLLECTION).UpdateOne(context.TODO(), filter, update)

	return u, err
}

//
func (r *TokenRepository) Find(id string) (*model.Token, error) {

	var token *model.Token

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	err := GetStoreDB().Collection(COLLECTION).FindOne(context.TODO(), filter).Decode(&token)

	return token, err
}

//
func (r *TokenRepository) FindAll() ([]*model.Token, error) {

	var tokens []*model.Token

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

		tokens = append(tokens, &elem)
	}

	return tokens, err
}

//
func (r *TokenRepository) Delete(id string) (int64, error) {

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	result, err := GetStoreDB().Collection(COLLECTION).DeleteOne(context.TODO(), filter)

	return result.DeletedCount, err
}
