# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
from gevent import monkey

from .Address import Address
from .City import City
from .GetUsersReqBody import GetUsersReqBody
from .http_client import HTTPClient
from .users_service import UsersService

monkey.patch_all()


BASE_URI = "http://api.jumpscale.com/v3"


class Client:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def _get_services():
        return ("users",)

    def __init__(self, base_uri=BASE_URI):
        http_client = HTTPClient(base_uri)
        self.users = UsersService(http_client)
        self.close = http_client.close
