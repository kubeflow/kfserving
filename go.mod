module github.com/kubeflow/kfserving

go 1.13

require (
	cloud.google.com/go v0.47.0
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	contrib.go.opencensus.io/exporter/stackdriver v0.12.9-0.20191108183826-59d068f8d8ff
	github.com/PuerkitoBio/purell v1.1.1
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
	github.com/aws/aws-sdk-go v1.28.0
	github.com/beorn7/perks v1.0.1
	github.com/census-instrumentation/opencensus-proto v0.2.1
	github.com/cloudevents/sdk-go v0.9.2
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/distribution v2.7.1+incompatible
	github.com/emicklei/go-restful v2.11.0+incompatible
	github.com/evanphx/json-patch v4.5.0+incompatible
	github.com/fatih/color v1.9.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/getkin/kin-openapi v0.2.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.1
	github.com/go-openapi/jsonpointer v0.19.3
	github.com/go-openapi/jsonreference v0.19.3
	github.com/go-openapi/spec v0.19.4
	github.com/go-openapi/swag v0.19.5
	github.com/gobuffalo/flect v0.2.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/groupcache v0.0.0-20191002201903-404acd9df4cc
	github.com/golang/protobuf v1.3.2
	github.com/google/go-cmp v0.3.1
	github.com/google/go-containerregistry v0.0.0-20190910142231-b02d448a3705
	github.com/google/gofuzz v1.0.0
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.3.1
	github.com/hashicorp/golang-lru v0.5.3
	github.com/hpcloud/tail v1.0.0
	github.com/imdario/mergo v0.3.8
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/json-iterator/go v1.1.8
	github.com/mailru/easyjson v0.7.0
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/mattn/go-colorable v0.1.4
	github.com/mattn/go-isatty v0.0.11
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/onsi/ginkgo v1.10.2
	github.com/onsi/gomega v1.7.0
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.0.0
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4
	github.com/prometheus/common v0.7.0
	github.com/prometheus/procfs v0.0.5
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/tensorflow/tensorflow v1.13.1
	go.opencensus.io v0.22.1
	go.uber.org/atomic v1.4.0
	go.uber.org/multierr v1.2.0
	go.uber.org/zap v1.11.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20191021144547-ec77196f6094
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20191023065245-6d3f0bb11be5
	golang.org/x/tools v0.0.0-20191022213345-0bbdf54effa2
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
	gomodules.xyz/jsonpatch v2.0.1+incompatible
	google.golang.org/api v0.15.0
	google.golang.org/appengine v1.6.5
	google.golang.org/genproto v0.0.0-20200108215221-bd8f9a0ef82f
	google.golang.org/grpc v1.26.0
	gopkg.in/inf.v0 v0.9.1
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
	gopkg.in/yaml.v2 v2.2.4
	gopkg.in/yaml.v3 v3.0.0-20191120175047-4206685974f2
	istio.io/api v0.0.0-20191115173247-e1a1952e5b81
	istio.io/client-go v0.0.0-20191120150049-26c62a04cdbc
	istio.io/gogo-genproto v0.0.0-20191029161641-f7d19ec0141d
	k8s.io/api v0.17.0
	k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
	k8s.io/apimachinery v0.17.0
	k8s.io/apiserver v0.0.0-20190918200908-1e17798da8c1
	k8s.io/client-go v0.17.0
	k8s.io/code-generator v0.15.8-beta.1
	k8s.io/component-base v0.17.0
	k8s.io/gengo v0.0.0-20191010091904-7fa3014cb28f
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
	k8s.io/utils v0.0.0-20191114184206-e782cd3c129f
	knative.dev/pkg v0.0.0-20191217184203-cf220a867b3d
	knative.dev/serving v0.11.0
	sigs.k8s.io/controller-runtime v0.3.0
	sigs.k8s.io/controller-tools v0.2.2
	sigs.k8s.io/testing_frameworks v0.1.2
	sigs.k8s.io/yaml v1.1.0
)

//replace gopkg.in/fsnotify.v1 v1.4.7 => github.com/fsnotify/fsnotify v1.4.7
