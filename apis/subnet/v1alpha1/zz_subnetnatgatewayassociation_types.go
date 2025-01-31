/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubnetNatGatewayAssociationObservation struct {
}

type SubnetNatGatewayAssociationParameters struct {

	// +kubebuilder:validation:Required
	NatGatewayID *string `json:"natGatewayId" tf:"nat_gateway_id,omitempty"`

	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// SubnetNatGatewayAssociationSpec defines the desired state of SubnetNatGatewayAssociation
type SubnetNatGatewayAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetNatGatewayAssociationParameters `json:"forProvider"`
}

// SubnetNatGatewayAssociationStatus defines the observed state of SubnetNatGatewayAssociation.
type SubnetNatGatewayAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetNatGatewayAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetNatGatewayAssociation is the Schema for the SubnetNatGatewayAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SubnetNatGatewayAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetNatGatewayAssociationSpec   `json:"spec"`
	Status            SubnetNatGatewayAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetNatGatewayAssociationList contains a list of SubnetNatGatewayAssociations
type SubnetNatGatewayAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetNatGatewayAssociation `json:"items"`
}

// Repository type metadata.
var (
	SubnetNatGatewayAssociationKind             = "SubnetNatGatewayAssociation"
	SubnetNatGatewayAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: SubnetNatGatewayAssociationKind}.String()
	SubnetNatGatewayAssociationKindAPIVersion   = SubnetNatGatewayAssociationKind + "." + GroupVersion.String()
	SubnetNatGatewayAssociationGroupVersionKind = GroupVersion.WithKind(SubnetNatGatewayAssociationKind)
)

func init() {
	SchemeBuilder.Register(&SubnetNatGatewayAssociation{}, &SubnetNatGatewayAssociationList{})
}
