/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/iboware/bot-detect/config"
	"github.com/iboware/bot-detect/pkg/consumer"
	"github.com/iboware/bot-detect/pkg/detector"
	"github.com/iboware/bot-detect/pkg/ioutil"
	"github.com/iboware/bot-detect/pkg/publisher"
	"github.com/spf13/cobra"
)

func newServiceCmd() *cobra.Command {
	cfg := config.Config{}
	// serviceCmd represents the service command
	var serviceCmd = &cobra.Command{
		Use:   "service",
		Short: "Starts the service",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Initialize Pubsub Client
			client, err := pubsub.NewClient(ctx, cfg.ProjectID)
			if err != nil {
				log.Fatalf("pubsub.NewClient: %v", err)
			}
			defer client.Close()

			// Initialize and configure subscription
			sub := client.Subscription(cfg.SubscriptionID)
			sub.ReceiveSettings.Synchronous = false
			sub.ReceiveSettings.NumGoroutines = 16
			sub.ReceiveSettings.MaxOutstandingMessages = 8

			// Initialize and configure publisher
			pub := publisher.NewGCPPublisher(client, cfg.OutboundTopicID)

			// download and parse ip block list file
			tempPath := fmt.Sprintf("%s/%s", os.TempDir(), "blocklist.csv")
			if err := ioutil.DownloadFile(tempPath, cfg.IPBlockListURL); err != nil {
				log.Fatal("could not download ip block list file")
			}
			file, err := os.Open(tempPath)
			if err != nil {
				log.Fatal("could not open ip block list file")
			}
			blockListMap, err := ioutil.ParseIPBlockList(file)
			if err != nil {
				log.Fatal("could not parse ip block list file")
			}
			if err := file.Close(); err != nil {
				log.Fatal("could not close ip block list file")

			}

			// Initialize detective
			detective := detector.NewDetector(pub, blockListMap, cfg.OutboundTopicID)

			// Initialize and start the consumer
			consumer := consumer.NewGCPConsumer(client, sub, detective)
			if err := consumer.Start(ctx); err != nil {
				log.Fatal("consumer stopped unexpectedly", err)
			}
		},
	}
	cfg.RegisterFlags(serviceCmd.Flags())
	return serviceCmd
}
