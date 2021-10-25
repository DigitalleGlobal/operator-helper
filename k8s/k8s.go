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

package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	k8Labels "k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

// EnvVarPodIP holds the POD's IP
const EnvVarPodIP = "POD_IP"
const EnvVarEnvoySidecarStatus = "ENVOY_SIDECAR_STATUS"

// See https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/#labels
const (
	// LabelAppName defines name of the application e.g postgres
	LabelAppName = "app.kubernetes.io/name"
	// LabelAppInstance defines the unique name identifying the instance of an application
	LabelAppInstance = "app.kubernetes.io/instance"
	// LabelAppVersion defines the app label
	LabelAppVersion = "app.kubernetes.io/version"
	// LabelAppVComponent defines the component within the architecture e.g database
	LabelAppVComponent = "app.kubernetes.io/component"
	// LabelAppManagedBy defines the managed-by label
	LabelAppManagedBy = "app.kubernetes.io/managed-by"
	// LabelAppPartOf defines the managed-by label
	LabelAppPartOf = "app.kubernetes.io/part-of"
)

// ContainerShellCommand is helper factory method to create the shell command
func ContainerShellCommand() []string {
	return []string{
		"sh",
		"-c",
	}
}

// WaitForPodsToTerminate wait for all the pods matching the labels to terminate
func WaitForPodsToTerminate(k8sClient client.Client, namespace string, labels map[string]string) (err error) {
	listOptions := &client.ListOptions{
		Namespace:     namespace,
		LabelSelector: k8Labels.SelectorFromSet(labels),
	}
	err = wait.Poll(5*time.Second, 5*time.Minute, func() (done bool, err error) {
		podList := &v1.PodList{}
		err = k8sClient.List(context.TODO(), podList, listOptions)
		if err != nil {
			return false, err
		}
		if len(podList.Items) == 0 {
			return true, nil
		}
		return false, nil
	})
	return err
}
