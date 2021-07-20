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

import v1 "k8s.io/api/core/v1"

const (
	// DefaultStartupProbeInitialDelaySeconds is the default  initial delay or the startup probe
	DefaultStartupProbeInitialDelaySeconds = 10

	// DefaultStartupProbePeriodSeconds is the default period for the startup probe
	DefaultStartupProbePeriodSeconds = 10

	// DefaultStartupProbeFailureThreshold is the default failure threshold  for the startup probe
	DefaultStartupProbeFailureThreshold = 10

	// DefaultStartupProbeSuccessThreshold is the default success threshold  for the startup probe
	DefaultStartupProbeSuccessThreshold = 1

	// DefaultStartupProbeTimeoutSeconds is the default timeout for the startup probe
	DefaultStartupProbeTimeoutSeconds = 10

	// DefaultReadinessProbeInitialDelaySeconds is the default  initial delay or the readiness probe
	DefaultReadinessProbeInitialDelaySeconds = 10

	// DefaultReadinessProbePeriodSeconds is the default period for the readiness probe
	DefaultReadinessProbePeriodSeconds = 10

	// DefaultReadinessProbeFailureThreshold is the default failure threshold  for the readiness probe
	DefaultReadinessProbeFailureThreshold = 3

	// DefaultReadinessProbeSuccessThreshold is the default success threshold  for the readiness probe
	DefaultReadinessProbeSuccessThreshold = 1

	// DefaultReadinessProbeTimeoutSeconds is the default timeout for the readiness probe
	DefaultReadinessProbeTimeoutSeconds = 10

	// DefaultLivenessProbeInitialDelaySeconds is the default initial delay for the liveness probe
	DefaultLivenessProbeInitialDelaySeconds = 10

	// DefaultLivenessProbePeriodSeconds is the default period for the liveness probe
	DefaultLivenessProbePeriodSeconds = 10

	// DefaultLivenessProbeFailureThreshold is the default failure threshold for the liveness probe
	DefaultLivenessProbeFailureThreshold = 3

	// DefaultLivenessProbeSuccessThreshold is the default success threshold for the liveness probe
	DefaultLivenessProbeSuccessThreshold = 1

	// DefaultLivenessProbeTimeoutSeconds is the default timeout for the liveness probe
	DefaultLivenessProbeTimeoutSeconds = 3
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

func (in *Probe) ToK8sProbe(handler v1.Handler) v1.Probe {
	return v1.Probe{
		Handler:             handler,
		InitialDelaySeconds: in.InitialDelaySeconds,
		PeriodSeconds:       in.PeriodSeconds,
		SuccessThreshold:    in.SuccessThreshold,
		FailureThreshold:    in.FailureThreshold,
		TimeoutSeconds:      in.TimeoutSeconds,
	}
}

// SetDefault set the default values
func (in *Probes) SetDefault() (changed bool) {
	if in.Startup == nil {
		changed = true
		in.Startup = &Probe{}
	}
	if startupDefault(in.Startup) {
		changed = true
	}
	if in.Readiness == nil {
		changed = true
		in.Readiness = &Probe{}
	}
	if readinessDefault(in.Readiness) {
		changed = true
	}
	if in.Liveness == nil {
		changed = true
		in.Liveness = &Probe{}
	}
	if livenessDefault(in.Liveness) {
		changed = true
	}
	return changed
}

func startupDefault(probe *Probe) bool {
	return probeDefault(probe,
		DefaultStartupProbeInitialDelaySeconds,
		DefaultStartupProbePeriodSeconds,
		DefaultStartupProbeFailureThreshold,
		DefaultStartupProbeSuccessThreshold,
		DefaultStartupProbeTimeoutSeconds)
}

func livenessDefault(probe *Probe) bool {
	return probeDefault(probe,
		DefaultLivenessProbeInitialDelaySeconds,
		DefaultLivenessProbePeriodSeconds,
		DefaultLivenessProbeFailureThreshold,
		DefaultLivenessProbeSuccessThreshold,
		DefaultLivenessProbeTimeoutSeconds)
}

func readinessDefault(probe *Probe) bool {
	return probeDefault(probe,
		DefaultReadinessProbeInitialDelaySeconds,
		DefaultReadinessProbePeriodSeconds,
		DefaultReadinessProbeFailureThreshold,
		DefaultReadinessProbeSuccessThreshold,
		DefaultReadinessProbeTimeoutSeconds)
}

func probeDefault(probe *Probe, delaySec, periodSec, failureThreshold, successThreshold, timeoutSeconds int32) (changed bool) {
	if probe.InitialDelaySeconds == 0 {
		changed = true
		probe.InitialDelaySeconds = delaySec
	}
	if probe.PeriodSeconds == 0 {
		changed = true
		probe.PeriodSeconds = periodSec
	}
	if probe.FailureThreshold == 0 {
		changed = true
		probe.FailureThreshold = failureThreshold
	}
	if probe.SuccessThreshold == 0 {
		changed = true
		probe.SuccessThreshold = successThreshold
	}
	if probe.TimeoutSeconds == 0 {
		changed = true
		probe.TimeoutSeconds = timeoutSeconds
	}
	return
}
