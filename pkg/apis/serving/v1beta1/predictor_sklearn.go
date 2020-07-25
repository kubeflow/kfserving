package v1beta1

import (
	"fmt"
	"github.com/kubeflow/kfserving/pkg/constants"
	v1 "k8s.io/api/core/v1"
)

// SKLearnSpec defines arguments for configuring SKLearn model serving.
type SKLearnSpec struct {
	// Contains fields shared across all predictors
	PredictorExtensionSpec `json:",inline"`
}

// Validate returns an error if invalid
func (k *SKLearnSpec) Validate() error {
	return nil
}

// Default sets defaults on the resource
func (k *SKLearnSpec) Default(config *InferenceServicesConfig) {
	k.Container.Name = constants.InferenceServiceContainerName
	if k.RuntimeVersion == "" {
		k.RuntimeVersion = config.Predictors.SKlearn.DefaultGpuImageVersion
	}
	setResourceRequirementDefaults(&k.Resources)
}

// GetContainer transforms the resource into a container spec
func (k *SKLearnSpec) GetContainer(modelName string, config *InferenceServicesConfig) *v1.Container {
	arguments := []string{
		fmt.Sprintf("%s=%s", constants.ArgumentModelName, modelName),
		fmt.Sprintf("%s=%s", constants.ArgumentModelDir, constants.DefaultModelLocalMountPath),
		fmt.Sprintf("%s=%s", constants.ArgumentHttpPort, constants.InferenceServiceDefaultHttpPort),
	}
	/*if parallelism != 0 {
		arguments = append(arguments, fmt.Sprintf("%s=%s", constants.ArgumentWorkers, strconv.Itoa(parallelism)))
	}*/
	if k.Container.Image == "" {
		k.Container.Image = config.Predictors.SKlearn.ContainerImage + ":" + k.RuntimeVersion
	}
	k.Container.Name = constants.InferenceServiceContainerName
	k.Container.Args = arguments
	return &k.Container
}

func (k *SKLearnSpec) GetStorageUri() *string {
	return k.StorageURI
}
