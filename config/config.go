package config

import "github.com/spf13/pflag"

type Config struct {
	ProjectID       string
	SubscriptionID  string
	OutboundTopicID string
	IPBlockListURL  string
}

// RegisterFlags adds the configuration flags to the given flag set.
func (c *Config) RegisterFlags(f *pflag.FlagSet) {
	f.StringVarP(&c.ProjectID, "projectId", "p", "", "Project ID")
	f.StringVarP(&c.SubscriptionID, "subscriptionId", "s", "", "Subscription ID")
	f.StringVarP(&c.OutboundTopicID, "outboundTopicId", "t", "", "Outbound Topic ID")
	f.StringVarP(&c.IPBlockListURL, "ipBlockListUrl", "u", "https://antoinevastel.com/data/avastel-longtime-bot-ips.txt", "Ip Blocklist for the bot detection")
}
