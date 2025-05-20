package drawdata

const (
	DefaultDiagramColor = "#FFFFFF"
	Margin              = 4
	LineWidth           = 2
)

type Diagram struct {
	Margin       int           `json:"margin"`
	LineWidth    int           `json:"lineWidth"`
	Color        string        `json:"color"`
	Gadgets      []Gadget      `json:"gadgets"`
	Associations []Association `json:"associations"`
}
