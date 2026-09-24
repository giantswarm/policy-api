/*
Copyright 2026.

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
	"encoding/json"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func testStatus() PolicyExceptionStatus {
	return PolicyExceptionStatus{
		ObservedGeneration: 3,
		Conditions: []metav1.Condition{{
			Type:               PolicyExceptionReady,
			Status:             metav1.ConditionFalse,
			Reason:             ReasonNameTaken,
			Message:            "name taken",
			LastTransitionTime: metav1.Unix(1700000000, 0).Rfc3339Copy(),
		}},
		GeneratedExceptions: []GeneratedException{{
			APIVersion: "policies.kyverno.io/v1",
			Namespace:  "policy-exceptions",
			Name:       "gs-foo",
		}},
		UnresolvedPolicies:     []string{"missing-policy"},
		UnsupportedTargetKinds: []string{"Pod/exec"},
	}
}

func TestPolicyExceptionStatusJSONRoundTrip(t *testing.T) {
	in := PolicyException{Status: testStatus()}

	data, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, key := range []string{
		`"observedGeneration":3`,
		`"conditions":[`,
		`"generatedExceptions":[{"apiVersion":"policies.kyverno.io/v1","namespace":"policy-exceptions","name":"gs-foo"}]`,
		`"unresolvedPolicies":["missing-policy"]`,
		`"unsupportedTargetKinds":["Pod/exec"]`,
	} {
		if !strings.Contains(string(data), key) {
			t.Errorf("marshalled JSON lacks %s: %s", key, data)
		}
	}

	var out PolicyException
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !equality.Semantic.DeepEqual(in.Status, out.Status) {
		t.Errorf("round trip changed status:\n got %+v\nwant %+v", out.Status, in.Status)
	}
}

func TestPolicyExceptionZeroStatusMarshalsToEmptyObject(t *testing.T) {
	data, err := json.Marshal(PolicyExceptionStatus{})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if string(data) != "{}" {
		t.Errorf("empty status marshalled to %s, want {}", data)
	}
}

func TestPolicyExceptionDeepCopyDoesNotAliasStatus(t *testing.T) {
	in := &PolicyException{Status: testStatus()}
	out := in.DeepCopy()

	out.Status.Conditions[0].Reason = ReasonReconciled
	out.Status.GeneratedExceptions[0].Name = "other-name"
	out.Status.UnresolvedPolicies[0] = "other-policy"
	out.Status.UnsupportedTargetKinds[0] = "Pod/attach"

	if !equality.Semantic.DeepEqual(in.Status, testStatus()) {
		t.Errorf("mutating the copy changed the original: %+v", in.Status)
	}
}
