// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package billing

import (
	"sort"
)

// mapper pairs related conversion/merging functions.
type mapper struct {
	serviceDisplayName string
}

// mappers maps terraform resource types (i.e. `google_project`) into
// a slice of service names and ids.
//
// Modelling of relationships:
// terraform resources to CAI assets as []mapperFuncs:
// 1:1 = [mapper{convert: convertAbc}]                  (len=1)
// 1:N = [mapper{convert: convertAbc}, ...]             (len=N)
// N:1 = [mapper{convert: convertAbc, merge: mergeAbc}] (len=1)
func mappers() map[string]mapper {
	return map[string]mapper{
		// TODO: Use a generated mapping once it lands in the conversion library.
		"sqladmin.googleapis.com/Instance": {serviceDisplayName: "Cloud SQL"},
		"compute.googleapis.com/Instance":  {serviceDisplayName: "Compute Engine"},
	}
}

// SupportedTerraformBillingResources returns a sorted list of terraform resource names.
func SupportedTerraformBillingResources() []string {
	fns := mappers()
	list := make([]string, 0, len(fns))
	for k := range fns {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}
