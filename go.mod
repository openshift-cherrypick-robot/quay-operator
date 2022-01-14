module github.com/quay/quay-operator

go 1.12

require (
	github.com/go-logr/logr v1.2.0
	github.com/go-redis/redis/v8 v8.11.4 // indirect
	github.com/kube-object-storage/lib-bucket-provisioner v0.0.0-20210311161930-4bea5edaff58
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.17.0
	github.com/openshift/api v3.9.0+incompatible
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.53.1
	github.com/quay/clair/config v1.1.0
	github.com/quay/config-tool v0.1.9
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/sjson v1.2.3
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/kustomize/api v0.10.1
	sigs.k8s.io/kustomize/kyaml v0.13.0
	sigs.k8s.io/yaml v1.3.0
)
