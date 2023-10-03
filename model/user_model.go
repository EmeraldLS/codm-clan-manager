package model

type User struct { // Player Registration
	PlayerCode   int    `json:"player_code,omitempty" bson:"player_code,omitempty"`
	PlayerID     string `json:"player_id,omitempty" bson:"player_id,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	RegisteredAt string `json:"registered_at,omitempty" bson:"registered_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
