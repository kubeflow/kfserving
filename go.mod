module github.com/kubeflow/kfserving

go 1.13

require (
	cloud.google.com/go/storage v1.11.0
	github.com/aws/aws-sdk-go v1.37.1
	github.com/cloudevents/sdk-go v1.2.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/getkin/kin-openapi v0.2.0
	github.com/go-logr/logr v0.4.0
	github.com/go-openapi/spec v0.20.2
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6
	github.com/google/uuid v1.2.0
	github.com/googleapis/google-cloud-go-testing v0.0.0-20200911160855-bcd43fbb19e8
	github.com/json-iterator/go v1.1.11
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.17.0
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	google.golang.org/api v0.36.0
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	istio.io/api v0.0.0-20200715212100-dbf5277541ef
	istio.io/client-go v0.0.0-20201005161859-d8818315d678
	istio.io/gogo-genproto v0.0.0-20191029161641-f7d19ec0141d // indirect
	k8s.io/api v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
	knative.dev/networking v0.0.0-20210615114921-e291c8011a20
	knative.dev/pkg v0.0.0-20210616195222-841aa7369ca1
	knative.dev/serving v0.23.0
	sigs.k8s.io/controller-runtime v0.8.0
)
