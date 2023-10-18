package apache

import (
	"errors"
	"strings"

	"github.com/yakawa/AccessLogParser/AccessLog"
)

type ApacheParser struct {
	QuoteString string
	Format      []AccessLog.LogField
}

func NewApacheParser() *ApacheParser {
	commonFmt := []AccessLog.LogField{AccessLog.RemoteHost, AccessLog.Ident, AccessLog.RemoteUser, AccessLog.AcceptTime, AccessLog.RequestLine, AccessLog.StatusCode, AccessLog.ResponseBodyByte}
	return &ApacheParser{
		QuoteString: "\"",
		Format:      commonFmt,
	}
}

func (ap *ApacheParser) Parse(line string) (*AccessLog.Log, error) {
	l := &AccessLog.Log{}

	return l, nil
}

func (ap *ApacheParser) splitField(line string) ([]string, error) {
	fields := []string{}
	if line == "" {
		return fields, nil
	}

	arr := strings.Split(line, " ")
	field := ""
	inQuoted := false
	inDatetime := false
	for _, w := range arr {
		if inQuoted == false {
			// [24/Sep/2023:03:12:29 +0000]
			if inDatetime == false && strings.HasPrefix(w, "[") {
				inDatetime = true
				field = w
				continue
			}
			if inDatetime == true && strings.HasSuffix(w, "]") {
				field += " "
				field += w
				fields = append(fields, field)
				field = ""
				inDatetime = false
				continue
			} else if inDatetime == true {
				return nil, errors.New("Unknown Time Format")
			}
			if strings.HasPrefix(w, ap.QuoteString) && strings.HasSuffix(w, ap.QuoteString) {
				field = strings.TrimPrefix(w, ap.QuoteString)
				field = strings.TrimSuffix(field, ap.QuoteString)
			} else if strings.HasPrefix(w, ap.QuoteString) {
				inQuoted = true
				field = strings.TrimPrefix(w, ap.QuoteString)
				continue
			} else if strings.HasSuffix(w, ap.QuoteString) {
				return nil, errors.New("Broken Quoted String Found")
			} else {
				field = w
			}
			fields = append(fields, field)
			field = ""
		} else {
			if strings.HasPrefix(w, ap.QuoteString) {
				return nil, errors.New("Broken Quoted String Found")
			} else if strings.HasSuffix(w, ap.QuoteString) {
				field += " "
				field += strings.TrimSuffix(w, ap.QuoteString)
				inQuoted = false
				fields = append(fields, field)
				field = ""
			} else {
				field += " "
				field += w
			}
		}
	}
	if field != "" {
		return nil, errors.New("Unterminated Quote String")
	}
	return fields, nil
}
