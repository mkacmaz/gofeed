package cap

import (
	"strings"
	"testing"
)

const cap12XML = `<?xml version="1.0" encoding="UTF-8"?>
<cap:alert xmlns:cap="urn:oasis:names:tc:emergency:cap:1.2">
  <cap:identifier>test-cap-12</cap:identifier>
  <cap:sender>sender@example.com</cap:sender>
  <cap:sent>2024-01-01T00:00:00-00:00</cap:sent>
  <cap:status>Actual</cap:status>
  <cap:msgType>Alert</cap:msgType>
  <cap:scope>Public</cap:scope>
  <cap:info>
    <cap:category>Met</cap:category>
    <cap:event>Storm Warning</cap:event>
    <cap:urgency>Immediate</cap:urgency>
    <cap:severity>Severe</cap:severity>
    <cap:certainty>Observed</cap:certainty>
    <cap:headline>Severe Storm Warning</cap:headline>
    <cap:description>A severe storm is approaching.</cap:description>
  </cap:info>
</cap:alert>`

const cap11XML = `<?xml version="1.0" encoding="UTF-8"?>
<cap:alert xmlns:cap="urn:oasis:names:tc:emergency:cap:1.1">
  <cap:identifier>test-cap-11</cap:identifier>
  <cap:sender>eliot.christian@meteoswiss.ch</cap:sender>
  <cap:sent>2012-10-20T08:30:00-00:00</cap:sent>
  <cap:status>Actual</cap:status>
  <cap:msgType>Alert</cap:msgType>
  <cap:scope>Public</cap:scope>
  <cap:info>
    <cap:category>Infra</cap:category>
    <cap:event>power failure</cap:event>
    <cap:urgency>Immediate</cap:urgency>
    <cap:severity>Minor</cap:severity>
    <cap:certainty>Observed</cap:certainty>
    <cap:headline>Electrical power failure at Geneva</cap:headline>
    <cap:description>Geneva is experiencing power failure.</cap:description>
    <cap:instruction>Remain calm.</cap:instruction>
    <cap:area>
      <cap:areaDesc>Geneva, airport to lake and river</cap:areaDesc>
    </cap:area>
  </cap:info>
</cap:alert>`

const invalidXML = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0"><channel><title>Not a CAP feed</title></channel></rss>`

func TestParseCAP12(t *testing.T) {
	parser := &Parser{}
	result, err := parser.Parse(strings.NewReader(cap12XML))
	if err != nil {
		t.Fatalf("unexpected error parsing CAP 1.2: %v", err)
	}

	alert, ok := result.(*Alert)
	if !ok {
		t.Fatal("result is not *Alert")
	}

	if alert.Version != "1.2" {
		t.Errorf("expected version 1.2, got %q", alert.Version)
	}
	if alert.Identifier != "test-cap-12" {
		t.Errorf("expected identifier test-cap-12, got %q", alert.Identifier)
	}
	if alert.Status != "Actual" {
		t.Errorf("expected status Actual, got %q", alert.Status)
	}
	if len(alert.Info) != 1 {
		t.Fatalf("expected 1 info block, got %d", len(alert.Info))
	}
	if alert.Info[0].Headline != "Severe Storm Warning" {
		t.Errorf("expected headline 'Severe Storm Warning', got %q", alert.Info[0].Headline)
	}
}

func TestParseCAP11(t *testing.T) {
	parser := &Parser{}
	result, err := parser.Parse(strings.NewReader(cap11XML))
	if err != nil {
		t.Fatalf("unexpected error parsing CAP 1.1: %v", err)
	}

	alert, ok := result.(*Alert)
	if !ok {
		t.Fatal("result is not *Alert")
	}

	if alert.Version != "1.1" {
		t.Errorf("expected version 1.1, got %q", alert.Version)
	}
	if alert.Identifier != "test-cap-11" {
		t.Errorf("expected identifier test-cap-11, got %q", alert.Identifier)
	}
	if alert.Sender != "eliot.christian@meteoswiss.ch" {
		t.Errorf("expected sender eliot.christian@meteoswiss.ch, got %q", alert.Sender)
	}
	if alert.Status != "Actual" {
		t.Errorf("expected status Actual, got %q", alert.Status)
	}
	if alert.Scope != "Public" {
		t.Errorf("expected scope Public, got %q", alert.Scope)
	}
	if len(alert.Info) != 1 {
		t.Fatalf("expected 1 info block, got %d", len(alert.Info))
	}
	if alert.Info[0].Headline != "Electrical power failure at Geneva" {
		t.Errorf("expected headline 'Electrical power failure at Geneva', got %q", alert.Info[0].Headline)
	}
	if alert.Info[0].Description != "Geneva is experiencing power failure." {
		t.Errorf("expected description 'Geneva is experiencing power failure.', got %q", alert.Info[0].Description)
	}
	if alert.Info[0].Instruction != "Remain calm." {
		t.Errorf("expected instruction 'Remain calm.', got %q", alert.Info[0].Instruction)
	}
	if len(alert.Info[0].Area) != 1 {
		t.Fatalf("expected 1 area, got %d", len(alert.Info[0].Area))
	}
	if alert.Info[0].Area[0].AreaDesc != "Geneva, airport to lake and river" {
		t.Errorf("expected areaDesc 'Geneva, airport to lake and river', got %q", alert.Info[0].Area[0].AreaDesc)
	}
}

func TestParseInvalidReturnsError(t *testing.T) {
	parser := &Parser{}
	_, err := parser.Parse(strings.NewReader(invalidXML))
	if err == nil {
		t.Fatal("expected error parsing non-CAP XML, got nil")
	}
}
