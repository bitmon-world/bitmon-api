package types

type Elements struct {
	ID          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Image       string `bson:"image" json:"image"`
}
