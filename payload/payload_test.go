package payload

import "testing"

func Test_X(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{{"simple", false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := X(); (&result == nil) != tt.wantErr {
				t.Error("result is null")
			}
		})
	}
}
