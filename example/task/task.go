package main

import (
	"context"
	"fmt"
	"github.com/whioue/cloud-go-sdk-V1.0/cloudgo/service"
	"github.com/whioue/cloud-go-sdk-V1.0/domain/v1/task/dto"
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

	Detail()
}

func Create() {
	c := context.Background()

	args := dto.TaskCreateArgs{
		Name: "Hym3",
	}

	reply, err := client.ApiV1().Task().Create(c, args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}

func List() {
	c := context.Background()

	args := dto.TaskListArgs{
		Page:     1,
		PageSize: 10,
	}

	reply, err := client.ApiV1().Task().List(c, args)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range reply.List {
		fmt.Println(k, v.Name)
	}
}

func Delete() {
	c := context.Background()

	args := dto.TaskDeleteArgs{
		ID: 1,
	}

	err := client.ApiV1().Task().Delete(c, args)
	if err != nil {
		fmt.Println(err)
	}
}

func Detail() {
	c := context.Background()

	args := dto.TaskDetailArgs{
		ID: 10,
	}

	reply, err := client.ApiV1().Task().Detail(c, args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}
