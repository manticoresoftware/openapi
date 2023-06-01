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

import java.math.BigDecimal;
/**
 * Hello world!
 *
 */

public class App 
{
	public static void testSearchRequestBuild(ApiClient client, IndexApi indexApi, SearchApi searchApi, UtilsApi utilsApi)
	{
		try {
			utilsApi.sql("DROP TABLE IF EXISTS movies", true);
		    utilsApi.sql("CREATE TABLE IF NOT EXISTS movies (title text, plot text, year integer, rating float, code multi)", true);
		    
		    List<String> docs = Arrays.asList(
				"{\"insert\": {\"index\" : \"movies\", \"id\" : 1, \"doc\" : {\"title\" : \"Star Trek 2: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2002, \"rating\": 6.4, \"code\": [1,2,3]}}}",
		        "{\"insert\": {\"index\" : \"movies\", \"id\" : 2, \"doc\" : {\"title\" : \"Star Trek 1: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2001, \"rating\": 6.5, \"code\": [1,12,3]}}}",
		        "{\"insert\": {\"index\" : \"movies\", \"id\" : 3, \"doc\" : {\"title\" : \"Star Trek 3: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2003, \"rating\": 6.6, \"code\": [11,2,3]}}}",
		        "{\"insert\": {\"index\" : \"movies\", \"id\" : 4, \"doc\" : {\"title\" : \"Star Trek 4: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2003, \"rating\": 6, \"code\": [1,2,4]}}}"					        	
		    );
	
			BulkResponse res = indexApi.bulk( String.join("\n", docs) );
			System.out.println(res);
			
			Map<String,Object> query = new HashMap<String,Object>();
	        //query.put("match_all", new HashMap<String,Object>());
	        query.put("query_string", "Star");
	        SearchRequest searchRequest = new SearchRequest();
	        searchRequest.setIndex("movies");
	        searchRequest.setQuery(query);
			searchRequest.setLimit(10);
			searchRequest.setTrackScores(false);

			SearchResponse searchResponse = searchApi.search(searchRequest);

			Map<String,Object> options = new HashMap<String,Object>();
			options.put("cutoff", 5);
			options.put("ranker", "bm25");
			searchRequest.setOptions(options);
			
			searchRequest.setSource("title");
			
			searchResponse = searchApi.search(searchRequest);
	        
	        List<String> includes = new ArrayList<String>( Arrays.asList("title", "year", "rating") );
			List<String> excludes = Arrays.asList("code");
			SourceByRules source = new SourceByRules();
			source.setIncludes(includes);
			source.setExcludes(excludes);
			searchRequest.setSource(source);
			
			searchResponse = searchApi.search(searchRequest);

			List<Object> sort = new ArrayList<Object>( Arrays.asList("year") );
			SortOrder sort2 = new SortOrder();
			sort2.setAttr("rating");
			sort2.setOrder(SortOrder.OrderEnum.ASC);
			sort.add(sort2);
			SortMVA sort3 = new SortMVA();
			sort3.setAttr("code");
			sort3.setOrder(SortMVA.OrderEnum.DESC);
			sort3.setMode(SortMVA.ModeEnum.MAX);
			sort.add(sort3);
			searchRequest.setSort(sort);

			searchResponse = searchApi.search(searchRequest);
			
			Map<String,String> expr = new HashMap<String,String>();
			expr.put("expr1", "min(year,2900)");
			List<Object> expressions = new ArrayList<Object>();
			expressions.add(expr);
        	expressions.add( 
        		new HashMap<String,String>(){{ 
		        	put("expr2","max(year,2100)");
		
				}}
        	);
        	searchRequest.setExpressions(expressions);

        	searchResponse = searchApi.search(searchRequest);
        	
        	source = (SourceByRules)searchRequest.getSource();
        	includes.addAll( Arrays.asList("expr1", "expr2") );
        	source.setIncludes(includes);
        	searchRequest.setSource(source);

        	searchResponse = searchApi.search(searchRequest);
        				        					
	        Aggregation agg = new Aggregation();
	        agg.setName("agg1");
	        agg.setField("year");
        	agg.setSize(10);
			searchRequest.setAggs( new ArrayList<Aggregation>( Arrays.asList(agg) ) );

			searchResponse = searchApi.search(searchRequest);
			
			agg = new Aggregation();
	        agg.setName("agg2");
	        agg.setField("rating");
			List<Aggregation> aggs = searchRequest.getAggs();
			aggs.add(agg);
			searchRequest.setAggs(aggs);

			searchResponse = searchApi.search(searchRequest);
			
			Highlight highlight = new Highlight();
        	highlight.setFieldnames( Arrays.asList("title") );
        	highlight.setPostTags("</post_tag>");
    	    highlight.setEncoder(Highlight.EncoderEnum.DEFAULT);
	        highlight.setSnippetBoundary(Highlight.SnippetBoundaryEnum.SENTENCE);
        	searchRequest.setHighlight(highlight);

        	searchResponse = searchApi.search(searchRequest);

			HighlightField highlightField = new HighlightField();
			highlightField.setName("title");
			highlightField.setLimit(5);
			List<HighlightField> highlightFields = new ArrayList<HighlightField>( Arrays.asList(highlightField) );
			highlight.setFields(highlightFields);
			searchRequest.setHighlight(highlight);

			searchResponse = searchApi.search(searchRequest);
			
			highlightField = new HighlightField();
			highlightField.setName("plot");
			highlightField.setLimitWords(10);
			highlightFields = highlight.getFields();
        	highlightFields.add(highlightField);
        	highlight.setFields(highlightFields);
        	
        	searchRequest.setHighlight(highlight);

        	searchResponse = searchApi.search(searchRequest);

        	HashMap<String,Object> highlightQuery = new HashMap<String,Object>(){{
	            put("match", new HashMap<String,Object>(){{
	            	put("*","Star");
	        	}});
	        }};
        	highlight.setHighlightQuery(highlightQuery);

        	searchResponse = searchApi.search(searchRequest);
						
			QueryFilter queryFilter = new QueryFilter();
			queryFilter.setQueryString("Star Trek 2");								
			searchRequest.setFulltextFilter(queryFilter);

			searchResponse = searchApi.search(searchRequest);
			
			MatchFilter matchFilter = new MatchFilter();
			matchFilter.setQueryString("Nemesis");
			matchFilter.setQueryFields("title");	
			searchRequest.setFulltextFilter(matchFilter);

			searchResponse = searchApi.search(searchRequest);
			
			MatchPhraseFilter matchPhraseFilter = new MatchPhraseFilter();
			matchPhraseFilter.setQueryPhrase("Star Trek 2");
			matchPhraseFilter.setQueryFields("title");
			searchRequest.setFulltextFilter(matchPhraseFilter);

			searchResponse = searchApi.search(searchRequest);
			
			MatchOpFilter matchOpFilter = new MatchOpFilter();
			matchOpFilter.setQueryString("Enterprise test");
			matchOpFilter.setQueryFields("title,plot");
			matchOpFilter.setOperator(MatchOpFilter.OperatorEnum.OR);
			searchRequest.setFulltextFilter(matchOpFilter);

			searchResponse = searchApi.search(searchRequest);
			
			EqualsFilter equalsFilter = new EqualsFilter();
			equalsFilter.setField("year");
			equalsFilter.setValue(2003);
			searchRequest.setAttrFilter(equalsFilter);

			searchResponse = searchApi.search(searchRequest);
			
			InFilter inFilter = new InFilter();
			inFilter.setField("year");
			inFilter.setValues( Arrays.asList(2001, 2002) );
	        searchRequest.setAttrFilter(inFilter);

	        searchResponse = searchApi.search(searchRequest);
			
			RangeFilter rangeFilter = new RangeFilter();
			rangeFilter.setField("year");
			rangeFilter.setLte(BigDecimal.valueOf(2001));
			rangeFilter.setGte(BigDecimal.valueOf(0));
			searchRequest.setAttrFilter(rangeFilter);

			searchResponse = searchApi.search(searchRequest);

			rangeFilter.setField("rating");
			rangeFilter.setGt(BigDecimal.valueOf(1.5));
			searchRequest.setAttrFilter(rangeFilter);

			searchResponse = searchApi.search(searchRequest);

			GeoDistanceFilter geoFilter = new GeoDistanceFilter();
			GeoDistanceFilterLocationAnchor locAnchor = new GeoDistanceFilterLocationAnchor();
			locAnchor.setLat(BigDecimal.valueOf(10));
			locAnchor.setLon(BigDecimal.valueOf(20));
			geoFilter.setLocationAnchor(locAnchor);
			geoFilter.setLocationSource("year,rating");
			geoFilter.setDistanceType(GeoDistanceFilter.DistanceTypeEnum.ADAPTIVE);
			geoFilter.setDistance("100km");
			searchRequest.setAttrFilter(geoFilter);

			searchResponse = searchApi.search(searchRequest);
			
			BoolFilter boolFilter = new BoolFilter();
			equalsFilter = new EqualsFilter();
			equalsFilter.setField("year");
			equalsFilter.setValue(2003);
			boolFilter.setMust( new ArrayList<Object>( Arrays.asList(equalsFilter) ) );
			rangeFilter = new RangeFilter();
			rangeFilter.setField("rating");
			rangeFilter.setLte(BigDecimal.valueOf(6));
			List<Object> mustFilter = boolFilter.getMust();
			mustFilter.add(rangeFilter);
			boolFilter.setMust(mustFilter);
			searchRequest.setAttrFilter(boolFilter);

			searchResponse = searchApi.search(searchRequest);

			equalsFilter = new EqualsFilter();
			equalsFilter.setField("year");
			equalsFilter.setValue(2001);
			boolFilter.setMustNot( Arrays.asList(equalsFilter) );
			searchRequest.setAttrFilter(boolFilter);

			searchResponse = searchApi.search(searchRequest);
			
			boolFilter = new BoolFilter();
			
			MatchFilter fulltextFilter = new MatchFilter();
			fulltextFilter.setQueryString("Star");
			fulltextFilter.setQueryFields("title");
			BoolFilter nestedBoolFilter = new BoolFilter();
			equalsFilter = new EqualsFilter();
			equalsFilter.setField("rating");
			equalsFilter.setValue(6.5);
			nestedBoolFilter.setShould( Arrays.asList(equalsFilter, fulltextFilter) );
			boolFilter.setMust( Arrays.asList(nestedBoolFilter) );
			searchRequest.setAttrFilter(boolFilter);
											
			searchResponse = searchApi.search(searchRequest);
	  
	        System.out.println("Search tests finished");
	        
		} catch (ApiException e) {
	      System.err.println("Exception when testing a build of searchRequest");
	      System.err.println("Status code: " + e.getCode());
	      System.err.println("Reason: " + e.getResponseBody());
	      System.err.println("Response headers: " + e.getResponseHeaders());
	      e.printStackTrace();
	    }  
	}
 
