/*
Copyright 2021 The Kubernetes Authors.

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

// From https://github.com/kubernetes/kubernetes/blob/
// 0712728ee9caab4d04ec0a751855aca8c891b690/pkg/generated/openapi/cmd/models-schema/main.go

// Package main provides the entry point for the policy-api controller
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/validation/spec"

	openapi "github.com/giantswarm/policy-api/pkg/generated/openapi"
)

// Outputs openAPI schema JSON containing the schema definitions in zz_generated.openapi.go.
func main() {
	err := output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %v\n", err)
		os.Exit(1)
	}
}

func output() error {
	refFunc := func(name string) spec.Ref {
		return spec.MustCreateRef(fmt.Sprintf("#/definitions/%s", friendlyName(name)))
	}

	defs := openapi.GetOpenAPIDefinitions(refFunc)

	schemaDefs := make(map[string]spec.Schema, len(defs))
	for k, v := range defs {
		// Replace top-level schema with v2 if a v2 schema is embedded
		// so that the output of this program is always in OpenAPI v2.
		// This is done by looking up an extension that marks the embedded v2
		// schema, and, if the v2 schema is found, make it the resulting schema for
		// the type.
		if schema, ok := v.Schema.Extensions[common.ExtensionV2Schema]; ok {
			if v2Schema, isOpenAPISchema := schema.(spec.Schema); isOpenAPISchema {
				schemaDefs[friendlyName(k)] = v2Schema
				continue
			}
		}

		schemaDefs[friendlyName(k)] = v.Schema
	}
	data, err := json.Marshal(&spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Definitions: schemaDefs,
			Info: &spec.Info{
				InfoProps: spec.InfoProps{
					Title:   "Gateway API",
					Version: "unversioned",
				},
			},
			Swagger: "2.0",
		},
	})
	if err != nil {
		return fmt.Errorf("error serializing api definitions: %w", err)
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		return fmt.Errorf("error writing api definitions: %w", err)
	}

	return nil
}

// From vendor/k8s.io/apiserver/pkg/endpoints/openapi/openapi.go
func friendlyName(name string) string {
	nameParts := strings.Split(name, "/")
	// Reverse first part. e.g., io.k8s... instead of k8s.io...
	if len(nameParts) > 0 && strings.Contains(nameParts[0], ".") {
		parts := strings.Split(nameParts[0], ".")
		for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
			parts[i], parts[j] = parts[j], parts[i]
		}
		nameParts[0] = strings.Join(parts, ".")
	}
	return strings.Join(nameParts, ".")
}
