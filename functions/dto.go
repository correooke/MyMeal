
package dto

type DishDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"desc"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

type CategoryDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Image       string    `json:"image"`
	Dishes      []DishDTO `json:"dishes"`
}

type RestoDTO struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"desc"`
	Address       string         `json:"address"`
	Configuration map[string]string `json:"config"`
	MainImage     string         `json:"mainImage"`
	DesignType    string         `json:"designType"`
	Currency      string         `json:"currency"`
	Owner         string         `json:"owner"`
	Menues        map[string][]string `json:"menu"`
	Categories    []CategoryDTO  `json:"categories"`
}
