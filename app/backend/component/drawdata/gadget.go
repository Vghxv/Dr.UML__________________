package drawdata

type Gadget struct {
	GadgetType int           `json:"gadgetType"`
	X          int           `json:"x"`
	Y          int           `json:"y"`
	Layer      int           `json:"layer"`
	Height     int           `json:"height"`
	Width      int           `json:"width"`
	Color      int           `json:"color"`
	Attributes [][]Attribute `json:"attributes"`
}
