package main

// import "github.com/aws/aws-sdk-go/service/s3"
// import "github.com/aws/aws-sdk-go/aws"
// import "strings"
// import "fmt"

// func main() {
// 	client := s3.New(nil)
// 	result, err := client.PutObject(&s3.PutObjectInput{
// 		Bucket: aws.String("my-bucket"),
// 		Key:    aws.String("my-key"),
// 		Body:   strings.NewReader("Hello from AWS!!"),
// 	})
// 	if err != nil {
// 		fmt.Println("")
// 	} else {
// 		fmt.Println("result is:", *result)
// 	}
// }
