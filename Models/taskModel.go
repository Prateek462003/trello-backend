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

func (t *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
	aux := &struct {
		Image *string `json:"image,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Image != nil {
		t.Image = sql.NullString{String: *aux.Image, Valid: true}
	} else {
		t.Image = sql.NullString{String: "", Valid: false}
	}

	return nil
}
