package Attribute

type Attribute struct {
	content string
	size    int
	style   TextStyle
}

func (att *Attribute) GetContent() string {
	return att.content
}

func (att *Attribute) SetContent(content string) {
	att.content = content
}

func (att *Attribute) GetSize() int {
	return att.size
}

func (att *Attribute) SetSize(size int) {
	att.size = size
}

func (att *Attribute) GetStyle() TextStyle {
	return att.style
}

func (att *Attribute) SetStyle(style TextStyle) {
	att.style = style
}

func (att *Attribute) SetBold(value bool) {
	if value {
		att.style |= Bold // Set the bold bit
	} else {
		att.style &^= Bold // Clear the bold bit
	}
}

func (att *Attribute) SetItalic(value bool) {
	if value {
		att.style |= Italic // Set the italic bit
	} else {
		att.style &^= Italic // Clear the italic bit
	}
}

func (att *Attribute) SetUnderline(value bool) {
	if value {
		att.style |= Underline // Set the underline bit
	} else {
		att.style &^= Underline // Clear the underline bit
	}
}

func (att *Attribute) IsBold() bool {
	return att.style&Bold != 0
}

func (att *Attribute) IsItalic() bool {
	return att.style&Italic != 0
}

func (att *Attribute) IsUnderline() bool {
	return att.style&Underline != 0
}

func (att *Attribute) Copy() *Attribute {
	return &Attribute{
		content: att.content,
		size:    att.size,
		style:   att.style,
	}
}
