package nag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// AllTimeLeadersGrids wraps request to and response from alltimeleadergrids endpoint.
type AllTimeLeadersGrids struct {
	*Client
	LeagueID   string
	PerMode    string
	SeasonType string
	TopX       int

	Response *Response
}

// NewAllTimeLeadersGrids creates a default AllTimeLeadersGrids instance.
func NewAllTimeLeadersGrids() *AllTimeLeadersGrids {
	return &AllTimeLeadersGrids{
		Client:     NewDefaultClient(),
		LeagueID:   "00",
		PerMode:    "Totals",
		SeasonType: "Regular Season",
		TopX:       10,
	}
}

// Get sends a GET request to alltimeleadersgrids endpoint.
func (c *AllTimeLeadersGrids) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/alltimeleadersgrids", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("LeagueID", c.LeagueID)
	q.Add("PerMode", c.PerMode)
	q.Add("SeasonType", c.SeasonType)
	q.Add("TopX", strconv.Itoa(c.TopX))
	req.URL.RawQuery = q.Encode()

	b, err := c.Do(req)
	if err != nil {
		return err
	}

	var res Response
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	c.Response = &res
	return nil
}

// JSON formats response from alltimeleadersgrids into JSON
func (c *AllTimeLeadersGrids) JSON() (json.RawMessage, error) {
	return GetJSON(*c.Response)
}
