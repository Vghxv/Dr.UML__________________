package drawdata

const (
	Margin    = 4
	LineWidth = 2
)

type Components struct {
	Margin    int      `json:"margin"`
	LineWidth int      `json:"lineWidth"`
	Gadgets   []Gadget `json:"components"`
	// Associations []Association `json:"associations"`
}
