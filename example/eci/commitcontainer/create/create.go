package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	commitcontainer "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/commitcontainer"
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
		eci.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &commitcontainer.CreateCommitContainerTaskRequest{
		// Fill in the request parameters
		RegionId:         "b342b77ef26b11ecb0ac0242ac110002",
		TenantId:         "xxx",
		ContainerGroupId: "eci-xxx",
		ContainerName:    "container-x",
		Image: &commitcontainer.ImageInfo{
			Repository: "docker.io/library/nginx",
			Tag:        "latest",
			Message:    "xxx",
			Author:     "admin@example.org",
		},
		Registry: &commitcontainer.RegistryInfo{
			Registry: "docker.io",
			Username: "xxx",
			// base64 encoded
			Password: "cGFzc3dvcmQ=",
		},
	}
	resp, raw, err := client.CommitContainerTask().CreateCommitContainerTask(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
