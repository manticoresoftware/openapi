/*
 * Manticore Search Client
 *
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 5.0.0
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

using ManticoreSearch.Client;
using ManticoreSearch.Api;
// uncomment below to import models
//using ManticoreSearch.Model;

namespace ManticoreSearch.Test.Api
{
    /// <summary>
    ///  Class for testing UtilsApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    public class UtilsApiTests : IDisposable
    {
        private UtilsApi instance;

        public UtilsApiTests()
        {
            instance = new UtilsApi();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of UtilsApi
        /// </summary>
        [Fact]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsType' UtilsApi
            //Assert.IsType<UtilsApi>(instance);
        }

        /// <summary>
        /// Test Sql
        /// </summary>
        [Fact]
        public void SqlTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string body = null;
            //bool? rawResponse = null;
            //var response = instance.Sql(body, rawResponse);
            //Assert.IsType<List<Object>>(response);
        }
    }
}
