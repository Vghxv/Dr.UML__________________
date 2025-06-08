package attribute

import (
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
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
		att.UpdateDrawData()
		return nil
	})
	att.UpdateDrawData()
	return att, nil
}

func FromSavedAssAttribute(savedAssAtt utils.SavedAtt) (*AssAttribute, duerror.DUError) {
	ass := &AssAttribute{
		Attribute: Attribute{
			content:  savedAssAtt.Content,
			size:     savedAssAtt.Size,
			style:    Textstyle(savedAssAtt.Style),
			fontFile: savedAssAtt.FontFile,
		},
		ratio: savedAssAtt.Ratio,
	}
	ass.UpdateDrawData()
	return ass, nil
}

func (att *AssAttribute) ToSavedAssAttribute() utils.SavedAtt {
	return utils.SavedAtt{
		Content:  att.content,
		Size:     att.size,
		Style:    int(att.style),
		FontFile: att.fontFile,
		Ratio:    att.ratio,
	}
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
	att.UpdateDrawData()
	return nil
}

func (att *AssAttribute) UpdateDrawData() {
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
