package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/tag"
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

func BuildUnbindTagRequest() *tag.UnbindTagRequest {
	return &tag.UnbindTagRequest{
		RegionId:     "b342b77ef26b11ecb0ac0242ac110002",
		ResourceId:   "eci-xvcv485qgvxnlhzz",
		ResourceType: "ContainerGroup",
		TagKeys:      []string{"key"},
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
		return
	}

	ctx := context.Background()
	req := BuildUnbindTagRequest()
	resp, raw, err := client.Tag().UnbindTag(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)

}
