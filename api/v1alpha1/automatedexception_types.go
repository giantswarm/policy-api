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

// AutomatedExceptionStatus defines the observed state of AutomatedException
type AutomatedExceptionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=autopolex
// +kubebuilder:printcolumn:name="Namespaces",type=string,JSONPath=`.spec.targets[].namespaces`
// +kubebuilder:printcolumn:name="Kind",type=string,JSONPath=`.spec.targets[].kind`
// +kubebuilder:printcolumn:name="Names",type=string,JSONPath=`.spec.targets[].names`
// +kubebuilder:printcolumn:name="Policies",type=string,JSONPath=`.spec.policies`

// AutomatedException is the Schema for the automatedexceptions API
type AutomatedException struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PolicyExceptionSpec      `json:"spec,omitempty"`
	Status AutomatedExceptionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AutomatedExceptionList contains a list of AutomatedException
type AutomatedExceptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutomatedException `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AutomatedException{}, &AutomatedExceptionList{})
}
