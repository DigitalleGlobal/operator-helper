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

package operator

import (
	"fmt"
	"github.com/monimesl/operator-helper/reconciler"
	"github.com/monimesl/operator-helper/webhook"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"log"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Start configures
func Start(mgr manager.Manager) error {
	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		return err
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		return err
	}
	return mgr.Start(ctrl.SetupSignalHandler())
}

// BootOrDie configures...
func BootOrDie(config *rest.Config, options ctrl.Options, getReconcilers func() []reconciler.Reconciler, getRuntimeObjs func() []runtime.Object) {
	if err := Boot(config, options, getReconcilers, getRuntimeObjs); err != nil {
		log.Fatal(err)
	}
}

// Boot configures...
func Boot(config *rest.Config, options ctrl.Options, getReconcilers func() []reconciler.Reconciler, getRuntimeObjs func() []runtime.Object) error {
	mgr, err := manager.New(config, options)
	if err != nil {
		return fmt.Errorf("manager create error: %w", err)
	}
	if getRuntimeObjs != nil {
		if err = webhook.Configure(mgr, getRuntimeObjs()...); err != nil {
			return fmt.Errorf("webhook config error: %w", err)
		}
	}
	if getReconcilers != nil {
		if err = reconciler.Configure(mgr, getReconcilers()...); err != nil {
			return fmt.Errorf("reconciler config error: %w", err)
		}
	}
	if err = Start(mgr); err != nil {
		return fmt.Errorf("operator start error: %w", err)
	}
	return nil
}
