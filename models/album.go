package models

import (
	"encoding/json"
	"time"
)

type Album struct {
	ID     uint    			   `json:"id" gorm:"primary_key"`
	Title  string  			   `json:"title"`
	Artist string              `json:"artist"`
	Price  float64             `json:"price"`
}


type Interaction struct {
    ID         string          `gorm:"type:uuid;primaryKey" json:"id"`
    CreatedAt  time.Time       `gorm:"index" json:"created_at"`
    UpdatedAt  time.Time       `gorm:"index" json:"updated_at"`
    Settings   json.RawMessage `json:"settings" swaggertype:"object"`
    Messages   []Message       `gorm:"foreignKey:InteractionID;references:ID" json:"messages"`
}


type Message struct {
    ID             string      `gorm:"type:uuid;primaryKey" json:"id"`
    CreatedAt      time.Time   `gorm:"index" json:"created_at"`
    Role           string      `json:"role"`
    Content        string      `json:"content"`
    InteractionID  string      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"interaction_id"`
    Interaction    Interaction `gorm:"foreignKey:ID" json:"-"`
}

func (Interaction) TableName() string {
    return "interaction"
}

func (Message) TableName() string {
	return "message"
}