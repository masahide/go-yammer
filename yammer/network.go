package yammer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/go-querystring/query"
	"github.com/masahide/go-yammer/schema"
)

// GetNetworksOptions is options
type GetNetworksOptions struct {
	IncludeSuspended             bool `url:"include_suspended"`
	ExcludeOwnMessagesFromUnseen bool `url:"exclude_own_messages_from_unseen"`
}

// GetNetworks is https://www.yammer.com/api/v1/networks/current.json
// https://developer.yammer.com/docs/networkscurrentjson
func (c *Client) GetNetworks(options GetNetworksOptions) ([]schema.Network, error) {
	querystring, _ := query.Values(options)

	url := fmt.Sprintf("/api/v1/networks/current.json?%s", querystring.Encode())
	resp, err := c.sendRequest(nil, "GET", url)
	if err != nil {
		return []schema.Network{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []schema.Network{}, err
	}

	var network []schema.Network
	err = json.Unmarshal(body, &network)
	if err != nil {
		return []schema.Network{}, err
	}

	return network, nil
}
