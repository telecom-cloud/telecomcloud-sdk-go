package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/isuite"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/isuite/framework"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://ccse-global.ctapi.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func dumpHttpMiddleware(next cli.Endpoint) cli.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {
		log.Printf("proxy-request  method %s path %s header %s body %s", req.Method(), string(req.Host())+string(req.Path()), req.Header.String(), string(req.Body()))
		err = next(ctx, req, resp)
		if err != nil {
			return err
		}
		log.Printf("proxy-response status %d  header %s body %s", resp.StatusCode(), string(resp.Header.Header()), string(resp.Body()))
		return nil
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []isuite.Option{
		isuite.WithClientConfig(config),
		isuite.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
		isuite.WithClientMiddleware(dumpHttpMiddleware),
	}

	client, err := isuite.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()

	req := &framework.ListFrameworkTagsRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		ClusterId:   "",
		FrameworkId: "",
		PageNum:     1,
		PageSize:    1000,
	}

	resp, raw, err := client.Framework().ListFrameworkTags(ctx, req)
	if err != nil {
		fmt.Printf("list tags err: %v\n", err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
