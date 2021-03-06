var ManticoreSearchApi = require('manticore_search_api');
 
var client= new ManticoreSearchApi.ApiClient()
client.basePath="http://127.0.0.1:6368"
var api = new ManticoreSearchApi.IndexApi(client) 
var utilsapi = new ManticoreSearchApi.UtilsApi(client)
var searchapi = new ManticoreSearchApi.SearchApi(client)


/* test insert */

api.insert( {
                'index': 'movies',

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
        }, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('Inserted: ' );
	console.log(data);
  }
});

/* test replace */

api.replace({
                'index': 'movies',
				'id':1400,
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
        }, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('Replaced: ' );
	console.log(data);
  }
});
		
/* test delete */
		
api.callDelete({
                'index': 'movies',
                'id':1400
        }, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('DELETE: ' );
	console.log(data);
  }
});
		
/* test bulk */
let bulk = [{"insert": {"index": "movies", "doc": {"plot": "A secret team goes to North Pole", "rating": 9.5, "language": [2, 3], "title": "This is an older movie", "lon": 51.99, "meta": {"keywords":["travel","ice"],"genre":["adventure"]}, "year": 1950, "lat": 60.4, "advise": "PG-13"}}},{"delete": {"index": "movies","id":1400}}];
let bulk_nd = "";
for (const item of bulk) {
		bulk_nd += JSON.stringify(item) + "\n";
}	
api.bulk(bulk_nd, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('BULK: ' );
	console.log(data);
  }
});	

/* test sql -  normal select */

utilsapi.sql('query=select * from movies',  (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('SQL select: ');
	console.log(data);
  }
});

/* test sql - non-select */

utilsapi.sql('mode=raw&query=SHOW TABLES',  (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('SQL DDL: ');
	console.log(data);
  }
});

/* test search */

searchapi.search({
            'index':'movies',
            'query':{'bool':{'must':[{'query_string':' movie'}]}},
            'script_fields':{'myexpr':{'script':{'inline':'IF(rating>8,1,0)'}}},
             'sort':[{'myexpr':'desc'},{'_score':'desc'}],
             'profile': true
            },  (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('Search: ');
	console.log(data);
  }
});

/* test percolate */

searchapi.percolate('test_pq',{
            
             "query":{"percolate":{"document":{"content":"sample content"}}}
            },  (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('Percolate with hit: ');
	console.log(data);
  }
});
searchapi.percolate('test_pq',{
            
             "query":{"percolate":{"document":{"content":"not found"}}}
            },  (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('Percolate no hit: ');
	console.log(data);
  }
});
	