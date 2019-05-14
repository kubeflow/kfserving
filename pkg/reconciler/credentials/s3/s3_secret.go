/*
Copyright 2019 kubeflow.org.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package s3

import (
	"github.com/kubeflow/kfserving/pkg/constants"
	"k8s.io/api/core/v1"
)

const (
	AWSAccessKeyId         = "AWS_ACCESS_KEY_ID"
	AWSSecretAccessKey     = "AWS_SECRET_ACCESS_KEY"
	AWSAccessKeyIdName     = "awsAccessKeyID"
	AWSSecretAccessKeyName = "awsSecretAccessKey"
	AWSEndpointUrl         = "AWS_ENDPOINT_URL"
	AWSRegion              = "AWS_REGION"
	S3Endpoint             = "S3_ENDPOINT"
	S3UseHttps             = "S3_USE_HTTPS"
	S3VerifySSL            = "S3_VERIFY_SSL"
)

func BuildSecretEnvs(secret *v1.Secret) []v1.EnvVar {
	envs := []v1.EnvVar{
		{
			Name: AWSAccessKeyId,
			ValueFrom: &v1.EnvVarSource{
				SecretKeyRef: &v1.SecretKeySelector{
					LocalObjectReference: v1.LocalObjectReference{
						Name: secret.Name,
					},
					Key: AWSAccessKeyIdName,
				},
			},
		},
		{
			Name: AWSSecretAccessKey,
			ValueFrom: &v1.EnvVarSource{
				SecretKeyRef: &v1.SecretKeySelector{
					LocalObjectReference: v1.LocalObjectReference{
						Name: secret.Name,
					},
					Key: AWSSecretAccessKeyName,
				},
			},
		},
	}

	if s3Endpoint, ok := secret.Annotations[constants.KFServiceS3SecretEndpointAnnotation]; ok {
		s3EndpointUrl := "https://" + s3Endpoint
		if s3UseHttps, ok := secret.Annotations[constants.KFServiceS3SecretHttpsAnnotation]; ok {
			if s3UseHttps == "0" {
				s3EndpointUrl = "http://" + secret.Annotations[constants.KFServiceS3SecretEndpointAnnotation]
			}
			envs = append(envs, v1.EnvVar{
				Name:  S3UseHttps,
				Value: s3UseHttps,
			})
		}
		envs = append(envs, v1.EnvVar{
			Name:  S3Endpoint,
			Value: s3Endpoint,
		})
		envs = append(envs, v1.EnvVar{
			Name:  AWSEndpointUrl,
			Value: s3EndpointUrl,
		})
	}

	if s3Region, ok := secret.Annotations[constants.KFServiceS3SecretRegionAnnotation]; ok {
		envs = append(envs, v1.EnvVar{
			Name:  AWSRegion,
			Value: s3Region,
		})
	}

	if val, ok := secret.Annotations[constants.KFServiceS3SecretSSLAnnotation]; ok {
		envs = append(envs, v1.EnvVar{
			Name:  S3VerifySSL,
			Value: val,
		})
	}
	return envs
}
