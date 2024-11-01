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
import com.manticoresearch.client.model.JoinCond;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.manticoresearch.client.JSON;


/**
 * JoinOn
 */
@JsonPropertyOrder({
  JoinOn.JSON_PROPERTY_RIGHT,
  JoinOn.JSON_PROPERTY_LEFT,
  JoinOn.JSON_PROPERTY_OPERATOR
})
@JsonTypeName("joinOn")
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-10-28T14:42:59.426983397Z[Etc/UTC]", comments = "Generator version: 7.3.0-SNAPSHOT")
public class JoinOn {
  public static final String JSON_PROPERTY_RIGHT = "right";
  private JoinCond right;

  public static final String JSON_PROPERTY_LEFT = "left";
  private JoinCond left;

  /**
   * Gets or Sets operator
   */
  public enum OperatorEnum {
    EQ("eq");

    private String value;

    OperatorEnum(String value) {
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
    public static OperatorEnum fromValue(String value) {
      for (OperatorEnum b : OperatorEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_OPERATOR = "operator";
  private OperatorEnum operator;

  public JoinOn() { 
  }

  public JoinOn right(JoinCond right) {
    this.right = right;
    return this;
  }

  /**
   * Get right
   * @return right
   */
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_RIGHT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JoinCond getRight() {
    return right;
  }


  @JsonProperty(JSON_PROPERTY_RIGHT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setRight(JoinCond right) {
    this.right = right;
  }


  public JoinOn left(JoinCond left) {
    this.left = left;
    return this;
  }

  /**
   * Get left
   * @return left
   */
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_LEFT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JoinCond getLeft() {
    return left;
  }


  @JsonProperty(JSON_PROPERTY_LEFT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setLeft(JoinCond left) {
    this.left = left;
  }


  public JoinOn operator(OperatorEnum operator) {
    this.operator = operator;
    return this;
  }

  /**
   * Get operator
   * @return operator
   */
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_OPERATOR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public OperatorEnum getOperator() {
    return operator;
  }


  @JsonProperty(JSON_PROPERTY_OPERATOR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setOperator(OperatorEnum operator) {
    this.operator = operator;
  }


  /**
   * Return true if this joinOn object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    JoinOn joinOn = (JoinOn) o;
    return Objects.equals(this.right, joinOn.right) &&
        Objects.equals(this.left, joinOn.left) &&
        Objects.equals(this.operator, joinOn.operator);
  }

  @Override
  public int hashCode() {
    return Objects.hash(right, left, operator);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class JoinOn {\n");
    sb.append("    right: ").append(toIndentedString(right)).append("\n");
    sb.append("    left: ").append(toIndentedString(left)).append("\n");
    sb.append("    operator: ").append(toIndentedString(operator)).append("\n");
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

