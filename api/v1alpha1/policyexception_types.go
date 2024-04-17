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

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=gspolex

// PolicyException is the Schema for the policyexceptions API
// +k8s:openapi-gen=true
type PolicyException struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PolicyExceptionSpec `json:"spec,omitempty"`
}

// PolicyExceptionSpec defines the desired state of PolicyException
type PolicyExceptionSpec struct {
	// Policies defines the list of policies to be excluded
	Policies []string `json:"policies"`

	// Targes defines the list of target workloads where the exceptions will be applied
	Targets []Target `json:"targets"`
}

// Target defines a resource to which a PolicyException applies
// +k8s:openapi-gen=true
type Target struct {
	// +listType=atomic
	Namespaces []string `json:"namespaces"`
	// +listType=atomic
	Names []string `json:"names"`
	Kind  string   `json:"kind"`
}

//+kubebuilder:object:root=true

// PolicyExceptionList contains a list of PolicyException
type PolicyExceptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyException `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PolicyException{}, &PolicyExceptionList{})
}
