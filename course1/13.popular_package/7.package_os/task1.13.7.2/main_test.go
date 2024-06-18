package main

import (
	"io"

	"os"

	"strings"

	"testing"
)

func TestWriteFile(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantFd  string
		wantErr bool
	}{
		{"case1", args{strings.NewReader("hello world")}, "hello world", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := "C:\\Users\\Eug\\Desktop\\text1.txt"
			fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				t.Errorf("%#v", err)
			}
			err = WriteFile(tt.args.data, fd)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_, err = os.Open(filePath)
			if err != nil {
				t.Errorf("couldn't open created file: %#v", err)
			}
		})
	}
}
