// Copyright Contributors to the Open Cluster Management project

/*
Copyright 2021.

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

package mcewebhook

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	backplanev1 "github.com/stolostron/backplane-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	// "k8s.io/apimachinery/pkg/runtime/schema"
	// ctrl "sigs.k8s.io/controller-runtime"
	"github.com/stolostron/backplane-operator/pkg/utils"
	// "time"
	//+kubebuilder:scaffold:imports
)

var _ = Describe("BackplaneConfig controller", func() {
	Context("Legacy clean up tasks", func() {
		It("Removes the legacy CLC Prometheus configuration", func() {
			By("creating the backplane config with nonexistant secret")
			createCtx := context.Background()
			// timeout := time.Second * 60
			// interval := time.Millisecond * 250
			utils.SetDeployOnOCP(false)
			// Create target namespace
			err := k8sClient.Create(context.Background(), &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
				Spec: corev1.NamespaceSpec{},
			})
			if err != nil && !errors.IsAlreadyExists(err) {
				Expect(err).To(BeNil())
			}
			backplaneConfig := &backplanev1.MultiClusterEngine{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "multicluster.openshift.io/v1",
					Kind:       "MultiClusterEngine",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
				Spec: backplanev1.MultiClusterEngineSpec{
					TargetNamespace: "test",
					ImagePullSecret: "testsecret",
				},
			}

			Expect(k8sClient.Create(createCtx, backplaneConfig)).Should(Succeed())

			testWH := backplanev1.ValidatingWebhook("test")
			Expect(k8sClient.Create(createCtx, testWH)).Should(Succeed())
			// Eventually(func() error {
			// 	ctx := context.Background()
			// 	u := &unstructured.Unstructured{}
			// 	u.SetName("multiclusterengines.multicluster.openshift.io")
			// 	u.SetNamespace("test")
			// 	u.SetGroupVersionKind(schema.GroupVersionKind{
			// 		Group:   "admissionregistration.k8s.io",
			// 		Kind:    "ValidatingWebhookConfiguration",
			// 		Version: "v1",
			// 	})
			// 	return k8sClient.Delete(ctx, u)
			// }, timeout, interval).Should(Succeed())

			// _, err := webhookReconciler.Reconcile(createCtx, ctrl.Result{})
			// Expect(err).ToBe(nil)

		})
	})
})
