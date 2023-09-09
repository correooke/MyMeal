package model

type Dish struct {
	ID          string  `bson:"_id,omitempty" json:"id"`
	Name        string  `bson:"name" json:"name"`
	Price       float64 `bson:"price" json:"price"`
	Description string  `bson:"desc" json:"desc"`
}
