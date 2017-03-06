package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// )

// func exitWithError(err error) {
// 	fmt.Fprintln(os.Stderr, err)
// 	os.Exit(1)
// }

// func main() {
// 	cfg := Config{}
// 	if err := cfg.Load(); err != nil {
// 		exitWithError(fmt.Errorf("failed to load config, %v", err))
// 	}

// 	// Create the config specifying the Region for the DynamoDB table.
// 	// If Config.Region is not set the region must come from the shared
// 	// config or AWS_REGION environment variable.
// 	awsRegion := "us-west-2"
// 	awscfg := &aws.Config{Region: aws.String(awsRegion)}
// 	if len(cfg.Region) > 0 {
// 		awscfg.WithRegion(cfg.Region)
// 	}

// 	// Create the session that the DynamoDB service will use.
// 	sess, err := session.NewSession(awscfg)
// 	if err != nil {
// 		exitWithError(fmt.Errorf("failed to create session, %v", err))
// 	}

// 	// Create the DynamoDB service client to make the query request with.
// 	svc := dynamodb.New(sess)

// 	// Build the query input parameters
// 	// params := &dynamodb.ScanInput{
// 	// 	TableName: aws.String(cfg.Table),
// 	// }
// 	// if cfg.Limit > 0 {
// 	// 	params.Limit = aws.Int64(cfg.Limit)
// 	// }

// 	// Make the DynamoDB Query API call
// 	// 	result, err := svc.Scan(params)
// 	// 	if err != nil {
// 	// 		exitWithError(fmt.Errorf("failed to make Query API call, %v", err))
// 	// 	}

// 	// 	items := []Item{}

// 	// 	// Unmarshal the Items field in the result value to the Item Go type.
// 	// 	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
// 	// 	if err != nil {
// 	// 		exitWithError(fmt.Errorf("failed to unmarshal Query result items, %v", err))
// 	// 	}

// 	// 	// Print out the items returned
// 	// 	for i, item := range items {
// 	// 		fmt.Printf("%d: Key: %d, Desc: %s\n", i, item.Key, item.Desc)
// 	// 		fmt.Printf("\tNum Data Values: %d\n", len(item.Data))
// 	// 		for k, v := range item.Data {
// 	// 			fmt.Printf("\t- %q: %v\n", k, v)
// 	// 		}
// 	// 	}
// 	// }

// 	//Start:List Tables
// 	// params := &dynamodb.ListTablesInput{
// 	// 	Limit: aws.Int64(2),
// 	// }
// 	// resp, err := svc.ListTables(params)

// 	// if err != nil {
// 	// 	// Print the error, cast err to awserr.Error to get the Code and
// 	// 	// Message from an error.
// 	// 	fmt.Println(err.Error())
// 	// 	return
// 	// }

// 	// // Pretty-print the response data.
// 	// fmt.Println(resp)
// 	//End List tables

// 	// Start create table
// 	params := &dynamodb.CreateTableInput{
// 		TableName: aws.String("Thread"),
// 		AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 			{
// 				AttributeName: aws.String("ForumName"),
// 				AttributeType: aws.String("S"),
// 			},
// 			{
// 				AttributeName: aws.String("Subject"),
// 				AttributeType: aws.String("S"),
// 			},
// 		},
// 		KeySchema: []*dynamodb.KeySchemaElement{
// 			{
// 				AttributeName: aws.String("ForumName"),
// 				KeyType:       aws.String("HASH"),
// 			},
// 			{
// 				AttributeName: aws.String("Subject"),
// 				KeyType:       aws.String("RANGE"),
// 			},
// 		},
// 		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 			ReadCapacityUnits:  aws.Int64(10),
// 			WriteCapacityUnits: aws.Int64(5),
// 		},
// 	}
// 	resp, err := svc.CreateTable(params)

// 	if err != nil {
// 		// Print the error, cast err to awserr.Error to get the Code and
// 		// Message from an error.
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// Pretty-print the response data.
// 	fmt.Println(resp)
// 	// Endcreate table
// }

// type Item struct {
// 	Key  int
// 	Desc string
// 	Data map[string]interface{}
// }

// type Config struct {
// 	Table  string // required
// 	Region string // optional
// 	Limit  int64  // optional

// }

// func (c *Config) Load() error {
// 	flag.Int64Var(&c.Limit, "limit", 0, "Limit is the max items to be returned, 0 is no limit")
// 	flag.StringVar(&c.Table, "table", "Music12", "Table to Query on")
// 	flag.StringVar(&c.Region, "region", "us-west-2", "AWS Region the table is in")
// 	flag.Parse()

// 	if len(c.Table) == 0 {
// 		flag.PrintDefaults()
// 		return fmt.Errorf("table name is required.")
// 	}

// 	return nil
// }
