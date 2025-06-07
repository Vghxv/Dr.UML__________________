package utils

type fileType int

const (
	FiletypeDiagram    fileType = 0b0001
	FiletypeSubmodule  fileType = 0b0010
	ClassDiagram                = 0b0011
	UseCaseDiagram              = 0b101
	SequenceDiagram             = 0b1001
	SupportedFiletypes          = FiletypeDiagram | FiletypeSubmodule | ClassDiagram | UseCaseDiagram | SequenceDiagram
)

type SavedAtt struct {
	Content  string  `json:"content"`
	Size     int     `json:"size"`
	Style    int     `json:"style"`
	FontFile string  `json:"fontFile"`
	Ratio    float64 `json:"ratio,omitempty"`
}

type SavedGad struct {
	GadgetType int        `json:"GadgetType"`
	Point      string     `json:"point"`
	Layer      int        `json:"layer"`
	Color      string     `json:string`
	Attributes []SavedAtt `json:"attributes"`
}

type SavedAss struct {
	AssType    int        `json:"assType"`
	Layer      int        `json:"layer"`
	Parents    []int      `json:"parents"`
	StartPoint string     `json:"startPoint"`
	EndPoint   string     `json:"endPoint"`
	Attributes []SavedAtt `json:"attributes"`
}

type SavedFile struct {
	Filetype     fileType   `json:"filetype"`
	LastEdit     string     `json:"lastEdit"`
	Gadgets      []SavedGad `json:"Gadgets"`
	Associations []SavedAss `json:"Associations"`
}
