/*
Copyright 2024.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PolicyConfigSpec defines the desired state of PolicyConfig
type PolicyConfigSpec struct {
	PolicyName  string `json:"policyName,omitempty"`
	PolicyState string `json:"policyState,omitempty"`
}

// PolicyConfig is the Schema for the PolicyConfigs API
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=gspolconfig,scope=Cluster
// +k8s:openapi-gen=true
type PolicyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PolicyConfigSpec `json:"spec,omitempty"`
}

// PolicyConfigList contains a list of PolicyConfigs
// +kubebuilder:object:root=true
type PolicyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PolicyConfig{}, &PolicyConfigList{})
}
