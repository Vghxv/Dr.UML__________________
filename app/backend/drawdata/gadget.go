package drawdata

// a default gadget color (grey)
const DefaultGadgetColor = "#808080"

type Gadget struct {
	GadgetType int           `json:"gadgetType"`
	X          int           `json:"x"`
	Y          int           `json:"y"`
	Layer      int           `json:"layer"`
	Height     int           `json:"height"`
	Width      int           `json:"width"`
	Color      string        `json:"color"`
	IsSelected bool          `json:"isSelected"`
	Attributes [][]Attribute `json:"attributes"`
}
