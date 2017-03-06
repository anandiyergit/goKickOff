package main

import (
	"fmt"
	"os"
)

func EnvProviders() {
	port := os.Getenv("p")
	awsAccessKeyID := os.Getenv("aws_access_key_id")
	awsSecretAccessKey := os.Getenv("aws_secret_access_key")
	endpoint := os.Getenv("endpoint")
	fmt.Println("Port:" + port + "Key id:" + awsAccessKeyID + "access key:" + awsSecretAccessKey + "endpoint:" + endpoint)

}
