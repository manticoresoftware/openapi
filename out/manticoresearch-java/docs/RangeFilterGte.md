

# RangeFilterGte

## oneOf schemas
* [BigDecimal](BigDecimal.md)
* [Long](Long.md)
* [String](String.md)

NOTE: this class is nullable.

## Example
```java
// Import classes:
import com.manticoresearch.client.model.RangeFilterGte;
import com.manticoresearch.client.model.BigDecimal;
import com.manticoresearch.client.model.Long;
import com.manticoresearch.client.model.String;

public class Example {
    public static void main(String[] args) {
        RangeFilterGte exampleRangeFilterGte = new RangeFilterGte();

        // create a new BigDecimal
        BigDecimal exampleBigDecimal = new BigDecimal();
        // set RangeFilterGte to BigDecimal
        exampleRangeFilterGte.setActualInstance(exampleBigDecimal);
        // to get back the BigDecimal set earlier
        BigDecimal testBigDecimal = (BigDecimal) exampleRangeFilterGte.getActualInstance();

        // create a new Long
        Long exampleLong = new Long();
        // set RangeFilterGte to Long
        exampleRangeFilterGte.setActualInstance(exampleLong);
        // to get back the Long set earlier
        Long testLong = (Long) exampleRangeFilterGte.getActualInstance();

        // create a new String
        String exampleString = new String();
        // set RangeFilterGte to String
        exampleRangeFilterGte.setActualInstance(exampleString);
        // to get back the String set earlier
        String testString = (String) exampleRangeFilterGte.getActualInstance();
    }
}
```


