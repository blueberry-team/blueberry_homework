package entities

import "time"

type NameEntity struct {
	Name string `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
}
