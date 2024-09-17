

# RangeFilterLtx

## oneOf schemas
* [BigDecimal](BigDecimal.md)
* [Long](Long.md)
* [String](String.md)

## Example
```java
// Import classes:
import com.manticoresearch.client.model.RangeFilterLtx;
import com.manticoresearch.client.model.BigDecimal;
import com.manticoresearch.client.model.Long;
import com.manticoresearch.client.model.String;

public class Example {
    public static void main(String[] args) {
        RangeFilterLtx exampleRangeFilterLtx = new RangeFilterLtx();

        // create a new BigDecimal
        BigDecimal exampleBigDecimal = new BigDecimal();
        // set RangeFilterLtx to BigDecimal
        exampleRangeFilterLtx.setActualInstance(exampleBigDecimal);
        // to get back the BigDecimal set earlier
        BigDecimal testBigDecimal = (BigDecimal) exampleRangeFilterLtx.getActualInstance();

        // create a new Long
        Long exampleLong = new Long();
        // set RangeFilterLtx to Long
        exampleRangeFilterLtx.setActualInstance(exampleLong);
        // to get back the Long set earlier
        Long testLong = (Long) exampleRangeFilterLtx.getActualInstance();

        // create a new String
        String exampleString = new String();
        // set RangeFilterLtx to String
        exampleRangeFilterLtx.setActualInstance(exampleString);
        // to get back the String set earlier
        String testString = (String) exampleRangeFilterLtx.getActualInstance();
    }
}
```


