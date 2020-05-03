from alibiexplainer.anchor_tabular import AnchorTabular
import kfserving
import os
import dill
from sklearnserver.model import SKLearnModel
from alibi.datasets import fetch_adult
from alibi.api.interfaces import Explanation
import numpy as np
import json
from .utils import Predictor

ADULT_EXPLAINER_URI = "gs://seldon-models/sklearn/income/alibi/0.4.0"
ADULT_MODEL_URI = "gs://seldon-models/sklearn/income/model"
EXPLAINER_FILENAME = "explainer.dill"

def test_anchor_tabular():
    os.environ.clear()
    alibi_model = os.path.join(kfserving.Storage.download(ADULT_EXPLAINER_URI), EXPLAINER_FILENAME)
    with open(alibi_model, 'rb') as f:
        skmodel = SKLearnModel("adult", ADULT_MODEL_URI)
        skmodel.load()
        predictor = Predictor(skmodel)
        alibi_model = dill.load(f)
        anchor_tabular = AnchorTabular(predictor.predict_fn,alibi_model)
        adult = fetch_adult()
        X_test = adult.data[30001:, :]
        np.random.seed(0)
        explanation : Explanation = anchor_tabular.explain(X_test[0:1].tolist())
        exp_json = json.loads(explanation.to_json())
        assert exp_json["data"]["anchor"][0] == 'Age <= 28.00'
        assert exp_json["data"]["anchor"][1] == 'Marital Status = Never-Married'