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

namespace ManticoreSearch.Model
{
    /// <summary>
    /// Response object containing the results of a search request
    /// </summary>
    [DataContract(Name = "searchResponse")]
    public partial class SearchResponse : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="SearchResponse" /> class.
        /// </summary>
        /// <param name="">Time taken to execute the search.</param>
        /// <param name="">Indicates whether the search operation timed out.</param>
        /// <param name="">Aggregated search results grouped by the specified criteria.</param>
        /// <param name="">.</param>
        /// <param name="">Profile information about the search execution, if profiling is enabled.</param>
        /// <param name="">Warnings encountered during the search operation.</param>
        public SearchResponse(int  = default(int), bool  = default(bool), Object  = default(Object), SearchResponseHits  = default(SearchResponseHits), Object  = default(Object), Object  = default(Object))
        {
            this.Took = ;
            this.TimedOut = ;
            this.Aggregations = ;
            this.Hits = ;
            this.Profile = ;
            this.Warning = ;
        }

        /// <summary>
        /// Time taken to execute the search
        /// </summary>
        /// <value>Time taken to execute the search</value>
        [DataMember(Name = "took", EmitDefaultValue = false)]
        public int Took { get; set; }

        /// <summary>
        /// Indicates whether the search operation timed out
        /// </summary>
        /// <value>Indicates whether the search operation timed out</value>
        [DataMember(Name = "timed_out", EmitDefaultValue = true)]
        public bool TimedOut { get; set; }

        /// <summary>
        /// Aggregated search results grouped by the specified criteria
        /// </summary>
        /// <value>Aggregated search results grouped by the specified criteria</value>
        [DataMember(Name = "aggregations", EmitDefaultValue = false)]
        public Object Aggregations { get; set; }

        /// <summary>
        /// Gets or Sets Hits
        /// </summary>
        [DataMember(Name = "hits", EmitDefaultValue = false)]
        public SearchResponseHits Hits { get; set; }

        /// <summary>
        /// Profile information about the search execution, if profiling is enabled
        /// </summary>
        /// <value>Profile information about the search execution, if profiling is enabled</value>
        [DataMember(Name = "profile", EmitDefaultValue = false)]
        public Object Profile { get; set; }

        /// <summary>
        /// Warnings encountered during the search operation
        /// </summary>
        /// <value>Warnings encountered during the search operation</value>
        [DataMember(Name = "warning", EmitDefaultValue = false)]
        public Object Warning { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class SearchResponse {\n");
            sb.Append("  Took: ").Append(Took).Append("\n");
            sb.Append("  TimedOut: ").Append(TimedOut).Append("\n");
            sb.Append("  Aggregations: ").Append(Aggregations).Append("\n");
            sb.Append("  Hits: ").Append(Hits).Append("\n");
            sb.Append("  Profile: ").Append(Profile).Append("\n");
            sb.Append("  Warning: ").Append(Warning).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
