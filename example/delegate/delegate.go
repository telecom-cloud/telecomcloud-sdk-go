package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	iam "github.com/telecom-cloud/telecomcloud-sdk-go/service/iam"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/iam/types/delegate"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = ""
	regionId      = ""
	openapiConfig *config.OpenapiConfig
	options       []iam.Option
	client        iam.ClientSet
	err           error
	ctx           context.Context
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_CRS_DOMAIN")
	regionId = os.Getenv("CRS_REGIONID")
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

	options = []iam.Option{
		iam.WithClientConfig(openapiConfig),
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
		Name:      "",
		AccountId: "0",
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
