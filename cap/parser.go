package cap

import (
	"encoding/xml"
	"io"
)

// Parser is a CAP Parser
type Parser struct{}

// Parse parses a CAP XML feed into a cap.Alert
func (cp *Parser) Parse(feed io.Reader) (interface{}, error) {
	alert := &Alert{}
	decoder := xml.NewDecoder(feed)
	err := decoder.Decode(alert)
	if err != nil {
		return nil, err
	}
	return alert, nil
}
