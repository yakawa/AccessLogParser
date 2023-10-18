package common

import (
	"errors"
	"strings"
)

//go:generate stringer -type=HttpMethodType
type HttpMethodType int

const (
	HTTP_UNKNOWN HttpMethodType = iota
	HTTP_GET
	HTTP_POST
	HTTP_PUT
	HTTP_HEAD
	HTTP_DELETE
	HTTP_OPTIONS
	HTTP_TRACE
	HTTP_CONNECT
)

//go:generate stringer -type=HttpProtocolType
type HttpProtocolType int

const (
	HTTP_PROTOCOL_UNKNOWN HttpProtocolType = iota
	HTTP_0_9
	HTTP_1_0
	HTTP_1_1
	HTTP_2
	HTTP_3
)

type HttpRequest struct {
	Protocol       HttpProtocolType
	ProtocolString string
	Method         HttpMethodType
	MethodString   string
	URI            string
}

func ParseRequestLine(line string) (*HttpRequest, error) {
	request := strings.Split(line, " ")
	req := &HttpRequest{}
	if len(request) != 3 {
		return nil, errors.New("Unknown Request Line")
	}
	switch strings.ToUpper(request[0]) {
	case "GET":
		req.Method = HTTP_GET
	case "POST":
		req.Method = HTTP_POST
	case "PUT":
		req.Method = HTTP_PUT
	case "HEAD":
		req.Method = HTTP_HEAD
	case "DELETE":
		req.Method = HTTP_DELETE
	case "OPTIONS":
		req.Method = HTTP_OPTIONS
	case "TRACE":
		req.Method = HTTP_TRACE
	case "CONNECT":
		req.Method = HTTP_CONNECT
	default:
		req.Method = HTTP_UNKNOWN
	}
	req.MethodString = request[0]

	switch strings.ToUpper(request[2]) {
	case "HTTP/0.9":
		req.Protocol = HTTP_0_9
	case "HTTP/1.0":
		req.Protocol = HTTP_1_0
	case "HTTP/1.1":
		req.Protocol = HTTP_1_1
	case "HTTP/2":
		req.Protocol = HTTP_2
	case "HTTP/3":
		req.Protocol = HTTP_3
	default:
		req.Protocol = HTTP_PROTOCOL_UNKNOWN
	}
	req.ProtocolString = request[2]

	req.URI = request[1]
	return req, nil
}
