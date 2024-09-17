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
import com.manticoresearch.client.model.Aggregation;
import com.manticoresearch.client.model.Highlight;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.manticoresearch.client.JSON;


/**
 * Request object for search operation
 */
@JsonPropertyOrder({
  SearchRequest.JSON_PROPERTY_INDEX,
  SearchRequest.JSON_PROPERTY_QUERY,
  SearchRequest.JSON_PROPERTY_FULLTEXT_FILTER,
  SearchRequest.JSON_PROPERTY_ATTR_FILTER,
  SearchRequest.JSON_PROPERTY_LIMIT,
  SearchRequest.JSON_PROPERTY_OFFSET,
  SearchRequest.JSON_PROPERTY_MAX_MATCHES,
  SearchRequest.JSON_PROPERTY_SORT,
  SearchRequest.JSON_PROPERTY_AGGS,
  SearchRequest.JSON_PROPERTY_EXPRESSIONS,
  SearchRequest.JSON_PROPERTY_HIGHLIGHT,
  SearchRequest.JSON_PROPERTY_SOURCE,
  SearchRequest.JSON_PROPERTY_OPTIONS,
  SearchRequest.JSON_PROPERTY_PROFILE,
  SearchRequest.JSON_PROPERTY_TRACK_SCORES
})
@JsonTypeName("searchRequest")
@JsonIgnoreProperties(ignoreUnknown = true)
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-08-07T13:45:53.763550451Z[Etc/UTC]")
public class SearchRequest {
  public static final String JSON_PROPERTY_INDEX = "index";
  private String index = "";

  public static final String JSON_PROPERTY_QUERY = "query";
  private Object query;

  public static final String JSON_PROPERTY_FULLTEXT_FILTER = "fulltext_filter";
  private Object fulltextFilter;

  public static final String JSON_PROPERTY_ATTR_FILTER = "attr_filter";
  private Object attrFilter;

  public static final String JSON_PROPERTY_LIMIT = "limit";
  private Integer limit;

  public static final String JSON_PROPERTY_OFFSET = "offset";
  private Integer offset;

  public static final String JSON_PROPERTY_MAX_MATCHES = "max_matches";
  private Integer maxMatches;

  public static final String JSON_PROPERTY_SORT = "sort";
  private List<Object> sort;

  public static final String JSON_PROPERTY_AGGS = "aggs";
  private Map<String, Aggregation> aggs = new HashMap<>();

  public static final String JSON_PROPERTY_EXPRESSIONS = "expressions";
  private Map<String, String> expressions = new HashMap<>();

  public static final String JSON_PROPERTY_HIGHLIGHT = "highlight";
  private Highlight highlight;

  public static final String JSON_PROPERTY_SOURCE = "_source";
  private Object source;

  public static final String JSON_PROPERTY_OPTIONS = "options";
  private Map<String, Object> options = new HashMap<>();

  public static final String JSON_PROPERTY_PROFILE = "profile";
  private Boolean profile;

  public static final String JSON_PROPERTY_TRACK_SCORES = "track_scores";
  private Boolean trackScores;

  public SearchRequest() { 
  }

  public SearchRequest index(String index) {
    this.index = index;
    return this;
  }

   /**
   * Get index
   * @return index
  **/
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


  public SearchRequest query(Object query) {
    this.query = query;
    return this;
  }

   /**
   * Get query
   * @return query
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_QUERY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Object getQuery() {
    return query;
  }


  @JsonProperty(JSON_PROPERTY_QUERY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setQuery(Object query) {
    this.query = query;
    this.fulltextFilter = null; 
    this.attrFilter = null;
  }


  public SearchRequest fulltextFilter(Object fulltextFilter) {
    this.fulltextFilter = fulltextFilter;
    return this;
  }

   /**
   * Get fulltextFilter
   * @return fulltextFilter
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_FULLTEXT_FILTER)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Object getFulltextFilter() {
    return fulltextFilter;
  }


  @JsonProperty(JSON_PROPERTY_FULLTEXT_FILTER)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setFulltextFilter(Object fulltextFilter) {
    this.fulltextFilter = fulltextFilter;
  }


  public SearchRequest attrFilter(Object attrFilter) {
    this.attrFilter = attrFilter;
    return this;
  }

   /**
   * Get attrFilter
   * @return attrFilter
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_ATTR_FILTER)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Object getAttrFilter() {
    return attrFilter;
  }


  @JsonProperty(JSON_PROPERTY_ATTR_FILTER)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setAttrFilter(Object attrFilter) {
    this.attrFilter = attrFilter;
  }


  public SearchRequest limit(Integer limit) {
    this.limit = limit;
    return this;
  }

   /**
   * Get limit
   * @return limit
  **/
  @jakarta.annotation.Nullable
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


