package drawdata

type Association struct {
	AssType    int            `json:"assType"`
	Layer      int            `json:"layer"`
	StartX     int            `json:"startX"`
	StartY     int            `json:"startY"`
	EndX       int            `json:"endX"`
	EndY       int            `json:"endY"`
	DeltaX     int            `json:"deltaX"`
	DeltaY     int            `json:"deltaY"`
	Attributes []AssAttribute `json:"attributes"`
}
