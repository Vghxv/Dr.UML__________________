package drawdata

type Attribute struct {
	Content   string `json:"content"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	FontSize  int    `json:"fontSize"`
	FontStyle int    `json:"fontStyle"`
	FontFile  string `json:"fontFile"`
}
