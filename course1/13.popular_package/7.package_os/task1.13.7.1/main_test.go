package main

import (
	"os"

	"testing"
)

func TestWriteFile(t *testing.T) {
	type args struct {
		filePath string
		data     []byte
		perm     os.FileMode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{"C:\\Users\\Eug\\Desktop\\file.txt", []byte("hello world"), os.FileMode(0644)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.filePath, tt.args.data, tt.args.perm); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
