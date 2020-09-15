package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/fcosta-td/terraform-provider-cassandra/cassandra"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}
