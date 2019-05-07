import pytest
import kfserving


class DummyModel(kfserving.KFModel):
    def __init__(self, name):
        self.name = name
        self.ready = False

    def load(self):
        self.ready = True

    def preprocess(self, inputs):
        return inputs

    def postprocess(self, outputs):
        return outputs

    def predict(self, inputs):
        return inputs


class TestTFHttpServer(object):

    @pytest.fixture(scope="class")
    def app(self):
        import kfserving
        model = DummyModel("TestModel")
        model.load()
        server = kfserving.KFServer()
        server.register_model(model)
        return server.createApplication()

    async def test_liveness(self,http_server_client):
        resp = await http_server_client.fetch('/')
        assert resp.code == 200

    async def test_protocol(self, http_server_client):
        resp = await http_server_client.fetch('/protocol')
        assert resp.code == 200
        assert resp.body == b"tensorflow.http"

    async def test_model(self, http_server_client):
        resp = await http_server_client.fetch('/models/TestModel')
        assert resp.code == 200

    async def test_predict(selfself, http_server_client):
        resp = await http_server_client.fetch('/models/TestModel:predict',method="POST",body=b'{"instances":[[1,2]]}')
        assert resp.code == 200
        assert resp.body == b"{'predictions': [[1, 2]]}"

class TestSeldonHttpServer(object):

    @pytest.fixture(scope="class")
    def app(self):
        import kfserving
        model = DummyModel("TestModelSeldon")
        model.load()
        server = kfserving.KFServer(protocol=kfserving.server.SELDON_HTTP_PROTOCOL)
        server.register_model(model)
        return server.createApplication()

    async def test_liveness(self,http_server_client):
        resp = await http_server_client.fetch('/')
        assert resp.code == 200

    async def test_protocol(self, http_server_client):
        resp = await http_server_client.fetch('/protocol')
        assert resp.code == 200
        assert resp.body == b"seldon.http"

    async def test_model(self, http_server_client):
        resp = await http_server_client.fetch('/models/TestModelSeldon:predict',method="POST",body=b'{"data":{"ndarray":[[1,2]]}}')
        assert resp.code == 200




