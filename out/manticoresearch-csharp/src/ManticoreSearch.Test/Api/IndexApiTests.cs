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
    ///  Class for testing IndexApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    public class IndexApiTests : IDisposable
    {
    	private IndexApi instance;
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
            instance = new IndexApi(httpClient, config, httpClientHandler);
        }
                
        private object CheckTest(string testName)
        {
            Dictionary<string,Func<Object>> classTests;
            if (implementedTests.TryGetValue("IndexApi", out classTests))
            {
                Func<Object> test;    
                if (classTests.TryGetValue(testName, out test))
                {
                    return test();
                }
            }
            return null;
        }     
        
        public IndexApiTests()
        {
            implementedTests = new Dictionary<string, Dictionary<string,Func<Object>>>()
            {
                {
                "SearchApi", 
                    new Dictionary<string, Func<Object>>()
                    {
                    	{ "SearchTest", () => 
                            {
                            	var utilsApi = new UtilsApi();
                        		utilsApi.Sql("DROP TABLE IF EXISTS movies", true);
					        	utilsApi.Sql("CREATE TABLE IF NOT EXISTS movies (title text, plot text, year integer, rating float, code multi)", true);
					        	
					        	string[] docs = {
									"{\"insert\": {\"index\" : \"movies\", \"id\" : 1, \"doc\" : {\"title\" : \"Star Trek 2: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2002, \"rating\": 6.4, \"code\": [1,2,3]}}}",
						            "{\"insert\": {\"index\" : \"movies\", \"id\" : 2, \"doc\" : {\"title\" : \"Star Trek 1: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2001, \"rating\": 6.5, \"code\": [1,12,3]}}}",
						            "{\"insert\": {\"index\" : \"movies\", \"id\" : 3, \"doc\" : {\"title\" : \"Star Trek 3: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2003, \"rating\": 6.6, \"code\": [11,2,3]}}}",
						            "{\"insert\": {\"index\" : \"movies\", \"id\" : 4, \"doc\" : {\"title\" : \"Star Trek 4: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2003, \"rating\": 6.5, \"code\": [1,2,4]}}}"					        	
						        };
					        						        	
	                			var indexApi = new IndexApi(httpClient, config, httpClientHandler);
	            				var res = indexApi.Bulk(string.Join("\n", docs));
	            				
	            				var searchApi = new SearchApi();

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
        													
	        					var includes = new List<string> {"title", "year"};
        						var excludes = new List<string> {"code"};
        						searchRequest.Source = new SourceByRules(includes, excludes);

        						searchRes = searchApi.Search(searchRequest);

	        					searchRequest.Sort = new List<Object> {"year"};
	        					var sort2 = new SortOrder("rating", SortOrder.OrderEnum.Asc);
	        					searchRequest.Sort.Add(sort2);
	        					var sort3 = new SortMVA("code", SortMVA.OrderEnum.Desc, SortMVA.ModeEnum.Max);
	        					searchRequest.Sort.Add(sort3);

	        					searchRes = searchApi.Search(searchRequest);
	        					
	        					searchRequest.Expressions = new Dictionary<string, string> { {"expr", "min(year,2900)"} };
	        					includes.Add("expr");
					        	searchRequest.Expressions["expr2"] = "max(year,2100)";
					        	includes.Add("expr2");
					        	searchRequest.Source = new SourceByRules(includes, excludes);

					        	searchRes = searchApi.Search(searchRequest);
					        	
					        	var aggTerms = new AggregationTerms("year", 10);
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
					        	
					        	searchRequest.AttrFilter = new EqualsFilter("year", 2003);
					        	
								searchRes = searchApi.Search(searchRequest);

					        	var inFilter = new InFilter("year", new List<Object> {2001, 2002});
					        	var addValues = new List<Object> {10,11};
					    	    inFilter.Values.AddRange(addValues);
						        searchRequest.AttrFilter = inFilter;
	        
	        					searchRes = searchApi.Search(searchRequest);

	        					var rangeFilter = new RangeFilter("year");
								rangeFilter.Lte = 2002;
								rangeFilter.Gte = 0;
								searchRequest.AttrFilter = rangeFilter;

								searchRes = searchApi.Search(searchRequest);

								var geoFilter = new GeoDistanceFilter();
								var locAnchor = new GeoDistanceFilterLocationAnchor(10, 20);
								geoFilter.LocationAnchor = locAnchor;
								geoFilter.LocationSource = "year,rating";
								geoFilter.DistanceType = GeoDistanceFilter.DistanceTypeEnum.Adaptive;
								geoFilter.Distance = "100km";
        						searchRequest.AttrFilter = geoFilter;

								searchRes = searchApi.Search(searchRequest);

        						var boolFilter = new BoolFilter();
        						boolFilter.Must = new List<Object> { new EqualsFilter("year", 2001) };
        						rangeFilter = new RangeFilter("rating");
								rangeFilter.Lte = 20;
        						boolFilter.Must.Add(rangeFilter);
        						searchRequest.AttrFilter = boolFilter;

								searchRes = searchApi.Search(searchRequest);
        	
        						boolFilter.MustNot = new List<Object> { new EqualsFilter("year", 2001) };
								searchRequest.AttrFilter = boolFilter;

								searchRes = searchApi.Search(searchRequest);
								
								var fulltextFilter = new MatchFilter("Star", "title");
        						var nestedBoolFilter = new BoolFilter();
        						nestedBoolFilter.Should = new List<Object> { new EqualsFilter("rating", 6.5), fulltextFilter };
        						boolFilter.Must = new List<Object> {nestedBoolFilter};
            					searchRequest.AttrFilter = boolFilter;

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
            					var utilsApi = new UtilsApi();
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
            					var utilsApi = new UtilsApi();
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
        /// Test an instance of IndexApi
        /// </summary>
        [Fact]
        public void InstanceTest()
        {
            Assert.IsType<IndexApi>(instance);
        }

        /// <summary>
        /// Test Bulk
        /// </summary>
        [Fact]
        public void BulkTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string body = null;
            //var response = instance.Bulk(body);
			object response = this.CheckTest( System.Reflection.MethodBase.GetCurrentMethod().Name );
            if (response != null)
            {
            	Assert.IsType<BulkResponse>(response);
            }
        }

        /// <summary>
        /// Test Delete
        /// </summary>
        [Fact]
        public void DeleteTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //DeleteDocumentRequest deleteDocumentRequest = null;
            //var response = instance.Delete(deleteDocumentRequest);
			object response = this.CheckTest( System.Reflection.MethodBase.GetCurrentMethod().Name );
            if (response != null)
            {
            	Assert.IsType<DeleteResponse>(response);
            }
        }

        /// <summary>
        /// Test Insert
        /// </summary>
        [Fact]
        public void InsertTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //InsertDocumentRequest insertDocumentRequest = null;
            //var response = instance.Insert(insertDocumentRequest);
			object response = this.CheckTest( System.Reflection.MethodBase.GetCurrentMethod().Name );
            if (response != null)
            {
            	Assert.IsType<SuccessResponse>(response);
            }
        }

        /// <summary>
        /// Test Replace
        /// </summary>
        [Fact]
        public void ReplaceTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //InsertDocumentRequest insertDocumentRequest = null;
            //var response = instance.Replace(insertDocumentRequest);
			object response = this.CheckTest( System.Reflection.MethodBase.GetCurrentMethod().Name );
            if (response != null)
            {
            	Assert.IsType<SuccessResponse>(response);
            }
        }

        /// <summary>
        /// Test Update
        /// </summary>
        [Fact]
        public void UpdateTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //UpdateDocumentRequest updateDocumentRequest = null;
            //var response = instance.Update(updateDocumentRequest);
			object response = this.CheckTest( System.Reflection.MethodBase.GetCurrentMethod().Name );
            if (response != null)
            {
            	Assert.IsType<UpdateResponse>(response);
            }
        }

        /// <summary>
        /// Test Update_0
        /// </summary>
        [Fact]
        public void Update_0Test()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string index = null;
            //decimal id = null;
            //ReplaceDocumentRequest replaceDocumentRequest = null;
            //var response = instance.Update_0(index, id, replaceDocumentRequest);
			object response = this.CheckTest( System.Reflection.MethodBase.GetCurrentMethod().Name );
            if (response != null)
            {
            	Assert.IsType<UpdateResponse>(response);
            }
        }
    }
}
