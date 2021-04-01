package model

// Character entity model
type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Gender  string `json:"gender"`
}

// TODO: Add origin, location, image and episodes
