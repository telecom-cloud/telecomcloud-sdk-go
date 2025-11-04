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
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/isuite/training"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://ccse-global.ctapi.ctyun.cn/ccse"
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
		// isuite.WithClientMiddleware(dumpHttpMiddleware),
	}

	client, err := isuite.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()

	req := &training.ListTrainingRequest{
		RegionId:  "bb9fdb42056f11eda1610242ac110002",
		ClusterId: "",
		Namespace: "default",
	}

	resp, raw, err := client.Training().ListTraining(ctx, req)
	if err != nil {
		fmt.Printf("list training err: %v\n", err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
