package attribute

import (
	"os"

	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

// Attribute represents a configurable textual element with content, size, and style properties expressed as Textstyle.
type Attribute struct {
	content          string
	size             int
	style            Textstyle
	fontFile         string
	drawData         drawdata.Attribute
	updateParentDraw func() duerror.DUError
}

func NewAttribute(content string) (*Attribute, duerror.DUError) {
	// TODO
	att := &Attribute{
		content:  content,
		size:     drawdata.DefaultAttributeFontSize,
		style:    drawdata.DefaultAttributeFontStyle,
		fontFile: os.Getenv("APP_ROOT") + drawdata.DefaultAttributeFontFile,
	}
	if err := att.updateDrawData(); err != nil {
		return nil, err
	}
	return att, nil
}

// GetContent retrieves the content of the Attribute as a string along with an error if applicable.
func (att *Attribute) GetContent() string {
	return att.content
}

// SetContent updates the content field of the Attribute instance if the provided content is not empty.
func (att *Attribute) SetContent(content string) duerror.DUError {
	att.content = content
	return att.updateDrawData()
}

// GetSize returns the size of the attribute and an error if the size is negative.
func (att *Attribute) GetSize() int {
	if att.size < 0 {
		return 0
	}
	return att.size
}

// SetSize sets the size of the attribute. Returns an error if the size is negative.
func (att *Attribute) SetSize(size int) duerror.DUError {
	if size < 0 {
		return duerror.NewInvalidArgumentError("size cannot be negative")
	}
	att.size = size
	return att.updateDrawData()
}

// GetStyle returns the Textstyle of the Attribute and a possible DUError. It retrieves the current style applied.
func (att *Attribute) GetStyle() Textstyle {
	return att.style
}

// SetStyle sets the style attribute for the text. Returns an error if the style contains unsupported flags.
func (att *Attribute) SetStyle(style Textstyle) duerror.DUError {
	if style & ^supportedTextStyleFlags != 0 {
		return duerror.NewInvalidArgumentError("style contains unsupported flags")
	}
	att.style = style
	return att.updateDrawData()
}

// SetBold sets or clears the bold style for the attribute based on the provided boolean value. Returns an error if any occurs.
func (att *Attribute) SetBold(value bool) duerror.DUError {
	if value {
		att.style |= Bold // Set the bold bit
	} else {
		att.style &^= Bold // Clear the bold bit
	}
	return att.updateDrawData()
}

// SetItalic sets or unsets the italic style for the Attribute based on the value provided. Returns duerror.DUError if any error occurs.
func (att *Attribute) SetItalic(value bool) duerror.DUError {
	if value {
		att.style |= Italic // Set the italic bit
	} else {
		att.style &^= Italic // Clear the italic bit
	}
	return att.updateDrawData()
}

// SetUnderline modifies the underline property of the Attribute by setting or clearing the underline bit in its style field.
func (att *Attribute) SetUnderline(value bool) duerror.DUError {
	if value {
		att.style |= Underline // Set the underline bit
	} else {
		att.style &^= Underline // Clear the underline bit
	}
	return att.updateDrawData()
}

// SetFontFile sets the font file path for the Attribute and updates the drawData accordingly.
// Returns an error if the file path is invalid.
func (att *Attribute) SetFontFile(fontFile string) duerror.DUError {
	if err := utils.ValidateFilePath(fontFile); err != nil {
		return duerror.NewInvalidArgumentError("invalid font file path")
	}
	att.fontFile = fontFile
	return att.updateDrawData()
}

// IsBold checks if the bold style is applied to the attribute and returns a boolean along with an error if any occurs.
func (att *Attribute) IsBold() (bool, duerror.DUError) {
	return att.style&Bold != 0, nil
}

// IsItalic checks if the Italic style flag is set in the Attribute's style and returns a boolean and an error if any.
func (att *Attribute) IsItalic() (bool, duerror.DUError) {
	return att.style&Italic != 0, nil
}

// IsUnderline determines whether the underline style is applied to the attribute and returns an error if any occurs.
func (att *Attribute) IsUnderline() (bool, duerror.DUError) {
	return att.style&Underline != 0, nil
}

// Copy creates and returns a deep copy of the Attribute with identical content, size, and style. It returns an error if any occurs.
func (att *Attribute) Copy() (*Attribute, duerror.DUError) {
	return &Attribute{
		content: att.content,
		size:    att.size,
		style:   att.style,
	}, nil
}

func (att *Attribute) GetDrawData() drawdata.Attribute {
	return att.drawData
}

func (att *Attribute) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	if update == nil {
		return duerror.NewInvalidArgumentError("update function is nil")
	}
	att.updateParentDraw = update
	return nil
}

func (att *Attribute) updateDrawData() duerror.DUError {
	if att == nil {
		return duerror.NewInvalidArgumentError("attribute is nil")
	}

	height, width, err := utils.GetTextSize(att.content, att.size, att.fontFile)
	if err != nil {
		return err
	}

	// Validate inputs
	if height < 0 || width < 0 {
		return duerror.NewInvalidArgumentError("height and width must be non-negative")
	}

	att.drawData.Content = att.content
	att.drawData.Height = height
	att.drawData.Width = width
	att.drawData.FontSize = att.size
	att.drawData.FontStyle = int(att.style)
	att.drawData.FontFile = att.fontFile

	if att.updateParentDraw == nil {
		return nil
	}
	return att.updateParentDraw()
}
