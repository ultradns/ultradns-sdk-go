package udnssdk

import (
	"fmt"
	"net/http"
	"time"

)

// ZoneService provides access to Zone resources
type ZoneService struct {
	client *Client
}

// Zone wraps an Zone resource
type Zone struct {
	Properties struct {
		Name                 string    `json:"name"`
		AccountName          string    `json:"accountName`
		Type                 string    `json:"type"`
		DnssecStatus         string    `json:"dnssecStatus"`
		Status               string    `json:"status"`
		Owner                string    `json:"owner"`
		ResourceRecordCount  int       `json:"resourceRecordCount"`
		LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	} `json:"properties"`
}

// ZoneListDTO wraps a list of Zone resources
type ZoneListDTO struct {
	Zones      []Zone     `json:"zones"`
	Queryinfo  QueryInfo  `json:"queryInfo"`
	Resultinfo ResultInfo `json:"resultInfo"`
}

// ZoneKey collects the identifiers of a Zone
type ZoneKey struct {
	Zone        string
	AccountName string
}

// URI generates the URI for an Zone
func (k ZoneKey) URI() string {
	uri := fmt.Sprintf("zones/?&q=name:%s", k.Zone)
	if k.AccountName != "" {
		uri += fmt.Sprintf("+account_name:%s", k.AccountName)
	}
	return uri
}

// QueryURI generates the query URI for an Zone and offset
func (k ZoneKey) QueryURI(offset int, limit int) string {

	return fmt.Sprintf("%s&offset=%d&limit=%d", k.URI(), offset, limit)
}

// SelectWithOffset requests zone rrsets by ZoneKey & optional offset
func (s *ZoneService) SelectWithOffset(k ZoneKey, offset int, limit int) ([]Zone, ResultInfo, *http.Response, error) {
	var zoneld ZoneListDTO

	uri := k.QueryURI(offset,limit)
	res, err := s.client.get(uri, &zoneld)

	zones := []Zone{}
	for _, zone := range zoneld.Zones {
		zones = append(zones, zone)
	}
	return zones, zoneld.Resultinfo, res, err
}