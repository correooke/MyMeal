package model

type Category struct {
	ID          string   `bson:"_id,omitempty" json:"id"`
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"desc" json:"desc"`
	Image       string   `bson:"image" json:"image"`
	Dishes      []string `bson:"dishes" json:"dishes"`
}
