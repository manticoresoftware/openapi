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

from manticoresearch.models.agg_composite_source import AggCompositeSource

class TestAggCompositeSource(unittest.TestCase):
    """AggCompositeSource unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AggCompositeSource:
        """Test AggCompositeSource
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AggCompositeSource`
        """
        model = AggCompositeSource()
        if include_optional:
            return AggCompositeSource(
                terms = manticoresearch.models.agg_composite_term.aggCompositeTerm(
                    field = 'field1', )
            )
        else:
            return AggCompositeSource(
                terms = manticoresearch.models.agg_composite_term.aggCompositeTerm(
                    field = 'field1', ),
        )
        """

    def testAggCompositeSource(self):
        """Test AggCompositeSource"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
