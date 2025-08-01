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

from manticoresearch.models.search_response_hits import SearchResponseHits

class TestSearchResponseHits(unittest.TestCase):
    """SearchResponseHits unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SearchResponseHits:
        """Test SearchResponseHits
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SearchResponseHits`
        """
        model = SearchResponseHits()
        if include_optional:
            return SearchResponseHits(
                max_score = 56,
                total = 56,
                total_relation = '',
                hits = [
                    manticoresearch.models.hits_hits.hitsHits(
                        _id = 56, 
                        _score = 56, 
                        _source = manticoresearch.models._source._source(), 
                        _knn_dist = 1.337, 
                        highlight = manticoresearch.models.highlight.highlight(), 
                        table = '', 
                        _type: = '', 
                        fields = manticoresearch.models.fields.fields(), )
                    ]
            )
        else:
            return SearchResponseHits(
        )
        """

    def testSearchResponseHits(self):
        """Test SearchResponseHits"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
