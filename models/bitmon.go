package models

import (
	"context"
	"github.com/bitmon-world/bitmon-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*

	Database Structure:
	mons/ <- General monsters information
		ID/
 		   	GeneralMon
	elements/ <- In-game elements
		ID/
			ElementINfo
	modifiers/
		adv/ <- adventure algorithm modifiers

*/

type BitmonDBModel struct {
	db *mongo.Database
}

func (model *BitmonDBModel) GetMonsList() (data []types.GeneralMon, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("mons")
	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer func() {
		_ = cur.Close(ctx)
	}()
	for cur.Next(ctx) {
		var mon types.GeneralMon
		err = cur.Decode(&data)
		if err != nil {
			return
		}
		data = append(data, mon)
	}
	return
}

func (model *BitmonDBModel) GetGenMon(id string) (data types.GeneralMon, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("mons")
	filter := bson.D{{Key: "id", Value: id}}
	err = col.FindOne(ctx, filter).Decode(&data)
	return
}

func (model *BitmonDBModel) AddMon(mon types.GeneralMon) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("mons")
	_, err := col.UpdateOne(ctx, bson.D{}, bson.D{{Key: mon.ID, Value: mon}})
	return err
}

func (model *BitmonDBModel) GetElementsList() (data []types.Elements, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("elements")
	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer func() {
		_ = cur.Close(ctx)
	}()
	for cur.Next(ctx) {
		var element types.Elements
		err = cur.Decode(&element)
		if err != nil {
			return
		}
		data = append(data, element)
	}
	return
}

func (model *BitmonDBModel) GetElement(id string) (data types.Elements, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("elements")
	filter := bson.D{{Key: "id", Value: id}}
	err = col.FindOne(ctx, filter).Decode(&data)
	return
}

func (model *BitmonDBModel) AddElement(data types.Elements) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("elements")
	upsert := true
	_, err := col.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data}, &options.UpdateOptions{Upsert: &upsert})
	return err
}

func (model *BitmonDBModel) GetAdventureModifiers() (data types.AdventureModifiers, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("modifiers")
	filter := bson.D{{Key: "id", Value: "adv"}}
	err = col.FindOne(ctx, filter).Decode(&data)
	return
}

func NewDBModel(dbUri string, dbName string) *BitmonDBModel {
	opts := options.Client()
	false := false
	opts.RetryWrites = &false
	client, err := mongo.NewClient(opts.ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	return &BitmonDBModel{db: client.Database(dbName)}
}
