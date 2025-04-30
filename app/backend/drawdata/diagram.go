package drawdata

const (
	DefaultDiagramColor = 0xFFFFFF
)

type Diagram struct {
	Color      int        `json:"color"`
	Components Components `json:"components"`
}
