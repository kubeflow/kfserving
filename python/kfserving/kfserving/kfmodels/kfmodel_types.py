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

import enum
import os
from typing import Optional, Tuple


class KFModelTypes(enum.Enum):
    Sklearn = "sklearn"
    Xgboost = "xgboost"
    Pytorch = "pytorch"


MODEL_EXTENSIONS = {
    KFModelTypes.Sklearn: [".joblib", ".pkl", ".pickle"],
    KFModelTypes.Xgboost: [".bst"],
    KFModelTypes.Pytorch: [".pt"]
}


def get_kfmodel_type(model_name: str, model_dir: str) -> Tuple[Optional[KFModelTypes], str]:
    for model_type in MODEL_EXTENSIONS:
        for extension in MODEL_EXTENSIONS[model_type]:
            path = os.path.join(model_dir, model_name + extension)
            if os.path.exists(path):
                return model_type, path
    return None, ""
