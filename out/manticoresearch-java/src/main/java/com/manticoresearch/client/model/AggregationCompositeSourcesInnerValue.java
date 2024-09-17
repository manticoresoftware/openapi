/*
 * Manticore Search Client
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 3.3.1
 * Contact: info@manticoresearch.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.manticoresearch.client.model;

import java.util.Objects;
import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import com.manticoresearch.client.model.AggregationCompositeSourcesInnerValueTerms;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.manticoresearch.client.JSON;


/**
 * AggregationCompositeSourcesInnerValue
 */
@JsonPropertyOrder({
  AggregationCompositeSourcesInnerValue.JSON_PROPERTY_TERMS
})
@JsonTypeName("aggregation_composite_sources_inner_value")
@JsonIgnoreProperties(ignoreUnknown = true)
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-08-07T13:45:53.763550451Z[Etc/UTC]")
public class AggregationCompositeSourcesInnerValue {
  public static final String JSON_PROPERTY_TERMS = "terms";
  private AggregationCompositeSourcesInnerValueTerms terms;

  public AggregationCompositeSourcesInnerValue() { 
  }

  public AggregationCompositeSourcesInnerValue terms(AggregationCompositeSourcesInnerValueTerms terms) {
    this.terms = terms;
    return this;
  }

   /**
   * Get terms
   * @return terms
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_TERMS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public AggregationCompositeSourcesInnerValueTerms getTerms() {
    return terms;
  }


  @JsonProperty(JSON_PROPERTY_TERMS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setTerms(AggregationCompositeSourcesInnerValueTerms terms) {
    this.terms = terms;
  }


  /**
   * Return true if this aggregation_composite_sources_inner_value object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AggregationCompositeSourcesInnerValue aggregationCompositeSourcesInnerValue = (AggregationCompositeSourcesInnerValue) o;
    return Objects.equals(this.terms, aggregationCompositeSourcesInnerValue.terms);
  }

  @Override
  public int hashCode() {
    return Objects.hash(terms);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AggregationCompositeSourcesInnerValue {\n");
    sb.append("    terms: ").append(toIndentedString(terms)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

