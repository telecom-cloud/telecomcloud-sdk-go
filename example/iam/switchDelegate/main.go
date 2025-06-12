package main

import (
	"context"
	"fmt"
	"log"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/protocol"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	iam "github.com/telecom-cloud/telecomcloud-sdk-go/service/iam"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/iam/types/delegate"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = "https://ctiam-global.ctapi-internal-test.ctyun.cn:21443"
	openapiConfig *config.OpenapiConfig
	options       []iam.Option
	client        iam.ClientSet
	err           error
	ctx           context.Context
	accountId     = ""
	policyId      = ""
	delegateName  = ""
)

func init() {
	// accessKey = os.Getenv("CTAPI_AK")
	// secretKey = os.Getenv("CTAPI_AK")
	// baseDomain = os.Getenv("CTAPI_DOMAIN")
	// accountId = os.Getenv("ACCOUNT_ID")
	// policyId = os.Getenv("POLICY_ID")
}

func main() {
	ctx = context.Background()
	openapiConfig = &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options = []iam.Option{
		iam.WithClientConfig(openapiConfig),
		iam.WithClientMiddleware(dumpHttpMiddleware),
	}
	client, err = iam.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	switchDelegateRequest()
}

func switchDelegateRequest() {
	req := &delegate.SwitchDelegateRequest{
		Name:      delegateName,
		AccountId: accountId,
		ValidTime: 60 * 12,
		SupportS3: true,
	}
	resp, raw, err := client.Delegate().SwitchDelegate(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("resp.AppKey: %s ak: %s \n", resp.AppKey, accessKey)
	sk, err := Decode(resp.AppKey, accessKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ak: %s sk: %s \n", resp.AppId, sk)

	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}

func dumpHttpMiddleware(next cli.Endpoint) cli.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {
		log.Println("proxy-request", "method", req.Method(), "path", req.Path(), "header", string(req.Header.Header()), "body", string(req.Body()))
		err = next(ctx, req, resp)
		if err != nil {
			return err
		}
		log.Println(ctx, "proxy-response", "status", resp.StatusCode(), "header", string(resp.Header.Header()), "body", string(resp.Body()))
		return nil
	}
}
