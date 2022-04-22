/*
 * Manticore Search Client
 * Copyright (c) 2020-2021, Manticore Software LTD (https://manticoresearch.com)
 *
 * All rights reserved
 */
 
package com.manticore.clienttest;
import com.manticoresearch.client.ApiClient;
import com.manticoresearch.client.ApiException;
import com.manticoresearch.client.Configuration;
import com.manticoresearch.client.model.*;
import com.manticoresearch.client.api.IndexApi;
import com.manticoresearch.client.api.UtilsApi;
import com.manticoresearch.client.api.SearchApi;
 import java.lang.reflect.Method;
import org.json.JSONException;
import org.json.JSONObject;

import java.util.*;
import java.lang.reflect.Type;
/**
 * Hello world!
 *
 */

public class App 
{
 
    public static void main( String[] args )
    {
    ApiClient client = Configuration.getDefaultApiClient();
    client.setBasePath("http://manticoresearch-manticore:6368");


    IndexApi indexApi = new IndexApi(client);
    SearchApi searchApi = new SearchApi(client);
    UtilsApi utilsApi = new UtilsApi(client);
  /*  String body = "{\"insert\": {\"index\": \"movies\", \"id\": 0,\"doc\": {\"plot\": \"A secret team goes to North Pole\", \"rating\": 9.5, \"language\": [2, 3], \"title\": \"This is an older movie\", \"lon\": 51.99, \"meta\": {\"keywords\":[\"travel\",\"ice\"],\"genre\":[\"adventure\"]}, \"year\": 1950, \"lat\": 60.4, \"advise\": \"PG-13\"}}}" +"\n"+"{\"insert\": {\"index\": \"movies\", \"id\": 564343,\"doc\": {\"plot\": \" team goes to North Pole\", \"rating\": 9.5, \"language\": [2, 3], \"title\": \"This is a newer movie\", \"lon\": 51.99, \"meta\": {\"keywords\":[\"travel\",\"ice\"],\"genre\":[\"adventure\"]}, \"year\": 1950, \"lat\": 60.4, \"advise\": \"PG-13\"}}}" +"\n"+"{\"delete\":{\"index\" : \"movies\", \"id\" : 701}}"+"\n"; 
  */
    try {
        String sql ="DROP TABLE IF EXISTS products";
        Object sqlresult = utilsApi.sql(sql, true);
        System.out.println(sqlresult);
        
        sqlresult = utilsApi.sql("CREATE TABLE IF NOT EXISTS products (title text, price float, sizes multi, meta json, coeff float, tags1 multi, tags2 multi)", true);
        System.out.println(sqlresult);        

        sql ="SELECT * FROM products";
        sqlresult = utilsApi.sql(sql, false);
        System.out.println(sqlresult);        
        
        sql ="SELECT * FROM products";
        sqlresult = utilsApi.sql(sql, false);
        System.out.println(sqlresult);
        
        sql ="TRUNCATE TABLE products";
        sqlresult = utilsApi.sql(sql, true);
        System.out.println(sqlresult);        
        
        
        
                
        InsertDocumentRequest newdoc = new InsertDocumentRequest();
        HashMap<String,Object> doc = new HashMap<String,Object>(){{
            put("title","first");
            put("tags1",new int[] {4,2,1,3});
        }};
        newdoc.index("products").id(1L).setDoc(doc); 
        sqlresult = indexApi.insert(newdoc);
        Map<String,Object> query = new HashMap<String,Object>();
        query.put("match_all",null);
        SearchRequest searchRequest = new SearchRequest();
        searchRequest.setIndex("products");
        searchRequest.setQuery(query);
        HashMap<String,Object> highlight = new HashMap<String,Object>(){{
            put("fields",new String[] {"title"});
            
        }};
        searchRequest.setHighlight(highlight);
        SearchResponse searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );
        
        sql ="TRUNCATE TABLE products";
        sqlresult = utilsApi.sql(sql);
        System.out.println(sqlresult); 
        
         newdoc = new InsertDocumentRequest();
        doc = new HashMap<String,Object>(){{
            put("title","Crossbody Bag with Tassel");
            put("price",19.85);
        }};
        
        
        HashMap<String,Object> filters = new HashMap<String,Object>(){{
                  put("must",
                  new HashMap<String,Object>(){{
                        put("equals",new HashMap<String,Integer>(){{
                            put("author_id",123);
                        }});
                        put("in",
                          new HashMap<String,Object>(){{
                                  put("forum_id",new int[] {1,3,7});
                          }});
                  }});
        }};
        
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        query.put("bool",filters);
        searchRequest = new SearchRequest();
        searchRequest.setIndex("forum");
        searchRequest.setQuery(query);
        searchRequest.setSort(new ArrayList<Object>(){{
            add(new HashMap<String,String>(){{ put("post_date","desc");}});
        }});
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );
        List<Object> hits  = searchResponse.getHits().getHits(); 
        for(Object hit:hits) {
          HashMap<String,? > map = (HashMap<String,?>) hit;
          HashMap<String,? > source = (HashMap<String,?>) map.get("_source");
         System.out.println("id: "+map.get("_id")+" | title: "+source.get("title"));
        }        
        
        
        
