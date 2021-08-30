/*
Copyright 2021 Jose Nuñez.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type Core5gStatus string

const (
	Core5gInitial   Core5gStatus = "Initial"
	Core5gRendering Core5gStatus = "Rendering"
	Core5gCreating  Core5gStatus = "Creating"
	Core5gRendered  Core5gStatus = "Rendered"
)

// Core5gSpec defines the desired state of Core5g
type Core5gSpec struct {
	Image string `json:"image"`
}

// Core5g is the Schema for the core5gs API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Core5g struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Core5gSpec   `json:"spec,omitempty"`
	Status Core5gStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Core5gList contains a list of Core5g
type Core5gList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Core5g `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Core5g{}, &Core5gList{})
}
