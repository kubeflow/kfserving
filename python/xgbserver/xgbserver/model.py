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

import kfserving
import xgboost as xgb
from xgboost import XGBModel
import os
from typing import Dict

BOOSTER_FILE = "model.bst"


class XGBoostModel(kfserving.KFModel):
    def __init__(self, name: str, model_dir: str, nthread: int, method: str,
                 booster: XGBModel = None):
        super().__init__(name)
        self.name = name
        self.model_dir = model_dir
        self.nthread = nthread
        self.method = method
        self._classifier = None
        if not booster is None:
            self._booster = booster
            self.ready = True

    def load(self) -> bool:
        model_file = os.path.join(
            kfserving.Storage.download(self.model_dir), BOOSTER_FILE)
        if self.method == "predict_proba":
            self._classifier = xgb.XGBClassifier()
            self._classifier.load_model(model_file)
        else:
            self._booster = xgb.Booster(params={"nthread": self.nthread},
                                        model_file=model_file)
        self.ready = True
        return self.ready

    def predict(self, request: Dict) -> Dict:
        try:
            # Use of list as input is deprecated see https://github.com/dmlc/xgboost/pull/3970
            instances = request["instances"]
            dmatrix = xgb.DMatrix(instances, nthread=self.nthread)
            if self.method == "predict":
                result: xgb.DMatrix = self._booster.predict(dmatrix)
            elif self.method == "predict_proba":
                result: xgb.DMatrix = self._classifier.predict_proba(instances)
            else:
                raise Exception("Not a valid prediction method: %s" % self.method)
            return {"predictions": result.tolist()}
        except Exception as e:
            raise Exception("Failed to predict %s" % e)
