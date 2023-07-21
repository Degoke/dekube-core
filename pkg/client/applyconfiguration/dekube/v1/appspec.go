/*
Copyright Adegoke Adewoye.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// AppSpecApplyConfiguration represents an declarative configuration of the AppSpec type for use
// with apply.
type AppSpecApplyConfiguration struct {
	Name        *string                         `json:"name,omitempty"`
	Image       *string                         `json:"image,omitempty"`
	Replicas    *int32                          `json:"replicas,omitempty"`
	Annotations *map[string]string              `json:"annotations,omitempty"`
	Environment *map[string]string              `json:"environment,omitempty"`
	Labels      *map[string]string              `json:"labels,omitempty"`
	Limits      *AppResourcesApplyConfiguration `json:"limits,omitempty"`
	Requests    *AppResourcesApplyConfiguration `json:"requests,omitempty"`
	Domain      *string                         `json:"domain,omitempty"`
	Port        *string                         `json:"port,omitempty"`
}

// AppSpecApplyConfiguration constructs an declarative configuration of the AppSpec type for use with
// apply.
func AppSpec() *AppSpecApplyConfiguration {
	return &AppSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithName(value string) *AppSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithImage(value string) *AppSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithReplicas(value int32) *AppSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithAnnotations sets the Annotations field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Annotations field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithAnnotations(value map[string]string) *AppSpecApplyConfiguration {
	b.Annotations = &value
	return b
}

// WithEnvironment sets the Environment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Environment field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithEnvironment(value map[string]string) *AppSpecApplyConfiguration {
	b.Environment = &value
	return b
}

// WithLabels sets the Labels field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Labels field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithLabels(value map[string]string) *AppSpecApplyConfiguration {
	b.Labels = &value
	return b
}

// WithLimits sets the Limits field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Limits field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithLimits(value *AppResourcesApplyConfiguration) *AppSpecApplyConfiguration {
	b.Limits = value
	return b
}

// WithRequests sets the Requests field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Requests field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithRequests(value *AppResourcesApplyConfiguration) *AppSpecApplyConfiguration {
	b.Requests = value
	return b
}

// WithDomain sets the Domain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Domain field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithDomain(value string) *AppSpecApplyConfiguration {
	b.Domain = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *AppSpecApplyConfiguration) WithPort(value string) *AppSpecApplyConfiguration {
	b.Port = &value
	return b
}
