package main

import (
	"context"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/region"
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
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &region.DescribeRegionRequest{
		RegionId: "b342b77ef26b11ecb0ac0242ac110002",
	}
	resp, raw, err := client.Region().DescribeRegion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

var (
	beforeMW0 = "BeforeMiddleware0"
	afterMW0  = "AfterMiddleware0"
	beforeMW1 = "BeforeMiddleware1"
	afterMW1  = "AfterMiddleware1"
)

func mockMW0(next cli.Endpoint) cli.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {
		req.BodyBuffer().WriteString(beforeMW0)
		err = next(ctx, req, resp)
		if err != nil {
			return err
		}
		req.BodyBuffer().WriteString(afterMW0)
		return nil
	}
}

func mockMW1(next cli.Endpoint) cli.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {
		req.BodyBuffer().WriteString(beforeMW1)
		err = next(ctx, req, resp)
		if err != nil {
			return err
		}
		req.BodyBuffer().WriteString(afterMW1)
		return nil
	}
}
