package model

// Menu represents a menu with its details
type Menu struct {
	ID         string   `json:"id" bson:"_id,omitempty"`
	Name       string   `json:"name"`
	StartTime  string   `json:"startTime"`
	EndTime    string   `json:"endTime"`
	Categories []string `json:"categoryIDs"`
}

func (m Menu) GetID() string {
	return m.ID
}
