package cap

// alertBase contains the common fields shared between CAP 1.1 and 1.2 alerts.
type alertBase struct {
	Identifier  string   `xml:"identifier"`
	Sender      string   `xml:"sender"`
	Sent        string   `xml:"sent"`
	Status      string   `xml:"status"`
	MsgType     string   `xml:"msgType"`
	Source      string   `xml:"source,omitempty"`
	Scope       string   `xml:"scope"`
	Restriction string   `xml:"restriction,omitempty"`
	Addresses   string   `xml:"addresses,omitempty"`
	Code        []string `xml:"code,omitempty"`
	Note        string   `xml:"note,omitempty"`
	References  string   `xml:"references,omitempty"`
	Incidents   string   `xml:"incidents,omitempty"`
	Info        []*Info  `xml:"info,omitempty"`
}

// Alert represents a CAP alert message (supports both 1.1 and 1.2).
// The Version field is set by the parser to indicate which CAP version was detected.
type Alert struct {
	XMLName struct{} `xml:"urn:oasis:names:tc:emergency:cap:1.2 alert"`
	alertBase
	// Version is the detected CAP version ("1.1" or "1.2"), set by the parser.
	Version string `xml:"-"`
}

// alert11 is used internally to parse CAP 1.1 documents.
type alert11 struct {
	XMLName struct{} `xml:"urn:oasis:names:tc:emergency:cap:1.1 alert"`
	alertBase
}

// Info represents the info element in a CAP alert
type Info struct {
	Language     string       `xml:"language,omitempty"`
	Category     []string     `xml:"category"`
	Event        string       `xml:"event"`
	ResponseType []string     `xml:"responseType,omitempty"`
	Urgency      string       `xml:"urgency"`
	Severity     string       `xml:"severity"`
	Certainty    string       `xml:"certainty"`
	Audience     string       `xml:"audience,omitempty"`
	EventCode    []*ValuePair `xml:"eventCode,omitempty"`
	Effective    string       `xml:"effective,omitempty"`
	Onset        string       `xml:"onset,omitempty"`
	Expires      string       `xml:"expires,omitempty"`
	SenderName   string       `xml:"senderName,omitempty"`
	Headline     string       `xml:"headline,omitempty"`
	Description  string       `xml:"description,omitempty"`
	Instruction  string       `xml:"instruction,omitempty"`
	Web          string       `xml:"web,omitempty"`
	Contact      string       `xml:"contact,omitempty"`
	Parameter    []*ValuePair `xml:"parameter,omitempty"`
	Resource     []*Resource  `xml:"resource,omitempty"`
	Area         []*Area      `xml:"area,omitempty"`
}

// ValuePair represents a name-value pair
type ValuePair struct {
	ValueName string `xml:"valueName"`
	Value     string `xml:"value"`
}

// Resource represents a resource element
type Resource struct {
	ResourceDesc string `xml:"resourceDesc"`
	MimeType     string `xml:"mimeType,omitempty"`
	Size         int    `xml:"size,omitempty"`
	URI          string `xml:"uri,omitempty"`
	DerefURI     string `xml:"derefUri,omitempty"`
	Digest       string `xml:"digest,omitempty"`
}

// Area represents a geographic area
type Area struct {
	AreaDesc string       `xml:"areaDesc"`
	Polygon  []string     `xml:"polygon,omitempty"`
	Circle   []string     `xml:"circle,omitempty"`
	Geocode  []*ValuePair `xml:"geocode,omitempty"`
	Altitude string       `xml:"altitude,omitempty"`
	Ceiling  string       `xml:"ceiling,omitempty"`
}
