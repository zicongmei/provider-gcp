/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/upbound/upjet/pkg/config"
)

var externalNameConfigs = map[string]config.ExternalName{

	// container
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_cluster": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/clusters/{{ .external_name }}"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.location }}/{{ .parameters.cluster }}/{{ .external_name }}"),

	// containeraws
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_aws_cluster": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/clusters/{{ .external_name }}"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_aws_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.location }}/{{ .parameters.cluster }}/{{ .external_name }}"),

	// containerazure
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_azure_cluster": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/clusters/{{ .external_name }}"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_azure_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.location }}/{{ .parameters.cluster }}/{{ .external_name }}"),
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/azureClients/my-client
	"google_container_azure_client": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/azureClients/{{ .external_name }}"),

	// gkehub
	//
	// Imported by using the following format: projects/{{project}}/locations/global/memberships/{{membership_id}}
	"google_gke_hub_membership": config.TemplatedStringAsIdentifier("membership_id", "projects/{{ .setup.configuration.project }}/locations/global/memberships/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/memberships/{{membership_id}} roles/viewer user:jane@example.com
	"google_gke_hub_membership_iam_member": config.IdentifierFromProvider,

	// Following are dependencies for Anthos

	// containerregistry
	//
	// This resource can not be imported. The resource will create a bucket if missing, and do nothing with any bucket it finds
	"google_container_registry": config.IdentifierFromProvider,

	// compute
	//
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
	"google_compute_subnetwork": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/subnetworks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/networks/{{name}}
	"google_compute_network": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/networks/{{ .external_name }}"),
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}

func externalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := externalNameConfigs[r.Name]; ok {
			r.Version = VersionV1Beta1
			r.ExternalName = e
		}
	}
}
