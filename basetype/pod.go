/*
 * Copyright 2021 - now, the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package basetype

import (
	v1 "k8s.io/api/core/v1"
)

// +k8s:openapi-gen=true
// +kubebuilder:object:generate=true

// PodConfig defines the configurations of a kubernetes pod
type PodConfig struct {

	// Env defines environment variables for the pod
	Env []v1.EnvVar `json:"env,omitempty"`

	// Affinity defines the pod's scheduling constraints
	Affinity v1.Affinity `json:"affinity,omitempty"`

	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Optional duration in seconds the pod may be active on the node relative to
	// StartTime before the system will actively try to mark it failed and kill associated containers.
	// Value must be a positive integer.
	ActiveDeadlineSeconds int64 `json:"activeDeadlineSeconds,omitempty"`
	// Restart policy for all containers within the pod.
	// One of Always, OnFailure, Never.
	// Default to Always.
	RestartPolicy v1.RestartPolicy `json:"restartPolicy,omitempty"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Some fields are also present in container.securityContext.  Field values of
	// container.securityContext take precedence over field values of PodSecurityContext.
	SecurityContext v1.PodSecurityContext `json:"securityContext,omitempty"`

	// Tolerations are attached to tolerates any taint that matches
	// the triple <key,value,effect> using the matching operator <operator>.
	Tolerations []v1.Toleration `json:"tolerations,omitempty"`

	// Labels defines the labels to attach to the broker pod
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations defines the annotations to attach to the pod
	Annotations map[string]string `json:"annotations,omitempty"`

	// TerminationGracePeriodSeconds is the duration in seconds after the processes running in the pod are sent
	// a termination signal and the time when the processes are forcibly halted with a kill signal.
	// Set this value longer than the expected cleanup time for your process.
	// Defaults to 30 seconds.
	// +optional
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`
}
