package models

import "github.com/kamva/mgm/v3"

type Todo struct {
	mgm.DefaultModel `bson:",inline"`

	// ID          string `json:"id" bson:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
