# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

from .Address import Address
from .City import City
from .GetUsersReqBody import GetUsersReqBody
from .http_client import HTTPClient
from .users_service import UsersService


class Client:

    _services = ("users",)

    def __init__(self, base_uri="http://api.jumpscale.com/v3"):
        http_client = HTTPClient(base_uri)
        self.users = UsersService(http_client)
        self.close = http_client.close
