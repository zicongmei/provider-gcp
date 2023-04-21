// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	"github.com/upbound/provider-gcp/config/compute"
	tjconfig "github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/registry/reference"
	"github.com/upbound/upjet/pkg/types/name"

	"github.com/upbound/provider-gcp/config/container"
	"github.com/upbound/provider-gcp/config/containeraws"
	"github.com/upbound/provider-gcp/config/containerazure"
	"github.com/upbound/provider-gcp/config/gkehub"
)

const (
	resourcePrefix = "gcp"
	modulePath     = "github.com/upbound/provider-gcp"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

var skipList = []string{}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		tjconfig.WithDefaultResourceOptions(
			groupOverrides(),
			externalNameConfig(),
			defaultVersion(),
			externalNameConfigurations(),
			descriptionOverrides(),
		),
		tjconfig.WithRootGroup("gcp.zicong.io"),
		tjconfig.WithShortName("gcp"),
		// Comment out the following line to generate all resources.
		tjconfig.WithIncludeList(resourcesWithExternalNameConfig()),
		tjconfig.WithReferenceInjectors([]tjconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		tjconfig.WithSkipList(skipList),
		tjconfig.WithFeaturesPackage("internal/features"),
	)

	for _, configure := range []func(provider *tjconfig.Provider){

		containeraws.Configure,
		containerazure.Configure,
		container.Configure,
		compute.Configure,
		gkehub.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// resourcesWithExternalNameConfig returns the list of resources that have external
// name configured in ExternalNameConfigs table.
func resourcesWithExternalNameConfig() []string {
	l := make([]string, len(externalNameConfigs))
	i := 0
	for n := range externalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}

func init() {
	// GCP specific acronyms

	// Todo(turkenh): move to Terrajet?
	name.AddAcronym("idp", "IdP")
	name.AddAcronym("oauth", "OAuth")
}
