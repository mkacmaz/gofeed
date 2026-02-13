package cap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

// Parser is a CAP Parser
type Parser struct{}

// Parse parses a CAP XML feed into a cap.Alert.
// It supports both CAP 1.2 and CAP 1.1 namespaces.
func (cp *Parser) Parse(feed io.Reader) (interface{}, error) {
	data, err := io.ReadAll(feed)
	if err != nil {
		return nil, err
	}

	// Try CAP 1.2 first
	alert12 := &Alert{}
	if err := xml.Unmarshal(data, alert12); err == nil && alert12.Identifier != "" {
		alert12.Version = "1.2"
		return alert12, nil
	}

	// Try CAP 1.1
	a11 := &alert11{}
	if err := xml.Unmarshal(data, a11); err == nil && a11.Identifier != "" {
		alert := &Alert{}
		alert.alertBase = a11.alertBase
		alert.Version = "1.1"
		return alert, nil
	}

	// Neither version matched; decode again with 1.2 to return the original error
	alert := &Alert{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(alert); err != nil {
		return nil, fmt.Errorf("failed to parse CAP feed (tried 1.2 and 1.1 namespaces): %s", err.Error())
	}
	return alert, nil
}
