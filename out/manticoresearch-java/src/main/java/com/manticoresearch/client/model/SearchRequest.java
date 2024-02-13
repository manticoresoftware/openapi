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
import com.manticoresearch.client.model.KnnSearchRequestByDocId;
import com.manticoresearch.client.model.KnnSearchRequestByVector;
import com.manticoresearch.client.model.SearchRequest;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.manticoresearch.client.JSON;

import com.fasterxml.jackson.core.type.TypeReference;

import jakarta.ws.rs.core.GenericType;
import jakarta.ws.rs.core.Response;
import java.io.IOException;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.JsonToken;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.MapperFeature;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.deser.std.StdDeserializer;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import com.manticoresearch.client.JSON;

@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-01-30T07:42:09.809929621Z[Etc/UTC]")
@JsonDeserialize(using = SearchRequest.SearchRequestDeserializer.class)
@JsonSerialize(using = SearchRequest.SearchRequestSerializer.class)
public class SearchRequest extends AbstractOpenApiSchema {
    private static final Logger log = Logger.getLogger(SearchRequest.class.getName());

    public static class SearchRequestSerializer extends StdSerializer<SearchRequest> {
        public SearchRequestSerializer(Class<SearchRequest> t) {
            super(t);
        }

        public SearchRequestSerializer() {
            this(null);
        }

        @Override
        public void serialize(SearchRequest value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
            jgen.writeObject(value.getActualInstance());
        }
    }

    public static class SearchRequestDeserializer extends StdDeserializer<SearchRequest> {
        public SearchRequestDeserializer() {
            this(SearchRequest.class);
        }

        public SearchRequestDeserializer(Class<?> vc) {
            super(vc);
        }

        @Override
        public SearchRequest deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
            JsonNode tree = jp.readValueAsTree();
            Object deserialized = null;
            boolean typeCoercion = ctxt.isEnabled(MapperFeature.ALLOW_COERCION_OF_SCALARS);
            int match = 0;
            JsonToken token = tree.traverse(jp.getCodec()).nextToken();
            // deserialize KnnSearchRequestByDocId
            try {
                boolean attemptParsing = true;
                // ensure that we respect type coercion as set on the client ObjectMapper
                if (KnnSearchRequestByDocId.class.equals(Integer.class) || KnnSearchRequestByDocId.class.equals(Long.class) || KnnSearchRequestByDocId.class.equals(Float.class) || KnnSearchRequestByDocId.class.equals(Double.class) || KnnSearchRequestByDocId.class.equals(Boolean.class) || KnnSearchRequestByDocId.class.equals(String.class)) {
                    attemptParsing = typeCoercion;
                    if (!attemptParsing) {
                        attemptParsing |= ((KnnSearchRequestByDocId.class.equals(Integer.class) || KnnSearchRequestByDocId.class.equals(Long.class)) && token == JsonToken.VALUE_NUMBER_INT);
                        attemptParsing |= ((KnnSearchRequestByDocId.class.equals(Float.class) || KnnSearchRequestByDocId.class.equals(Double.class)) && token == JsonToken.VALUE_NUMBER_FLOAT);
                        attemptParsing |= (KnnSearchRequestByDocId.class.equals(Boolean.class) && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE));
                        attemptParsing |= (KnnSearchRequestByDocId.class.equals(String.class) && token == JsonToken.VALUE_STRING);
                    }
                }
                if (attemptParsing) {
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(KnnSearchRequestByDocId.class);
                    // TODO: there is no validation against JSON schema constraints
                    // (min, max, enum, pattern...), this does not perform a strict JSON
                    // validation, which means the 'match' count may be higher than it should be.
                    match++;
                    log.log(Level.FINER, "Input data matches schema 'KnnSearchRequestByDocId'");
                }
            } catch (Exception e) {
                // deserialization failed, continue
                log.log(Level.FINER, "Input data does not match schema 'KnnSearchRequestByDocId'", e);
            }

