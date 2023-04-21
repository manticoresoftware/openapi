/*
 * Manticore Search Client
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 3.3.0
 * Contact: info@manticoresearch.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.manticoresearch.client.model;

import java.util.Objects;
import java.util.Arrays;
import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.manticoresearch.client.JSON;


/**
 * Query match expression with logical operator
 */
@JsonPropertyOrder({
  MatchOp.JSON_PROPERTY_QUERY_INFO
})
@JsonTypeName("matchOp")
@JsonIgnoreProperties(ignoreUnknown = true)
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-04-21T16:12:17.895346Z[Etc/UTC]")
public class MatchOp {
  public static final String JSON_PROPERTY_QUERY_INFO = "query_info";
  private Object queryInfo;

  public MatchOp() { 
  }

  public MatchOp queryInfo(Object queryInfo) {
    this.queryInfo = queryInfo;
    return this;
  }

   /**
   * Get queryInfo
   * @return queryInfo
  **/
  @javax.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_QUERY_INFO)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public Object getQueryInfo() {
    return queryInfo;
  }


  @JsonProperty(JSON_PROPERTY_QUERY_INFO)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setQueryInfo(Object queryInfo) {
    this.queryInfo = queryInfo;
  }


  /**
   * Return true if this matchOp object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MatchOp matchOp = (MatchOp) o;
    return Objects.equals(this.queryInfo, matchOp.queryInfo);
  }

  @Override
  public int hashCode() {
    return Objects.hash(queryInfo);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MatchOp {\n");
    sb.append("    queryInfo: ").append(toIndentedString(queryInfo)).append("\n");
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

