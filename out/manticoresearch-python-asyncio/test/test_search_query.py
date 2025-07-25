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

from manticoresearch.models.search_query import SearchQuery

class TestSearchQuery(unittest.TestCase):
    """SearchQuery unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SearchQuery:
        """Test SearchQuery
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SearchQuery`
        """
        model = SearchQuery()
        if include_optional:
            return SearchQuery(
                query_string = None,
                match = None,
                match_phrase = None,
                match_all = None,
                bool = manticoresearch.models.bool_filter.boolFilter(
                    must = [
                        null
                        ], 
                    must_not = [
                        null
                        ], 
                    should = [
                        null
                        ], ),
                equals = None,
                var_in = None,
                range = None,
                geo_distance = {
                    'key' : null
                    },
                highlight = None
            )
        else:
            return SearchQuery(
        )
        """

    def testSearchQuery(self):
        """Test SearchQuery"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
