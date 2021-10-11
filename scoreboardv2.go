package nag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ronaudinho/nag/params"
)

// ScoreBoardV2 wraps request to and response from scoreboardv2 endpoint.
type ScoreBoardV2 struct {
	*Client
	DayOffset int
	GameDate  string
	LeagueID  params.LeagueID

	Response *Response
}

// NewScoreBoardV2 creates a default ScoreBoardV2 instance.
func NewScoreBoardV2() *ScoreBoardV2 {
	return &ScoreBoardV2{
		Client:    NewDefaultClient(),
		DayOffset: 0,
		GameDate:  time.Now().Format("2006-01-02"),
		LeagueID:  params.DefaultLeagueID,
	}
}

// Get sends a GET request to scoreboardv2 endpoint.
func (c *ScoreBoardV2) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/scoreboardv2", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("DayOffset", strconv.Itoa(c.DayOffset))
	q.Add("GameDate", c.GameDate)
	q.Add("LeagueID", string(c.LeagueID))
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

// JSON formats response from scoreboardv2 into JSON
func (c *ScoreBoardV2) JSON() (json.RawMessage, error) {
	return GetJSON(*c.Response)
}
