package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	cg "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/containergroup"
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

func BuildCreateContainerGroupRequest() *cg.CreateContainerGroupRequest {
	return &cg.CreateContainerGroupRequest{
		RegionId:           "b342b77ef26b11ecb0ac0242ac110002",
		ContainerGroupName: "eci-test",
		Cpu:                1,
		Memory:             2,
		AzInfo: []*cg.AzInfo{
			{
				AzName: "cn-xinan1-1A",
			},
		},
		RestartPolicy:   "Never",
		VpcId:           "vpc-yehu9qjjol",
		VSwitchId:       "subnet-z73ymwzk87",
		SecurityGroupId: "sg-0mpilsidy4",
		//HostAliases: []*cg.HostAlias{
		//	{},
		//},
		Tags: []*cg.Tag{
			{
				Key:   "key",
				Value: "val",
			},
		},
		Containers: []*cg.Container{
			{
				Name:            "nginx",
				Cpu:             1,
				Memory:          2,
				Image:           "nginx:latest",
				ImagePullPolicy: "IfNotPresent",
				EnvironmentVar: []*cg.EnvironmentVar{
					{
						Key:   "Mode",
						Value: "Proxy",
					},
				},
				ReadinessProbe: &cg.Probe{
					Exec: &cg.Exec{
						Command: []string{
							"curl",
							"localhost:80",
						},
					},
					InitialDelaySeconds: 10,
					TimeoutSeconds:      10,
					PeriodSeconds:       10,
					SuccessThreshold:    1,
					FailureThreshold:    1,
				},
				VolumeMount: []*cg.VolumeMount{
					{
						Name:      "test-volume",
						MountPath: "/test-volume",
					},
				},
			},
		},
		Volumes: []*cg.Volume{
			{
				Name: "test-volume",
				/*EmptyDirVolume*/
				Type: "EmptyDirVolume",
				EmptyDirVolume: &cg.EmptyDirVolume{
					Medium:    "Memory",
					SizeLimit: 100,
				},
				/*ConfigFileVolume*/
				//Type: "ConfigFileVolume",
				//ConfigFileVolume: &cg.ConfigFileVolume{
				//	DefaultMode: 755,
				//	ConfigFileToPaths: []*cg.ConfigFileToPath{
				//		{
				//			Mode:    755,
				//			Path:    "test/nginx.conf",
				//			Content: "d29vZHM=",
				//		},
				//	},
				//},
				/*NasVolume*/
				//Type: "NasVolume",
				//NasVolume: &cg.NasVolume{
				//	SourcePath: "127.0.0.1:/mnt/sfs_cap",
				//	ReadOnly:   true,
				//},
				/*DiskVolume*/
				//Type: "DiskVolume",
				//DiskVolume: &cg.DiskVolume{
				//	VolumeHandle: "3b534328-d121-11ee-b8a2-76f00f27cfe3",
				//	FsType:       "xfs",
				//},
				/*ZosVolume*/
				//Type: "ZosVolume",
				//ZosVolume: &cg.ZosVolume{
				//	Bucket:     "bucket-9999",
				//	Url:        "127.0.0.1:80",
				//	Credential: "YWs6c2s=",
				//	ReadOnly:   true,
				//},
			},
		},
		HostName: "eci-test",
	}
}

func CreateContainerGroup(ctx context.Context, cli eci.ContainerGroupClient) (*cg.CreateContainerGroupResponse, *protocol.Response, error) {
	req := BuildCreateContainerGroupRequest()
	return cli.CreateContainerGroup(ctx, req)
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
	resp, raw, err := CreateContainerGroup(ctx, client.ContainerGroup())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
