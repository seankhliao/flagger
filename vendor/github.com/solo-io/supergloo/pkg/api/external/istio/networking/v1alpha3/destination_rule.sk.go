// Code generated by solo-kit. DO NOT EDIT.

package v1alpha3

import (
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewDestinationRule(namespace, name string) *DestinationRule {
	destinationrule := &DestinationRule{}
	destinationrule.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return destinationrule
}

func (r *DestinationRule) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *DestinationRule) SetStatus(status core.Status) {
	r.Status = status
}

func (r *DestinationRule) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Host,
		r.TrafficPolicy,
		r.Subsets,
		r.ConfigScope,
	)
}

type DestinationRuleList []*DestinationRule
type DestinationrulesByNamespace map[string]DestinationRuleList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list DestinationRuleList) Find(namespace, name string) (*DestinationRule, error) {
	for _, destinationRule := range list {
		if destinationRule.GetMetadata().Name == name {
			if namespace == "" || destinationRule.GetMetadata().Namespace == namespace {
				return destinationRule, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find destinationRule %v.%v", namespace, name)
}

func (list DestinationRuleList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, destinationRule := range list {
		ress = append(ress, destinationRule)
	}
	return ress
}

func (list DestinationRuleList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, destinationRule := range list {
		ress = append(ress, destinationRule)
	}
	return ress
}

func (list DestinationRuleList) Names() []string {
	var names []string
	for _, destinationRule := range list {
		names = append(names, destinationRule.GetMetadata().Name)
	}
	return names
}

func (list DestinationRuleList) NamespacesDotNames() []string {
	var names []string
	for _, destinationRule := range list {
		names = append(names, destinationRule.GetMetadata().Namespace+"."+destinationRule.GetMetadata().Name)
	}
	return names
}

func (list DestinationRuleList) Sort() DestinationRuleList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list DestinationRuleList) Clone() DestinationRuleList {
	var destinationRuleList DestinationRuleList
	for _, destinationRule := range list {
		destinationRuleList = append(destinationRuleList, resources.Clone(destinationRule).(*DestinationRule))
	}
	return destinationRuleList
}

func (list DestinationRuleList) Each(f func(element *DestinationRule)) {
	for _, destinationRule := range list {
		f(destinationRule)
	}
}

func (list DestinationRuleList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *DestinationRule) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace DestinationrulesByNamespace) Add(destinationRule ...*DestinationRule) {
	for _, item := range destinationRule {
		byNamespace[item.GetMetadata().Namespace] = append(byNamespace[item.GetMetadata().Namespace], item)
	}
}

func (byNamespace DestinationrulesByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace DestinationrulesByNamespace) List() DestinationRuleList {
	var list DestinationRuleList
	for _, destinationRuleList := range byNamespace {
		list = append(list, destinationRuleList...)
	}
	return list.Sort()
}

func (byNamespace DestinationrulesByNamespace) Clone() DestinationrulesByNamespace {
	cloned := make(DestinationrulesByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &DestinationRule{}

// Kubernetes Adapter for DestinationRule

func (o *DestinationRule) GetObjectKind() schema.ObjectKind {
	t := DestinationRuleCrd.TypeMeta()
	return &t
}

func (o *DestinationRule) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*DestinationRule)
}

var DestinationRuleCrd = crd.NewCrd("networking.istio.io",
	"destinationrules",
	"networking.istio.io",
	"v1alpha3",
	"DestinationRule",
	"destinationrule",
	false,
	&DestinationRule{})
