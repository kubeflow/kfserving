# Copyright 2020 kubeflow.org.
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

# coding: utf-8

"""
    KFServing

    Python SDK for KFServing  # noqa: E501

    The version of the OpenAPI document: v0.1
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kfserving.configuration import Configuration


class V1beta1Batcher(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'max_batch_size': 'int',
        'max_latency': 'int',
        'timeout': 'int'
    }

    attribute_map = {
        'max_batch_size': 'maxBatchSize',
        'max_latency': 'maxLatency',
        'timeout': 'timeout'
    }

    def __init__(self, max_batch_size=None, max_latency=None, timeout=None, local_vars_configuration=None):  # noqa: E501
        """V1beta1Batcher - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._max_batch_size = None
        self._max_latency = None
        self._timeout = None
        self.discriminator = None

        if max_batch_size is not None:
            self.max_batch_size = max_batch_size
        if max_latency is not None:
            self.max_latency = max_latency
        if timeout is not None:
            self.timeout = timeout

    @property
    def max_batch_size(self):
        """Gets the max_batch_size of this V1beta1Batcher.  # noqa: E501

        Specifies the max number of requests to trigger a batch  # noqa: E501

        :return: The max_batch_size of this V1beta1Batcher.  # noqa: E501
        :rtype: int
        """
        return self._max_batch_size

    @max_batch_size.setter
    def max_batch_size(self, max_batch_size):
        """Sets the max_batch_size of this V1beta1Batcher.

        Specifies the max number of requests to trigger a batch  # noqa: E501

        :param max_batch_size: The max_batch_size of this V1beta1Batcher.  # noqa: E501
        :type: int
        """

        self._max_batch_size = max_batch_size

    @property
    def max_latency(self):
        """Gets the max_latency of this V1beta1Batcher.  # noqa: E501

        Specifies the max latency to trigger a batch  # noqa: E501

        :return: The max_latency of this V1beta1Batcher.  # noqa: E501
        :rtype: int
        """
        return self._max_latency

    @max_latency.setter
    def max_latency(self, max_latency):
        """Sets the max_latency of this V1beta1Batcher.

        Specifies the max latency to trigger a batch  # noqa: E501

        :param max_latency: The max_latency of this V1beta1Batcher.  # noqa: E501
        :type: int
        """

        self._max_latency = max_latency

    @property
    def timeout(self):
        """Gets the timeout of this V1beta1Batcher.  # noqa: E501

        Specifies the timeout of a batch  # noqa: E501

        :return: The timeout of this V1beta1Batcher.  # noqa: E501
        :rtype: int
        """
        return self._timeout

    @timeout.setter
    def timeout(self, timeout):
        """Sets the timeout of this V1beta1Batcher.

        Specifies the timeout of a batch  # noqa: E501

        :param timeout: The timeout of this V1beta1Batcher.  # noqa: E501
        :type: int
        """

        self._timeout = timeout

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1beta1Batcher):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1beta1Batcher):
            return True

        return self.to_dict() != other.to_dict()
