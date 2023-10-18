package AccessLogParser

import (
	"github.com/yakawa/AccessLogParser/AccessLog"
	"github.com/yakawa/AccessLogParser/apache"
)

type Parser interface {
	Parse(string) (*AccessLog.Log, error)
}

type Parsers struct {
	Apache *apache.ApacheParser
}

func Parse(line string) {
}
