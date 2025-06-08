package attribute

import (
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils/duerror"
)

// AssAttribute represents an attribute specific to associations with a ratio property
type AssAttribute struct {
	Attribute
	ratio                 float64
	assDD                 drawdata.AssAttribute
	updateParentDrawOuter func() duerror.DUError
}

// NewAssAttribute creates a new AssAttribute instance with the specified ratio
// It returns an error if the ratio is not between 0 and 1
func NewAssAttribute(ratio float64, content string) (*AssAttribute, duerror.DUError) {
	if ratio < 0 || ratio > 1 {
		return nil, duerror.NewInvalidArgumentError("ratio should be between 0 and 1")
	}
	tmp, err := NewAttribute(content)
	if err != nil {
		return nil, err
	}
	att := &AssAttribute{
		Attribute: *tmp,
		ratio:     ratio,
	}
	att.Attribute.RegisterUpdateParentDraw(func() duerror.DUError {
		att.updateDrawData()
		return nil
	})
	att.updateDrawData()
	return att, nil
}

// GetRatio retrieves the ratio value of the AssAttribute
func (att *AssAttribute) GetRatio() float64 {
	return att.ratio
}

func (att *AssAttribute) GetDrawData() drawdata.AssAttribute {
	return att.assDD
}

// SetRatio returns an error if the ratio is not between 0 and 1
// It returns an error if the ratio is not between 0 and 1
func (att *AssAttribute) SetRatio(ratio float64) duerror.DUError {
	if ratio < 0 || ratio > 1 {
		return duerror.NewInvalidArgumentError("ratio should be between 0 and 1")
	}
	att.ratio = ratio
	att.updateDrawData()
	return nil
}

func (att *AssAttribute) updateDrawData() {
	att.assDD.Content = att.content
	att.assDD.FontSize = att.size
	att.assDD.FontStyle = int(att.style)
	att.assDD.FontFile = att.getFontFileBase()
	att.assDD.Ratio = att.ratio
	if att.updateParentDrawOuter != nil {
		att.updateParentDrawOuter()
	}
}

func (att *AssAttribute) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	if update == nil {
		return duerror.NewInvalidArgumentError("update function is nil")
	}
	att.updateParentDrawOuter = update
	return nil
}
