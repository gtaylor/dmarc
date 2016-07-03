package aggregate

import (
	"encoding/xml"
	log "github.com/Sirupsen/logrus"
	"io"
	"strconv"
	"time"
)

// Top-level feedback tag.
type Feedback struct {
	ReportMetadata  ReportMetadata  `xml:"report_metadata"`
	PolicyPublished PolicyPublished `xml:"policy_published"`
	Records         []Record        `xml:"record"`
}

// feedback/report_metadata
type ReportMetadata struct {
	OrgName          string    `xml:"org_name"`
	Email            string    `xml:"email"`
	ExtraContactInfo string    `xml:"extra_contact_info"`
	ReportID         string    `xml:"report_id"`
	DateRange        DateRange `xml:"date_range"`
}

// feedback/report_metadata/date_range
type DateRange struct {
	Begin UTCTime `xml:"begin"`
	End   UTCTime `xml:"end"`
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
	Row         Row         `xml:"row"`
	Identifiers Identifiers `xml:"identifiers"`
	AuthResults AuthResults `xml:"auth_results"`
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
	DKIM        string `xml:"dkim"`
	SPF         string `xml:"spf"`
}

// feedback/record/row/identifiers
type Identifiers struct {
	HeaderFrom string `xml:"header_from"`
}

// feedback/record/row/auth_results
type AuthResults struct {
	DKIM DKIM `xml:"dkim"`
	SPF  SPF  `xml:"spf"`
}

// feedback/record/row/auth_results/dkim
type DKIM struct {
	Domain      string `xml:"domain"`
	Result      string `xml:"result"`
	HumanResult string `xml:"human_result"`
	Selector	string `xml:"selector"`
}

// feedback/record/row/auth_results/dkim
type SPF struct {
	Domain string `xml:"domain"`
	Result string `xml:"result"`
	Scope  string `xml:"scope"`
}

// Custom type for parsing the timestamped date ranges.
type UTCTime struct {
	time.Time
}

func (c *UTCTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var val string
	d.DecodeElement(&val, &start)
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Errorf("Unable to parse timestamp (%s): %s", val, err)
		return err
	}
	tm := time.Unix(i, 0)
	*c = UTCTime{tm.UTC()}
	return nil
}

func ParseAggregateReportXML(xmlFileReader io.Reader) (*Feedback, error) {
	fb := &Feedback{}
	err := xml.NewDecoder(xmlFileReader).Decode(fb)
	if err != nil {
		return nil, err
	}
	return fb, nil
}
