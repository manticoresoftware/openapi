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
 * Query match expression
 */
@JsonPropertyOrder({
  MatchPhraseFilter.JSON_PROPERTY_QUERY_PHRASE,
  MatchPhraseFilter.JSON_PROPERTY_QUERY_FIELDS
})
@JsonTypeName("matchPhraseFilter")
@JsonIgnoreProperties(ignoreUnknown = true)
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-12-18T11:24:55.908019234Z[Etc/UTC]")
public class MatchPhraseFilter {
  public static final String JSON_PROPERTY_QUERY_PHRASE = "query_phrase";
  private String queryPhrase;

  public static final String JSON_PROPERTY_QUERY_FIELDS = "query_fields";
  private String queryFields;

  public MatchPhraseFilter() { 
  }

  public MatchPhraseFilter queryPhrase(String queryPhrase) {
    this.queryPhrase = queryPhrase;
    return this;
  }

   /**
   * Get queryPhrase
   * @return queryPhrase
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_QUERY_PHRASE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getQueryPhrase() {
    return queryPhrase;
  }


  @JsonProperty(JSON_PROPERTY_QUERY_PHRASE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setQueryPhrase(String queryPhrase) {
    this.queryPhrase = queryPhrase;
  }


  public MatchPhraseFilter queryFields(String queryFields) {
    this.queryFields = queryFields;
    return this;
  }

   /**
   * Get queryFields
   * @return queryFields
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_QUERY_FIELDS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getQueryFields() {
    return queryFields;
  }


  @JsonProperty(JSON_PROPERTY_QUERY_FIELDS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setQueryFields(String queryFields) {
    this.queryFields = queryFields;
  }


  /**
   * Return true if this matchPhraseFilter object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MatchPhraseFilter matchPhraseFilter = (MatchPhraseFilter) o;
    return Objects.equals(this.queryPhrase, matchPhraseFilter.queryPhrase) &&
        Objects.equals(this.queryFields, matchPhraseFilter.queryFields);
  }

  @Override
  public int hashCode() {
    return Objects.hash(queryPhrase, queryFields);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MatchPhraseFilter {\n");
    sb.append("    queryPhrase: ").append(toIndentedString(queryPhrase)).append("\n");
    sb.append("    queryFields: ").append(toIndentedString(queryFields)).append("\n");
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

