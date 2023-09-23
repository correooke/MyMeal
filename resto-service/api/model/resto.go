package model

type Resto struct {
	ID            string            `bson:"_id,omitempty" json:"id"`
	Name          string            `bson:"name" json:"name"`
	Description   string            `bson:"desc" json:"desc"`
	Configuration map[string]string `bson:"config" json:"config"`
	MainImage     string            `bson:"mainImage" json:"mainImage"`
	DesignType    string            `bson:"designType" json:"designType"`
	Categories    []string          `bson:"categories" json:"categories"`
	Currency      string            `bson:"currency" json:"currency"`
}

func (r Resto) GetID() string {
	return r.ID
}
