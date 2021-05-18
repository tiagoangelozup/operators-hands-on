package controllers

import (
	installv1alpha1 "github.com/example/hello-kubernetes-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewHelloKubernetesDeployment(hk *installv1alpha1.HelloKubernetes) *appsv1.Deployment {
	labels := map[string]string{"app": "hello-kubernetes"}
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      hk.Name,
			Namespace: hk.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "paulbouwer/hello-kubernetes:1.10",
						Name:  "hello-kubernetes",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8080,
							Name:          "http",
						}},
						Env: []corev1.EnvVar{
							{
								Name:  "CONTAINER_IMAGE",
								Value: "paulbouwer/hello-kubernetes:1.10",
							},
							{
								Name: "KUBERNETES_NAMESPACE",
								ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{
									FieldPath: "metadata.namespace",
								}},
							},
							{
								Name: "KUBERNETES_POD_NAME",
								ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{
									FieldPath: "metadata.name",
								}},
							},
							{
								Name: "KUBERNETES_NODE_NAME",
								ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{
									FieldPath: "spec.nodeName",
								}},
							},
						},
					}},
				},
			},
		},
	}
	return dep
}
