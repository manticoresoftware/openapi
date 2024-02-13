

# SearchRequest

## oneOf schemas
* [KnnSearchRequestByDocId](KnnSearchRequestByDocId.md)
* [KnnSearchRequestByVector](KnnSearchRequestByVector.md)
* [SearchRequest](SearchRequest.md)

## Example
```java
// Import classes:
import com.manticoresearch.client.model.SearchRequest;
import com.manticoresearch.client.model.KnnSearchRequestByDocId;
import com.manticoresearch.client.model.KnnSearchRequestByVector;
import com.manticoresearch.client.model.SearchRequest;

public class Example {
    public static void main(String[] args) {
        SearchRequest exampleSearchRequest = new SearchRequest();

        // create a new KnnSearchRequestByDocId
        KnnSearchRequestByDocId exampleKnnSearchRequestByDocId = new KnnSearchRequestByDocId();
        // set SearchRequest to KnnSearchRequestByDocId
        exampleSearchRequest.setActualInstance(exampleKnnSearchRequestByDocId);
        // to get back the KnnSearchRequestByDocId set earlier
        KnnSearchRequestByDocId testKnnSearchRequestByDocId = (KnnSearchRequestByDocId) exampleSearchRequest.getActualInstance();

        // create a new KnnSearchRequestByVector
        KnnSearchRequestByVector exampleKnnSearchRequestByVector = new KnnSearchRequestByVector();
        // set SearchRequest to KnnSearchRequestByVector
        exampleSearchRequest.setActualInstance(exampleKnnSearchRequestByVector);
        // to get back the KnnSearchRequestByVector set earlier
        KnnSearchRequestByVector testKnnSearchRequestByVector = (KnnSearchRequestByVector) exampleSearchRequest.getActualInstance();

        // create a new SearchRequest
        SearchRequest exampleSearchRequest = new SearchRequest();
        // set SearchRequest to SearchRequest
        exampleSearchRequest.setActualInstance(exampleSearchRequest);
        // to get back the SearchRequest set earlier
        SearchRequest testSearchRequest = (SearchRequest) exampleSearchRequest.getActualInstance();
    }
}
```


