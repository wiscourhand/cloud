package handler

import (
	"github.com/cloudwego/kitex/client/genericclient"
)

var clientMap = new(map[string]genericclient.Client)

func addClient(name string, content string, includes map[string]string) {

}
