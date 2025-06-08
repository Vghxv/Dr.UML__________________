package attribute

import "testing"

func TestNewAssAttribute(t *testing.T) {
	t.Run("Invalid ratio", func(t *testing.T) {
		_, err := NewAssAttribute(-1, "")
		if err == nil {
			t.Errorf("expected error")
		}
		_, err = NewAssAttribute(2, "")
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("Success", func(t *testing.T) {
		ratio := 0.5
		content := "test content"
		att, err := NewAssAttribute(ratio, content)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if att.GetContent() != content {
			t.Errorf("unexpected content: %v, got: %v", content, att.GetContent())
		}
		if att.GetRatio() != ratio {
			t.Errorf("unexpected content: %v, got: %v", ratio, att.GetRatio())
		}
		dd := att.GetDrawData()
		if dd.Ratio != ratio {
			t.Errorf("incorrect drawdata")
		}
	})
}

func TestSetRatio(t *testing.T) {
	att, _ := NewAssAttribute(0.5, "")
	t.Run("Invalid ratio", func(t *testing.T) {
		err := att.SetRatio(-1)
		if err == nil {
			t.Errorf("expected error")
		}
		err = att.SetRatio(2)
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("Success", func(t *testing.T) {
		newRatio := 0.69
		err := att.SetRatio(newRatio)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if att.GetRatio() != newRatio {
			t.Errorf("expect ratio: %v, got %v", newRatio, att.GetRatio())
		}
	})
}
