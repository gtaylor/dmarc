package tests

import (
	"github.com/gtaylor/dmarc/parse/aggregate"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

// Tests the dmarc.org.sample.xml
func TestDmarcOrgSample(t *testing.T) {
	xmlFile, err := os.Open("samples/dmarc.org.sample.xml")
	if err != nil {
		t.Errorf("Failure opening dmarc.org.sample.xml: %v", err)
	}
	defer xmlFile.Close()
	r, err := aggregate.ParseAggregateReportXML(xmlFile)
	if err != nil {
		t.Errorf("Error while parsing dmarc.org.sample.xml: %v", err)
	}

	rMd := &r.ReportMetadata
	assert.Equal(t, rMd.OrgName, "acme.com")
	assert.Equal(t, rMd.Email, "noreply-dmarc-support@acme.com")
	assert.Equal(t, rMd.ExtraContactInfo, "http://acme.com/dmarc/support")
	assert.Equal(t, rMd.ReportID, "9391651994964116463")

	expectedBegin := aggregate.UTCTime{time.Date(2012, 4, 28, 0, 0, 0, 0, time.UTC)}
	assert.Equal(t, rMd.DateRange.Begin, expectedBegin)
	expectedEnd := aggregate.UTCTime{time.Date(2012, 4, 28, 23, 59, 59, 0, time.UTC)}
	assert.Equal(t, rMd.DateRange.End, expectedEnd)

	polPub := &r.PolicyPublished
	assert.Equal(t, polPub.Domain, "example.com")
	assert.Equal(t, polPub.AlignDKIM, "r")
	assert.Equal(t, polPub.AlignSPF, "r")
	assert.Equal(t, polPub.Policy, "none")
	assert.Equal(t, polPub.SubdomainPolicy, "none")
	assert.Equal(t, polPub.Percentage, 100)

	assert.Equal(t, len(r.Records), 1)
	rec := &r.Records[0]
	assert.Equal(t, rec.Row.SourceIP, "72.150.241.94")
	assert.Equal(t, rec.Row.Count, 2)
	assert.Equal(t, rec.Row.PolicyEvaluated.Disposition, "none")
	assert.Equal(t, rec.Row.PolicyEvaluated.DKIM, "fail")
	assert.Equal(t, rec.Row.PolicyEvaluated.SPF, "pass")

	assert.Equal(t, rec.Identifiers.HeaderFrom, "example.com")
	assert.Equal(t, rec.AuthResults.DKIM.Domain, "example.com")
	assert.Equal(t, rec.AuthResults.DKIM.Result, "fail")
	assert.Equal(t, rec.AuthResults.DKIM.HumanResult, "")

	assert.Equal(t, rec.AuthResults.SPF.Domain, "example.com")
	assert.Equal(t, rec.AuthResults.SPF.Result, "pass")
}
