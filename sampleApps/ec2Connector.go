package main

// import (
// 	"fmt"
// 	"log"
// 	"strings"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/ec2"
// )

// // This example will list instances with a filter
// //
// // Usage:
// // filter_ec2_by_tag <name_filter>
// func main() {
// 	sess, err := session.NewSession()
// 	if err != nil {
// 		log.Fatalf("failed to create session %v\n", err)
// 	}

// 	nameFilter := "Remote1"
// 	awsRegion := "us-west-2"
// 	svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
// 	fmt.Printf("listing instances with tag %v in: %v\n", nameFilter, awsRegion)
// 	params := &ec2.DescribeInstancesInput{
// 		Filters: []*ec2.Filter{
// 			{
// 				Name: aws.String("tag:Name"),
// 				Values: []*string{
// 					aws.String(strings.Join([]string{"*", nameFilter, "*"}, "")),
// 				},
// 			},
// 		},
// 	}
// 	resp, err := svc.DescribeInstances(params)
// 	if err != nil {
// 		fmt.Println("there was an error listing instances in", awsRegion, err.Error())
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Printf("%+v\n", *resp)
// }