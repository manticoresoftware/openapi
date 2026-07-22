
# coding: utf-8
# Manticore Search Client
# Copyright (c) 2020-2021, Manticore Software LTD (https://manticoresearch.com)
#
# All rights reserved

from __future__ import absolute_import
from pprint import pprint
import unittest
import json
import manticoresearch
from manticoresearch.api.index_api import IndexApi  # noqa: E501
from manticoresearch.models import *
from manticoresearch.rest import ApiException
from unittest import IsolatedAsyncioTestCase
from urllib.parse import quote
from time import sleep
import sys

class TestManualApi(IsolatedAsyncioTestCase):

    def setUp(self):
        self.configuration = manticoresearch.Configuration(
            host = "http://localhost:9408"
        )
        
    def tearDown(self):

        pass

    async def utils_api_test(self, client=None):
        utilsApi = manticoresearch.UtilsApi(client)
        await utilsApi.sql('query=DROP TABLE IF EXISTS movies')
        res = await utilsApi.sql("CREATE TABLE IF NOT EXISTS movies (title text, plot text, _year integer, rating float, code multi) min_infix_len='2'")
        self.assertEqual(res.to_dict(), [{'total': 0, 'error': '', 'warning': ''}])
        print("Utils API tests successful")

    async def index_api_test(self, client=None):
        indexApi = manticoresearch.IndexApi(client)
        docs = [ \
            {"insert": {"table" : "movies", "id" : 1, "doc" : {"title" : "Star Trek 2: Nemesis", "plot": "The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.", "_year": 2002, "rating": 6.4, "code": [1,2,3]}}}, \
            {"insert": {"table" : "movies", "id" : 2, "doc" : {"title" : "Star Trek 1: Nemesis", "plot": "The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.", "_year": 2001, "rating": 6.5, "code": [1,12,3]}}},
            {"insert": {"table" : "movies", "id" : 3, "doc" : {"title" : "Star Trek 3: Nemesis", "plot": "The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.", "_year": 2003, "rating": 6.6, "code": [11,2,3]}}}, \
            {"insert": {"table" : "movies", "id" : 4, "doc" : {"title" : "Star Trek 4: Nemesis", "plot": "The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.", "_year": 2003, "rating": 6.5, "code": [1,2,4]}}},
        ]        
        res = await indexApi.bulk('\n'.join(map(json.dumps,docs)))
        res_json = res.to_dict()
        self.assertEqual(res_json['items'][0]['bulk']['created'], 4)
        print("Index API tests successful")

    async def search_api_test(self, client=None):
        searchApi = manticoresearch.SearchApi(client)
        search_request = SearchRequest(
            table="movies",
        )
        matchFilter = QueryFilter() 
        matchFilter.match = {"title":"4"}
        mustCond = [ matchFilter ]
        boolFilter = BoolFilter(must=mustCond)
        searchQuery = SearchQuery(bool={"must": [ {"match": {"title":"4"}}] })
        search_request.query = searchQuery

        searchApi = manticoresearch.SearchApi(client)
        res = await searchApi.search(search_request)
        self.assertEqual(res.hits.hits[0].id, 4)

        search_request = {"table":"movies","query":{"bool": {"must": [ {"match": {"title":"4"}}] }}}

        res = await searchApi.search(search_request)
        print(res.hits.hits[0].id)
        self.assertEqual(res.hits.hits[0].id, 4)

        autocomplete_request = {"table":"movies","query": "Romul","options": {"fuzziness": 0, "layouts": "us,uk"} }

        res = await searchApi.autocomplete(autocomplete_request)
        self.assertEqual(res[0]['total'], 2)
        self.assertEqual(res[0]['data'][0]['query'], 'romulan')
        self.assertEqual(res[0]['data'][1]['query'], 'romulus')
        
        print("Search API tests successful")

    async def test_manual(self):
        async with manticoresearch.ApiClient(self.configuration) as client:
            await self.utils_api_test(client)
            await self.index_api_test(client)
            await self.search_api_test(client)
            
        print("\nTests finished")
if __name__ == '__main__':
    unittest.main()
