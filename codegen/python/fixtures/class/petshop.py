# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

"""
Auto-generated class for petshop
"""
from six import string_types

from . import client_support
from .Cat import Cat


class petshop(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def _get_schema():
        return {
            "cats": {"type": [Cat], "required": True},
            "name": {"type": string_types, "required": True},
        }

    @staticmethod
    def create(**kwargs):
        """
        :type cats: list[Cat]
        :type name: string_types
        :rtype: petshop
        """

        return petshop(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError("No data or kwargs present")

        class_name = "petshop"
        data = json or kwargs

        # set attributes
        data_types = [Cat]
        self.cats = client_support.set_property(
            "cats", data, data_types, False, [], True, True, class_name
        )
        data_types = [string_types]
        self.name = client_support.set_property(
            "name", data, data_types, False, [], False, True, class_name
        )

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
