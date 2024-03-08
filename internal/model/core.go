package model

type Pagination struct {
	Items    interface{} `json:"items"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
}

type Board struct {
	Name               string         `json:"name"`
	CreatedAt          int            `json:"created_at"`
	Shape              [][]int        `json:"shape"`
	Palette            []PaletteEntry `json:"palette"`
	MaxAvailablePixels int            `json:"max_available_pixels"`
}

type PaletteEntry struct {
	Name  string `json:"name"`
	Color int    `json:"color"`
}

type Placement struct {
	Position int `json:"position"`
	Color    int `json:"color"`
	Modified int `json:"modified"`
}

// HTTP Responses

type InfoResponse struct {
	Name       string   `json:"name"`
	Version    string   `json:"version"`
	Source     string   `json:"source"`
	Extensions []string `json:"extensions"`
}

type AccessResponse struct {
	Permissions []string `json:"permissions"`
}
