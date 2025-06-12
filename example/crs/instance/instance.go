package main

import (
	"context"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	crs "github.com/telecom-cloud/telecomcloud-sdk-go/service/crs"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/crs/types/instance"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = ""
	regionId      = ""
	openapiConfig *config.OpenapiConfig
	options       []crs.Option
	client        crs.ClientSet
	err           error
	ctx           context.Context
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_CRS_DOMAIN")
	//regionId = os.Getenv("CRS_REGIONID")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	ctx = context.Background()
	openapiConfig = &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options = []crs.Option{
		crs.WithClientConfig(openapiConfig),
		crs.WithClientMiddleware(mockMW0),
	}
	client, err = crs.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	// getAuthorizationToken()
	getDelegateuser()
}

func getAuthorizationToken() {
	req := &instance.GetAuthorizationTokenRequest{
		Auth:       "all",
		RegionId:   regionId,
		InstanceId: 0,
	}
	resp, raw, err := client.Instance().GetAuthorizationToken(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

func getDelegateuser() {
	ctx := context.Background()
	req := &instance.GetDelegateUsernameRequest{
		RegionId: regionId,
	}
	resp, raw, err := client.Instance().GetDelegateUsername(ctx, req)
	if err != nil {
		fmt.Printf("failed %s", err.Error())
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
		// req.BodyBuffer().WriteString(beforeMW0)
		fmt.Printf("request query: %s\n", string(req.QueryString()))
		err = next(ctx, req, resp)
		if err != nil {
			return err
		}
		req.BodyBuffer().WriteString(afterMW0)
		return nil
	}
}
