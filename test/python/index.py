from __future__ import print_function

import time
import openapi_client
from openapi_client.rest import ApiException
from pprint import pprint
import json
import urllib
def create_jsonlines(original):

    if isinstance(original, str):
        original = json.loads(original)

    return '\n'.join([json.dumps(original[outer_key], sort_keys=True) 
                      for outer_key in sorted(original.keys(),
                                              key=lambda x: int(x))])
# Defining the host is optional and defaults to https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "127.0.0.1:6368"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    index_instance = openapi_client.IndexApi(api_client)
    search_instance = openapi_client.SearchApi(api_client)
    api_instance = openapi_client.UtilsApi(api_client)
    search_request = openapi_client.SearchRequest(
    
        index = 'movies',
        query= {
            'match_phrase' : {
                'movie_title' : 'star trek nemesis',
            }
        },
        sort= [{'_score':'desc'},{'id':'asc'}],
        script_fields= {'myexpr':{'script':{'inline':'IF(price<10,1,0)'}}},
        highlight = {'fields':{'title','content'}},
        limit = 12,
        offset =100,
        source=['title','content','cat_id'],
        profile= True
    
   )
    try:
        sql = api_instance.sql('mode=raw&query=select * from movies');
        pprint(sql)
        exit()
        search_result = search_instance.search({
            'index':'movies',
            'query':{'bool':{'must':[{'query_string':'new movie'}]}},
            'script_fields':{'myexpr':{'script':{'inline':'IF(rating>8,1,0)'}}},
             'sort':[{'myexpr':'desc'},{'_score':'desc'}],
             'profile': True
            })
        pprint(search_result.hits)
        insert = index_instance.insert({
                'index': 'movies',
                'id':700,
                'doc' : {
                    'title':'This is an old movie',
                    'plot': 'A secret team goes to North Pole',
                    'year':1950,
                    'rating':9.5,
                    'lat': 60.4,
                    'lon': 51.99,
                    'advise': 'PG-13',
                    'meta': '{"keywords":{"travel","ice"},"genre":{"adventure"}}',
                    'language':[2,3]
                }
        })
        pprint(insert)        
        insert = index_instance.insert({
                'index': 'movies',
                'id':701,
                'doc' : {
                    'title':'This is an old movie',
                    'plot': 'A secret team goes to North Pole',
                    'year':1950,
                    'rating':9.5,
                    'lat': 60.4,
                    'lon': 51.99,
                    'advise': 'PG-13',
                    'meta': '{"keywords":{"travel","ice"},"genre":{"adventure"}}',
                    'language':[2,3]
                }
        })
        pprint(insert)                
        insert = index_instance.insert({
                'index': 'movies',
                'doc' : {
                    'title':'This is a new movie',
                    'plot': 'A secret team goes to North Pole',
                    'year':2020,
                    'rating':9.5,
                    'lat': 60.4,
                    'lon': 51.99,
                    'advise': 'PG-13',
                    'meta': '{"keywords":{"travel","ice"},"genre":{"adventure"}}',
                    'language':[2,3]
                }
        })
        pprint(insert)
        replace = index_instance.replace({
                'index': 'movies',
                'doc' : {
                    'title':'This is a new movie',
                    'plot': 'A secret team goes to North Pole',
                    'year':2020,
                    'rating':9.1,
                    'lat': 60.4,
                    'lon': 51.99,
                    'advise': 'PG-13',
                    'meta': '{"keywords":{"travel","ice"},"genre":{"adventure"}}',
                    'language':[2,3]
                }
        })
        pprint(replace)        
        update = index_instance.update({
                'index': 'movies',
                'doc' : {
                    'rating':9.49
                },
                'query': {'bool':{'must':[{'query_string':'new movie'}]}}
        })
        pprint(update)            
        update = index_instance.update({
                'index': 'movies',
                'doc' : {
                    'rating':9.49
                },
                'id':100
        })
        pprint(update)        
        delete = index_instance.delete({
                'index': 'movies',
                'query': {'bool':{'must':[{'query_string':'new movie'}]}}
        })
        pprint(delete)            
        delete = index_instance.delete({
                'index': 'movies',
                'id':100
        })
        pprint(delete)    

        bulk = index_instance.bulk(
        '{"insert": {"index": "movies", "doc": {"plot": "A secret team goes to North Pole", "rating": 9.5, "language": [2, 3], "title": "This is an older movie", "lon": 51.99, "meta": {"keywords":["travel","ice"],"genre":["adventure"]}, "year": 1950, "lat": 60.4, "advise": "PG-13"}}}'+
        "\n"+
        '{"delete": {"index": "movies","id":700}}' )
        pprint(bulk)        
       
    except ApiException as e:
        print("Exception : %s\n" % e)
