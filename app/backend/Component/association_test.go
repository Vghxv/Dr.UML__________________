package component

import (
	"Dr.uml/backend/utils"
	"testing"

	"Dr.uml/backend/component/attribute"
	"Dr.uml/backend/utils/duerror"
)

func Test_Association_GetAssType(t *testing.T) {
	tests := []struct {
		name   string
		ass    *Association
		expect AssociationType
	}{
		{"valid type", &Association{assType: 2}, 2},
		{"default type", &Association{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ass.GetAssType(); got != tt.expect {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
		})
	}
}

func Test_Association_GetAttributes(t *testing.T) {
	att := &attribute.AssAttribute{}
	tests := []struct {
		name    string
		ass     *Association
		expect  []*attribute.AssAttribute
		wantErr bool
	}{
		{"with attributes", &Association{attributes: []*attribute.AssAttribute{att}}, []*attribute.AssAttribute{att}, false},
		{"no attributes", &Association{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ass.GetAttributes()
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch: got %v, want error %v", err, tt.wantErr)
			}
			if len(got) != len(tt.expect) {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
		})
	}
}

func Test_Association_GetLayer(t *testing.T) {
	tests := []struct {
		name    string
		ass     *Association
		want    int
		wantErr bool
	}{
		{
			name:    "valid layer 5",
			ass:     &Association{layer: 5},
			want:    5,
			wantErr: false,
		},
		{
			name:    "default layer 0",
			ass:     &Association{},
			want:    0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ass.GetLayer()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Association_SetParentStartEnd(t *testing.T) {
	gadget := &Gadget{}
	tests := []struct {
		name      string
		ass       *Association
		getMethod func(*Association) (*Gadget, duerror.DUError)
		setMethod func(*Association, *Gadget) duerror.DUError
		gadget    *Gadget
		wantErr   bool
	}{
		{"get valid start", &Association{parents: [2]*Gadget{gadget, gadget}}, func(a *Association) (*Gadget, duerror.DUError) { return a.GetParentStart(), nil }, nil, gadget, false},
		{"get valid end", &Association{parents: [2]*Gadget{gadget, gadget}}, func(a *Association) (*Gadget, duerror.DUError) { return a.GetParentEnd(), nil }, nil, gadget, false},
		{"set valid start", &Association{}, nil, func(a *Association, g *Gadget) duerror.DUError { return a.SetParentStart(g) }, gadget, false},
		{"set valid end", &Association{}, nil, func(a *Association, g *Gadget) duerror.DUError { return a.SetParentEnd(g) }, gadget, false},
		{"set nil start", &Association{}, nil, func(a *Association, g *Gadget) duerror.DUError { return a.SetParentStart(nil) }, nil, true},
		{"set nil end", &Association{}, nil, func(a *Association, g *Gadget) duerror.DUError { return a.SetParentEnd(nil) }, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.getMethod != nil {
				got, err := tt.getMethod(tt.ass)
				if (err != nil) != tt.wantErr {
					t.Errorf("error mismatch: got %v, want error %v", err, tt.wantErr)
				}
				if got != tt.gadget {
					t.Errorf("expected %v, got %v", tt.gadget, got)
				}
			}
			if tt.setMethod != nil {
				err := tt.setMethod(tt.ass, tt.gadget)
				if (err != nil) != tt.wantErr {
					t.Errorf("error mismatch: got %v, want error %v", err, tt.wantErr)
				}
			}
		})
	}
}

func Test_Association_SetAssType(t *testing.T) {
	tests := []struct {
		name    string
		ass     *Association
		assType AssociationType
		wantErr bool
	}{
		{"set valid type", &Association{}, 1, false},
		{"set invalid type", &Association{}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ass.SetAssType(tt.assType)
			if tt.ass.GetAssType() != tt.assType {
				t.Errorf("expected association type %v, got %v", tt.assType, tt.ass.GetAssType())
			}
		})
	}
}

func Test_Association_Cover(t *testing.T) {
	tests := []struct {
		name    string
		ass     *Association
		point   utils.Point
		want    bool
		wantErr bool
	}{
		{
			name: "point inside",
			ass: &Association{
				parents: [2]*Gadget{
					{point: utils.Point{X: 0, Y: 0}},
					{point: utils.Point{X: 10, Y: 10}},
				},
			},
			point:   utils.Point{X: 5, Y: 5},
			want:    false, /*TODO: Change after the func is implemented*/
			wantErr: false,
		},
		{
			name: "point outside",
			ass: &Association{
				parents: [2]*Gadget{
					{point: utils.Point{X: 0, Y: 0}},
					{point: utils.Point{X: 10, Y: 10}},
				},
			},
			point:   utils.Point{X: 20, Y: 20},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ass.Cover(tt.point)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cover() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Association_SetLayer(t *testing.T) {
	tests := []struct {
		name    string
		ass     *Association
		layer   int
		wantErr bool
	}{
		{
			name:    "set valid layer",
			ass:     &Association{},
			layer:   5,
			wantErr: false,
		},
		{
			name:    "set negative layer",
			ass:     &Association{},
			layer:   -1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ass.SetLayer(tt.layer)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetLayer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				got, _ := tt.ass.GetLayer()
				if got != tt.layer {
					t.Errorf("Layer not set correctly, got = %v, want %v", got, tt.layer)
				}
			}
		})
	}
}

func Test_Association_UpdateDrawData(t *testing.T) {
	tests := []struct {
		name    string
		ass     *Association
		wantErr bool
	}{
		{
			name:    "update with nil association",
			ass:     nil,
			wantErr: true,
		},
		{
			name:    "orphan ass",
			ass:     &Association{},
			wantErr: true,
		},
		{
			name: "valid update",
			ass: &Association{
				parents: [2]*Gadget{
					{point: utils.Point{X: 0, Y: 0}},
					{point: utils.Point{X: 10, Y: 10}},
				},
			},
			wantErr: false,
		},
		{
			name: "nil att update",
			ass: &Association{
				parents: [2]*Gadget{
					{point: utils.Point{X: 0, Y: 0}},
					{point: utils.Point{X: 10, Y: 10}},
				},
				attributes: []*attribute.AssAttribute{nil},
			},
			wantErr: true,
		},
		{
			name: "valid with attributes",
			ass: &Association{
				parents: [2]*Gadget{
					{point: utils.Point{X: 0, Y: 0}},
					{point: utils.Point{X: 10, Y: 10}},
				},
				attributes: []*attribute.AssAttribute{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "valid update" {
				attr, _ := attribute.NewAssAttribute(0.2)
				tt.ass.attributes = append(tt.ass.attributes, attr)
			}
			err := tt.ass.UpdateDrawData()
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateDrawData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Association_GetDrawData(t *testing.T) {
	tests := []struct {
		name    string
		ass     *Association
		wantErr bool
	}{
		{
			name:    "get from valid association",
			ass:     &Association{},
			wantErr: false,
		},
		{
			name:    "get from nil association",
			ass:     nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ass != nil {
				_, err := tt.ass.GetDrawData()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetDrawData() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func Test_Association_AddAttribute(t *testing.T) {
	att := &attribute.AssAttribute{}
	tests := []struct {
		name    string
		ass     *Association
		att     *attribute.AssAttribute
		wantErr bool
	}{
		{"add valid attribute", &Association{}, att, false},
		{"add nil attribute", &Association{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ass.AddAttribute(tt.att)
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch: got %v, want error %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Association_RemoveAttribute(t *testing.T) {
	att := &attribute.AssAttribute{}
	tests := []struct {
		name    string
		ass     *Association
		index   int
		wantErr bool
	}{
		{"valid remove", &Association{attributes: []*attribute.AssAttribute{att}}, 0, false},
		{"invalid index", &Association{attributes: []*attribute.AssAttribute{att}}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ass.RemoveAttribute(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch: got %v, want error %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Association_MoveAttribute(t *testing.T) {
	att := &attribute.AssAttribute{}
	tests := []struct {
		name    string
		ass     *Association
		index   int
		ratio   float64
		wantErr bool
	}{
		{"valid move", &Association{attributes: []*attribute.AssAttribute{att}}, 0, 0.5, false},
		{"invalid index", &Association{attributes: []*attribute.AssAttribute{att}}, 1, 0.5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ass.MoveAttribute(tt.index, tt.ratio)
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch: got %v, want error %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewAssociation(t *testing.T) {
	gadget := &Gadget{}
	tests := []struct {
		name    string
		parents [2]*Gadget
		assType AssociationType
		wantErr bool
	}{
		{"valid association", [2]*Gadget{gadget, gadget}, 1, false},
		{"nil parent", [2]*Gadget{nil, gadget}, 1, true},
		{"invalid assType", [2]*Gadget{gadget, gadget}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAssociation(tt.parents, tt.assType)
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch: got %v, want error %v", err, tt.wantErr)
			}
		})
	}
}
