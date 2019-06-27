package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PacketMachineProviderConfig contains Config for Packet machines.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PacketMachineProviderConfig struct {
	metav1.TypeMeta `json:",inline"`

	ProjectID    string   `json:"projectid"`
	Facilities   []string `json:"facilities"`
	InstanceType string   `json:"instanceType"`
	Tags         []string `json:"tags,omitempty"`
	OS           string   `json:"os,omitempty"`
	BillingCycle string   `json:"billingCycle,omitempty"`
}
