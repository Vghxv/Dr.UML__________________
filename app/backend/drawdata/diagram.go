package drawdata

const (
	DefaultDiagramColor = 0xFFFFFF
	Margin              = 4
	LineWidth           = 2
)

type Diagram struct {
	Margin       int           `json:"margin"`
	LineWidth    int           `json:"lineWidth"`
	Color        int           `json:"color"`
	Gadgets      []Gadget      `json:"gadgets"`
	Associations []Association `json:"associations"`
}
