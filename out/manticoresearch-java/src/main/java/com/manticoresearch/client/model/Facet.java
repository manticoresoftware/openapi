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
 * Query FACET expression
 */
@JsonPropertyOrder({
  Facet.JSON_PROPERTY_ATTR,
  Facet.JSON_PROPERTY_ALIAS,
  Facet.JSON_PROPERTY_LIMIT
})
@JsonTypeName("facet")
@JsonIgnoreProperties(ignoreUnknown = true)
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-05-28T11:55:17.809597Z[Etc/UTC]")
public class Facet {
  public static final String JSON_PROPERTY_ATTR = "attr";
  private String attr;

  public static final String JSON_PROPERTY_ALIAS = "alias";
  private String alias;

  public static final String JSON_PROPERTY_LIMIT = "limit";
  private Integer limit;

  public Facet() { 
  }

  public Facet attr(String attr) {
    this.attr = attr;
    return this;
  }

   /**
   * The name of an attribute to facet
   * @return attr
  **/
  @javax.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_ATTR)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getAttr() {
    return attr;
  }


  @JsonProperty(JSON_PROPERTY_ATTR)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setAttr(String attr) {
    this.attr = attr;
  }


  public Facet alias(String alias) {
    this.alias = alias;
    return this;
  }

   /**
   * Facet alias
   * @return alias
  **/
  @javax.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_ALIAS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getAlias() {
    return alias;
  }


  @JsonProperty(JSON_PROPERTY_ALIAS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setAlias(String alias) {
    this.alias = alias;
  }


  public Facet limit(Integer limit) {
    this.limit = limit;
    return this;
  }

   /**
   * The number of facet values to return
   * @return limit
  **/
  @javax.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_LIMIT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getLimit() {
    return limit;
  }


  @JsonProperty(JSON_PROPERTY_LIMIT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setLimit(Integer limit) {
    this.limit = limit;
  }


  /**
   * Return true if this facet object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Facet facet = (Facet) o;
    return Objects.equals(this.attr, facet.attr) &&
        Objects.equals(this.alias, facet.alias) &&
        Objects.equals(this.limit, facet.limit);
  }

  @Override
  public int hashCode() {
    return Objects.hash(attr, alias, limit);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Facet {\n");
    sb.append("    attr: ").append(toIndentedString(attr)).append("\n");
    sb.append("    alias: ").append(toIndentedString(alias)).append("\n");
    sb.append("    limit: ").append(toIndentedString(limit)).append("\n");
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
