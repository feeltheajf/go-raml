# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

"""
Auto-generated class for Catanimal
"""
from six import string_types

from . import client_support
from .EnumCity import EnumCity


class Catanimal(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def _get_schema():
        return {
            "cities": {"type": [EnumCity], "required": True},
            "colours": {"type": [string_types], "required": True},
            "kind": {"type": string_types, "required": True},
            "name": {"type": string_types, "required": False},
        }

    @staticmethod
    def create(**kwargs):
        """
        :type cities: list[EnumCity]
        :type colours: list[string_types]
        :type kind: string_types
        :type name: string_types
        :rtype: Catanimal
        """

        return Catanimal(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError("No data or kwargs present")

        class_name = "Catanimal"
        data = json or kwargs

        # set attributes
        data_types = [EnumCity]
        self.cities = client_support.set_property(
            "cities", data, data_types, False, [], True, True, class_name
        )
        data_types = [string_types]
        self.colours = client_support.set_property(
            "colours", data, data_types, False, [], True, True, class_name
        )
        data_types = [string_types]
        self.kind = client_support.set_property(
            "kind", data, data_types, False, [], False, True, class_name
        )
        data_types = [string_types]
        self.name = client_support.set_property(
            "name", data, data_types, False, [], False, False, class_name
        )

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
