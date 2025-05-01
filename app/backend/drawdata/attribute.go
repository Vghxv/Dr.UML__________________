package drawdata

const (
	DefaultAttributeFontSize  = 12
	DefaultAttributeFontStyle = 0
	DefaultAttributeFontFile  = "/assets/Inkfree.ttf"
)

type Attribute struct {
	Content   string `json:"content"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	FontSize  int    `json:"fontSize"`
	FontStyle int    `json:"fontStyle"`
	FontFile  string `json:"fontFile"`
}