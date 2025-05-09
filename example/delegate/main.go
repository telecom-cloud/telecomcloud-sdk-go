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
	baseDomain    = "https://crs-global.ctapi.ctyun.cn"
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
	domain := os.Getenv("CTAPI_IAM_DOMAIN")
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
	CreateDelegate()
	// SwitchDelegate()
}

func CreateDelegate() {
	req := &delegate.CreateAutomateDelegateRoleRequest{
		AccountId: "",
		Name:      "",
		PolicyIds: []string{""},
	}
	resp, raw, err := client.Delegate().CreateAutomateDelegateRole(ctx, req)
	if raw != nil {
		fmt.Printf("raw: %v\n", string(raw.Body()))
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: resp %v\n", resp)
}

func SwitchDelegate() {
	req := &delegate.SwitchDelegateRequest{
		AccountId: "",
		Name:      "",
		ValidTime: 60 * 12,
	}
	resp, raw, err := client.Delegate().SwitchDelegate(ctx, req)
	if raw != nil {
		fmt.Printf("raw: %v\n", string(raw.Body()))
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: resp %v\n", resp)
}
