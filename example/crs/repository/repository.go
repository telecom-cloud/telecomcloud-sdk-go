package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	crs "github.com/telecom-cloud/telecomcloud-sdk-go/service/crs"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/crs/types/repository"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = "https://crs-global.ctapi.ctyun.cn"
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

	options = []crs.Option{
		crs.WithClientConfig(openapiConfig),
	}
	client, err = crs.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	list()
	get()
	create()
}

func list() {
	req := &repository.ListRepositoryRequest{
		RegionId:      regionId,
		InstanceId:    0,
		NamespaceName: "",
		OrderBy:       "",
		Order:         "",
		PageNum:       1,
		PageSize:      10,
	}
	resp, raw, err := client.Repository().ListRepository(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

func get() {
	req := &repository.GetRepositoryRequest{
		RegionId:     regionId,
		InstanceId:   0,
		RepositoryId: 0,
	}
	resp, raw, err := client.Repository().GetRepository(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

func create() {
	req := &repository.CreateRepositoryRequest{
		RegionId:       regionId,
		InstanceId:     0,
		NamespaceName:  "",
		RepositoryName: "",
		IsPublic:       true,
	}
	resp, raw, err := client.Repository().CreateRepository(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
