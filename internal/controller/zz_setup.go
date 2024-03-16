// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	serviceaccount "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/iam/serviceaccount"
	serviceaccountapikey "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/iam/serviceaccountapikey"
	serviceaccountiambinding "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/iam/serviceaccountiambinding"
	serviceaccountiammember "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/iam/serviceaccountiammember"
	serviceaccountiampolicy "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/iam/serviceaccountiampolicy"
	serviceaccountkey "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/iam/serviceaccountstaticaccesskey"
	group "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/group"
	groupiammember "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/groupiammember"
	groupmembership "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/groupmembership"
	organizationiambinding "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/organizationiambinding"
	organizationiammember "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/organizationiammember"
	osloginsettings "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/osloginsettings"
	samlfederation "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/samlfederation"
	samlfederationuseraccount "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/samlfederationuseraccount"
	usersshkey "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/organizationmanager/usersshkey"
	providerconfig "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/providerconfig"
	cloud "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/resourcemanager/cloud"
	cloudiambinding "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/resourcemanager/cloudiambinding"
	cloudiammember "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/resourcemanager/cloudiammember"
	folder "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/resourcemanager/folder"
	folderiambinding "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/resourcemanager/folderiambinding"
	folderiammember "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/resourcemanager/folderiammember"
	folderiampolicy "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/resourcemanager/folderiampolicy"
	address "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/vpc/address"
	defaultsecuritygroup "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/vpc/defaultsecuritygroup"
	gateway "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/vpc/gateway"
	network "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/vpc/network"
	routetable "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/vpc/routetable"
	securitygroup "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/vpc/securitygrouprule"
	subnet "github.com/frost2sam/provider-upjet-yc-vpc/internal/controller/vpc/subnet"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		serviceaccount.Setup,
		serviceaccountapikey.Setup,
		serviceaccountiambinding.Setup,
		serviceaccountiammember.Setup,
		serviceaccountiampolicy.Setup,
		serviceaccountkey.Setup,
		serviceaccountstaticaccesskey.Setup,
		group.Setup,
		groupiammember.Setup,
		groupmembership.Setup,
		organizationiambinding.Setup,
		organizationiammember.Setup,
		osloginsettings.Setup,
		samlfederation.Setup,
		samlfederationuseraccount.Setup,
		usersshkey.Setup,
		providerconfig.Setup,
		cloud.Setup,
		cloudiambinding.Setup,
		cloudiammember.Setup,
		folder.Setup,
		folderiambinding.Setup,
		folderiammember.Setup,
		folderiampolicy.Setup,
		address.Setup,
		defaultsecuritygroup.Setup,
		gateway.Setup,
		network.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
