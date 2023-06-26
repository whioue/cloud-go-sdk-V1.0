package main

import (
	"context"
	"fmt"
	"github.com/whioue/cloud-go-sdk-V1.0/cloudgo/service"
	"github.com/whioue/cloud-go-sdk-V1.0/domain/v1/train/dto"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
	"time"
)

var client *service.CloudClient

func main() {
	//cloudConfig := "/Users/haminyg/Desktop/cloud/cloud-go-sdk/cloud.yaml"
	var err error
	config := &rest.Config{
		Host:    "http://127.0.0.1:8100",
		APIPath: "",
		ContentConfig: rest.ContentConfig{
			GroupVersion: "v1/cloud",
		},
		Timeout:       4 * time.Second,
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
	}

	client, err = service.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	Start()
}

func Start() {
	c := context.Background()

	err := client.ApiV1().Train().Start(c, dto.TrainStartArgs{TaskID: 1})
	if err != nil {
		fmt.Println(err)
	}
}