            // deserialize KnnSearchRequestByVector
            try {
                boolean attemptParsing = true;
                // ensure that we respect type coercion as set on the client ObjectMapper
                if (KnnSearchRequestByVector.class.equals(Integer.class) || KnnSearchRequestByVector.class.equals(Long.class) || KnnSearchRequestByVector.class.equals(Float.class) || KnnSearchRequestByVector.class.equals(Double.class) || KnnSearchRequestByVector.class.equals(Boolean.class) || KnnSearchRequestByVector.class.equals(String.class)) {
                    attemptParsing = typeCoercion;
                    if (!attemptParsing) {
                        attemptParsing |= ((KnnSearchRequestByVector.class.equals(Integer.class) || KnnSearchRequestByVector.class.equals(Long.class)) && token == JsonToken.VALUE_NUMBER_INT);
                        attemptParsing |= ((KnnSearchRequestByVector.class.equals(Float.class) || KnnSearchRequestByVector.class.equals(Double.class)) && token == JsonToken.VALUE_NUMBER_FLOAT);
                        attemptParsing |= (KnnSearchRequestByVector.class.equals(Boolean.class) && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE));
                        attemptParsing |= (KnnSearchRequestByVector.class.equals(String.class) && token == JsonToken.VALUE_STRING);
                    }
                }
                if (attemptParsing) {
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(KnnSearchRequestByVector.class);
                    // TODO: there is no validation against JSON schema constraints
                    // (min, max, enum, pattern...), this does not perform a strict JSON
                    // validation, which means the 'match' count may be higher than it should be.
                    match++;
                    log.log(Level.FINER, "Input data matches schema 'KnnSearchRequestByVector'");
                }
            } catch (Exception e) {
                // deserialization failed, continue
                log.log(Level.FINER, "Input data does not match schema 'KnnSearchRequestByVector'", e);
            }

            // deserialize SearchRequest
            try {
                boolean attemptParsing = true;
                // ensure that we respect type coercion as set on the client ObjectMapper
                if (SearchRequest.class.equals(Integer.class) || SearchRequest.class.equals(Long.class) || SearchRequest.class.equals(Float.class) || SearchRequest.class.equals(Double.class) || SearchRequest.class.equals(Boolean.class) || SearchRequest.class.equals(String.class)) {
                    attemptParsing = typeCoercion;
                    if (!attemptParsing) {
                        attemptParsing |= ((SearchRequest.class.equals(Integer.class) || SearchRequest.class.equals(Long.class)) && token == JsonToken.VALUE_NUMBER_INT);
                        attemptParsing |= ((SearchRequest.class.equals(Float.class) || SearchRequest.class.equals(Double.class)) && token == JsonToken.VALUE_NUMBER_FLOAT);
                        attemptParsing |= (SearchRequest.class.equals(Boolean.class) && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE));
                        attemptParsing |= (SearchRequest.class.equals(String.class) && token == JsonToken.VALUE_STRING);
                    }
                }
                if (attemptParsing) {
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(SearchRequest.class);
                    // TODO: there is no validation against JSON schema constraints
                    // (min, max, enum, pattern...), this does not perform a strict JSON
                    // validation, which means the 'match' count may be higher than it should be.
                    match++;
                    log.log(Level.FINER, "Input data matches schema 'SearchRequest'");
                }
            } catch (Exception e) {
                // deserialization failed, continue
                log.log(Level.FINER, "Input data does not match schema 'SearchRequest'", e);
            }

            if (match == 1) {
                SearchRequest ret = new SearchRequest();
                ret.setActualInstance(deserialized);
                return ret;
            }
            throw new IOException(String.format("Failed deserialization for SearchRequest: %d classes match result, expected 1", match));
        }

        /**
         * Handle deserialization of the 'null' value.
         */
        @Override
        public SearchRequest getNullValue(DeserializationContext ctxt) throws JsonMappingException {
            throw new JsonMappingException(ctxt.getParser(), "SearchRequest cannot be null");
        }
    }

    // store a list of schema names defined in oneOf
    public static final Map<String, GenericType> schemas = new HashMap<>();

    public SearchRequest() {
        super("oneOf", Boolean.FALSE);
    }

    public SearchRequest(KnnSearchRequestByDocId o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    public SearchRequest(KnnSearchRequestByVector o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    public SearchRequest(SearchRequest o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    static {
        schemas.put("KnnSearchRequestByDocId", new GenericType<KnnSearchRequestByDocId>() {
        });
        schemas.put("KnnSearchRequestByVector", new GenericType<KnnSearchRequestByVector>() {
        });
        schemas.put("SearchRequest", new GenericType<SearchRequest>() {
        });
        JSON.registerDescendants(SearchRequest.class, Collections.unmodifiableMap(schemas));
    }

    @Override
    public Map<String, GenericType> getSchemas() {
        return SearchRequest.schemas;
    }

    /**
     * Set the instance that matches the oneOf child schema, check
     * the instance parameter is valid against the oneOf child schemas:
     * KnnSearchRequestByDocId, KnnSearchRequestByVector, SearchRequest
     *
     * It could be an instance of the 'oneOf' schemas.
     * The oneOf child schemas may themselves be a composed schema (allOf, anyOf, oneOf).
     */
    @Override
    public void setActualInstance(Object instance) {
        if (JSON.isInstanceOf(KnnSearchRequestByDocId.class, instance, new HashSet<>())) {
            super.setActualInstance(instance);
            return;
        }

        if (JSON.isInstanceOf(KnnSearchRequestByVector.class, instance, new HashSet<>())) {
            super.setActualInstance(instance);
            return;
        }

        if (JSON.isInstanceOf(SearchRequest.class, instance, new HashSet<>())) {
            super.setActualInstance(instance);
            return;
        }

        throw new RuntimeException("Invalid instance type. Must be KnnSearchRequestByDocId, KnnSearchRequestByVector, SearchRequest");
    }

    /**
     * Get the actual instance, which can be the following:
     * KnnSearchRequestByDocId, KnnSearchRequestByVector, SearchRequest
     *
     * @return The actual instance (KnnSearchRequestByDocId, KnnSearchRequestByVector, SearchRequest)
     */
    @Override
    public Object getActualInstance() {
        return super.getActualInstance();
    }

    /**
     * Get the actual instance of `KnnSearchRequestByDocId`. If the actual instance is not `KnnSearchRequestByDocId`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `KnnSearchRequestByDocId`
     * @throws ClassCastException if the instance is not `KnnSearchRequestByDocId`
     */
    public KnnSearchRequestByDocId getKnnSearchRequestByDocId() throws ClassCastException {
        return (KnnSearchRequestByDocId)super.getActualInstance();
    }

    /**
     * Get the actual instance of `KnnSearchRequestByVector`. If the actual instance is not `KnnSearchRequestByVector`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `KnnSearchRequestByVector`
     * @throws ClassCastException if the instance is not `KnnSearchRequestByVector`
     */
    public KnnSearchRequestByVector getKnnSearchRequestByVector() throws ClassCastException {
        return (KnnSearchRequestByVector)super.getActualInstance();
    }

    /**
     * Get the actual instance of `SearchRequest`. If the actual instance is not `SearchRequest`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `SearchRequest`
     * @throws ClassCastException if the instance is not `SearchRequest`
     */
    public SearchRequest getSearchRequest() throws ClassCastException {
        return (SearchRequest)super.getActualInstance();
    }

}

