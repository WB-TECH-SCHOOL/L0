package models

import "encoding/json"

type Order struct {
	ID   string          `json:"id"`
	Data json.RawMessage `json:"data"`
}
