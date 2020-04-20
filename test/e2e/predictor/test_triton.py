# Copyright 2019 kubeflow.org.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import pytest
from kubernetes import client

from kfserving import KFServingClient
from kfserving import constants
from kfserving import V1alpha2EndpointSpec
from kfserving import V1alpha2PredictorSpec
from kfserving import V1alpha2TransformerSpec
from kfserving import V1alpha2TritonSpec
from kfserving import V1alpha2CustomSpec
from kfserving import V1alpha2InferenceServiceSpec
from kfserving import V1alpha2InferenceService
from kubernetes.client import V1ResourceRequirements
from kubernetes.client import V1Container
from kubernetes.client import V1EnvVar
from ..common.utils import predict
from ..common.utils import KFSERVING_TEST_NAMESPACE

api_version = constants.KFSERVING_GROUP + '/' + constants.KFSERVING_VERSION
KFServing = KFServingClient(config_file="~/.kube/config")


@pytest.mark.flaky(reruns=3)
def test_triton():
    service_name = 'isvc-triton'
    default_endpoint_spec = V1alpha2EndpointSpec(
        predictor=V1alpha2PredictorSpec(
            min_replicas=1,
            triton=V1alpha2TritonSpec(
                storage_uri='gs://kfserving-samples/models/triton/bert',
                resources=V1ResourceRequirements(
                    requests={'cpu': '1', 'memory': '16Gi', 'nvidia.com/gpu': '1'},
                    limits={'cpu': '1', 'memory': '16Gi', 'nvidia.com/gpu': '1'}))),
        transformer=V1alpha2TransformerSpec(
            min_replicas=1,
            custom=V1alpha2CustomSpec(
                container=V1Container(
                  image='gcr.io/kubeflow-ci/kfserving/bert-transformer:latest',
                  name='kfserving-container',
                  env=[
                      V1EnvVar(name="STORAGE_URI", value="gs://kfserving-samples/models/triton/bert-transformer")
                  ],
                  resources=V1ResourceRequirements(
                    requests={'cpu': '100m', 'memory': '1Gi'},
                    limits={'cpu': '100m', 'memory': '1Gi'})))))

    isvc = V1alpha2InferenceService(api_version=api_version,
                                    kind=constants.KFSERVING_KIND,
                                    metadata=client.V1ObjectMeta(
                                        name=service_name, namespace=KFSERVING_TEST_NAMESPACE,
                                        annotations={'serving.kubeflow.org/gke-accelerator': 'nvidia-tesla-k80'}),
                                    spec=V1alpha2InferenceServiceSpec(default=default_endpoint_spec))

    KFServing.create(isvc)
    try:
        KFServing.wait_isvc_ready(service_name, namespace=KFSERVING_TEST_NAMESPACE)
    except RuntimeError as e:
        print(KFServing.api_instance.get_namespaced_custom_object("serving.knative.dev", "v1", KFSERVING_TEST_NAMESPACE,
                                                                  "services", service_name + "-predictor-default"))
        print(KFServing.api_instance.get_namespaced_custom_object("serving.knative.dev", "v1", KFSERVING_TEST_NAMESPACE,
                                                                  "services", service_name + "-transformer-default"))
        deployments = KFServing.app_api.list_namespaced_deployment(KFSERVING_TEST_NAMESPACE,
                                                                   label_selector='serving.kubeflow.org/inferenceservice={}'.
                                                                   format(service_name))
        for deployment in deployments.items:
            print(deployment)
        raise e
    prediction = predict(service_name, './data/qa.json')
    assert(prediction["predictions"] == "John F. Kennedy")
    KFServing.delete(service_name, KFSERVING_TEST_NAMESPACE)
