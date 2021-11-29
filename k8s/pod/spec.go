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

import (
	"context"
	"fmt"
	"github.com/monimesl/operator-helper/basetype"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewSpec(cfg basetype.PodConfig, volumes []v1.Volume, initContainers []v1.Container, containers []v1.Container) v1.PodSpec {
	return v1.PodSpec{
		Volumes:                       volumes,
		InitContainers:                initContainers,
		Containers:                    containers,
		RestartPolicy:                 cfg.Spec.RestartPolicy,
		TerminationGracePeriodSeconds: cfg.Spec.TerminationGracePeriodSeconds,
		ActiveDeadlineSeconds:         cfg.Spec.ActiveDeadlineSeconds,
		DNSPolicy:                     cfg.Spec.DNSPolicy,
		NodeSelector:                  cfg.Spec.NodeSelector,
		ServiceAccountName:            cfg.Spec.ServiceAccountName,
		NodeName:                      cfg.Spec.NodeName,
		SecurityContext:               cfg.Spec.SecurityContext,
		Affinity:                      cfg.Spec.Affinity,
		Tolerations:                   cfg.Spec.Tolerations,
		PriorityClassName:             cfg.Spec.PriorityClassName,
		Priority:                      cfg.Spec.Priority,
		PreemptionPolicy:              cfg.Spec.PreemptionPolicy,
		Overhead:                      cfg.Spec.Overhead,
	}
}

func NewMetadata(cfg basetype.PodConfig, name, generateName string, labels, annotations map[string]string) metav1.ObjectMeta {
	metadata := cfg.ObjectMeta
	metadata.Name = name
	metadata.GenerateName = generateName
	if metadata.Labels == nil {
		metadata.Labels = map[string]string{}
	}
	if metadata.Annotations == nil {
		metadata.Annotations = map[string]string{}
	}
	for label, value := range labels {
		metadata.Labels[label] = value
	}
	for annotation, value := range annotations {
		metadata.Annotations[annotation] = value
	}
	return metadata
}

// IsReady checks if the pod is ready
func IsReady(pod *v1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == v1.PodReady && condition.Status == v1.ConditionTrue {
			return true
		}
	}
	return false
}

// ListAllWithMatchingLabels list the pods matching the labels
func ListAllWithMatchingLabels(cl client.Client, namespace string, labels map[string]string) (*v1.PodList, error) {
	selector, err := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
		MatchLabels: labels,
	})
	if err != nil {
		return nil, fmt.Errorf("error on creating selector from label selector: %w", err)
	}
	list := &v1.PodList{}
	listOpts := &client.ListOptions{
		Namespace:     namespace,
		LabelSelector: selector,
	}
	err = cl.List(context.TODO(), list, listOpts)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ListAllWithMatchingLabelsByReadiness list the pods matching the labels
func ListAllWithMatchingLabelsByReadiness(cl client.Client, namespace string, labels map[string]string) (ready []v1.Pod, unready []v1.Pod, err error) {
	pods, err0 := ListAllWithMatchingLabels(cl, namespace, labels)
	if err0 != nil {
		err = err0
		return
	}
	for _, pod := range pods.Items {
		if IsReady(&pod) {
			ready = append(ready, pod)
		} else {
			unready = append(unready, pod)
		}
	}
	return
}
