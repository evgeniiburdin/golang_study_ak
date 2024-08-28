package main

import "testing"

func TestExecBin(t *testing.T) {
	type args struct {
		binPath string
		args    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{
			"C:\\Users\\Eug\\go\\src\\student.vkusvill.ru\\evgeniiburdin\\go-course\\golang_study_ak" +
				"\\course1\\11.type_composite\\10.composite_interface\\task1.11.10.2\\main.exe",
			nil,
		}, "Hello, World!\n15\n16.5\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecBin(tt.args.binPath, tt.args.args...); got != tt.want {
				t.Errorf("ExecBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
