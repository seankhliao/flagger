// Code generated by solo-kit. DO NOT EDIT.

package v1

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

func NewUpstream(namespace, name string) *Upstream {
	upstream := &Upstream{}
	upstream.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return upstream
}

func (r *Upstream) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Upstream) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Upstream) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.UpstreamSpec,
		r.DiscoveryMetadata,
	)
}

type UpstreamList []*Upstream
type UpstreamsByNamespace map[string]UpstreamList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list UpstreamList) Find(namespace, name string) (*Upstream, error) {
	for _, upstream := range list {
		if upstream.GetMetadata().Name == name {
			if namespace == "" || upstream.GetMetadata().Namespace == namespace {
				return upstream, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find upstream %v.%v", namespace, name)
}

func (list UpstreamList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, upstream := range list {
		ress = append(ress, upstream)
	}
	return ress
}

func (list UpstreamList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, upstream := range list {
		ress = append(ress, upstream)
	}
	return ress
}

func (list UpstreamList) Names() []string {
	var names []string
	for _, upstream := range list {
		names = append(names, upstream.GetMetadata().Name)
	}
	return names
}

func (list UpstreamList) NamespacesDotNames() []string {
	var names []string
	for _, upstream := range list {
		names = append(names, upstream.GetMetadata().Namespace+"."+upstream.GetMetadata().Name)
	}
	return names
}

func (list UpstreamList) Sort() UpstreamList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list UpstreamList) Clone() UpstreamList {
	var upstreamList UpstreamList
	for _, upstream := range list {
		upstreamList = append(upstreamList, resources.Clone(upstream).(*Upstream))
	}
	return upstreamList
}

func (list UpstreamList) Each(f func(element *Upstream)) {
	for _, upstream := range list {
		f(upstream)
	}
}

func (list UpstreamList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Upstream) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace UpstreamsByNamespace) Add(upstream ...*Upstream) {
	for _, item := range upstream {
		byNamespace[item.GetMetadata().Namespace] = append(byNamespace[item.GetMetadata().Namespace], item)
	}
}

func (byNamespace UpstreamsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace UpstreamsByNamespace) List() UpstreamList {
	var list UpstreamList
	for _, upstreamList := range byNamespace {
		list = append(list, upstreamList...)
	}
	return list.Sort()
}

func (byNamespace UpstreamsByNamespace) Clone() UpstreamsByNamespace {
	cloned := make(UpstreamsByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &Upstream{}

// Kubernetes Adapter for Upstream

func (o *Upstream) GetObjectKind() schema.ObjectKind {
	t := UpstreamCrd.TypeMeta()
	return &t
}

func (o *Upstream) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Upstream)
}

var UpstreamCrd = crd.NewCrd("gloo.solo.io",
	"upstreams",
	"gloo.solo.io",
	"v1",
	"Upstream",
	"us",
	false,
	&Upstream{})
