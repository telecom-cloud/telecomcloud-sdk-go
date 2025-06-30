package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	containergroup "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/containergroup"
	"net/http"
	"os"
	"time"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://eci-global.ctapi.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_ECI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
		eci.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &containergroup.ExecCommandRequest{
		// Fill in the request parameters
		ContainerGroupId: "eci-vtyz7dxommh6wna4",
		ContainerName:    "container-1",
		Command:          []string{"/bin/sh"},
		TTY:              true,
		Stdin:            true,
		Sync:             false,
		RegionId:         "bb9fdb42056f11eda1610242ac110002",
	}
	resp, raw, err := client.ContainerGroup().ExecCommand(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)

	if err := execContainerWebsocket(resp.RequestId, resp.WebSocketUri); err != nil {
		fmt.Println(err)
	}
}

func execContainerWebsocket(requestId string, url string) error {
	header := http.Header{}
	header.Add("X-Request-Id", requestId)
	wsconn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return err
	}
	defer wsconn.Close()

	go func() {
		for {
			_, b, err := wsconn.ReadMessage()
			if err != nil {
				fmt.Printf("read message failed:%v\n", err)
				break
			}
			os.Stdout.Write(b)
		}

	}()

	cmd := containergroup.ExecMessage{
		Command: "ls\r\n",
	}

	cmdbytes, _ := json.Marshal(cmd)
	err = wsconn.WriteMessage(websocket.TextMessage, cmdbytes)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 10)
	return nil
}
