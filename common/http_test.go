package common

import (
	"reflect"
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	tests := []struct {
		Line    string
		Wants   *HttpRequest
		IsError bool
	}{
		{
			Line: "GET /index.html HTTP/1.1",
			Wants: &HttpRequest{
				Protocol:       HTTP_1_1,
				ProtocolString: "HTTP/1.1",
				Method:         HTTP_GET,
				MethodString:   "GET",
				URI:            "/index.html"},
			IsError: false,
		},
	}

	for _, tt := range tests {
		t.Run("ParseRequest", func(t *testing.T) {
			got, err := ParseRequestLine(tt.Line)
			if err != nil && tt.IsError == false {
				t.Errorf("Error is not nil")
				t.Log(err)
			}
			if !reflect.DeepEqual(got, tt.Wants) {
				t.Errorf("Failed: Wants (%v), but Got: (%v)\n", tt.Wants, got)
			}
		})
	}
}
