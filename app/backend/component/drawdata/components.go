package drawdata

const (
	Margin    = 4
	LineWidth = 2
)

type Components struct {
	Margin     int         `json:"margin"`
	LineWidth  int         `json:"lineWidth"`
	Components []Component `json:"components"`
}

// TODO: can be either Gadget or Association
type Component interface {
}
