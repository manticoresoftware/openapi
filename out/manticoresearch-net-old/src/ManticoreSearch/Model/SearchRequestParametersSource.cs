/*
 * Manticore Search Client
 *
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 5.0.0
 * Contact: info@manticoresearch.com
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using FileParameter = ManticoreSearch.Client.FileParameter;
using OpenAPIDateConverter = ManticoreSearch.Client.OpenAPIDateConverter;
using System.Reflection;

namespace ManticoreSearch.Model
{
    /// <summary>
    /// SearchRequestParametersSource
    /// </summary>
    [JsonConverter(typeof(SearchRequestParametersSourceJsonConverter))]
    [DataContract(Name = "searchRequestParameters__source")]
    public partial class SearchRequestParametersSource : AbstractOpenAPISchema, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequestParametersSource" /> class
        /// with the <see cref="List{String}" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of List&lt;string&gt;.</param>
        public SearchRequestParametersSource(List<string> actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequestParametersSource" /> class
        /// with the <see cref="string" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of string.</param>
        public SearchRequestParametersSource(string actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequestParametersSource" /> class
        /// with the <see cref="SourceByRules" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of SourceByRules.</param>
        public SearchRequestParametersSource(SourceByRules actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }


        private Object _actualInstance;

        /// <summary>
        /// Gets or Sets ActualInstance
        /// </summary>
        public override Object ActualInstance
        {
            get
            {
                return _actualInstance;
            }
            set
            {
                if (value.GetType() == typeof(List<string>) || value is List<string>)
                {
                    this._actualInstance = value;
                }
                else if (value.GetType() == typeof(SourceByRules) || value is SourceByRules)
                {
                    this._actualInstance = value;
                }
                else if (value.GetType() == typeof(string) || value is string)
                {
                    this._actualInstance = value;
                }
                else
                {
                    throw new ArgumentException("Invalid instance found. Must be the following types: List<string>, SourceByRules, string");
                }
            }
        }

        /// <summary>
        /// Get the actual instance of `List&lt;string&gt;`. If the actual instance is not `List&lt;string&gt;`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of List&lt;string&gt;</returns>
        public List<string> GetList()
        {
            return (List<string>)this.ActualInstance;
        }

        /// <summary>
        /// Get the actual instance of `string`. If the actual instance is not `string`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of string</returns>
        public string GetString()
        {
            return (string)this.ActualInstance;
        }

        /// <summary>
        /// Get the actual instance of `SourceByRules`. If the actual instance is not `SourceByRules`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of SourceByRules</returns>
        public SourceByRules GetSourceByRules()
        {
            return (SourceByRules)this.ActualInstance;
        }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class SearchRequestParametersSource {\n");
            sb.Append("  ActualInstance: ").Append(this.ActualInstance).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public override string ToJson()
        {
            return JsonConvert.SerializeObject(this.ActualInstance, SearchRequestParametersSource.SerializerSettings);
        }

        /// <summary>
        /// Converts the JSON string into an instance of SearchRequestParametersSource
        /// </summary>
        /// <param name="jsonString">JSON string</param>
        /// <returns>An instance of SearchRequestParametersSource</returns>
        public static SearchRequestParametersSource FromJson(string jsonString)
        {
            SearchRequestParametersSource newSearchRequestParametersSource = null;

            if (string.IsNullOrEmpty(jsonString))
            {
                return newSearchRequestParametersSource;
            }
            int match = 0;
            List<string> matchedTypes = new List<string>();

            try
            {
                // if it does not contains "AdditionalProperties", use SerializerSettings to deserialize
                if (typeof(List<string>).GetProperty("AdditionalProperties") == null)
                {
                    newSearchRequestParametersSource = new SearchRequestParametersSource(JsonConvert.DeserializeObject<List<string>>(jsonString, SearchRequestParametersSource.SerializerSettings));
                }
                else
                {
                    newSearchRequestParametersSource = new SearchRequestParametersSource(JsonConvert.DeserializeObject<List<string>>(jsonString, SearchRequestParametersSource.AdditionalPropertiesSerializerSettings));
                }
                matchedTypes.Add("List<string>");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(string.Format("Failed to deserialize `{0}` into List<string>: {1}", jsonString, exception.ToString()));
            }

            try
            {
                // if it does not contains "AdditionalProperties", use SerializerSettings to deserialize
                if (typeof(SourceByRules).GetProperty("AdditionalProperties") == null)
                {
                    newSearchRequestParametersSource = new SearchRequestParametersSource(JsonConvert.DeserializeObject<SourceByRules>(jsonString, SearchRequestParametersSource.SerializerSettings));
                }
                else
                {
                    newSearchRequestParametersSource = new SearchRequestParametersSource(JsonConvert.DeserializeObject<SourceByRules>(jsonString, SearchRequestParametersSource.AdditionalPropertiesSerializerSettings));
                }
                matchedTypes.Add("SourceByRules");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(string.Format("Failed to deserialize `{0}` into SourceByRules: {1}", jsonString, exception.ToString()));
            }

            try
            {
                // if it does not contains "AdditionalProperties", use SerializerSettings to deserialize
                if (typeof(string).GetProperty("AdditionalProperties") == null)
                {
                    newSearchRequestParametersSource = new SearchRequestParametersSource(JsonConvert.DeserializeObject<string>(jsonString, SearchRequestParametersSource.SerializerSettings));
                }
                else
                {
                    newSearchRequestParametersSource = new SearchRequestParametersSource(JsonConvert.DeserializeObject<string>(jsonString, SearchRequestParametersSource.AdditionalPropertiesSerializerSettings));
                }
                matchedTypes.Add("string");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(string.Format("Failed to deserialize `{0}` into string: {1}", jsonString, exception.ToString()));
            }

            if (match == 0)
            {
                throw new InvalidDataException("The JSON string `" + jsonString + "` cannot be deserialized into any schema defined.");
            }
            else if (match > 1)
            {
                throw new InvalidDataException("The JSON string `" + jsonString + "` incorrectly matches more than one schema (should be exactly one match): " + String.Join(",", matchedTypes));
            }

            // deserialization is considered successful at this point if no exception has been thrown.
            return newSearchRequestParametersSource;
        }


        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

    /// <summary>
    /// Custom JSON converter for SearchRequestParametersSource
    /// </summary>
    public class SearchRequestParametersSourceJsonConverter : JsonConverter
    {
        /// <summary>
        /// To write the JSON string
        /// </summary>
        /// <param name="writer">JSON writer</param>
        /// <param name="value">Object to be converted into a JSON string</param>
        /// <param name="serializer">JSON Serializer</param>
        public override void WriteJson(JsonWriter writer, object value, JsonSerializer serializer)
        {
            writer.WriteRawValue((string)(typeof(SearchRequestParametersSource).GetMethod("ToJson").Invoke(value, null)));
        }

        /// <summary>
        /// To convert a JSON string into an object
        /// </summary>
        /// <param name="reader">JSON reader</param>
        /// <param name="objectType">Object type</param>
        /// <param name="existingValue">Existing value</param>
        /// <param name="serializer">JSON Serializer</param>
        /// <returns>The object converted from the JSON string</returns>
        public override object ReadJson(JsonReader reader, Type objectType, object existingValue, JsonSerializer serializer)
        {
            switch(reader.TokenType) 
            {
                case JsonToken.String: 
                    return new SearchRequestParametersSource(Convert.ToString(reader.Value));
                case JsonToken.StartObject:
                    return SearchRequestParametersSource.FromJson(JObject.Load(reader).ToString(Formatting.None));
                case JsonToken.StartArray:
                    return SearchRequestParametersSource.FromJson(JArray.Load(reader).ToString(Formatting.None));
                default:
                    return null;
            }
        }

        /// <summary>
        /// Check if the object can be converted
        /// </summary>
        /// <param name="objectType">Object type</param>
        /// <returns>True if the object can be converted</returns>
        public override bool CanConvert(Type objectType)
        {
            return false;
        }
    }

}
