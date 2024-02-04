package model

type Resto struct {
	ID            string              `bson:"_id,omitempty" json:"id"`
	Name          string              `bson:"name" json:"name"`
	Description   string              `bson:"desc" json:"desc"`
	Address       string              `bson:"address" json:"address"`
	Configuration map[string]string   `bson:"config" json:"config"`
	MainImage     string              `bson:"mainImage" json:"mainImage"`
	DesignType    string              `bson:"designType" json:"designType"`
	Currency      string              `bson:"currency" json:"currency"`
	Owner         string              `bson:"owner" json:"owner"`
	Menues        map[string][]string `bson:"menu" json:"menu"`
}

func (r Resto) GetID() string {
	return r.ID
}
