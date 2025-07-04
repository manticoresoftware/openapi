# coding: utf-8

"""
    Manticore Search Client

    Сlient for Manticore Search. 

    The version of the OpenAPI document: 5.0.0
    Contact: info@manticoresearch.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from manticoresearch.models.bool_filter import BoolFilter

class TestBoolFilter(unittest.TestCase):
    """BoolFilter unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> BoolFilter:
        """Test BoolFilter
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `BoolFilter`
        """
        model = BoolFilter()
        if include_optional:
            return BoolFilter(
                must = [
                    null
                    ],
                must_not = [
                    null
                    ],
                should = [
                    null
                    ]
            )
        else:
            return BoolFilter(
        )
        """

    def testBoolFilter(self):
        """Test BoolFilter"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
