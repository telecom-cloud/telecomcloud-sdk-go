package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/ebs"
	ebstype "github.com/telecom-cloud/telecomcloud-sdk-go/service/ebs/types/ebs"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://ebs-global.ctapi.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_ECI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func BuildGetEbsInfoRequest() *ebstype.GetEbsInfoRequest {
	return &ebstype.GetEbsInfoRequest{
		RegionId: "bb9fdb42056f11eda1610242ac110002",
		DiskId:   "feca101b-7702-493e-821e-3710d40471b2",
	}

}
func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []ebs.Option{
		ebs.WithClientConfig(config),
	}
	client, err := ebs.NewClientSet(baseDomain, options...)
	if err != nil {
		return
	}

	ctx := context.Background()
	req := BuildGetEbsInfoRequest()
	resp, raw, err := client.Ebs().GetEbsInfo(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)

}
