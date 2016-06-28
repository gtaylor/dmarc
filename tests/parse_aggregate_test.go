package tests

import (
	"testing"
	"github.com/gtaylor/dmarc/parse/aggregate"
	"os"
)

func TestParseAggregate(t *testing.T) {
	xmlFile, err := os.Open("samples/dmarc.org.sample.xml")
	if err != nil {
		t.Errorf("os error: %v\n", err)
	}
	defer xmlFile.Close()
	aggregate.ParseAggregateReportXML(xmlFile)
}
