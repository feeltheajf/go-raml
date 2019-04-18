# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

from .City import City


class UsersService:

    _methods = (
        "users_userId_address_addressId_get",
        "users_userId_delete",
        "getuserid",
        "users_userId_post",
        "users_delete",
        "get_users",
        "option_users",
        "create_users",
    )

    def __init__(self, client):
        self.client = client

    def users_userId_address_addressId_get(
        self,
        addressId,
        userId,
        headers=None,
        query_params=None,
        content_type="application/json",
    ):
        """
        get address id
        of address
        It is method for GET /users/{userId}/address/{addressId}
        """
        if query_params is None:
            query_params = {}

        uri = (
            self.client.base_url + "/users/" + userId + "/address/" + addressId
        )
        return self.client.get(uri, None, headers, query_params, content_type)

    def users_userId_delete(
        self,
        userId,
        headers=None,
        query_params=None,
        content_type="application/json",
    ):
        """
        It is method for DELETE /users/{userId}
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/users/" + userId
        return self.client.delete(
            uri, None, headers, query_params, content_type
        )

    def getuserid(
        self,
        userId,
        headers=None,
        query_params=None,
        content_type="application/json",
    ):
        """
        get id
        It is method for GET /users/{userId}
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/users/" + userId
        return self.client.get(uri, None, headers, query_params, content_type)

    @property
    def _users_userId_post_data_types(self):
        """It is data schema for POST /users/{userId}"""
        return ()

    def users_userId_post(
        self,
        data,
        userId,
        headers=None,
        query_params=None,
        content_type="application/json",
    ):
        """
        post without request body
        It is method for POST /users/{userId}
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/users/" + userId
        return self.client.post(uri, data, headers, query_params, content_type)

    @property
    def _users_delete_data_types(self):
        """It is data schema for DELETE /users"""
        return (City,)

    def users_delete(
        self,
        data,
        headers=None,
        query_params=None,
        content_type="application/json",
    ):
        """
        delete with request body
        It is method for DELETE /users
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/users"
        return self.client.delete(
            uri, data, headers, query_params, content_type
        )

    @property
    def _get_users_data_types(self):
        """It is data schema for GET /users"""
        return ()

    def get_users(
        self,
        data,
        headers=None,
        query_params=None,
        content_type="application/json",
    ):
        """
        First line of comment.
        Second line of comment
        It is method for GET /users
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/users"
        return self.client.get(uri, data, headers, query_params, content_type)

    def option_users(
        self, headers=None, query_params=None, content_type="application/json"
    ):
        """
        It is method for OPTIONS /users
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/users"
        return self.client.session.options(
            uri, None, headers, query_params, content_type
        )

    @property
    def _create_users_data_types(self):
        """It is data schema for POST /users"""
        return (City,)

    def create_users(
        self,
        data,
        headers=None,
        query_params=None,
        content_type="application/json",
    ):
        """
        create users
        It is method for POST /users
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/users"
        return self.client.post(uri, data, headers, query_params, content_type)
