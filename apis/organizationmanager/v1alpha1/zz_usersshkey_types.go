// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserSSHKeyInitParameters struct {

	// Data of the user ssh key.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// User ssh key will be no longer valid after expiration timestamp.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// Name of the user ssh key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Organization that the user ssh key belongs to.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Subject that the user ssh key belongs to.
	SubjectID *string `json:"subjectId,omitempty" tf:"subject_id,omitempty"`
}

type UserSSHKeyObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Data of the user ssh key.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// User ssh key will be no longer valid after expiration timestamp.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the user ssh key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Organization that the user ssh key belongs to.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Subject that the user ssh key belongs to.
	SubjectID *string `json:"subjectId,omitempty" tf:"subject_id,omitempty"`
}

type UserSSHKeyParameters struct {

	// Data of the user ssh key.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// User ssh key will be no longer valid after expiration timestamp.
	// +kubebuilder:validation:Optional
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// Name of the user ssh key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Organization that the user ssh key belongs to.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Subject that the user ssh key belongs to.
	// +kubebuilder:validation:Optional
	SubjectID *string `json:"subjectId,omitempty" tf:"subject_id,omitempty"`
}

// UserSSHKeySpec defines the desired state of UserSSHKey
type UserSSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserSSHKeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider UserSSHKeyInitParameters `json:"initProvider,omitempty"`
}

// UserSSHKeyStatus defines the observed state of UserSSHKey.
type UserSSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserSSHKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserSSHKey is the Schema for the UserSSHKeys API. Allows management of User Ssh Keys within an existing Yandex.Cloud Organization and Subject.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type UserSSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.data) || (has(self.initProvider) && has(self.initProvider.data))",message="spec.forProvider.data is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.organizationId) || (has(self.initProvider) && has(self.initProvider.organizationId))",message="spec.forProvider.organizationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subjectId) || (has(self.initProvider) && has(self.initProvider.subjectId))",message="spec.forProvider.subjectId is a required parameter"
	Spec   UserSSHKeySpec   `json:"spec"`
	Status UserSSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserSSHKeyList contains a list of UserSSHKeys
type UserSSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserSSHKey `json:"items"`
}

// Repository type metadata.
var (
	UserSSHKey_Kind             = "UserSSHKey"
	UserSSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserSSHKey_Kind}.String()
	UserSSHKey_KindAPIVersion   = UserSSHKey_Kind + "." + CRDGroupVersion.String()
	UserSSHKey_GroupVersionKind = CRDGroupVersion.WithKind(UserSSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&UserSSHKey{}, &UserSSHKeyList{})
}
