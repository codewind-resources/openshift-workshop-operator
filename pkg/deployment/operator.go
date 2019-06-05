package deployment

import (
	cloudnativev1alpha1 "github.com/redhat/cloud-native-workshop-operator/pkg/apis/cloudnative/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewAnsibleOperatorDeployment(cr *cloudnativev1alpha1.Workshop, name string, namespace string, image string, serviceAccountName string) *appsv1.Deployment {
	labels := GetLabels(cr, name)

	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{MatchLabels: labels},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: serviceAccountName,
					Containers: []corev1.Container{
						{
							Name:            "ansible",
							Image:           image,
							ImagePullPolicy: corev1.PullAlways,
							Command: []string{
								"/usr/local/bin/ao-logs",
								"/tmp/ansible-operator/runner",
								"stdout",
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "runner",
									MountPath: "/tmp/ansible-operator/runner",
									ReadOnly:  true,
								},
							},
						},
						{
							Name:            "operator",
							Image:           image,
							ImagePullPolicy: corev1.PullAlways,
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "runner",
									MountPath: "/tmp/ansible-operator/runner",
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "WATCH_NAMESPACE",
									Value: "",
								},
								{
									Name:  "OPERATOR_NAME",
									Value: name,
								},
								{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "runner",
						},
					},
				},
			},
		},
	}
}

func NewOperatorDeployment(cr *cloudnativev1alpha1.Workshop, name string, namespace string, image string, serviceAccountName string, commands []string) *appsv1.Deployment {
	labels := GetLabels(cr, name)

	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{MatchLabels: labels},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: serviceAccountName,
					Containers: []corev1.Container{
						{
							Name:            "operator",
							Image:           image,
							ImagePullPolicy: corev1.PullAlways,
							Ports: []corev1.ContainerPort{
								{
									Name:          "metrics",
									ContainerPort: 60000,
									Protocol:      "TCP",
								},
							},
							Command: commands,
							Env: []corev1.EnvVar{
								{
									Name:  "WATCH_NAMESPACE",
									Value: "",
								},
								{
									Name:  "OPERATOR_NAME",
									Value: name,
								},
								{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
