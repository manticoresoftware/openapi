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

from manticoresearch.models.agg_terms import AggTerms

class TestAggTerms(unittest.TestCase):
    """AggTerms unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AggTerms:
        """Test AggTerms
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AggTerms`
        """
        model = AggTerms()
        if include_optional:
            return AggTerms(
                field = 'field1',
                size = 1000
            )
        else:
            return AggTerms(
                field = 'field1',
        )
        """

    def testAggTerms(self):
        """Test AggTerms"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
