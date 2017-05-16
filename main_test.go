package main

import (
	"testing"
)

func TestLookupAWS(t *testing.T) {

	expectedOrg := "Amazon.com"

	data := lookup("35.157.173.31")
	if data.Org != expectedOrg {
		t.Error("Expected", expectedOrg, "got", data.Org)
	}
}

func TestLookupGoogle(t *testing.T) {

	expectedOrg := "Google Cloud"

	data := lookup("104.154.227.34")
	if data.Org != expectedOrg {
		t.Error("Expected", expectedOrg, "got", data.Org)
	}
}
