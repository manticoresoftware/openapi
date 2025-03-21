/*
Manticore Search Client

Testing IndexAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"fmt"
	"encoding/json"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/manticoresoftware/manticoresearch-go"
)

func Test_openapi_IndexAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Servers[0].URL = "http://localhost:9408"
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IndexAPIService Bulk", func(t *testing.T) {

		apiClient.UtilsAPI.Sql(context.Background()).Body("DROP TABLE IF EXISTS test").Execute()
		apiClient.UtilsAPI.Sql(context.Background()).Body("CREATE TABLE IF NOT EXISTS test (body text, title string)").Execute()
		
   		body := "{\"insert\": {\"table\": \"test\", \"id\": 1, \"doc\": {\"body\": \"test\", \"title\": \"test\"}}}" + "\n"
   		resp, httpRes, err := apiClient.IndexAPI.Bulk(context.Background()).Body(body).Execute()
   		fmt.Printf("test %v\n", resp)
   		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService Delete", func(t *testing.T) {
		fmt.Println("Test 1");
		indexDoc1 := map[string]interface{} {"id": 1, "title": "Title 1"}
		indexReq1 := manticoreclient.NewInsertDocumentRequest("test1", indexDoc)
		apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();

		indexDoc1 = map[string]interface{} {"id": 2, "title": "Title 2"}
		indexReq1 = manticoreclient.NewInsertDocumentRequest("test1", indexDoc)
		apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();

		searchRequest1 := manticoreclient.NewSearchRequest("test1")
		query1 := map[string]interface{} {"query_string": "Title"};
		searchRequest.SetQuery(query1);

		res1, _, _ := apiClient.SearchAPI.Search(context.Background()).SearchRequest(*searchRequest1).Execute()
		outRes1, outErr1 := json.Marshal(res1)
    	require.Nil(t, outErr1)
    	//fmt.Printf("%+v\n", string(outRes1[:]))

		apiClient.UtilsAPI.Sql(context.Background()).Body("DROP TABLE IF EXISTS products").Execute()
        apiClient.UtilsAPI.Sql(context.Background()).Body("CREATE TABLE IF NOT EXISTS products (title text, price float, sizes multi, meta json, coeff float, tags1 multi, tags2 multi)").Execute()
        indexDoc := map[string]interface{} {"title": "first", "tags1": []int{4, 2, 1, 3}}
    	indexReq := openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(1)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();

    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(2)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();
    	
    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(0)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();
    	
    	delReq := openapiclient.NewDeleteDocumentRequest("products")
    	matchExpr := map[string]interface{} {"*": "first"}
    	delQuery := map[string]interface{} {"match": matchExpr }
    	delReq.SetQuery(delQuery)
    	
    	delResp, httpRes, err := apiClient.IndexAPI.Delete(context.Background()).DeleteDocumentRequest(*delReq).Execute();
    	
    	require.Nil(t, err)
		require.NotNil(t, delResp)
		assert.Equal(t, 200, httpRes.StatusCode)
		
		outRes, outErr := json.Marshal(delResp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))
	})

	t.Run("Test IndexAPIService Insert", func(t *testing.T) {

		apiClient.UtilsAPI.Sql(context.Background()).Body("DROP TABLE IF EXISTS products").Execute()
        apiClient.UtilsAPI.Sql(context.Background()).Body("CREATE TABLE IF NOT EXISTS products (title text, price float, sizes multi, meta json, coeff float, tags1 multi, tags2 multi)").Execute()
        indexDoc := map[string]interface{} {"title": "first", "tags1": []int{4, 2, 1, 3}}
    	indexReq := openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(1)
    	resp, httpRes, err := apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();

    	require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

    	outRes, outErr := json.Marshal(resp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))
		
		indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(2)
    	resp, httpRes, err = apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();
    	
    	require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		
		outRes, outErr = json.Marshal(resp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))
    	
    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(0)
    	resp, httpRes, err = apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();
    	
    	require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		
		outRes, outErr = json.Marshal(resp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))

	})

	t.Run("Test IndexAPIService Replace", func(t *testing.T) {

		apiClient.UtilsAPI.Sql(context.Background()).Body("DROP TABLE IF EXISTS products").Execute()
        apiClient.UtilsAPI.Sql(context.Background()).Body("CREATE TABLE IF NOT EXISTS products (title text, price float, sizes multi, meta json, coeff float, tags1 multi, tags2 multi)").Execute()
        indexDoc := map[string]interface{} {"title": "first", "tags1": []int{4, 2, 1, 3}}
    	indexReq := openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(1)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()

    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(2)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()
    	
    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(0)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()
		

		replDoc := map[string]interface{} {"title": "second", "tags1": []int{5}}
		replReq := openapiclient.NewInsertDocumentRequest("products", replDoc)
		replReq.SetId(1)
		resp, httpRes, err := apiClient.IndexAPI.Replace(context.Background()).InsertDocumentRequest(*replReq).Execute()
    	
    	require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		
		outRes, outErr := json.Marshal(resp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))
	})

	t.Run("Test IndexAPIService Update", func(t *testing.T) {

		apiClient.UtilsAPI.Sql(context.Background()).Body("DROP TABLE IF EXISTS products").Execute()
        apiClient.UtilsAPI.Sql(context.Background()).Body("CREATE TABLE IF NOT EXISTS products (title text, price float, sizes multi, meta json, coeff float, tags1 multi, tags2 multi)").Execute()
        indexDoc := map[string]interface{} {"title": "first", "tags1": []int{4, 2, 1, 3}}
    	indexReq := openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(1)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()

    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(2)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()
    	
    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(0)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()

		updDoc := map[string]interface{} {}
		updReq := openapiclient.NewUpdateDocumentRequest("products", updDoc)
		resp, httpRes, err := apiClient.IndexAPI.Update(context.Background()).UpdateDocumentRequest(*updReq).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

		outRes, outErr := json.Marshal(resp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))
    	
    	updDoc = map[string]interface{} {"price":10}
		updReq = openapiclient.NewUpdateDocumentRequest("products", updDoc)
		updReq.SetId(2)
    	resp, httpRes, err = apiClient.IndexAPI.Update(context.Background()).UpdateDocumentRequest(*updReq).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

		outRes, outErr = json.Marshal(resp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))
    	
	})

	fmt.Println("Index tests finished");
}
