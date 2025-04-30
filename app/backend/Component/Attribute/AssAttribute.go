package attribute

import (
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils/duerror"
)

// AssAttribute represents an attribute specific to associations with a ratio property
type AssAttribute struct {
	Attribute
	ratio float64
	assDD drawdata.AssAttribute // not `drawData`
}

// NewAssAttribute creates a new AssAttribute instance with the specified ratio
// It returns an error if the ratio is not between 0 and 1
func NewAssAttribute(ratio float64) (*AssAttribute, duerror.DUError) {
	if ratio < 0 || ratio > 1 {
		return nil, duerror.NewInvalidArgumentError("ratio should be between 0 and 1")
	}
	return &AssAttribute{
		ratio: ratio,
	}, nil
}

// GetRatio retrieves the ratio value of the AssAttribute
func (att *AssAttribute) GetRatio() (float64, duerror.DUError) {
	return att.ratio, nil
}

func (att *AssAttribute) GetAssDD() drawdata.AssAttribute {
	return att.assDD
}

func (att *AssAttribute) SetContent(content string) duerror.DUError {
	if err := att.Attribute.SetContent(content); err != nil {
		return err
	}
	att.assDD.Content = content
	return nil
}

func (att *AssAttribute) SetSize(size int) duerror.DUError {
	if err := att.Attribute.SetSize(size); err != nil {
		return err
	}
	att.assDD.FontSize = size
	return nil
}

func (att *AssAttribute) SetStyle(style Textstyle) duerror.DUError {
	if err := att.Attribute.SetStyle(style); err != nil {
		return err
	}
	att.assDD.FontStyle = int(style)
	return nil
}

func (att *AssAttribute) SetFontFile(fontFile string) duerror.DUError {
	if err := att.Attribute.SetFontFile(fontFile); err != nil {
		return err
	}

	att.assDD.FontFile = fontFile
	return nil
}

// SetRatio returns an error if the ratio is not between 0 and 1
// It returns an error if the ratio is not between 0 and 1
func (att *AssAttribute) SetRatio(ratio float64) duerror.DUError {
	if ratio < 0 || ratio > 1 {
		return duerror.NewInvalidArgumentError("ratio should be between 0 and 1")
	}
	att.ratio = ratio
	att.assDD.Ratio = ratio
	return nil
}

func (att *AssAttribute) UpdateDrawData() {
	att.assDD.Content = att.content
	att.assDD.FontSize = att.size
	att.assDD.FontStyle = int(att.style)
	att.assDD.FontFile = att.fontFile
	att.assDD.Ratio = att.ratio
}
