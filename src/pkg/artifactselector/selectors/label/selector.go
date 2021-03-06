// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package label

import (
	"github.com/goharbor/harbor/src/pkg/artifactselector"
	"strings"
)

const (
	// Kind ...
	Kind = "label"
	// With labels
	With = "withLabels"
	// Without labels
	Without = "withoutLabels"
)

// selector is for label selector
type selector struct {
	// Pre defined pattern decorations
	// "with" or "without"
	decoration string
	// Label list
	labels []string
}

// Select candidates by the labels
func (s *selector) Select(artifacts []*artifactselector.Candidate) (selected []*artifactselector.Candidate, err error) {
	for _, art := range artifacts {
		if isMatched(s.labels, art.Labels, s.decoration) {
			selected = append(selected, art)
		}
	}

	return selected, nil
}

// New is factory method for list selector
func New(decoration string, pattern string) artifactselector.Selector {
	labels := make([]string, 0)
	if len(pattern) > 0 {
		labels = append(labels, strings.Split(pattern, ",")...)
	}

	return &selector{
		decoration: decoration,
		labels:     labels,
	}
}

// Check if the resource labels match the pattern labels
func isMatched(patternLbls []string, resLbls []string, decoration string) bool {
	hash := make(map[string]bool)

	for _, lbl := range resLbls {
		hash[lbl] = true
	}

	for _, lbl := range patternLbls {
		_, exists := hash[lbl]

		if decoration == Without && exists {
			return false
		}

		if decoration == With && !exists {
			return false
		}
	}

	return true
}
