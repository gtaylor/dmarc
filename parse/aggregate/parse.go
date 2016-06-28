package aggregate

import (
	"encoding/xml"
	log "github.com/Sirupsen/logrus"
	"io"
)

// Top-level feedback tag.
type Feedback struct {
	ReportMetadata  ReportMetadata  `xml:"report_metadata"`
	PolicyPublished PolicyPublished `xml:"policy_published"`
	Records         []Record        `xml:"record"`
}

// feedback/report_metadata
type ReportMetadata struct {
	Organization   string `xml:"org_name"`
	Email          string `xml:"email"`
	ReportID       string `xml:"report_id"`
	DateRangeBegin string `xml:"date_range>begin"`
	DateRangeEnd   string `xml:"date_range>end"`
}

// feedback/policy_published
type PolicyPublished struct {
	Domain          string `xml:"domain"`
	AlignDKIM       string `xml:"adkim"`
	AlignSPF        string `xml:"aspf"`
	Policy          string `xml:"p"`
	SubdomainPolicy string `xml:"sp"`
	Percentage      int    `xml:"pct"`
}

// feedback/record
type Record struct {
	Row        Row    `xml:"row"`
	Identifiers		Identifiers		`xml:"identifiers"`
	AuthResults	AuthResults	`xml:"auth_results"`
}

// feedback/record->row
type Row struct {
	Count           int             `xml:"count"`
	SourceIP        string          `xml:"source_ip"`
	PolicyEvaluated PolicyEvaluated `xml:"policy_evaluated"`
}

// feedback/record/row/policy_evaluated
type PolicyEvaluated struct {
	Disposition string `xml:"disposition"`
	EvalDKIM    string `xml:"dkim"`
	EvalSPF     string `xml:"spf"`
}

// feedback/record/row/identifiers
type Identifiers struct {
	HeaderFrom string `xml:"header_from"`
}

// feedback/record/row/auth_results
type AuthResults struct {
	DKIM	DKIM	`xml:"dkim"`
	SPF		SPF		`xml:"spf"`
}

// feedback/record/row/auth_results/dkim
type DKIM struct {
	Domain 		string	`xml:"domain"`
	Result 		string	`xml:"result"`
	HumanResult 		string	`xml:"human_result"`
}

// feedback/record/row/auth_results/dkim
type SPF struct {
	Domain 		string	`xml:"domain"`
	Result 		string	`xml:"result"`
}

func ParseAggregateReportXML(xmlFileReader io.Reader) {
	fb := &Feedback{}
	err := xml.NewDecoder(xmlFileReader).Decode(fb)
	if err != nil {
		log.Fatal(err)
	}
}
