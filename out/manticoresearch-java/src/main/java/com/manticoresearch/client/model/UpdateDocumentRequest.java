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
import com.manticoresearch.client.model.QueryFilter;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import com.fasterxml.jackson.annotation.JsonIgnore;
import org.openapitools.jackson.nullable.JsonNullable;
import java.util.NoSuchElementException;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.manticoresearch.client.JSON;


/**
 * Payload for updating a document or multiple documents in an index
 */
@JsonPropertyOrder({
  UpdateDocumentRequest.JSON_PROPERTY_INDEX,
  UpdateDocumentRequest.JSON_PROPERTY_CLUSTER,
  UpdateDocumentRequest.JSON_PROPERTY_DOC,
  UpdateDocumentRequest.JSON_PROPERTY_ID,
  UpdateDocumentRequest.JSON_PROPERTY_QUERY
})
@JsonTypeName("updateDocumentRequest")
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-10-21T07:19:33.210051324Z[Etc/UTC]", comments = "Generator version: 7.3.0-SNAPSHOT")
public class UpdateDocumentRequest {
  public static final String JSON_PROPERTY_INDEX = "index";
  private String index;

  public static final String JSON_PROPERTY_CLUSTER = "cluster";
  private String cluster;

  public static final String JSON_PROPERTY_DOC = "doc";
  private Object doc;

  public static final String JSON_PROPERTY_ID = "id";
  private Long id;

  public static final String JSON_PROPERTY_QUERY = "query";
  private JsonNullable<QueryFilter> query = JsonNullable.<QueryFilter>undefined();

  public UpdateDocumentRequest() { 
  }

  public UpdateDocumentRequest index(String index) {
    this.index = index;
    return this;
  }

  /**
   * Name of the document index
   * @return index
   */
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_INDEX)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getIndex() {
    return index;
  }


  @JsonProperty(JSON_PROPERTY_INDEX)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setIndex(String index) {
    this.index = index;
  }


  public UpdateDocumentRequest cluster(String cluster) {
    this.cluster = cluster;
    return this;
  }

  /**
   * Name of the document cluster
   * @return cluster
   */
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_CLUSTER)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getCluster() {
    return cluster;
  }


  @JsonProperty(JSON_PROPERTY_CLUSTER)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setCluster(String cluster) {
    this.cluster = cluster;
  }


  public UpdateDocumentRequest doc(Object doc) {
    this.doc = doc;
    return this;
  }

  /**
   * Object containing the document fields to update
   * @return doc
   */
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_DOC)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public Object getDoc() {
    return doc;
  }


  @JsonProperty(JSON_PROPERTY_DOC)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setDoc(Object doc) {
    this.doc = doc;
  }


  public UpdateDocumentRequest id(Long id) {
    this.id = id;
    return this;
  }

  /**
   * Document ID
   * @return id
   */
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_ID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Long getId() {
    return id;
  }


  @JsonProperty(JSON_PROPERTY_ID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setId(Long id) {
    this.id = id;
  }


  public UpdateDocumentRequest query(QueryFilter query) {
    this.query = JsonNullable.<QueryFilter>of(query);
    return this;
  }

  /**
   * Get query
   * @return query
   */
  @jakarta.annotation.Nullable
  @JsonIgnore

  public QueryFilter getQuery() {
        return query.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_QUERY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<QueryFilter> getQuery_JsonNullable() {
    return query;
  }
  
  @JsonProperty(JSON_PROPERTY_QUERY)
  public void setQuery_JsonNullable(JsonNullable<QueryFilter> query) {
    this.query = query;
  }

  public void setQuery(QueryFilter query) {
    this.query = JsonNullable.<QueryFilter>of(query);
  }


  /**
   * Return true if this updateDocumentRequest object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UpdateDocumentRequest updateDocumentRequest = (UpdateDocumentRequest) o;
    return Objects.equals(this.index, updateDocumentRequest.index) &&
        Objects.equals(this.cluster, updateDocumentRequest.cluster) &&
        Objects.equals(this.doc, updateDocumentRequest.doc) &&
        Objects.equals(this.id, updateDocumentRequest.id) &&
        equalsNullable(this.query, updateDocumentRequest.query);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(index, cluster, doc, id, hashCodeNullable(query));
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateDocumentRequest {\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    cluster: ").append(toIndentedString(cluster)).append("\n");
    sb.append("    doc: ").append(toIndentedString(doc)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    query: ").append(toIndentedString(query)).append("\n");
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

