package deployment

import (
	openshiftv1alpha1 "github.com/redhat/openshift-workshop-operator/pkg/apis/openshift/v1alpha1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCustomResourceDefinition(cr *openshiftv1alpha1.Workshop, name string, group string, kind string, listKind string, plural string, singular string, version string, shortNames []string, additionalPrinterColumns []apiextensionsv1beta1.CustomResourceColumnDefinition) *apiextensionsv1beta1.CustomResourceDefinition {
	labels := GetLabels(cr, name)
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CustomResourceDefinition",
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Version: version,
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:       kind,
				ListKind:   listKind,
				Plural:     plural,
				Singular:   singular,
				ShortNames: shortNames,
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
			AdditionalPrinterColumns: additionalPrinterColumns,
		},
	}
}
