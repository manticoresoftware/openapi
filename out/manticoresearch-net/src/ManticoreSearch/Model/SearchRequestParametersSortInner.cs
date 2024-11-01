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
    /// SearchRequestParametersSortInner
    /// </summary>
    [JsonConverter(typeof(SearchRequestParametersSortInnerJsonConverter))]
    [DataContract(Name = "searchRequestParameters_sort_inner")]
    public partial class SearchRequestParametersSortInner : AbstractOpenAPISchema, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequestParametersSortInner" /> class
        /// with the <see cref="string" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of string.</param>
        public SearchRequestParametersSortInner(string actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequestParametersSortInner" /> class
        /// with the <see cref="Dictionary{string, string}" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of Dictionary&lt;string, string&gt;.</param>
        public SearchRequestParametersSortInner(Dictionary<string, string> actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.") ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequestParametersSortInner" /> class
        /// with the <see cref="SortObject" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of SortObject.</param>
        public SearchRequestParametersSortInner(SortObject actualInstance)
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
                if (value.GetType() == typeof(Dictionary<string, string>) || value is Dictionary<string, string>)
                {
                    this._actualInstance = value;
                }
                else if (value.GetType() == typeof(SortObject) || value is SortObject)
                {
                    this._actualInstance = value;
                }
                else if (value.GetType() == typeof(string) || value is string)
                {
                    this._actualInstance = value;
                }
                else
                {
                    throw new ArgumentException("Invalid instance found. Must be the following types: Dictionary<string, string>, SortObject, string");
                }
            }
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
        /// Get the actual instance of `Dictionary&lt;string, string&gt;`. If the actual instance is not `Dictionary&lt;string, string&gt;`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of Dictionary&lt;string, string&gt;</returns>
        public Dictionary<string, string> GetDictionary()
        {
            return (Dictionary<string, string>)this.ActualInstance;
        }

        /// <summary>
        /// Get the actual instance of `SortObject`. If the actual instance is not `SortObject`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of SortObject</returns>
        public SortObject GetSortObject()
        {
            return (SortObject)this.ActualInstance;
        }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class SearchRequestParametersSortInner {\n");
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
            return JsonConvert.SerializeObject(this.ActualInstance, SearchRequestParametersSortInner.SerializerSettings);
        }

        /// <summary>
        /// Converts the JSON string into an instance of SearchRequestParametersSortInner
        /// </summary>
        /// <param name="jsonString">JSON string</param>
        /// <returns>An instance of SearchRequestParametersSortInner</returns>
        public static SearchRequestParametersSortInner FromJson(string jsonString)
        {
            SearchRequestParametersSortInner newSearchRequestParametersSortInner = null;

            if (string.IsNullOrEmpty(jsonString))
            {
                return newSearchRequestParametersSortInner;
            }
            int match = 0;
            List<string> matchedTypes = new List<string>();

            try
            {
                // if it does not contains "AdditionalProperties", use SerializerSettings to deserialize
                if (typeof(Dictionary<string, string>).GetProperty("AdditionalProperties") == null)
                {
                    newSearchRequestParametersSortInner = new SearchRequestParametersSortInner(JsonConvert.DeserializeObject<Dictionary<string, string>>(jsonString, SearchRequestParametersSortInner.SerializerSettings));
                }
                else
                {
                    newSearchRequestParametersSortInner = new SearchRequestParametersSortInner(JsonConvert.DeserializeObject<Dictionary<string, string>>(jsonString, SearchRequestParametersSortInner.AdditionalPropertiesSerializerSettings));
                }
                matchedTypes.Add("Dictionary<string, string>");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(string.Format("Failed to deserialize `{0}` into Dictionary<string, string>: {1}", jsonString, exception.ToString()));
            }

            try
            {
                // if it does not contains "AdditionalProperties", use SerializerSettings to deserialize
                if (typeof(SortObject).GetProperty("AdditionalProperties") == null)
                {
                    newSearchRequestParametersSortInner = new SearchRequestParametersSortInner(JsonConvert.DeserializeObject<SortObject>(jsonString, SearchRequestParametersSortInner.SerializerSettings));
                }
                else
                {
                    newSearchRequestParametersSortInner = new SearchRequestParametersSortInner(JsonConvert.DeserializeObject<SortObject>(jsonString, SearchRequestParametersSortInner.AdditionalPropertiesSerializerSettings));
                }
                matchedTypes.Add("SortObject");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(string.Format("Failed to deserialize `{0}` into SortObject: {1}", jsonString, exception.ToString()));
            }

            try
            {
                // if it does not contains "AdditionalProperties", use SerializerSettings to deserialize
                if (typeof(string).GetProperty("AdditionalProperties") == null)
                {
                    newSearchRequestParametersSortInner = new SearchRequestParametersSortInner(JsonConvert.DeserializeObject<string>(jsonString, SearchRequestParametersSortInner.SerializerSettings));
                }
                else
                {
                    newSearchRequestParametersSortInner = new SearchRequestParametersSortInner(JsonConvert.DeserializeObject<string>(jsonString, SearchRequestParametersSortInner.AdditionalPropertiesSerializerSettings));
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
            return newSearchRequestParametersSortInner;
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
    /// Custom JSON converter for SearchRequestParametersSortInner
    /// </summary>
    public class SearchRequestParametersSortInnerJsonConverter : JsonConverter
    {
        /// <summary>
        /// To write the JSON string
        /// </summary>
        /// <param name="writer">JSON writer</param>
        /// <param name="value">Object to be converted into a JSON string</param>
        /// <param name="serializer">JSON Serializer</param>
        public override void WriteJson(JsonWriter writer, object value, JsonSerializer serializer)
        {
            writer.WriteRawValue((string)(typeof(SearchRequestParametersSortInner).GetMethod("ToJson").Invoke(value, null)));
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
                    return new SearchRequestParametersSortInner(Convert.ToString(reader.Value));
                case JsonToken.String: 
                    return new SearchRequestParametersSortInner(Convert.ToString(reader.Value));
                case JsonToken.StartObject:
                    return SearchRequestParametersSortInner.FromJson(JObject.Load(reader).ToString(Formatting.None));
                case JsonToken.StartArray:
                    return SearchRequestParametersSortInner.FromJson(JArray.Load(reader).ToString(Formatting.None));
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