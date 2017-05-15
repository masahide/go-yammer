package cometd

import (
	"fmt"
	"log"
)

func (c *Client) Poll(messageChan chan *ConnectionResponse, stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Aborting.")
			return
		default:
			resp, err := c.connect()
			if err != nil {
				log.Println(err)
				close(messageChan)
				return
			}
			for _, r := range resp {
				if r != nil {
					messageChan <- r
				}
			}
		}
	}
}
