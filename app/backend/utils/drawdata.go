package utils

const (
	DrawDataMargin    = 4
	DrawDataLineWidth = 2
)

type DrawDataGadget struct {
	GadgetType string              `json:"gadgetType"`
	X          int                 `json:"x"`
	Y          int                 `json:"y"`
	Layer      int                 `json:"layer"`
	Height     int                 `json:"height"`
	Width      int                 `json:"width"`
	Color      int                 `json:"color"`
	Attributes []DrawDataAttribute `json:"attributes"`
}

type DrawDataAttribute struct {
	Content   string `json:"content"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	FontSize  int    `json:"fontSize"`
	FontStyle int    `json:"fontStyle"`
	FontFile  string `json:"fontFile"`
}
