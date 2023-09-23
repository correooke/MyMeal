package model

type Dish struct {
	ID          string  `bson:"_id,omitempty" json:"id"`
	Name        string  `bson:"name" json:"name"`
	Description string  `bson:"desc" json:"desc"`
	Image       string  `bson:"image" json:"image"`
	Price       float64 `bson:"price" json:"price"`
}

func (d Dish) GetID() string {
	return d.ID
}
