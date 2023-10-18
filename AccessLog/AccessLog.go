package AccessLog

import "time"

type LogField int

const (
	UnknownField LogField = iota
	RemoteAddr
	LocalAddr
	ResponseBodyByte
	ProcessTime
	FileName
	RemoteHost
	Protocol
	Ident
	Method
	ServerPort
	ClientPort
	ProcessID
	QueryString
	RequestLine
	StatusCode
	AcceptTime
	TotalTime
	RemoteUser
	ServerName
	CanocialServerName
	ConnectionStatus
	RecvByte
	SentByte
)

type Log struct {
	VirtualHost   string
	Host          string
	RemoteLogname string
	User          string
	Time          time.Time
	Request       string
	Status        int
	Size          uint64
	Referer       string
	UserAgent     string
	Forwarder     string
	RequestURI    string
	Protocol      string
	Method        string
	Original      string
}

func (l *Log) parseRequest() error {
	if l.RequestURI != "" || l.Protocol != "" || l.Method != "" {
		// Already parsed
		return nil
	}

	return nil
}
