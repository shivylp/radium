package sources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spy16/radium"
)

// NewRadium initializes a radium.Source implementation
// using the given radium server url
func NewRadium(serverURL string) *Radium {
	rad := &Radium{}
	rad.server = serverURL
	return rad
}

// Radium implements radium.Source interface using
// radium server as the source of reference
type Radium struct {
	server string
}

// Search makes a GET /search to the radium server and formats the
// response
func (rad Radium) Search(ctx context.Context, query radium.Query) ([]radium.Article, error) {
	urlObj, err := url.Parse(rad.server + "/search")
	if err != nil {
		return nil, err
	}

	urlQuery := urlObj.Query()
	urlQuery.Set("q", query.Text)
	for name, val := range query.Attribs {
		urlQuery.Set(name, val)
	}
	urlObj.RawQuery = urlQuery.Encode()

	req, _ := http.NewRequest(http.MethodGet, urlObj.String(), nil)
	req.Header.Set("User-Agent", "curl/7.54.0")
	req.WithContext(ctx)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	var rs []radium.Article
	err = json.NewDecoder(resp.Body).Decode(&rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
