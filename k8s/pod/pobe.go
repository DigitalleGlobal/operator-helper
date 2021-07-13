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

package pod

const (
	// DefaultStartupProbeInitialDelaySeconds is the default  initial delay or the startup probe
	DefaultStartupProbeInitialDelaySeconds = 10

	// DefaultStartupProbePeriodSeconds is the default period for the startup probe
	DefaultStartupProbePeriodSeconds = 10

	// DefaultStartupProbeFailureThreshold is the default failure threshold  for the startup probe
	DefaultStartupProbeFailureThreshold = 30

	// DefaultStartupProbeSuccessThreshold is the default success threshold  for the startup probe
	DefaultStartupProbeSuccessThreshold = 1

	// DefaultStartupProbeTimeoutSeconds is the default timeout for the startup probe
	DefaultStartupProbeTimeoutSeconds = 10

	// DefaultReadinessProbeInitialDelaySeconds is the default  initial delay or the readiness probe
	DefaultReadinessProbeInitialDelaySeconds = 20

	// DefaultReadinessProbePeriodSeconds is the default period for the readiness probe
	DefaultReadinessProbePeriodSeconds = 10

	// DefaultReadinessProbeFailureThreshold is the default failure threshold  for the readiness probe
	DefaultReadinessProbeFailureThreshold = 9

	// DefaultReadinessProbeSuccessThreshold is the default success threshold  for the readiness probe
	DefaultReadinessProbeSuccessThreshold = 1

	// DefaultReadinessProbeTimeoutSeconds is the default timeout for the readiness probe
	DefaultReadinessProbeTimeoutSeconds = 5

	// DefaultLivenessProbeInitialDelaySeconds is the default initial delay for the liveness probe
	DefaultLivenessProbeInitialDelaySeconds = 60

	// DefaultLivenessProbePeriodSeconds is the default period for the liveness probe
	DefaultLivenessProbePeriodSeconds = 15

	// DefaultLivenessProbeFailureThreshold is the default failure threshold for the liveness probe
	DefaultLivenessProbeFailureThreshold = 4

	// DefaultLivenessProbeSuccessThreshold is the default success threshold for the liveness probe
	DefaultLivenessProbeSuccessThreshold = 1

	// DefaultLivenessProbeTimeoutSeconds is the default timeout for the liveness probe
	DefaultLivenessProbeTimeoutSeconds = 5
)

// +k8s:openapi-gen=true
// +kubebuilder:object:generate=true

type Probes struct {
	// +optional
	Startup *Probe `json:"startup"`
	// +optional
	Liveness *Probe `json:"liveness"`
	// +optional
	Readiness *Probe `json:"readiness"`
}

type Probe struct {
	// +kubebuilder:validation:Minimum=0
	// +optional
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`
	// +kubebuilder:validation:Minimum=0
	// +optional
	PeriodSeconds int32 `json:"periodSeconds"`
	// +kubebuilder:validation:Minimum=0
	// +optional
	FailureThreshold int32 `json:"failureThreshold"`
	// +kubebuilder:validation:Minimum=0
	// +optional
	SuccessThreshold int32 `json:"successThreshold"`
	// +kubebuilder:validation:Minimum=0
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds"`
}

func (in *Probes) setDefault() (changed bool) {
	if in.Startup == nil {
		changed = true
		in.Startup = &Probe{}
		in.Startup.InitialDelaySeconds = DefaultStartupProbeInitialDelaySeconds
		in.Startup.PeriodSeconds = DefaultStartupProbePeriodSeconds
		in.Startup.FailureThreshold = DefaultStartupProbeFailureThreshold
		in.Startup.SuccessThreshold = DefaultStartupProbeSuccessThreshold
		in.Startup.TimeoutSeconds = DefaultStartupProbeTimeoutSeconds
	}
	if in.Readiness == nil {
		changed = true
		in.Readiness = &Probe{}
		in.Readiness.InitialDelaySeconds = DefaultReadinessProbeInitialDelaySeconds
		in.Readiness.PeriodSeconds = DefaultReadinessProbePeriodSeconds
		in.Readiness.FailureThreshold = DefaultReadinessProbeFailureThreshold
		in.Readiness.SuccessThreshold = DefaultReadinessProbeSuccessThreshold
		in.Readiness.TimeoutSeconds = DefaultReadinessProbeTimeoutSeconds
	}
	if in.Liveness == nil {
		changed = true
		in.Liveness = &Probe{}
		in.Liveness.InitialDelaySeconds = DefaultLivenessProbeInitialDelaySeconds
		in.Liveness.PeriodSeconds = DefaultLivenessProbePeriodSeconds
		in.Liveness.FailureThreshold = DefaultLivenessProbeFailureThreshold
		in.Liveness.SuccessThreshold = DefaultLivenessProbeSuccessThreshold
		in.Liveness.TimeoutSeconds = DefaultLivenessProbeTimeoutSeconds
	}

	return changed
}
