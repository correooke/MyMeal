
package model

type RestaurantGroup struct {
    ID          string   `bson:"_id,omitempty" json:"id"`
    Name        string   `bson:"name" json:"name"`
    Description string   `bson:"description" json:"description"`
    Headquarters string  `bson:"headquarters" json:"headquarters"`
    Brands      []string `bson:"brands" json:"brands"`
    Founded     string   `bson:"founded" json:"founded"`
    Website     string   `bson:"website" json:"website"`
}

func (rg RestaurantGroup) GetID() string {
    return rg.ID
}
