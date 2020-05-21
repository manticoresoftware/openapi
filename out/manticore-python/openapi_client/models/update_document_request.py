# coding: utf-8

"""
    Manticore Search API

    This is the API for Manticore Search HTTP protocol   # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: adrian.nuta@manticoresearch.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from openapi_client.configuration import Configuration


class UpdateDocumentRequest(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'index': 'str',
        'doc': 'object',
        'id': 'int',
        'query': 'object'
    }

    attribute_map = {
        'index': 'index',
        'doc': 'doc',
        'id': 'id',
        'query': 'query'
    }

    def __init__(self, index=None, doc=None, id=None, query=None, local_vars_configuration=None):  # noqa: E501
        """UpdateDocumentRequest - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._index = None
        self._doc = None
        self._id = None
        self._query = None
        self.discriminator = None

        self.index = index
        self.doc = doc
        if id is not None:
            self.id = id
        if query is not None:
            self.query = query

    @property
    def index(self):
        """Gets the index of this UpdateDocumentRequest.  # noqa: E501


        :return: The index of this UpdateDocumentRequest.  # noqa: E501
        :rtype: str
        """
        return self._index

    @index.setter
    def index(self, index):
        """Sets the index of this UpdateDocumentRequest.


        :param index: The index of this UpdateDocumentRequest.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and index is None:  # noqa: E501
            raise ValueError("Invalid value for `index`, must not be `None`")  # noqa: E501

        self._index = index

    @property
    def doc(self):
        """Gets the doc of this UpdateDocumentRequest.  # noqa: E501


        :return: The doc of this UpdateDocumentRequest.  # noqa: E501
        :rtype: object
        """
        return self._doc

    @doc.setter
    def doc(self, doc):
        """Sets the doc of this UpdateDocumentRequest.


        :param doc: The doc of this UpdateDocumentRequest.  # noqa: E501
        :type: object
        """
        if self.local_vars_configuration.client_side_validation and doc is None:  # noqa: E501
            raise ValueError("Invalid value for `doc`, must not be `None`")  # noqa: E501

        self._doc = doc

    @property
    def id(self):
        """Gets the id of this UpdateDocumentRequest.  # noqa: E501


        :return: The id of this UpdateDocumentRequest.  # noqa: E501
        :rtype: int
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this UpdateDocumentRequest.


        :param id: The id of this UpdateDocumentRequest.  # noqa: E501
        :type: int
        """
        if (self.local_vars_configuration.client_side_validation and
                id is not None and id < 1):  # noqa: E501
            raise ValueError("Invalid value for `id`, must be a value greater than or equal to `1`")  # noqa: E501

        self._id = id

    @property
    def query(self):
        """Gets the query of this UpdateDocumentRequest.  # noqa: E501


        :return: The query of this UpdateDocumentRequest.  # noqa: E501
        :rtype: object
        """
        return self._query

    @query.setter
    def query(self, query):
        """Sets the query of this UpdateDocumentRequest.


        :param query: The query of this UpdateDocumentRequest.  # noqa: E501
        :type: object
        """

        self._query = query

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, UpdateDocumentRequest):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, UpdateDocumentRequest):
            return True

        return self.to_dict() != other.to_dict()