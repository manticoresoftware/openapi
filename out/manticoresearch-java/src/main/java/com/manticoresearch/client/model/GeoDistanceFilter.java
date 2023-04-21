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
import com.manticoresearch.client.model.GeoDistanceFilterLocationAnchor;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.manticoresearch.client.JSON;


/**
 * Geo distance attribute filter
 */
@JsonPropertyOrder({
  GeoDistanceFilter.JSON_PROPERTY_LOCATION_ANCHOR,
  GeoDistanceFilter.JSON_PROPERTY_LOCATION_SOURCE,
  GeoDistanceFilter.JSON_PROPERTY_DISTANCE_TYPE,
  GeoDistanceFilter.JSON_PROPERTY_DISTANCE
})
@JsonTypeName("geoDistanceFilter")
@JsonIgnoreProperties(ignoreUnknown = true)
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-04-21T13:25:16.289613Z[Etc/UTC]")
public class GeoDistanceFilter {
  public static final String JSON_PROPERTY_LOCATION_ANCHOR = "location_anchor";
  private GeoDistanceFilterLocationAnchor locationAnchor;

  public static final String JSON_PROPERTY_LOCATION_SOURCE = "location_source";
  private String locationSource;

  /**
   * Gets or Sets distanceType
   */
  public enum DistanceTypeEnum {
    ADAPTIVE("adaptive"),
    
    HAVERSINE("haversine");

    private String value;

    DistanceTypeEnum(String value) {
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
    public static DistanceTypeEnum fromValue(String value) {
      for (DistanceTypeEnum b : DistanceTypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_DISTANCE_TYPE = "distance_type";
  private DistanceTypeEnum distanceType;

  public static final String JSON_PROPERTY_DISTANCE = "distance";
  private String distance;

  public GeoDistanceFilter() { 
  }

  public GeoDistanceFilter locationAnchor(GeoDistanceFilterLocationAnchor locationAnchor) {
    this.locationAnchor = locationAnchor;
    return this;
  }

   /**
   * Get locationAnchor
   * @return locationAnchor
  **/
  @javax.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_LOCATION_ANCHOR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public GeoDistanceFilterLocationAnchor getLocationAnchor() {
    return locationAnchor;
  }


  @JsonProperty(JSON_PROPERTY_LOCATION_ANCHOR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setLocationAnchor(GeoDistanceFilterLocationAnchor locationAnchor) {
    this.locationAnchor = locationAnchor;
  }


  public GeoDistanceFilter locationSource(String locationSource) {
    this.locationSource = locationSource;
    return this;
  }

   /**
   * Attribute containing latitude and longitude data
   * @return locationSource
  **/
  @javax.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_LOCATION_SOURCE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getLocationSource() {
    return locationSource;
  }


  @JsonProperty(JSON_PROPERTY_LOCATION_SOURCE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setLocationSource(String locationSource) {
    this.locationSource = locationSource;
  }


  public GeoDistanceFilter distanceType(DistanceTypeEnum distanceType) {
    this.distanceType = distanceType;
    return this;
  }

   /**
   * Get distanceType
   * @return distanceType
  **/
  @javax.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_DISTANCE_TYPE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public DistanceTypeEnum getDistanceType() {
    return distanceType;
  }


  @JsonProperty(JSON_PROPERTY_DISTANCE_TYPE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDistanceType(DistanceTypeEnum distanceType) {
    this.distanceType = distanceType;
  }


  public GeoDistanceFilter distance(String distance) {
    this.distance = distance;
    return this;
  }

   /**
   * Get distance
   * @return distance
  **/
  @javax.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_DISTANCE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getDistance() {
    return distance;
  }


  @JsonProperty(JSON_PROPERTY_DISTANCE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDistance(String distance) {
    this.distance = distance;
  }


  /**
   * Return true if this geoDistanceFilter object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GeoDistanceFilter geoDistanceFilter = (GeoDistanceFilter) o;
    return Objects.equals(this.locationAnchor, geoDistanceFilter.locationAnchor) &&
        Objects.equals(this.locationSource, geoDistanceFilter.locationSource) &&
        Objects.equals(this.distanceType, geoDistanceFilter.distanceType) &&
        Objects.equals(this.distance, geoDistanceFilter.distance);
  }

  @Override
  public int hashCode() {
    return Objects.hash(locationAnchor, locationSource, distanceType, distance);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GeoDistanceFilter {\n");
    sb.append("    locationAnchor: ").append(toIndentedString(locationAnchor)).append("\n");
    sb.append("    locationSource: ").append(toIndentedString(locationSource)).append("\n");
    sb.append("    distanceType: ").append(toIndentedString(distanceType)).append("\n");
    sb.append("    distance: ").append(toIndentedString(distance)).append("\n");
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

