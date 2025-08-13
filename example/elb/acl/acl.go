package main

import (
	"context"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/common/utils"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	elb "github.com/telecom-cloud/telecomcloud-sdk-go/service/elb"
	elbtypes "github.com/telecom-cloud/telecomcloud-sdk-go/service/elb/types/elb"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = "https://ctelb-global.ctapi.ctyun.cn"
	regionId      = ""
	openapiConfig *config.OpenapiConfig
	options       []elb.Option
	client        elb.ClientSet
	err           error
	ctx           context.Context
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_ELB_DOMAIN")
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

	options = []elb.Option{
		elb.WithClientConfig(openapiConfig),
		elb.WithClientMiddleware(mockMW0),
	}
	client, err = elb.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	createElbAcl()
	// getElbAcl("")
}

func createElbAcl() {
	requestId := utils.GetRandomString(32)
	req := &elbtypes.CreateElbACLRequest{
		ClientToken: requestId,
		RegionID:    regionId,
		Name:        "test-elb-acl",
		Description: "test-elb-acl",
		SourceIps:   []string{"127.0.0.1"},
	}
	resp, raw, err := client.Elb().CreateElbAcl(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)

}

func getElbAcl(id string) {
	req := &elbtypes.GetElbACLRequest{
		RegionID:        regionId,
		AccessControlID: id,
		Id:              id,
	}
	resp, raw, err := client.Elb().GetElbAcl(ctx, req)
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