        query = new HashMap<String,Object>();
        query.put("match", new HashMap<String,Object>(){{
                put("title","first");
            }});
        searchRequest = new SearchRequest();
        searchRequest.setIndex("forum");
        searchRequest.setQuery(query);
        searchResponse = searchApi.search(searchRequest);
        
        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("forum");
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        searchRequest.setQuery(query);
        Object expressions = new HashMap<String,Object>(){{
            put("ebs","abs(author_id-forum_id)");
        }};
        searchRequest.setExpressions(expressions);
        searchResponse = searchApi.search(searchRequest);
        
        newdoc.index("products").id(1L).setDoc( doc);
        sqlresult = indexApi.insert(newdoc);
        System.out.println(sqlresult);      
        
        doc = new HashMap<String,Object>(){{
            put("title","Crossbody Bag with Tassel");
        }};
        newdoc.index("products").id(2L).setDoc(doc); 
        sqlresult = indexApi.insert(newdoc);
        System.out.println(sqlresult);              
        
         doc = new HashMap<String,Object>(){{
            put("title","Yellow bag");
        }};
        newdoc.index("products").id(0L).setDoc(doc); 
        sqlresult = indexApi.insert(newdoc);
        System.out.println(sqlresult);             
        
        doc = new HashMap<String,Object>(){{
            put("title","Yellow bag");
        }};
        newdoc.index("products").id(0L).setDoc(doc); 
        sqlresult = indexApi.insert(newdoc);
        System.out.println(sqlresult);          
        
        String body = "{\"insert\": {\"index\" : \"products\", \"id\" : 3, \"doc\" : {\"title\" : \"Crossbody Bag with Tassel\", \"price\" : 19.85}}}" +"\n"+"{\"delete\":{\"index\" : \"movies\", \"id\" : 701}}"+"\n"+
    "{\"insert\": {\"index\" : \"products\", \"id\" : 4, \"doc\" : {\"title\" : \"microfiber sheet set\", \"price\" : 19.99}}}"+"\n"+
    "{\"insert\": {\"index\" : \"products\", \"id\" : 5, \"doc\" : {\"title\" : \"CPet Hair Remover Glove\", \"price\" : 7.99}}}"+"\n";         
        BulkResponse bulkresult = indexApi.bulk(body);
        System.out.println(bulkresult);
        
