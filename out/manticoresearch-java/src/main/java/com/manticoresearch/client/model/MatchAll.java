/*
 * Manticore Search Client
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 5.0.0
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
import com.manticoresearch.client.JSON;


/**
 * Filter helper object defining the &#39;match all&#39;&#39; condition
 */
@JsonPropertyOrder({
  MatchAll.JSON_PROPERTY_ALL
})
@JsonTypeName("match_all")
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2025-06-30T07:16:18.467791800Z[Etc/UTC]", comments = "Generator version: 7.14.0")
public class MatchAll {
  /**
   * Gets or Sets all
   */
  public enum AllEnum {
    u("{}");

    private String value;

    AllEnum(String value) {
      this.value = value;
    }

    @JsonValue
    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    @JsonCreator
    public static AllEnum fromValue(String value) {
      for (AllEnum b : AllEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_ALL = "_all";
  private AllEnum all;

  public MatchAll() { 
  }

  public MatchAll all(AllEnum all) {
    this.all = all;
    return this;
  }

  /**
   * Get all
   * @return all
   */
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_ALL)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public AllEnum getAll() {
    return all;
  }


  @JsonProperty(JSON_PROPERTY_ALL)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setAll(AllEnum all) {
    this.all = all;
  }


  /**
   * Return true if this match_all object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MatchAll matchAll = (MatchAll) o;
    return Objects.equals(this.all, matchAll.all);
  }

  @Override
  public int hashCode() {
    return Objects.hash(all);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MatchAll {\n");
    sb.append("    all: ").append(toIndentedString(all)).append("\n");
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

