/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/frost2sam/provider-upjet-yc-vpc/config/iam"
	"github.com/frost2sam/provider-upjet-yc-vpc/config/organizationmanager"
	"github.com/frost2sam/provider-upjet-yc-vpc/config/resourcemanager"
	"github.com/frost2sam/provider-upjet-yc-vpc/config/vpc"
)

const (
	resourcePrefix = "yandex-cloud"
	modulePath     = "github.com/frost2sam/provider-upjet-yc-vpc"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("yandex-cloud.upjet.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		resourcemanager.Configure,
		organizationmanager.Configure,
		iam.Configure,
		vpc.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