        // Search options test
        searchRequest = new SearchRequest();
        searchRequest.setIndex("products");
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        searchRequest.setQuery(query);
        Object options = new HashMap<String,Object>(){{
			put("max_matches", 2);
		}};
		searchRequest.setOptions(options);
		sqlresult = searchApi.search(searchRequest);
 		System.out.println(sqlresult);
        
        
        doc = new HashMap<String,Object>(){{
            put("title","Yellow bag");
            put("sizes",new int[]{40,41,42,43});
        }};
        newdoc.index("products").id(0L).setDoc(doc); 
        sqlresult =  indexApi.insert(newdoc);
        System.out.println(sqlresult);             

        doc = new HashMap<String,Object>(){{
            put("title","Yellow bag");
            put("meta", 
                      new HashMap<String,Object>(){{
                                  put("size",41);
                                  put("color","red");
                       }});

   
        }};
        newdoc.index("products").id(0L).setDoc(doc); 
        sqlresult =  indexApi.insert(newdoc);
        System.out.println(sqlresult);

        doc = new HashMap<String,Object>(){{
            put("title","document one");
            put("price",10);
 
        }};
        newdoc.index("products").id(0L).setDoc(doc); 
        sqlresult =  indexApi.insert(newdoc);
        System.out.println(sqlresult);
        
        body = "{\"replace\": {\"index\" : \"products\", \"id\" : 1, \"doc\" : {\"title\" : \"document one\"}}}"+"\n"+
    "{\"replace\": {\"index\" : \"products\", \"id\" : 2, \"doc\" : {\"title\" : \"document two\"}}}";         
        bulkresult = indexApi.bulk(body);
        System.out.println(bulkresult);
        
        UpdateDocumentRequest updatedoc = new UpdateDocumentRequest();
        doc = new HashMap<String,Object >(){{
            put("price",10);
 
        }};
        updatedoc.index("products").id(1L).setDoc(doc); 
        sqlresult =  indexApi.update(updatedoc);
        System.out.println(sqlresult);
        
        updatedoc = new UpdateDocumentRequest();
        doc = new HashMap<String,Object >(){{
            put("price",10);
            put("coeff",3465.23);
            put("tags1",new int[]{3,6,4});
            put("tags2",new int[]{});
            
        }};
        updatedoc.index("products").id(1L).setDoc(doc); 
        sqlresult =  indexApi.update(updatedoc);
        System.out.println(sqlresult);              
         
        doc = new HashMap<String,Object>(){{
            put("title","title");
            put("meta", 
                      new HashMap<String,Object>(){{
                                  put("tags",new int[]{1,2,3});
                       }});
 
        }};
        newdoc.index("products").id(100L).setDoc(doc);        
        sqlresult = indexApi.insert(newdoc);
        System.out.println(sqlresult);       
        
        updatedoc = new UpdateDocumentRequest();
        doc = new HashMap<String,Object >(){{
            put("meta.tags[0]",100);
 
        }};
        updatedoc.index("products").id(100L).setDoc(doc); 
        sqlresult =  indexApi.update(updatedoc);
        System.out.println(sqlresult);        
        
        updatedoc = new UpdateDocumentRequest();
        doc = new HashMap<String,Object>(){{
            put("tags1",new int[]{});

          
        }};
        updatedoc.index("products").id(1L).setDoc(doc); 
        sqlresult =  indexApi.update(updatedoc);
        System.out.println(sqlresult);      
        
        body = "{ \"update\" : { \"index\" : \"products\", \"doc\": { \"coeff\" : 1000 }, \"query\": { \"range\": { \"price\": { \"gte\": 1000 } } } }} "+"\n"+
    "{ \"update\" : { \"index\" : \"products\", \"doc\": { \"coeff\" : 0 }, \"query\": { \"range\": { \"price\": { \"lt\": 1000 } } } } }"+"\n";         
        bulkresult = indexApi.bulk(body);
        System.out.println(bulkresult);
        
        
        sql ="DROP TABLE IF EXISTS test_pq";
        sqlresult = utilsApi.sql(sql);
        sqlresult =   utilsApi.sql("create table test_pq(title text, price integer) type='pq'");
        System.out.println(sqlresult);    
        doc = new HashMap<String,Object>(){{
            put("query",new HashMap<String,Object >(){{
                put("q1","@title shoes");
                put("filters","price>5");
                put("tags",new String[] {"Loius Vuitton"});
 
        }});
        }};
        
        
        DeleteDocumentRequest deleteRequest = new DeleteDocumentRequest();
        query = new HashMap<String,Object>();
        query.put("match",new HashMap<String,Object>(){{
            put("*","dummy");
        }});
        deleteRequest.index("products").setQuery(query); 
        sqlresult = indexApi.delete(deleteRequest);
        System.out.println(sqlresult);      
        
