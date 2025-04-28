package Attribute

import "Dr.uml/backend/Utils"

type Attribute struct {
	content string
	size    int
	style   TextStyle
}

func (att *Attribute) GetContent() (string, Utils.DUError) {
	return att.content, nil
}

func (att *Attribute) SetContent(content string) Utils.DUError {
	if content == "" {
		return Utils.NewInvalidArgumentError("content cannot be empty")
	}
	att.content = content
	return nil
}

func (att *Attribute) GetSize() (int, Utils.DUError) {
	if att.size < 0 {
		return 0, Utils.NewInvalidArgumentError("size cannot be negative")
	}
	return att.size, nil
}

func (att *Attribute) SetSize(size int) Utils.DUError {
	if size < 0 {
		return Utils.NewInvalidArgumentError("size cannot be negative")
	}
	att.size = size
	return nil
}

func (att *Attribute) GetStyle() (TextStyle, Utils.DUError) {
	return att.style, nil
}

func (att *Attribute) SetStyle(style TextStyle) Utils.DUError {
	if style < 0 || style > 7 {
		return Utils.NewInvalidArgumentError("style should be between 0 and 7")
	}
	att.style = style
	return nil
}

func (att *Attribute) SetBold(value bool) Utils.DUError {
	if value {
		att.style |= Bold // Set the bold bit
	} else {
		att.style &^= Bold // Clear the bold bit
	}
	return nil
}

func (att *Attribute) SetItalic(value bool) Utils.DUError {
	if value {
		att.style |= Italic // Set the italic bit
	} else {
		att.style &^= Italic // Clear the italic bit
	}
	return nil
}

func (att *Attribute) SetUnderline(value bool) Utils.DUError {
	if value {
		att.style |= Underline // Set the underline bit
	} else {
		att.style &^= Underline // Clear the underline bit
	}
	return nil
}

func (att *Attribute) IsBold() (bool, Utils.DUError) {
	return att.style&Bold != 0, nil
}

func (att *Attribute) IsItalic() (bool, Utils.DUError) {
	return att.style&Italic != 0, nil
}

func (att *Attribute) IsUnderline() (bool, Utils.DUError) {
	return att.style&Underline != 0, nil
}

func (att *Attribute) Copy() (*Attribute, Utils.DUError) {
	return &Attribute{
		content: att.content,
		size:    att.size,
		style:   att.style,
	}, nil
}
