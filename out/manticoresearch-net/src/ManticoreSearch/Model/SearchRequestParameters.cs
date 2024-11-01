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
    /// SearchRequestParameters
    /// </summary>
    [DataContract(Name = "searchRequestParameters")]
    public partial class SearchRequestParameters : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequestParameters" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected SearchRequestParameters() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequestParameters" /> class.
        /// </summary>
        /// <param name="aggs">aggs.</param>
        /// <param name="expressions">expressions.</param>
        /// <param name="join">join.</param>
        /// <param name="highlight">highlight.</param>
        /// <param name="index">index (required).</param>
        /// <param name="limit">limit.</param>
        /// <param name="maxMatches">maxMatches.</param>
        /// <param name="offset">offset.</param>
        /// <param name="options">options.</param>
        /// <param name="profile">profile.</param>
        /// <param name="sort">sort.</param>
        /// <param name="source">source.</param>
        /// <param name="trackScores">trackScores.</param>
        public SearchRequestParameters(Aggregation aggs = default(Aggregation), Dictionary<string, string> expressions = default(Dictionary<string, string>), List<JoinInner> join = default(List<JoinInner>), Highlight highlight = default(Highlight), string index = default(string), int limit = default(int), int maxMatches = default(int), int offset = default(int), Object options = default(Object), bool profile = default(bool), List<SearchRequestParametersSortInner> sort = default(List<SearchRequestParametersSortInner>), SearchRequestParametersSource source = default(SearchRequestParametersSource), bool trackScores = default(bool))
        {
            // to ensure "index" is required (not null)
            if (index == null)
            {
                throw new ArgumentNullException("index is a required property for SearchRequestParameters and cannot be null");
            }
            this.Index = index;
            this.Aggs = aggs;
            this.Expressions = expressions;
            this.Join = join;
            this.Highlight = highlight;
            this.Limit = limit;
            this.MaxMatches = maxMatches;
            this.Offset = offset;
            this.Options = options;
            this.Profile = profile;
            this.Sort = sort;
            this.Source = source;
            this.TrackScores = trackScores;
        }

        /// <summary>
        /// Gets or Sets Aggs
        /// </summary>
        /// <example>{agg1&#x3D;{terms&#x3D;{field&#x3D;field1, size&#x3D;1000}, sort&#x3D;[{field1&#x3D;{order&#x3D;asc}}]}}</example>
        [DataMember(Name = "aggs", EmitDefaultValue = false)]
        public Aggregation Aggs { get; set; }

        /// <summary>
        /// Gets or Sets Expressions
        /// </summary>
        /// <example>{title_len&#x3D;crc32(title)}</example>
        [DataMember(Name = "expressions", EmitDefaultValue = false)]
        public Dictionary<string, string> Expressions { get; set; }

        /// <summary>
        /// Gets or Sets Join
        /// </summary>
        [DataMember(Name = "join", EmitDefaultValue = false)]
        public List<JoinInner> Join { get; set; }

        /// <summary>
        /// Gets or Sets Highlight
        /// </summary>
        [DataMember(Name = "highlight", EmitDefaultValue = false)]
        public Highlight Highlight { get; set; }

        /// <summary>
        /// Gets or Sets Index
        /// </summary>
        [DataMember(Name = "index", IsRequired = true, EmitDefaultValue = true)]
        public string Index { get; set; }

        /// <summary>
        /// Gets or Sets Limit
        /// </summary>
        [DataMember(Name = "limit", EmitDefaultValue = false)]
        public int Limit { get; set; }

        /// <summary>
        /// Gets or Sets MaxMatches
        /// </summary>
        [DataMember(Name = "max_matches", EmitDefaultValue = false)]
        public int MaxMatches { get; set; }

        /// <summary>
        /// Gets or Sets Offset
        /// </summary>
        [DataMember(Name = "offset", EmitDefaultValue = false)]
        public int Offset { get; set; }

        /// <summary>
        /// Gets or Sets Options
        /// </summary>
        [DataMember(Name = "options", EmitDefaultValue = false)]
        public Object Options { get; set; }

        /// <summary>
        /// Gets or Sets Profile
        /// </summary>
        [DataMember(Name = "profile", EmitDefaultValue = true)]
        public bool Profile { get; set; }

        /// <summary>
        /// Gets or Sets Sort
        /// </summary>
        /// <example>[]</example>
        [DataMember(Name = "sort", EmitDefaultValue = false)]
        public List<SearchRequestParametersSortInner> Sort { get; set; }

        /// <summary>
        /// Gets or Sets Source
        /// </summary>
        [DataMember(Name = "_source", EmitDefaultValue = false)]
        public SearchRequestParametersSource Source { get; set; }

        /// <summary>
        /// Gets or Sets TrackScores
        /// </summary>
        [DataMember(Name = "track_scores", EmitDefaultValue = true)]
        public bool TrackScores { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class SearchRequestParameters {\n");
            sb.Append("  Aggs: ").Append(Aggs).Append("\n");
            sb.Append("  Expressions: ").Append(Expressions).Append("\n");
            sb.Append("  Join: ").Append(Join).Append("\n");
            sb.Append("  Highlight: ").Append(Highlight).Append("\n");
            sb.Append("  Index: ").Append(Index).Append("\n");
            sb.Append("  Limit: ").Append(Limit).Append("\n");
            sb.Append("  MaxMatches: ").Append(MaxMatches).Append("\n");
            sb.Append("  Offset: ").Append(Offset).Append("\n");
            sb.Append("  Options: ").Append(Options).Append("\n");
            sb.Append("  Profile: ").Append(Profile).Append("\n");
            sb.Append("  Sort: ").Append(Sort).Append("\n");
            sb.Append("  Source: ").Append(Source).Append("\n");
            sb.Append("  TrackScores: ").Append(TrackScores).Append("\n");
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