        sqlresult =  utilsApi.sql("SHOW TABLES");
        System.out.println(sqlresult);  
        
        query = new HashMap<String,Object>();
query.put("query_string","@title way* @content hey");
searchRequest = new SearchRequest();
searchRequest.setIndex("forum");
searchRequest.setQuery(query);
searchRequest.setProfile(true);
searchRequest.setLimit(1);
searchRequest.setSource(new ArrayList<String>(){{
    add("*");
}});
sqlresult =    searchApi.search(searchRequest);
 System.out.println(sqlresult);  
 
InsertDocumentRequest docRequest = new InsertDocumentRequest();
 doc = new HashMap<String,Object>(){{
            put("title","document one");
            put("price",10);
}};
docRequest.index("products").id(1L).setDoc(doc); 
sqlresult = indexApi.replace(docRequest);
 System.out.println(sqlresult);  
 System.out.println("--------------------");  
 
body = "{\"replace\": {\"index\" : \"products\", \"id\" : 1, \"doc\" : {\"title\" : \"document one\"}}}" +"\n"+ 
    "{\"replace\": {\"index\" : \"products\", \"id\" : 2, \"doc\" : {\"title\" : \"document two\"}}}"+"\n" ;         
          bulkresult = indexApi.bulk(body);
        System.out.println(bulkresult);
        
        
/* percolate */
        sql ="DROP TABLE IF EXISTS products";
        sqlresult = utilsApi.sql(sql);
        sqlresult =   utilsApi.sql("create table products(title text, color string) type='pq'");
         System.out.println(sqlresult);
        doc = new HashMap<String,Object>(){{
            put("query",     "@title bag"               );
        }};
        newdoc = new InsertDocumentRequest();
        newdoc.index("products").setDoc(doc); 
        sqlresult =  indexApi.insert(newdoc);
        System.out.println(sqlresult);
        
        doc = new HashMap<String,Object>(){{
            put("query",     "@title shoes"               );
            put("filters",     "color='red'"               );
        }};
        newdoc = new InsertDocumentRequest();
        newdoc.index("products").setDoc(doc); 
        sqlresult =  indexApi.insert(newdoc);
        System.out.println(sqlresult);
        
        doc = new HashMap<String,Object>(){{
            put("query",     "@title shoes"               );
             put("filters",     "color IN ('blue', 'green')"               );
        }};
        newdoc = new InsertDocumentRequest();
        newdoc.index("products").setDoc(doc); 
        sqlresult =  indexApi.insert(newdoc);
        System.out.println(sqlresult);        
PercolateRequest percolateRequest = new PercolateRequest();
        
    query = new HashMap<String,Object>(){{
        put("percolate",new HashMap<String,Object >(){{
            put("document", new HashMap<String,Object >(){{ 
                put("title","what a nice bag");
            }});
        }});
    }};
    percolateRequest.query(query);
    sqlresult =    searchApi.percolate("products",percolateRequest);
 System.out.println(sqlresult);  
        
 percolateRequest = new PercolateRequest();
        
    query = new HashMap<String,Object>(){{
        put("percolate",new HashMap<String,Object >(){{
            put("documents", new ArrayList<Object>(){{
                    add(new HashMap<String,Object >(){{ 
                        put("title","nice pair of shoes");
                        put("color","blue");
                    }});
                    add(new HashMap<String,Object >(){{ 
                        put("title","beautiful bag");

                    }});
                    
                     }});
        }});
    }};
    percolateRequest.query(query);
    sqlresult =    searchApi.percolate("products",percolateRequest);
 System.out.println(sqlresult);  
 
