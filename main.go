package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func main() {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create Cost Explorer client
	client := costexplorer.NewFromConfig(cfg)

	// Define the time range (last 7 days)
	end := time.Now()
	start := end.AddDate(0, 0, -7)

	// Build the GetCostAndUsageInput
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(start.Format("2006-01-02")),
			End:   aws.String(end.Format("2006-01-02")),
		},
		Granularity: types.GranularityDaily,
		Metrics:     []string{"BlendedCost"}, // Can include "UsageQuantity", "UnblendedCost", etc.
	}

	// Query Cost Explorer API
	output, err := client.GetCostAndUsage(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to get cost and usage, %v", err)
	}

	// Print results
	for _, result := range output.ResultsByTime {
		fmt.Printf("Date: %s to %s\n", *result.TimePeriod.Start, *result.TimePeriod.End)
		for _, group := range result.Groups {
			fmt.Printf("  Group: %s\n", group.Keys)
			for _, metric := range group.Metrics {
				fmt.Println("poop")
				fmt.Println(&metric.Amount, &metric.Unit)
			}
		}
		fmt.Println()
	}
}
