package cometd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/masahide/go-yammer/schema"
)

type ConnectionRequest struct {
	Channel        string `json:"channel"`
	ConnectionType string `json:"connectionType"`
	Id             string `json:"id"`
	ClientId       string `json:"clientId"`
}

type ConnectionResponse struct {
	Channel string         `json:"channel"`
	Data    ConnectionData `json:"data"`
}

type ConnectionData struct {
	Data interface{} `json:"data"`
	Type string      `json:"type"`
	Feed *schema.MessageFeed
}

type MessageData struct {
	Data schema.MessageFeed `json:"data"`
	Type string             `json:"type"`
}

func (a *ConnectionData) UnmarshalJSON(b []byte) error {
	var d MessageData
	err := json.Unmarshal(b, &d)
	if err == nil {
		a.Type = d.Type
		a.Feed = &d.Data
		return nil
	}
	var cd ConnectionData
	err = json.Unmarshal(b, &cd)
	if err != nil {
		return err
	}
	a.Data = cd.Data
	a.Type = cd.Type
	return nil
}

func (c *Client) connect() ([]*ConnectionResponse, error) {
	resp, err := c.do(c.connectionRequest())
	if err != nil {
		return []*ConnectionResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []*ConnectionResponse{}, err
	}

	var connResp []*ConnectionResponse
	err = json.Unmarshal(body, &connResp)
	if err != nil {
		return []*ConnectionResponse{}, fmt.Errorf("err:%s, resBody:'%s'", err, body)
	}

	return connResp, nil
}

func (c *Client) connectionRequest() []*ConnectionRequest {
	return []*ConnectionRequest{
		&ConnectionRequest{
			Id:             c.nextRequestId(),
			ClientId:       c.clientId,
			Channel:        "/meta/connect",
			ConnectionType: "long-polling",
		},
	}
}
