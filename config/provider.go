/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"https://github.com/antonum/upjet-rediscloud/config/account"
	"https://github.com/antonum/upjet-rediscloud/config/subscription"
	"https://github.com/antonum/upjet-rediscloud/config/database"
)

const (
	resourcePrefix = "rediscloud"
	modulePath     = "github.com/redislabs/provider-rediscloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		account.Configure,
		database.Configure,
		subscription.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