  public SearchRequest offset(Integer offset) {
    this.offset = offset;
    return this;
  }

   /**
   * Get offset
   * @return offset
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_OFFSET)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getOffset() {
    return offset;
  }


  @JsonProperty(JSON_PROPERTY_OFFSET)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setOffset(Integer offset) {
    this.offset = offset;
  }


  public SearchRequest maxMatches(Integer maxMatches) {
    this.maxMatches = maxMatches;
    return this;
  }

   /**
   * Get maxMatches
   * @return maxMatches
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_MAX_MATCHES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getMaxMatches() {
    return maxMatches;
  }


  @JsonProperty(JSON_PROPERTY_MAX_MATCHES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setMaxMatches(Integer maxMatches) {
    this.maxMatches = maxMatches;
  }


  public SearchRequest sort(List<Object> sort) {
    this.sort = sort;
    return this;
  }

  public SearchRequest addSortItem(Object sortItem) {
    if (this.sort == null) {
      this.sort = new ArrayList<>();
    }
    this.sort.add(sortItem);
    return this;
  }

   /**
   * Get sort
   * @return sort
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_SORT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<Object> getSort() {
    return sort;
  }


  @JsonProperty(JSON_PROPERTY_SORT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSort(List<Object> sort) {
    this.sort = sort;
  }


  public SearchRequest aggs(Map<String, Aggregation> aggs) {
    this.aggs = aggs;
    return this;
  }

  public SearchRequest putAggsItem(String key, Aggregation aggsItem) {
    if (this.aggs == null) {
      this.aggs = new HashMap<>();
    }
    this.aggs.put(key, aggsItem);
    return this;
  }

   /**
   * Get aggs
   * @return aggs
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_AGGS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Map<String, Aggregation> getAggs() {
    return aggs;
  }


  @JsonProperty(JSON_PROPERTY_AGGS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setAggs(Map<String, Aggregation> aggs) {
    this.aggs = aggs;
  }


  public SearchRequest expressions(Map<String, String> expressions) {
    this.expressions = expressions;
    return this;
  }

  public SearchRequest putExpressionsItem(String key, String expressionsItem) {
    if (this.expressions == null) {
      this.expressions = new HashMap<>();
    }
    this.expressions.put(key, expressionsItem);
    return this;
  }

   /**
   * Get expressions
   * @return expressions
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_EXPRESSIONS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Map<String, String> getExpressions() {
    return expressions;
  }


  @JsonProperty(JSON_PROPERTY_EXPRESSIONS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setExpressions(Map<String, String> expressions) {
    this.expressions = expressions;
  }


  public SearchRequest highlight(Highlight highlight) {
    this.highlight = highlight;
    return this;
  }

   /**
   * Get highlight
   * @return highlight
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_HIGHLIGHT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Highlight getHighlight() {
    return highlight;
  }


  @JsonProperty(JSON_PROPERTY_HIGHLIGHT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setHighlight(Highlight highlight) {
    this.highlight = highlight;
  }


  public SearchRequest source(Object source) {
    this.source = source;
    return this;
  }

   /**
   * Get source
   * @return source
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_SOURCE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Object getSource() {
    return source;
  }


  @JsonProperty(JSON_PROPERTY_SOURCE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSource(Object source) {
    this.source = source;
  }


  public SearchRequest options(Map<String, Object> options) {
    this.options = options;
    return this;
  }

  public SearchRequest putOptionsItem(String key, Object optionsItem) {
    if (this.options == null) {
      this.options = new HashMap<>();
    }
    this.options.put(key, optionsItem);
    return this;
  }

   /**
   * Get options
   * @return options
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_OPTIONS)
  @JsonInclude(content = JsonInclude.Include.ALWAYS, value = JsonInclude.Include.USE_DEFAULTS)

  public Map<String, Object> getOptions() {
    return options;
  }


  @JsonProperty(JSON_PROPERTY_OPTIONS)
  @JsonInclude(content = JsonInclude.Include.ALWAYS, value = JsonInclude.Include.USE_DEFAULTS)
  public void setOptions(Map<String, Object> options) {
    this.options = options;
  }


  public SearchRequest profile(Boolean profile) {
    this.profile = profile;
    return this;
  }

   /**
   * Get profile
   * @return profile
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_PROFILE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Boolean getProfile() {
    return profile;
  }


  @JsonProperty(JSON_PROPERTY_PROFILE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setProfile(Boolean profile) {
    this.profile = profile;
  }


  public SearchRequest trackScores(Boolean trackScores) {
    this.trackScores = trackScores;
    return this;
  }

   /**
   * Get trackScores
   * @return trackScores
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_TRACK_SCORES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Boolean getTrackScores() {
    return trackScores;
  }


  @JsonProperty(JSON_PROPERTY_TRACK_SCORES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setTrackScores(Boolean trackScores) {
    this.trackScores = trackScores;
  }


  /**
   * Return true if this searchRequest object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchRequest searchRequest = (SearchRequest) o;
    return Objects.equals(this.index, searchRequest.index) &&
        Objects.equals(this.query, searchRequest.query) &&
        Objects.equals(this.fulltextFilter, searchRequest.fulltextFilter) &&
        Objects.equals(this.attrFilter, searchRequest.attrFilter) &&
        Objects.equals(this.limit, searchRequest.limit) &&
        Objects.equals(this.offset, searchRequest.offset) &&
        Objects.equals(this.maxMatches, searchRequest.maxMatches) &&
        Objects.equals(this.sort, searchRequest.sort) &&
        Objects.equals(this.aggs, searchRequest.aggs) &&
        Objects.equals(this.expressions, searchRequest.expressions) &&
        Objects.equals(this.highlight, searchRequest.highlight) &&
        Objects.equals(this.source, searchRequest.source) &&
        Objects.equals(this.options, searchRequest.options) &&
        Objects.equals(this.profile, searchRequest.profile) &&
        Objects.equals(this.trackScores, searchRequest.trackScores);
  }

  @Override
  public int hashCode() {
    return Objects.hash(index, query, fulltextFilter, attrFilter, limit, offset, maxMatches, sort, aggs, expressions, highlight, source, options, profile, trackScores);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchRequest {\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    query: ").append(toIndentedString(query)).append("\n");
    sb.append("    fulltextFilter: ").append(toIndentedString(fulltextFilter)).append("\n");
    sb.append("    attrFilter: ").append(toIndentedString(attrFilter)).append("\n");
    sb.append("    limit: ").append(toIndentedString(limit)).append("\n");
    sb.append("    offset: ").append(toIndentedString(offset)).append("\n");
    sb.append("    maxMatches: ").append(toIndentedString(maxMatches)).append("\n");
    sb.append("    sort: ").append(toIndentedString(sort)).append("\n");
    sb.append("    aggs: ").append(toIndentedString(aggs)).append("\n");
    sb.append("    expressions: ").append(toIndentedString(expressions)).append("\n");
    sb.append("    highlight: ").append(toIndentedString(highlight)).append("\n");
    sb.append("    source: ").append(toIndentedString(source)).append("\n");
    sb.append("    options: ").append(toIndentedString(options)).append("\n");
    sb.append("    profile: ").append(toIndentedString(profile)).append("\n");
    sb.append("    trackScores: ").append(toIndentedString(trackScores)).append("\n");
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

