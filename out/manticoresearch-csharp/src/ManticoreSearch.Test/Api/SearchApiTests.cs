/*
 * Manticore Search Client
 *
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 3.3.1
 * Contact: info@manticoresearch.com
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */

using System;
using System.IO;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Reflection;
using Xunit;
using System.Net.Http;
using System.Text.Json;

using ManticoreSearch.Client;
using ManticoreSearch.Api;
// uncomment below to import models
using ManticoreSearch.Model;

namespace ManticoreSearch.Test.Api
{
    /// <summary>
    ///  Class for testing SearchApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    public class SearchApiTests : IDisposable
    {
    	private SearchApi instance;
        private HttpClientHandler httpClientHandler;
        private HttpClient httpClient;
        private Configuration config;

        private Dictionary<string, Dictionary<string,Func<Object>>> implementedTests;

        private void InitTests()
        {
            config = new Configuration();
            config.BasePath = "http://localhost:9408";
            httpClient = new HttpClient();
            httpClientHandler = new HttpClientHandler();
            instance = new SearchApi(httpClient, config, httpClientHandler);
        }
                
        private object CheckTest(string testName)
        {
            Dictionary<string,Func<Object>> classTests;
            if (implementedTests.TryGetValue("SearchApi", out classTests))
            {
                Func<Object> test;    
                if (classTests.TryGetValue(testName, out test))
                {
                    return test();
                }
            }
            return null;
        }     
        
        public SearchApiTests()
        {
            implementedTests = new Dictionary<string, Dictionary<string,Func<Object>>>()
            {
                {
                "SearchApi", 
                    new Dictionary<string, Func<Object>>()
                    {
                    	{ "SearchTest", () => 
                            {
                            	var utilsApi = new UtilsApi(httpClient, config, httpClientHandler);
                        		utilsApi.Sql("DROP TABLE IF EXISTS movies", true);
					        	utilsApi.Sql("CREATE TABLE IF NOT EXISTS movies (title text, plot text, _year integer, rating float, code multi, type_vector float_vector knn_type='hnsw' knn_dims='3' hnsw_similarity='l2')", true);
					        	
					        	string[] docs = {
						            "{\"insert\": {\"index\" : \"movies\", \"id\" : 1, \"doc\" : {\"title\" : \"Star Trek 2: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"_year\": 2002, \"rating\": 6.4, \"code\": [1,2,3], \"type_vector\": [0.2, 1.4, -2.3]}}}",
				        			"{\"insert\": {\"index\" : \"movies\", \"id\" : 2, \"doc\" : {\"title\" : \"Star Trek 1: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"_year\": 2001, \"rating\": 6.5, \"code\": [1,12,3], \"type_vector\": [0.8, 0.4, 1.3]}}}",
				        			"{\"insert\": {\"index\" : \"movies\", \"id\" : 3, \"doc\" : {\"title\" : \"Star Trek 3: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"_year\": 2003, \"rating\": 6.6, \"code\": [11,2,3], \"type_vector\": [1.5, -1.0, 1.6]}}}",
				        			"{\"insert\": {\"index\" : \"movies\", \"id\" : 4, \"doc\" : {\"title\" : \"Star Trek 4: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"_year\": 2003, \"rating\": 6, \"code\": [1,2,4], \"type_vector\": [0.4, 2.4, 0.9]}}}"					        	
						        };
					        						        	
	                			var indexApi = new IndexApi(httpClient, config, httpClientHandler);
	            				var res = indexApi.Bulk(string.Join("\n", docs));
	            					            				
	            				var searchApi = new SearchApi(httpClient, config, httpClientHandler);

		            			var searchRequest = new SearchRequest("movies");
		            			var searchRes = searchApi.Search(searchRequest);

								object query =  new { query_string="Star" };
								searchRequest.Query = query;
								searchRequest.Limit = 10;

								searchRes = searchApi.Search(searchRequest);

								searchRequest.Options = new Dictionary<string, Object>();
								searchRequest.Options["cutoff"] = 5;
								searchRequest.Options["ranker"] = "bm25";
        						searchRequest.Source = "title";

        						searchRes = searchApi.Search(searchRequest);
        													
	        					var includes = new List<string> {"title", "_year"};
        						var excludes = new List<string> {"code"};
        						searchRequest.Source = new SourceByRules(includes, excludes);

        						searchRes = searchApi.Search(searchRequest);

	        					searchRequest.Sort = new List<Object> {"_year"};
	        					var sort2 = new SortOrder("rating", SortOrder.OrderEnum.Asc);
	        					searchRequest.Sort.Add(sort2);
	        					var sort3 = new SortMVA("code", SortMVA.OrderEnum.Desc, SortMVA.ModeEnum.Max);
	        					searchRequest.Sort.Add(sort3);

	        					searchRes = searchApi.Search(searchRequest);
	        					
	        					searchRequest.Expressions = new Dictionary<string, string> { {"expr", "min(_year,2900)"} };
	        					includes.Add("expr");
					        	searchRequest.Expressions["expr2"] = "max(_year,2100)";
					        	includes.Add("expr2");
					        	searchRequest.Source = new SourceByRules(includes, excludes);

					        	searchRes = searchApi.Search(searchRequest);
					        	
					        	var aggTerms = new AggregationTerms("_year", 10);
					        	var agg1 = new Aggregation(aggTerms);
        						searchRequest.Aggs = new Dictionary<string, Aggregation> { {"agg1", agg1} };
        						aggTerms = new AggregationTerms("rating"); 
        						var sortExpr = new AggregationSortInnerValue("asc");
        						var sort = new Dictionary<string, AggregationSortInnerValue> { { "rating", sortExpr} };
        						var aggSort = new List<Dictionary<string, AggregationSortInnerValue>> {sort};
        						searchRequest.Aggs["agg2"] = new Aggregation(aggTerms, aggSort);
	        					
        						searchRes = searchApi.Search(searchRequest);
					        	
					        	var highlight = new Highlight();
					        	highlight.Fieldnames = new List<string> {"title"};
 					        	highlight.PostTags = "</post_tag>";
 					    	    highlight.Encoder = Highlight.EncoderEnum.Default;
 						        highlight.SnippetBoundary = Highlight.SnippetBoundaryEnum.Sentence;
 					        	searchRequest.Highlight = highlight;

 					        	searchRes = searchApi.Search(searchRequest);
        	
					        	var highlightField = new HighlightField("title");
								highlightField.Limit = 5;
								highlight.Fields = new List<HighlightField> {highlightField};

								searchRes = searchApi.Search(searchRequest);
								
								var highlightField2 = new HighlightField("plot");
								highlightField2.LimitWords = 10;
					        	highlight.Fields.Add(highlightField2);
					        	searchRequest.Highlight = highlight;

					        	searchRes = searchApi.Search(searchRequest);

								highlight.HighlightQuery = new Dictionary<string, Object> {{ "match",  new Dictionary<string, Object> { { "*", "Star"} } }};

								searchRes = searchApi.Search(searchRequest);

					        	searchRequest.FulltextFilter = new QueryFilter("Star Trek 2");

					        	searchRes = searchApi.Search(searchRequest);

					        	searchRequest.FulltextFilter = new MatchFilter("Nemesis", "title");

					        	searchRes = searchApi.Search(searchRequest);

					        	searchRequest.FulltextFilter = new MatchPhraseFilter("Star Trek 2", "title");

					        	searchRes = searchApi.Search(searchRequest);

					        	searchRequest.FulltextFilter = new MatchOpFilter("Enterprise test", "title,plot", MatchOpFilter.OperatorEnum.Or);

					        	searchRes = searchApi.Search(searchRequest);
					        	
					        	searchRequest.AttrFilter = new EqualsFilter("_year", 2003);
					        	
								searchRes = searchApi.Search(searchRequest);

					        	var inFilter = new InFilter("_year", new List<Object> {2001, 2002});
					        	var addValues = new List<Object> {10,11};
					    	    inFilter.Values.AddRange(addValues);
						        searchRequest.AttrFilter = inFilter;
	        
	        					searchRes = searchApi.Search(searchRequest);

	        					var rangeFilter = new RangeFilter("_year");
								rangeFilter.Lte = 2002;
								rangeFilter.Gte = 0;
								searchRequest.AttrFilter = rangeFilter;

								searchRes = searchApi.Search(searchRequest);

								var geoFilter = new GeoDistanceFilter();
								var locAnchor = new GeoDistanceFilterLocationAnchor(10, 20);
								geoFilter.LocationAnchor = locAnchor;
								geoFilter.LocationSource = "_year,rating";
								geoFilter.DistanceType = GeoDistanceFilter.DistanceTypeEnum.Adaptive;
								geoFilter.Distance = "100km";
        						searchRequest.AttrFilter = geoFilter;

								searchRes = searchApi.Search(searchRequest);

        						var boolFilter = new BoolFilter();
        						boolFilter.Must = new List<Object> { new EqualsFilter("_year", 2001) };
        						rangeFilter = new RangeFilter("rating");
								rangeFilter.Lte = 20;
        						boolFilter.Must.Add(rangeFilter);
        						searchRequest.AttrFilter = boolFilter;

								searchRes = searchApi.Search(searchRequest);
        	
        						boolFilter.MustNot = new List<Object> { new EqualsFilter("_year", 2001) };
								searchRequest.AttrFilter = boolFilter;

								searchRes = searchApi.Search(searchRequest);
								
								var fulltextFilter = new MatchFilter("Star", "title");
        						var nestedBoolFilter = new BoolFilter();
        						nestedBoolFilter.Should = new List<Object> { new EqualsFilter("rating", 6.5), fulltextFilter };
        						boolFilter.Must = new List<Object> {nestedBoolFilter};
            					searchRequest.AttrFilter = boolFilter;

								searchRes = searchApi.Search(searchRequest);
        						
        						searchRequest = new SearchRequest("movies");
        						var knnQueryByVector = new KnnQueryByVector("type_vector", new List<Decimal> {(decimal)1.5, (decimal)-1.0, (decimal)1.6}, 5);
        						var searchRequestKnn = new SearchRequestKnn(knnQueryByVector);

        						searchRequest.Knn = searchRequestKnn;
        						searchRes = searchApi.Search(searchRequest);


								var knnQueryByDocId = new KnnQueryByDocId("type_vector", 2, 5);
        						searchRequestKnn = new SearchRequestKnn(knnQueryByDocId);

        						searchRequest.Knn = searchRequestKnn;
        						searchRes = searchApi.Search(searchRequest);
								
								var knnBoolFilter = new Dictionary<string, Object> {{ 
									"bool",  
									new Dictionary<string, Object> {{ 
										"must", 
										new List<Object> {
											new Dictionary<string, Object> {{ 
												"equals", 
												new Dictionary<string, Object> {{ "id", 3}}
											}}
										}
									}} 
								}};
								
								knnQueryByVector.Filter = knnBoolFilter;
        						searchRequestKnn = new SearchRequestKnn(knnQueryByVector);

        						searchRequest.Knn = searchRequestKnn;
        						searchRes = searchApi.Search(searchRequest);

								knnQueryByDocId.Filter = knnBoolFilter;
        						searchRequestKnn = new SearchRequestKnn(knnQueryByDocId);

        						searchRequest.Knn = searchRequestKnn;
        						searchRes = searchApi.Search(searchRequest);


                                return searchRes;
                            }
                        },
                    }
                 },
                 {
                 "IndexApi", 
                    new Dictionary<string, Func<Object>>()
                    {
                        { "InsertTest", () => 
                            {
                            	string body = "CREATE TABLE IF NOT EXISTS test (body text, title string)";
            					var utilsApi = new UtilsApi(httpClient, config, httpClientHandler);
            					utilsApi.Sql(body, true);
                                Dictionary<string, Object> doc = new Dictionary<string, Object>(); 
                                doc.Add("body", "test");
                                doc.Add("title", "test");
                                InsertDocumentRequest insertDocumentRequest = new InsertDocumentRequest(index: "test", id: 1, doc: doc);
                                insertDocumentRequest = new InsertDocumentRequest(index: "test", id: 2, doc: doc);
                                var obj = new IndexApi(httpClient, config, httpClientHandler);
                                return obj.Insert(insertDocumentRequest);
                            }
                        },
                        { "BulkTest", () => 
		                	{
		                		string body = "CREATE TABLE IF NOT EXISTS test (body text, title string)";
            					var utilsApi = new UtilsApi(httpClient, config, httpClientHandler);
            					utilsApi.Sql("DROP TABLE IF EXISTS test", true);
            					utilsApi.Sql(body, true);
		                		body = "{\"insert\": {\"index\": \"test\", \"id\": 1, \"doc\": {\"body\": \"test\", \"title\": \"test\"}}}" + "\n";
		                		var obj = new IndexApi(httpClient, config, httpClientHandler);
		            			return obj.Bulk(body);
		                	}
		                },
		                { "ReplaceTest", () => 
		                	{
								Dictionary<string, Object> doc = new Dictionary<string, Object>(); 
		            			doc.Add("body", "test 2");
		            			doc.Add("title", "test");
		            			InsertDocumentRequest insertDocumentRequest = new InsertDocumentRequest(index: "test", id: 1, doc: doc);
		            			var obj = new IndexApi(httpClient, config, httpClientHandler);
		            			return obj.Replace(insertDocumentRequest);
		                	}
		                },
		                { "UpdateTest", () => 
		                	{
								Dictionary<string, Object> doc = new Dictionary<string, Object>();
					            doc.Add("title", "test 2");
					            UpdateDocumentRequest updateDocumentRequest = new UpdateDocumentRequest(index: "test", id: 2, doc: doc);
					            var obj = new IndexApi(httpClient, config, httpClientHandler);
					            return obj.Update(updateDocumentRequest);
		                	}
		                },
		                { "DeleteTest", () => 
		                	{
								DeleteDocumentRequest deleteDocumentRequest = new DeleteDocumentRequest(index: "test", id: 1);
								var obj = new IndexApi(httpClient, config, httpClientHandler);
		            			return obj.Delete(deleteDocumentRequest);
		                	}
		                },
                    }
                }
            };

            InitTests();
            
        }

        public void Dispose()
        {
        }

        /// <summary>
        /// Test an instance of SearchApi
        /// </summary>
        [Fact]
        public void InstanceTest()
        {
            Assert.IsType<SearchApi>(instance);
        }

        /// <summary>
        /// Test Percolate
        /// </summary>
        [Fact]
        public void PercolateTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string index = null;
            //PercolateRequest percolateRequest = null;
            //var response = instance.Percolate(index, percolateRequest);
			object response = this.CheckTest( System.Reflection.MethodBase.GetCurrentMethod().Name );
            if (response != null)
            {
            	Assert.IsType<SearchResponse>(response);
            }
        }

        /// <summary>
        /// Test Search
        /// </summary>
        [Fact]
        public void SearchTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //SearchRequest searchRequest = null;
            //var response = instance.Search(searchRequest);
			object response = this.CheckTest( System.Reflection.MethodBase.GetCurrentMethod().Name );
            if (response != null)
            {
            	Assert.IsType<SearchResponse>(response);
            }
        }
    }
}
