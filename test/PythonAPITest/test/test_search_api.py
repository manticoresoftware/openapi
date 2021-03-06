# coding: utf-8

"""
    Manticore Search API

    This is the API for Manticore Search HTTP protocol   # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: info@manticoresearch.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import
from pprint import pprint
import unittest
import openapi_client
from openapi_client.api.search_api import SearchApi  # noqa: E501
from openapi_client.rest import ApiException
from parametrized_test_case import ParametrizedTestCase

class TestSearchApi(ParametrizedTestCase):
    """SearchApi unit test stubs"""

    def setUp(self):
        client = openapi_client.ApiClient(self.settings['configuration'])
        self.api = openapi_client.SearchApi(client)  # noqa: E501
        self.index_api = openapi_client.IndexApi(client)  # noqa: E501

    def tearDown(self):
        pass

    def test_percolate(self):
        """Test case for percolate

        Perform reverse search on a percolate index  # noqa: E501
        """
        
        req_body = {"query":{"percolate":{"document":{"content":"sample content"}}}} 
        api_resp = self.api.percolate("test_pq", req_body)
        res = {'hits': api_resp.hits.hits, 'total': api_resp.hits.total, 'profile': api_resp.profile, 'timed_out': api_resp.timed_out}
        expected_res = {'hits': [{'_id': '1', '_index': 'test_pq', '_score': '1', '_source': {'query': {'ql': '@content sample content'}}, '_type': 'doc', 'fields': {'_percolator_document_slot': [1]}}], 'total': 1, 'profile': None, 'timed_out': False}
        self.assertDictEqual(res, expected_res)
         
        req_body = {"query":{"percolate":{"document":{"content":"no match"}}}} 
        api_resp = self.api.percolate("test_pq", req_body)
        res = {'hits': api_resp.hits.hits, 'total': api_resp.hits.total, 'profile': api_resp.profile, 'timed_out': api_resp.timed_out}
        expected_res = {'hits': [], 'total': 0, 'profile': None, 'timed_out': False}
        self.assertDictEqual(res, expected_res)
        pass

    def test_search(self):
        """Test case for search

        Performs a search  # noqa: E501
        """
        req_body = {"index":"test","id":1} 
        api_resp = self.index_api.delete(req_body)
        insert_req_body = {"index":"test","id":1,"doc":{"content":"sample content","name":"test doc","cat":"10"}} 
        self.index_api.insert(insert_req_body)
        req_body = {"index":"test","query":{"match":{"content":"sample"}}} 
        api_resp = self.api.search(req_body)
        res = {'id': api_resp.hits.hits[0]['_id'], 'total': api_resp.hits.total, 'profile': api_resp.profile, 'timed_out': api_resp.timed_out}
        expected_res = {'id': '1', 'total': 1, 'profile': None, 'timed_out': False}
        self.assertDictEqual(res, expected_res)
        
        req_body = {"index":"test","query":{"match":{"content":"no match"}}} 
        api_resp = self.api.search(req_body)
        res = {'hits': api_resp.hits.hits, 'total': api_resp.hits.total, 'profile': api_resp.profile, 'timed_out': api_resp.timed_out}
        expected_res = {'hits': [], 'total': 0, 'profile': None, 'timed_out': False}
        self.assertDictEqual(res, expected_res)


if __name__ == '__main__':
    unittest.main()