    public static void main( String[] args )
    {
	    ApiClient client = Configuration.getDefaultApiClient();
	    client.setBasePath("http://manticoresearch-manticore:9308");
	    IndexApi indexApi = new IndexApi(client);
	    SearchApi searchApi = new SearchApi(client);
	    UtilsApi utilsApi = new UtilsApi(client);

		try {
			if (true) {
				testSearchRequestBuild(client, indexApi, searchApi, utilsApi);
				//return;
			}
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
	        query.put("match_all", new HashMap<String,Object>());
	        SearchRequest searchRequest = new SearchRequest();
	        searchRequest.setIndex("products");
	        searchRequest.setQuery(query);
	        Highlight highlight = new Highlight();
	        highlight.setFieldnames( Arrays.asList("title") );
	   
	        searchRequest.setHighlight(highlight);
	        SearchResponse searchResponse = searchApi.search(searchRequest);
	        System.out.println(searchResponse.toString() );
	        
	        sql ="TRUNCATE TABLE products";
	        sqlresult = utilsApi.sql(sql, true);
	        System.out.println(sqlresult); 
	        
	         newdoc = new InsertDocumentRequest();
	        doc = new HashMap<String,Object>(){{
	            put("title","Crossbody Bag with Tassel");
	            put("price",19.85);
	        }};
	        
	        /*
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
	        query.put("match_all", new HashMap<String,Object>());
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
	        query.put("match_all", new HashMap<String,Object>());
	        searchRequest.setQuery(query);
	        //Object expressions = new HashMap<String,Object>(){{
	        //    put("ebs","abs(author_id-forum_id)");
	        //}};
	        List<Object> expressions = Arrays.asList(
		        new HashMap<String,Object>(){{
		            put("ebs","abs(author_id-forum_id)");
		        }}
	        );
	        searchRequest.setExpressions(expressions);
	        searchResponse = searchApi.search(searchRequest);
	        */
	        List<Object> expressions = Arrays.asList(
		        new HashMap<String,Object>(){{
		            put("ebs","abs(author_id-forum_id)");
		        }}
	        );
	        
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
	        
	        String body = "{\"insert\": {\"index\" : \"products\", \"id\" : 3, \"doc\" : {\"title\" : \"Crossbody Bag with Tassel\", \"price\" : 19.85}}}" +"\n"+
	    "{\"insert\": {\"index\" : \"products\", \"id\" : 4, \"doc\" : {\"title\" : \"microfiber sheet set\", \"price\" : 19.99}}}"+"\n"+
	    "{\"insert\": {\"index\" : \"products\", \"id\" : 5, \"doc\" : {\"title\" : \"CPet Hair Remover Glove\", \"price\" : 7.99}}}"+"\n";         
	        BulkResponse bulkresult = indexApi.bulk(body);
	        System.out.println(bulkresult);
	        
	        // Search options test
	        searchRequest = new SearchRequest();
	        searchRequest.setIndex("products");
	        query = new HashMap<String,Object>();
	        query.put("match_all", new HashMap<String,Object>());
	        searchRequest.setQuery(query);
	        Map<String,Object> options = new HashMap<String,Object>(){{
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
	        sqlresult = utilsApi.sql(sql, true);
	        sqlresult =   utilsApi.sql("create table test_pq(title text, price integer) type='pq'", true);
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
	        
	        sqlresult =  utilsApi.sql("SHOW TABLES", true);
	        System.out.println(sqlresult);  
	        
	        /*
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
			sqlresult =  searchApi.search(searchRequest);
			System.out.println(sqlresult);
			*/  
	 
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
	        sqlresult = utilsApi.sql(sql, true);
	        sqlresult =   utilsApi.sql("create table products(title text, color string) type='pq'", true);
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
			
			PercolateRequestQuery percolateRequestQuery = new PercolateRequestQuery();
	        
		    query = new HashMap<String,Object>(){{
		        put("percolate",new HashMap<String,Object >(){{
		            put("document", new HashMap<String,Object >(){{ 
		                put("title","what a nice bag");
		            }});
		        }});
		    }};
		    
		    query = new HashMap<String,Object >(){{
		            put("document", new HashMap<String,Object >(){{ 
		                put("title","what a nice bag");
		            }});
			}};
		    percolateRequestQuery.setPercolate(query);
		    percolateRequest.query(percolateRequestQuery);
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
	    
		    query = new HashMap<String,Object >(){{
		            put("documents", new ArrayList<Object>(){{
		                    add(new HashMap<String,Object >(){{ 
		                        put("title","nice pair of shoes");
		                        put("color","blue");
		                    }});
		                    add(new HashMap<String,Object >(){{ 
		                        put("title","beautiful bag");
		
		                    }});
		                    
		                     }});
		        }};
	        
		    percolateRequestQuery.setPercolate(query);
			percolateRequest.query(percolateRequestQuery);
		    sqlresult = searchApi.percolate("products",percolateRequest);
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
		    
		    query = new HashMap<String,Object >(){{
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
		    }};
	    
		    percolateRequestQuery.setPercolate(query);
			percolateRequest.query(percolateRequestQuery);
	    
			sqlresult = searchApi.percolate("products",percolateRequest);
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
			System.out.println(result); 
			*/
	
	        sqlresult =   utilsApi.sql("SHOW THREADS", true);
	        System.out.println(sqlresult);  
	        sqlresult = utilsApi.sql("DROP TABLE IF  EXISTS books", true);
	        System.out.println(sqlresult);      
	        sqlresult = utilsApi.sql("CREATE TABLE IF NOT EXISTS books (title text, content text)", true);
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
	        highlight = new Highlight();
	        
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
	        highlight = new Highlight();
	        highlight.setLimit(50);
	        
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
	        highlight = new Highlight();
	        highlight.setFieldnames( Arrays.asList("content") );
	
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
	        highlight = new Highlight();
	
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
	        highlight = new Highlight();
	        highlight.setFieldnames( Arrays.asList("content", "title") );
	        highlight.setHighlightQuery(
				new HashMap<String,Object>(){{
	            	put("match", new HashMap<String,Object>(){{
	                	put("*","polite distance");
	                }});
				}} 
	        );
	        
	        searchRequest.setHighlight(highlight);
	        searchResponse = searchApi.search(searchRequest);
	        System.out.println(searchResponse.toString() );   
	        
	        Aggregation agg = new Aggregation();
	        agg.setName("release_year");
	        agg.setField("release_year");
	        agg.setSize(100);
	        
	        searchRequest = new SearchRequest();
	        searchRequest.setIndex("films");        
	        searchRequest.setLimit(0);
	        query = new HashMap<String,Object>();
	        query.put("match_all", new HashMap<String,Object>());
	        searchRequest.setQuery(query);
	        searchRequest.setAggs( Arrays.asList(agg) );
	        
	        
	        //searchResponse = searchApi.search(searchRequest);
	        //System.out.println(searchResponse.toString() );           
	        
	        agg = new Aggregation();
	        agg.setName("color");
	        agg.setField("meta.color");
	        agg.setSize(100);
	           
	        searchRequest = new SearchRequest();
	        searchRequest.setIndex("products2");        
	        searchRequest.setLimit(0);
	        query = new HashMap<String,Object>();
	        query.put("match_all", new HashMap<String,Object>());
	        searchRequest.setQuery(query);
	        searchRequest.setAggs( Arrays.asList(agg) );
	        //searchResponse = searchApi.search(searchRequest);
	        //System.out.println(searchResponse.toString() );  
	        
	        agg = new Aggregation();
	        agg.setName("color");
	        agg.setField("meta.color");
	        
	        searchRequest = new SearchRequest();
	        searchRequest.setIndex("products2");        
	        searchRequest.setLimit(0);
	        query = new HashMap<String,Object>();
	        query.put("match_all", new HashMap<String,Object>());
	        searchRequest.setQuery(query);
	        searchRequest.setAggs( Arrays.asList(agg) );
	        //searchResponse = searchApi.search(searchRequest);
	        //System.out.println(searchResponse.toString() );  
	        
	        Aggregation agg1 = new Aggregation();
	        agg1.setName("group_property");
	        agg1.setField("price");
	        Aggregation agg2 = new Aggregation();
	        agg2.setName("group_brand_id");
	        agg2.setField("brand_id");
			List<Aggregation> aggs = Arrays.asList(agg1, agg2);        
	                
	        searchRequest = new SearchRequest();
	        searchRequest.setIndex("facetdemo2");        
	        searchRequest.setLimit(5);
	        query = new HashMap<String,Object>();
	        query.put("match_all", new HashMap<String,Object>());
	        searchRequest.setQuery(query);
	        searchRequest.setAggs(aggs);
	        //searchResponse = searchApi.search(searchRequest);
	        //System.out.println(searchResponse.toString() );  
	        
	        
	        searchRequest = new SearchRequest();
	        expressions = Arrays.asList(
		        new HashMap<String,Object>(){{
		            put("price_range","INTERVAL(price,200,400,600,800)");
		        }}
	        );
	        searchRequest.setExpressions(expressions);
	        agg = new Aggregation();
	        agg.setName("group_property");
	        agg.setField("price_range");
	
	        searchRequest.setIndex("facetdemo2");        
	        searchRequest.setLimit(5);
	        query = new HashMap<String,Object>();
	        query.put("match_all", new HashMap<String,Object>());
	        searchRequest.setQuery(query);
	        searchRequest.setAggs( Arrays.asList(agg) );
	        //searchResponse = searchApi.search(searchRequest);
	        //System.out.println(searchResponse.toString() );  
	        
	        searchRequest = new SearchRequest();
	        agg1 = new Aggregation();
	        agg1.setName("group_property");
	        agg1.setField("price");
	        agg1.setSize(1);
	        agg2 = new Aggregation();
	        agg2.setName("group_brand_id");
	        agg2.setField("brand_id");
	        agg2.setSize(3);
			aggs = Arrays.asList(agg1, agg2);
	        searchRequest.setIndex("facetdemo2");        
	        searchRequest.setLimit(5);
	        query = new HashMap<String,Object>();
	        query.put("match_all", new HashMap<String,Object>());
	        searchRequest.setQuery(query);
	        searchRequest.setAggs(aggs);
	        //searchResponse = searchApi.search(searchRequest);
	        //System.out.println(searchResponse.toString() );  
	        
	        
	        searchRequest = new SearchRequest();
	        searchRequest.setIndex("books");
	        query = new HashMap<String,Object>();
	        query.put("match",new HashMap<String,Object>(){{
	            put("content","and first");
	        }});        
	        searchRequest.setQuery(query);
	        highlight = new Highlight();
	        highlight.setFieldnames( Arrays.asList("content","title") );
	        highlight.setHighlightQuery( 
            	new HashMap<String,Object>(){{
                	put("match", new HashMap<String,Object>(){{
                    	put("*","into three five");
                   	}});
                }}
			);
	        searchRequest.setHighlight(highlight);
	        searchResponse = searchApi.search(searchRequest);
	        
	        System.out.println("Tests finished");
	    } catch (ApiException e) {
	      System.err.println("Exception when calling IndexApi#bulk");
	      System.err.println("Status code: " + e.getCode());
	      System.err.println("Reason: " + e.getResponseBody());
	      System.err.println("Response headers: " + e.getResponseHeaders());
	      e.printStackTrace();
	    }
    }
}
