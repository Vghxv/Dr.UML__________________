package attribute

import (
	"Dr.uml/backend/component/drawdata"
	"Dr.uml/backend/utils/duerror"
)

// AssAttribute represents an attribute specific to associations with a ratio property
type AssAttribute struct {
	Attribute
	Ratio float64
	assDD drawdata.AssAttribute // not `drawData`
}

// NewAssAttribute creates a new AssAttribute instance with the specified ratio
// It returns an error if the ratio is not between 0 and 1
func NewAssAttribute(ratio float64) (*AssAttribute, duerror.DUError) {
	if ratio < 0 || ratio > 1 {
		return nil, duerror.NewInvalidArgumentError("ratio should be between 0 and 1")
	}
	return &AssAttribute{
		Ratio: ratio,
	}, nil
}

// GetRatio retrieves the ratio value of the AssAttribute
func (att *AssAttribute) GetRatio() (float64, duerror.DUError) {
	return att.Ratio, nil
}

func (att *AssAttribute) GetAssDD() drawdata.AssAttribute {
	return att.assDD
}

// SetRatio returns an error if the ratio is not between 0 and 1
// It returns an error if the ratio is not between 0 and 1
func (att *AssAttribute) SetRatio(ratio float64) duerror.DUError {
	if ratio < 0 || ratio > 1 {
		return duerror.NewInvalidArgumentError("ratio should be between 0 and 1")
	}
	att.Ratio = ratio
	return nil
}

func (att *AssAttribute) UpdateDrawData() {
	att.assDD.Content = att.content
	att.assDD.FontSize = att.size
	att.assDD.FontStyle = int(att.style)
	att.assDD.FontFile = att.fontFile
	att.assDD.Ratio = att.Ratio
}
