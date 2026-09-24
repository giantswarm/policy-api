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
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
//+kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
//+kubebuilder:printcolumn:name=Policies,type=string,JSONPath=`.status.conditions[?(@.type=="PoliciesResolved")].status`
//+kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// PolicyException is the Schema for the policyexceptions API
// +k8s:openapi-gen=true
type PolicyException struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PolicyExceptionSpec   `json:"spec,omitempty"`
	Status PolicyExceptionStatus `json:"status,omitempty"`
}

// PolicyExceptionSpec defines the desired state of PolicyException
type PolicyExceptionSpec struct {
	// Policies defines the list of policies to be excluded
	Policies []string `json:"policies"`

	// Targes defines the list of target workloads where the exceptions will be applied
	Targets []Target `json:"targets"`
}

// Condition types of a PolicyException, written by kyverno-policy-operator.
const (
	// PolicyExceptionReady is True when every Kyverno PolicyException generated for it is applied,
	// stale ones are removed, and every target is translated. It does not check that the listed
	// policies exist; PoliciesResolved reports that.
	PolicyExceptionReady = "Ready"
	// PolicyExceptionPoliciesResolved is False when a listed policy matches no CEL policy. It is
	// informational and does not affect Ready, so health checks on Ready (Flux wait, kstatus)
	// are not blocked by policies still being migrated.
	PolicyExceptionPoliciesResolved = "PoliciesResolved"
	// PolicyExceptionTargetsTranslated is False when a target cannot be expressed in a CEL exception.
	PolicyExceptionTargetsTranslated = "TargetsTranslated"
)

// Reasons of the Ready condition. When several checks fail, Ready takes the first failing
// reason in the order InvalidNamespace, NameTaken, LookupFailed, ApplyFailed, DeleteFailed,
// UnsupportedKind, and its message lists all of them.
const (
	ReasonReconciled       = "Reconciled"
	ReasonInvalidNamespace = "InvalidNamespace"
	ReasonNameTaken        = "NameTaken"
	ReasonLookupFailed     = "LookupFailed"
	ReasonApplyFailed      = "ApplyFailed"
	ReasonDeleteFailed     = "DeleteFailed"
)

// Reasons of the PoliciesResolved condition. PolicyNotFound wins over NotMigrated, and the
// message names the policies for each.
const (
	// ReasonResolved means every listed policy matches a CEL policy.
	ReasonResolved = "Resolved"
	// ReasonNotMigrated means a listed policy matches only a legacy ClusterPolicy. The legacy
	// exception covers it, and the CEL exception takes over once the policy is migrated.
	// This is expected during the migration.
	ReasonNotMigrated = "NotMigrated"
	// ReasonPolicyNotFound means a listed policy matches no policy at all, for example a typo,
	// a removed policy, or a policy not installed yet.
	ReasonPolicyNotFound = "PolicyNotFound"
)

// Reasons of the TargetsTranslated condition.
const (
	ReasonTranslated = "Translated"
	// ReasonUnsupportedKind means a target kind cannot be expressed in a CEL exception. It is a
	// reason of both TargetsTranslated and Ready.
	ReasonUnsupportedKind = "UnsupportedKind"
)

// PolicyExceptionStatus defines the observed state of PolicyException. kyverno-policy-operator
// writes it.
type PolicyExceptionStatus struct {
	// ObservedGeneration is the metadata.generation the status was computed for.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions are Ready, PoliciesResolved and TargetsTranslated.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// GeneratedExceptions are the Kyverno PolicyExceptions written for this PolicyException.
	// +listType=atomic
	// +optional
	GeneratedExceptions []GeneratedException `json:"generatedExceptions,omitempty"`

	// UnresolvedPolicies are listed policies that match no CEL policy: those that match only a
	// legacy ClusterPolicy and those that match no policy at all.
	// +listType=atomic
	// +optional
	UnresolvedPolicies []string `json:"unresolvedPolicies,omitempty"`

	// UnsupportedTargetKinds are target kinds left out of the CEL exception, such as "Pod/exec".
	// +listType=atomic
	// +optional
	UnsupportedTargetKinds []string `json:"unsupportedTargetKinds,omitempty"`
}

// GeneratedException references a Kyverno PolicyException generated for a PolicyException.
type GeneratedException struct {
	// APIVersion is policies.kyverno.io/v1 or kyverno.io/v2.
	APIVersion string `json:"apiVersion"`
	Namespace  string `json:"namespace"`
	Name       string `json:"name"`
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
	addKnownTypes(&PolicyException{}, &PolicyExceptionList{})
}
