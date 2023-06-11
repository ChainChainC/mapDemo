package common

import "net/http"

type ClientApp struct {
	Client *http.Client
}

var LocalClient *ClientApp

func NewClientApp() {
	LocalClient = &ClientApp{
		Client: &http.Client{},
	}
}

func (c *ClientApp) PostReq() {

}
