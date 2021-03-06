/*
Copyright 2015 The Kubernetes Authors.

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

package testgroup

import (
	"k8s.io/kubernetes/pkg/api"
	metav1 "k8s.io/kubernetes/pkg/apis/meta/v1"
)

// +genclient=true

type TestType struct {
	metav1.TypeMeta
	api.ObjectMeta
	Status TestTypeStatus
}

type TestTypeList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []TestType
}

type TestTypeStatus struct {
	Blah string
}
