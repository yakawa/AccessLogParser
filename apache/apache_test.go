package apache

import (
	"reflect"
	"testing"
)

func TestSplitAccessLog(t *testing.T) {
	tests := []struct {
		Line    string
		Wants   []string
		IsError bool
	}{
		{Line: `192.168.225.100 - - [24/Sep/2023:03:12:29 +0000] "GET /index.html HTTP/1.1" 404 233 "-" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7)" 80`,
			Wants:   []string{"192.168.225.100", "-", "-", "[24/Sep/2023:03:12:29 +0000]", "GET /index.html HTTP/1.1", "404", "233", "-", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7)", "80"},
			IsError: false,
		},
	}

	p := NewApacheParser()

	for _, tt := range tests {
		t.Run("SplitLog", func(t *testing.T) {
			got, err := p.splitField(tt.Line)
			if err != nil && tt.IsError == false {
				t.Errorf("Error is not nil")
				t.Log(err)
			}
			if len(got) != len(tt.Wants) {
				t.Errorf("Failed: Return length is mismatch Wants (%d), but Got: (%d)\n", len(tt.Wants), len(got))
			}
			if !reflect.DeepEqual(got, tt.Wants) {
				for i := 0; i < len(got); i++ {
					t.Logf("%d: \"%s\", \"%s\"\n", i, got[i], tt.Wants[i])
				}
				t.Errorf("Failed: Wants (%v), but Got: (%v)\n", tt.Wants, got)
			}
		})
	}
}
