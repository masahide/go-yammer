package yammer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/masahide/go-yammer/schema"
)

var ErrCurrentFailed = errors.New("Current user did not return a successful response code")

func (c *Client) Current() (*schema.User, error) {
	var res schema.User
	url := fmt.Sprintf("%s/api/v1/users/current.json", c.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &res, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	if c.DebugMode {
		debug(httputil.DumpRequestOut(req, true))
	}

	resp, err := c.connection.Do(req)
	if err != nil {
		return &res, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &res, err
	}

	if c.DebugMode {
		debug(httputil.DumpResponse(resp, true))
	}

	if resp.StatusCode != http.StatusOK {
		return &res, ErrCurrentFailed
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
