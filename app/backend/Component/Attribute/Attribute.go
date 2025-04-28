package Attribute

import "Dr.uml/backend/Utils"

// Attribute represents a configurable textual element with content, size, and style properties expressed as TextStyle.
type Attribute struct {
	content string
	size    int
	style   TextStyle
}

// GetContent retrieves the content of the Attribute as a string along with an error if applicable.
func (att *Attribute) GetContent() (string, Utils.DUError) {
	return att.content, nil
}

// SetContent updates the content field of the Attribute instance if the provided content is not empty.
func (att *Attribute) SetContent(content string) Utils.DUError {
	att.content = content
	return nil
}

// GetSize returns the size of the attribute and an error if the size is negative.
func (att *Attribute) GetSize() (int, Utils.DUError) {
	if att.size < 0 {
		return 0, Utils.NewInvalidArgumentError("size was somehow set to negative. Is there a memory editor running?")
	}
	return att.size, nil
}

// SetSize sets the size of the attribute. Returns an error if the size is negative.
func (att *Attribute) SetSize(size int) Utils.DUError {
	if size < 0 {
		return Utils.NewInvalidArgumentError("size cannot be negative")
	}
	att.size = size
	return nil
}

// GetStyle returns the TextStyle of the Attribute and a possible DUError. It retrieves the current style applied.
func (att *Attribute) GetStyle() (TextStyle, Utils.DUError) {
	return att.style, nil
}

// SetStyle sets the style attribute for the text. Returns an error if the style is not between 0 and 7.
func (att *Attribute) SetStyle(style TextStyle) Utils.DUError {
	if style < 0 || style > 7 {
		return Utils.NewInvalidArgumentError("style should be between 0 and 7")
	}
	att.style = style
	return nil
}

// SetBold sets or clears the bold style for the attribute based on the provided boolean value. Returns an error if any occurs.
func (att *Attribute) SetBold(value bool) Utils.DUError {
	if value {
		att.style |= Bold // Set the bold bit
	} else {
		att.style &^= Bold // Clear the bold bit
	}
	return nil
}

// SetItalic sets or unsets the italic style for the Attribute based on the value provided. Returns Utils.DUError if any error occurs.
func (att *Attribute) SetItalic(value bool) Utils.DUError {
	if value {
		att.style |= Italic // Set the italic bit
	} else {
		att.style &^= Italic // Clear the italic bit
	}
	return nil
}

// SetUnderline modifies the underline property of the Attribute by setting or clearing the underline bit in its style field.
func (att *Attribute) SetUnderline(value bool) Utils.DUError {
	if value {
		att.style |= Underline // Set the underline bit
	} else {
		att.style &^= Underline // Clear the underline bit
	}
	return nil
}

// IsBold checks if the bold style is applied to the attribute and returns a boolean along with an error if any occurs.
func (att *Attribute) IsBold() (bool, Utils.DUError) {
	return att.style&Bold != 0, nil
}

// IsItalic checks if the Italic style flag is set in the Attribute's style and returns a boolean and an error if any.
func (att *Attribute) IsItalic() (bool, Utils.DUError) {
	return att.style&Italic != 0, nil
}

// IsUnderline determines whether the underline style is applied to the attribute and returns an error if any occurs.
func (att *Attribute) IsUnderline() (bool, Utils.DUError) {
	return att.style&Underline != 0, nil
}

// Copy creates and returns a deep copy of the Attribute with identical content, size, and style. It returns an error if any occurs.
func (att *Attribute) Copy() (*Attribute, Utils.DUError) {
	return &Attribute{
		content: att.content,
		size:    att.size,
		style:   att.style,
	}, nil
}
