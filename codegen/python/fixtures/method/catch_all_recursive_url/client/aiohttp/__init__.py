# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

from .File import File
from .http_client import HTTPClient
from .tree_service import TreeService


class Client:

    _base_uri = "http://localhost:5000"
    _services = ("tree",)

    def __init__(self, loop, base_uri=_base_uri):
        http_client = HTTPClient(loop, base_uri)

        self.tree = TreeService(http_client)
        self.close = http_client.close
