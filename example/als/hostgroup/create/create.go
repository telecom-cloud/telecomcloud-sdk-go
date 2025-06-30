package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/als"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/als/types/hostgroup"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://ctlts-global.ctapi.ctyun.cn"
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
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []als.Option{
		als.WithClientConfig(config),
		als.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}

	client, err := als.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()
	req := &hostgroup.CreateHostGroupRequest{
		RegionId: "bb9fdb42056f11eda1610242ac110002",
		// 主机组名称。最小长度：1，最大长度：64
		HostGroupName: "test-hostGroup",
		Description:   "test-hostGroup",
		// 主机标识类型，sign或uuid
		HostIdentityType: "sign",
		// 自定义标识列表。当hostIdentityType=sign时，该参数必填
		SignList: []string{"test-hostGroup"},
		// 主机uuid列表。当hostIdentityType=uuid时该参数必填，当hostIdentityType=sign时不用填
		HostList: nil,
		// 1-云容器引擎，2-云主机，不填默认为云主机
		Type: 2,
	}

	resp, raw, err := client.HostGroup().CreateHostGroup(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
