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
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.manticoresearch.client.JSON;


/**
 * AggregationCompositeSourcesInnerValueTerms
 */
@JsonPropertyOrder({
  AggregationCompositeSourcesInnerValueTerms.JSON_PROPERTY_FIELD
})
@JsonTypeName("aggregation_composite_sources_inner_value_terms")
@JsonIgnoreProperties(ignoreUnknown = true)
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-08-07T13:45:53.763550451Z[Etc/UTC]")
public class AggregationCompositeSourcesInnerValueTerms {
  public static final String JSON_PROPERTY_FIELD = "field";
  private String field;

  public AggregationCompositeSourcesInnerValueTerms() { 
  }

  public AggregationCompositeSourcesInnerValueTerms field(String field) {
    this.field = field;
    return this;
  }

   /**
   * Name of attribute to aggregate by
   * @return field
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_FIELD)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getField() {
    return field;
  }


  @JsonProperty(JSON_PROPERTY_FIELD)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setField(String field) {
    this.field = field;
  }


  /**
   * Return true if this aggregation_composite_sources_inner_value_terms object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AggregationCompositeSourcesInnerValueTerms aggregationCompositeSourcesInnerValueTerms = (AggregationCompositeSourcesInnerValueTerms) o;
    return Objects.equals(this.field, aggregationCompositeSourcesInnerValueTerms.field);
  }

  @Override
  public int hashCode() {
    return Objects.hash(field);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AggregationCompositeSourcesInnerValueTerms {\n");
    sb.append("    field: ").append(toIndentedString(field)).append("\n");
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
