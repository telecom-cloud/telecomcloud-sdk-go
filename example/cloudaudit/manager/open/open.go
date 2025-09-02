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
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/cloudaudit"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/cloudaudit/types/manager"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://cloudaudit-global.ctapi.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: "",
		SecretKey: "",
	}
	options := []cloudaudit.Option{
		cloudaudit.WithClientConfig(config),
		cloudaudit.WithClientMiddleware(dumpHttpMiddleware),
		cloudaudit.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}

	client, err := cloudaudit.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()
	req := &manager.ManagerServiceRequest{
		RegionId:  "",
		AccountId: "",
		UserId:    "",
	}

	resp, raw, err := client.Manager().OpenService(ctx, req)
	if err != nil {
		fmt.Printf("OpenService err: %v\n", err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
	fmt.Println(resp.Data)
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
