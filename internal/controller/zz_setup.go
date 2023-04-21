/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	network "github.com/upbound/provider-gcp/internal/controller/compute/network"
	subnetwork "github.com/upbound/provider-gcp/internal/controller/compute/subnetwork"
	cluster "github.com/upbound/provider-gcp/internal/controller/container/cluster"
	nodepool "github.com/upbound/provider-gcp/internal/controller/container/nodepool"
	registry "github.com/upbound/provider-gcp/internal/controller/container/registry"
	clustercontaineraws "github.com/upbound/provider-gcp/internal/controller/containeraws/cluster"
	nodepoolcontaineraws "github.com/upbound/provider-gcp/internal/controller/containeraws/nodepool"
	client "github.com/upbound/provider-gcp/internal/controller/containerazure/client"
	clustercontainerazure "github.com/upbound/provider-gcp/internal/controller/containerazure/cluster"
	nodepoolcontainerazure "github.com/upbound/provider-gcp/internal/controller/containerazure/nodepool"
	membership "github.com/upbound/provider-gcp/internal/controller/gkehub/membership"
	membershipiammember "github.com/upbound/provider-gcp/internal/controller/gkehub/membershipiammember"
	providerconfig "github.com/upbound/provider-gcp/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		network.Setup,
		subnetwork.Setup,
		cluster.Setup,
		nodepool.Setup,
		registry.Setup,
		clustercontaineraws.Setup,
		nodepoolcontaineraws.Setup,
		client.Setup,
		clustercontainerazure.Setup,
		nodepoolcontainerazure.Setup,
		membership.Setup,
		membershipiammember.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
