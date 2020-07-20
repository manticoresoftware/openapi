package com.manticore.clienttest;
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.model.*;
import org.openapitools.client.api.IndexApi;
import org.openapitools.client.api.UtilsApi;
/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:6368");


    IndexApi apiInstance = new IndexApi(defaultClient);
    UtilsApi apiUtils = new UtilsApi(defaultClient);
    String body = "{\"insert\": {\"index\": \"movies\", \"id\": 12121,\"doc\": {\"plot\": \"A secret team goes to North Pole\", \"rating\": 9.5, \"language\": [2, 3], \"title\": \"This is an older movie\", \"lon\": 51.99, \"meta\": {\"keywords\":[\"travel\",\"ice\"],\"genre\":[\"adventure\"]}, \"year\": 1950, \"lat\": 60.4, \"advise\": \"PG-13\"}}}" +"\n"; 
    try {
      String sql_body ="mode=raw&query=SHOW TABLES";
      Object sqlresult = apiUtils.sql(sql_body);
      System.out.println(sqlresult);
     // BulkResponse result = apiInstance.bulk(body);
      //System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IndexApi#bulk");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
    }
}
