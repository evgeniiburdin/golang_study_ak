package main

import "testing"

func TestLgTV_GetModel(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"case1", fields{false, "LG XL-100500"}, "LG XL-100500"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LgTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			if got := l.GetModel(); got != tt.want {
				t.Errorf("GetModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLgTV_GetStatus(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"case1", fields{false, "LG XL-100500"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LgTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			if got := l.GetStatus(); got != tt.want {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLgTV_LGHub(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"case1", fields{true, "LG XL-100500"}, "LGHub"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LgTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			l.LGHub()
			result := l.LGHub()
			if result != tt.want {
				t.Errorf("LGHub() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestLgTV_switchOFF(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"case1", fields{true, "LG XL-100500"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LgTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			l.switchOFF()
			if l.status != tt.want {
				t.Errorf("switchOFF() = %v, want %v", l.status, tt.want)
			}
		})
	}
}

func TestLgTV_switchON(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"case1", fields{false, "LG XL-100500"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LgTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			l.switchON()
			if l.status != tt.want {
				t.Errorf("switchON() = %v, want %v", l.status, tt.want)
			}
		})
	}
}

func TestSamsungTV_GetModel(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"case1", fields{false, "Samsung XL-100500"}, "Samsung XL-100500"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SamsungTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			if got := s.GetModel(); got != tt.want {
				t.Errorf("GetModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSamsungTV_GetStatus(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"case1", fields{false, "Samsung XL-100500"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SamsungTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			if got := s.GetStatus(); got != tt.want {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSamsungTV_SamsungHub(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"case1", fields{true, "Samsung XL-100500"}, "SamsungHub"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SamsungTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			if got := s.SamsungHub(); got != tt.want {
				t.Errorf("SamsungHub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSamsungTV_switchOFF(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"case1", fields{true, "Samsung XL-100500"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SamsungTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			s.switchOFF()
			if s.status != tt.want {
				t.Errorf("switchOFF() = %v, want %v", s.status, tt.want)
			}
		})
	}
}

func TestSamsungTV_switchON(t *testing.T) {
	type fields struct {
		status bool
		model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"case1", fields{false, "Samsung XL-100500"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SamsungTV{
				status: tt.fields.status,
				model:  tt.fields.model,
			}
			s.switchON()
			if s.status != tt.want {
				t.Errorf("switchON() = %v, want %v", s.status, tt.want)
			}
		})
	}
}