 percolateRequest = new PercolateRequest();
        
    query = new HashMap<String,Object>(){{
        put("percolate",new HashMap<String,Object >(){{
            put("documents", new ArrayList<Object>(){{
                    add(new HashMap<String,Object >(){{ 
                        put("title","angry test");
                        put("gid",3);
                    }});
                    add(new HashMap<String,Object >(){{ 
                        put("title","filter test doc2");
                        put("gid",13);

                    }});
                    
                     }});
        }});
    }};
    percolateRequest.query(query);
    sqlresult =    searchApi.percolate("pq",percolateRequest);
 System.out.println(sqlresult);  
 
 
        /*
      String sql ="SHOW TABLES";
     Object sqlresult = utilsApi.sql(sql);
      System.out.println(sqlresult);
      InsertDocumentRequest newdoc = new InsertDocumentRequest();
       HashMap<String,Object> doc = new HashMap<String,Object>(){{
            put("title","Crossbody Bag with Tassel");
            put("price",19.85);
       }};
      newdoc.index("testjava").id(0L).setDoc(doc);

       sqlresult =  indexApi.insert(newdoc);
      System.out.println(sqlresult);
      
      SearchRequest searchRequest = new SearchRequest("{\"index\":\"forum\",\"query\":{\"match_all\":{},\"bool\":{\"must\":[{\"equals\":{\"author_id\":123}},{\"in\":{\"forum_id\":[1,3,7]}}]}},\"sort\":[{\"post_date\":\"desc\"}]}");
     
      HashMap<String,Object> filters =
      new HashMap<String,Object>(){{
                  put("must",
                  new HashMap<String,Object>(){{
                          put("equals",
                          new HashMap<String,Integer>(){{
                                  put("author_id",123);
                          }});

                          
                  }});
      }};
      
      Object  filters = new Object() {
        public Object must = new Object() {
        public Object equals = new Object(){
            public Integer author_id = 123;
        };
        };
      };
      
      Gson gfilter = new Gson();
      System.out.println(filters);
      searchRequest.index("forum").putQueryItem("bool",filters);
     
             Object  filters = new Object() {
        public Object must = new Object() {
        public Object equals = new Object(){
            public Integer author_id = 123;
        };
        };
      };
   
       HashMap<String,Object> filters =
      new HashMap<String,Object>(){{
                  put("must",
                  new HashMap<String,Object>(){{
                        put("match",new HashMap<String,String>(){{
                            put("*","one");
                        }});
                        put("equals",
                          new HashMap<String,Integer>(){{
                                  put("brand_id",10);
                          }});

                          
                  }});
      }};
      

   class SearchR  {
        public float took;
        public boolean timed_out;
        public Object hits;
 }

        Map<String,Object> query = new HashMap<String,Object>();
        query.put("bool",filters);
    SearchRequest searchRequest = new SearchRequest();
        searchRequest.setIndex("facetdemo2");
        searchRequest.setQuery(query);
       System.out.println( searchRequest.toString());
       SearchResponse searchResponse = searchApi.search(searchRequest);
      // SearchR sr = gson.fromJson(searchResponse,SearchR.class);
       System.out.println(searchResponse.toString() );
      List<Object> hits  = searchResponse.getHits().getHits(); 
      Method[] methods;
      for(Object hit:hits) {

          HashMap<String,? > map = (HashMap<String,?>) hit;
          HashMap<String,? > source = (HashMap<String,?>) map.get("_source");
         System.out.println("id: "+map.get("_id")+" | title: "+source.get("title"));
         
         
          
      }
     
      /*
      List<Object> hits = searchResponse.getHits().getHits();
       System.out.println(hits);
       for(Object hit:hits) {
 
           System.out.println(hit);
       }
      
       BulkResponse result = indexApi.bulk(body);
       System.out.println(result); */
        sqlresult =   utilsApi.sql("SHOW THREADS");
        System.out.println(sqlresult);  
        sqlresult = utilsApi.sql("DROP TABLE IF  EXISTS books");
        System.out.println(sqlresult);      
        sqlresult = utilsApi.sql("CREATE TABLE IF NOT EXISTS books (title text, content text)");
        System.out.println(sqlresult);      
        
        
        body = "{\"insert\": {\"index\" : \"books\", \"id\" : 4, \"doc\" : {\"title\" : \"Book four\", \"content\" : \"Don`t try to compete in childishness, said Bliss.\"}}}" +"\n"+
    "{\"insert\": {\"index\" : \"books\", \"id\" : 3, \"doc\" : {\"title\" : \"Book three\", \"content\" :\"Trevize whispered, \\\"It gets infantile pleasure out of display. I`d love to knock it down.\\\"\"}}}"+"\n"+
    "{\"insert\": {\"index\" : \"books\", \"id\" : 5, \"doc\" : {\"title\" : \"Books two\", \"content\" :\"A door opened before them, revealing a small room. Bander said, \\\"Come, half-humans, I want to show you how we live.\\\"\"}}}"+"\n"+
    "{\"insert\": {\"index\" : \"books\", \"id\" : 1, \"doc\" : {\"title\" : \"Books one\", \"content\" :\"They followed Bander. The robots remained at a polite distance, but their presence was a constantly felt threat. Bander ushered all three into the room. One of the robots followed as well. Bander gestured the other robots away and entered itself. The door closed behind it. \"}}}"+"\n";
         bulkresult = indexApi.bulk(body);
        System.out.println(bulkresult);
        
        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("books");
        query = new HashMap<String,Object>();
        query.put("match",new HashMap<String,Object>(){{
            put("*","try");
        }});        
        searchRequest.setQuery(query);
        highlight = new HashMap<String,Object>(){{
           
            
        }};
        searchRequest.setHighlight(highlight);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );   
        
        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("books");
        query = new HashMap<String,Object>();
        query.put("match",new HashMap<String,Object>(){{
            put("*","try|gets|down|said");
        }});        
        searchRequest.setQuery(query);
        highlight = new HashMap<String,Object>(){{
           put("limit",50);
            
        }};
        searchRequest.setHighlight(highlight);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );   
        
        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("books");
        query = new HashMap<String,Object>();
        query.put("match",new HashMap<String,Object>(){{
            put("*","one|robots");
        }});        
        searchRequest.setQuery(query);
        highlight = new HashMap<String,Object>(){{
           put("fields",new String[] {"content"});
          
            
        }};
        searchRequest.setHighlight(highlight);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );   

        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("books");
        query = new HashMap<String,Object>();
        query.put("match",new HashMap<String,Object>(){{
            put("*","one|robots");
        }});        
        searchRequest.setQuery(query);
        highlight = new HashMap<String,Object>(){{
                     
        }};
        searchRequest.setHighlight(highlight);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );   

        searchRequest = new SearchRequest();
        searchRequest.setIndex("books");
        query = new HashMap<String,Object>();
        query.put("match",new HashMap<String,Object>(){{
            put("*","one|robots");
        }});        
        searchRequest.setQuery(query);
        highlight = new HashMap<String,Object>(){{
            put("fields",new String[] {"content","title"});
            put("highlight_query", 
                      new HashMap<String,Object>(){{
                                  put("match", new HashMap<String,Object>(){{
                                    put("*","polite distance");
                            }});
                       }});
        }};
        searchRequest.setHighlight(highlight);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );   
        
        HashMap<String,Object> aggs = new HashMap<String,Object>(){{
            put("release_year", new HashMap<String,Object>(){{ 
                put("terms", new HashMap<String,Object>(){{ 
                        put("field","release_year");
                        put("size",100);
                }});
            }});
        }};
        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("films");        
        searchRequest.setLimit(0);
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        searchRequest.setQuery(query);
        searchRequest.setAggs(aggs);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );           
        
        
        
   
        
       aggs = new HashMap<String,Object>(){{
            put("color", new HashMap<String,Object>(){{ 
                put("sizes", new HashMap<String,Object>(){{ 
                        put("field","meta.color");
                        put("size",100);

                }});
            }});
        }};
        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("products2");        
        searchRequest.setLimit(0);
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        searchRequest.setQuery(query);
        searchRequest.setAggs(aggs);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );  
        
        aggs = new HashMap<String,Object>(){{
            put("color", new HashMap<String,Object>(){{ 
                put("sizes", new HashMap<String,Object>(){{ 
                        put("field","meta.color");
                        put("size",100);

                }});
            }});
        }};
        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("products2");        
        searchRequest.setLimit(0);
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        searchRequest.setQuery(query);
        searchRequest.setAggs(aggs);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );  
        
        
        aggs = new HashMap<String,Object>(){{
            put("group_property", new HashMap<String,Object>(){{ 
                put("sizes", new HashMap<String,Object>(){{ 
                        put("field","price");
                       

                }});
            }});
            put("group_brand_id", new HashMap<String,Object>(){{ 
                put("sizes", new HashMap<String,Object>(){{ 
                        put("field","brand_id");
                       

                }});
            }});            
        }};
        
        searchRequest = new SearchRequest();
        searchRequest.setIndex("facetdemo2");        
        searchRequest.setLimit(5);
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        searchRequest.setQuery(query);
        searchRequest.setAggs(aggs);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );  
        
        
        searchRequest = new SearchRequest();
        expressions = new HashMap<String,Object>(){{
            put("price_range","INTERVAL(price,200,400,600,800)");
        }};
        searchRequest.setExpressions(expressions);
        aggs = new HashMap<String,Object>(){{
            put("group_property", new HashMap<String,Object>(){{ 
                put("sizes", new HashMap<String,Object>(){{ 
                        put("field","price_range");
                       

                }});
            }});
   
        }};
        searchRequest.setIndex("facetdemo2");        
        searchRequest.setLimit(5);
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        searchRequest.setQuery(query);
        searchRequest.setAggs(aggs);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );  
        
        
        searchRequest = new SearchRequest();
        aggs = new HashMap<String,Object>(){{
            put("group_property", new HashMap<String,Object>(){{ 
                put("sizes", new HashMap<String,Object>(){{ 
                        put("field","price");
                        put("size",1);
                       

                }});
            }});
            put("group_brand_id", new HashMap<String,Object>(){{ 
                put("sizes", new HashMap<String,Object>(){{ 
                        put("field","brand_id");
                        put("size",3);
                       

                }});
            }});            
        }};
        searchRequest.setIndex("facetdemo2");        
        searchRequest.setLimit(5);
        query = new HashMap<String,Object>();
        query.put("match_all",null);
        searchRequest.setQuery(query);
        searchRequest.setAggs(aggs);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );  
        
        
        /*
        searchRequest = new SearchRequest();
        searchRequest.setIndex("books");
        query = new HashMap<String,Object>();
        query.put("match",new HashMap<String,Object>(){{
            put("content","and first");
        }});        
        searchRequest.setQuery(query);
        highlight = new HashMap<String,Object>(){{
            put("fields",new String[] {"content","title"});
            put("highlight_query", 
                      new HashMap<String,Object>(){{
                                  put("match", new HashMap<String,Object>(){{
                                    put("*","into three five");
                            }});
                       }});
        }};
        searchRequest.setHighlight(highlight);
        searchResponse = searchApi.search(searchRequest);
        System.out.println(searchResponse.toString() );   
        */
        
    } catch (ApiException e) {
      System.err.println("Exception when calling IndexApi#bulk");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
    }
}
