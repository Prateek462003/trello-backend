package models

import (
	"database/sql"
	"encoding/json"
)

type Task struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Image       sql.NullString `json:"image"`
	ActivityId  int            `json:"activity_id"`
}

func (t Task) MarshalJSON() ([]byte, error) {
	type Alias Task
	return json.Marshal(&struct {
		Image *string `json:"image,omitempty"`
		Alias
	}{
		Image: func() *string {
			if t.Image.Valid {
				return &t.Image.String
			}
			return nil
		}(),
		Alias: (Alias)(t),
	})
}
