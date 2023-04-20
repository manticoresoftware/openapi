/*
 * Manticore Search Client
 *
 * Low-level client for Manticore Search. 
 *
 * The version of the OpenAPI document: 1.0.0
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
    /// Payload for search operation
    /// </summary>
    [DataContract(Name = "searchRequest")]
    public partial class SearchRequest : IEquatable<SearchRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected SearchRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="SearchRequest" /> class.
        /// </summary>
        /// <param name="index">index (required) (default to &quot;&quot;).</param>
        /// <param name="query">query.</param>
        /// <param name="fulltextFilter">fulltextFilter.</param>
        /// <param name="attrFilter">attrFilter.</param>
        /// <param name="limit">limit.</param>
        /// <param name="offset">offset.</param>
        /// <param name="maxMatches">maxMatches.</param>
        /// <param name="sort">sort.</param>
        /// <param name="sortOld">sortOld.</param>
        /// <param name="aggs">aggs.</param>
        /// <param name="expressions">expressions.</param>
        /// <param name="highlight">highlight.</param>
        /// <param name="source">source.</param>
        /// <param name="options">options.</param>
        /// <param name="profile">profile.</param>
        /// <param name="trackScores">trackScores.</param>
        public SearchRequest(string index = "", Object query = default(Object), Object fulltextFilter = default(Object), Object attrFilter = default(Object), int limit = default(int), int offset = default(int), int maxMatches = default(int), List<Object> sort = default(List<Object>), List<Object> sortOld = default(List<Object>), List<Aggregation> aggs = default(List<Aggregation>), List<Object> expressions = default(List<Object>), Highlight highlight = default(Highlight), Object source = default(Object), Dictionary<string, Object> options = default(Dictionary<string, Object>), bool profile = default(bool), bool trackScores = default(bool))
        {
            // to ensure "index" is required (not null)
            if (index == null)
            {
                throw new ArgumentNullException("index is a required property for SearchRequest and cannot be null");
            }
            this.Index = index;
            this.Query = query;
            this.FulltextFilter = fulltextFilter;
            this.AttrFilter = attrFilter;
            this.Limit = limit;
            this.Offset = offset;
            this.MaxMatches = maxMatches;
            this.Sort = sort;
            this.SortOld = sortOld;
            this.Aggs = aggs;
            this.Expressions = expressions;
            this.Highlight = highlight;
            this.Source = source;
            this.Options = options;
            this.Profile = profile;
            this.TrackScores = trackScores;
        }

        /// <summary>
        /// Gets or Sets Index
        /// </summary>
        [DataMember(Name = "index", IsRequired = true, EmitDefaultValue = false)]
        public string Index { get; set; }

        Object _Query;
        /// <summary>
        /// Gets or Sets Query
        /// </summary>
        [DataMember(Name = "query", EmitDefaultValue = false)]
        public Object Query 
        {
         	get { return _Query; } 
        	set { _Query = value; this.FulltextFilter = null; this.AttrFilter = null; }  
        }

        /// <summary>
        /// Gets or Sets FulltextFilter
        /// </summary>
        [DataMember(Name = "fulltext_filter", EmitDefaultValue = false)]
        public Object FulltextFilter { get; set; }

        /// <summary>
        /// Gets or Sets AttrFilter
        /// </summary>
        [DataMember(Name = "attr_filter", EmitDefaultValue = false)]
        public Object AttrFilter { get; set; }

        /// <summary>
        /// Gets or Sets Limit
        /// </summary>
        [DataMember(Name = "limit", EmitDefaultValue = false)]
        public int Limit { get; set; }

        /// <summary>
        /// Gets or Sets Offset
        /// </summary>
        [DataMember(Name = "offset", EmitDefaultValue = false)]
        public int Offset { get; set; }

        /// <summary>
        /// Gets or Sets MaxMatches
        /// </summary>
        [DataMember(Name = "max_matches", EmitDefaultValue = false)]
        public int MaxMatches { get; set; }

        /// <summary>
        /// Gets or Sets Sort
        /// </summary>
        [DataMember(Name = "sort", EmitDefaultValue = false)]
        public List<Object> Sort { get; set; }

        /// <summary>
        /// Gets or Sets SortOld
        /// </summary>
        [DataMember(Name = "sort_old", EmitDefaultValue = false)]
        public List<Object> SortOld { get; set; }

        /// <summary>
        /// Gets or Sets Aggs
        /// </summary>
        [DataMember(Name = "aggs", EmitDefaultValue = false)]
        public List<Aggregation> Aggs { get; set; }

        /// <summary>
        /// Gets or Sets Expressions
        /// </summary>
        [DataMember(Name = "expressions", EmitDefaultValue = false)]
        public List<Object> Expressions { get; set; }

        /// <summary>
        /// Gets or Sets Highlight
        /// </summary>
        [DataMember(Name = "highlight", EmitDefaultValue = false)]
        public Highlight Highlight { get; set; }

        /// <summary>
        /// Gets or Sets Source
        /// </summary>
        [DataMember(Name = "source", EmitDefaultValue = false)]
        public Object Source { get; set; }

        /// <summary>
        /// Gets or Sets Options
        /// </summary>
        [DataMember(Name = "options", EmitDefaultValue = false)]
        public Dictionary<string, Object> Options { get; set; }

        /// <summary>
        /// Gets or Sets Profile
        /// </summary>
        [DataMember(Name = "profile", EmitDefaultValue = true)]
        public bool Profile { get; set; }

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
            sb.Append("class SearchRequest {\n");
            sb.Append("  Index: ").Append(Index).Append("\n");
            sb.Append("  Query: ").Append(Query).Append("\n");
            sb.Append("  FulltextFilter: ").Append(FulltextFilter).Append("\n");
            sb.Append("  AttrFilter: ").Append(AttrFilter).Append("\n");
            sb.Append("  Limit: ").Append(Limit).Append("\n");
            sb.Append("  Offset: ").Append(Offset).Append("\n");
            sb.Append("  MaxMatches: ").Append(MaxMatches).Append("\n");
            sb.Append("  Sort: ").Append(Sort).Append("\n");
            sb.Append("  SortOld: ").Append(SortOld).Append("\n");
            sb.Append("  Aggs: ").Append(Aggs).Append("\n");
            sb.Append("  Expressions: ").Append(Expressions).Append("\n");
            sb.Append("  Highlight: ").Append(Highlight).Append("\n");
            sb.Append("  Source: ").Append(Source).Append("\n");
            sb.Append("  Options: ").Append(Options).Append("\n");
            sb.Append("  Profile: ").Append(Profile).Append("\n");
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
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as SearchRequest);
        }

        /// <summary>
        /// Returns true if SearchRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of SearchRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(SearchRequest input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Index == input.Index ||
                    (this.Index != null &&
                    this.Index.Equals(input.Index))
                ) && 
                (
                    this.Query == input.Query ||
                    (this.Query != null &&
                    this.Query.Equals(input.Query))
                ) && 
                (
                    this.FulltextFilter == input.FulltextFilter ||
                    (this.FulltextFilter != null &&
                    this.FulltextFilter.Equals(input.FulltextFilter))
                ) && 
                (
                    this.AttrFilter == input.AttrFilter ||
                    (this.AttrFilter != null &&
                    this.AttrFilter.Equals(input.AttrFilter))
                ) && 
                (
                    this.Limit == input.Limit ||
                    this.Limit.Equals(input.Limit)
                ) && 
                (
                    this.Offset == input.Offset ||
                    this.Offset.Equals(input.Offset)
                ) && 
                (
                    this.MaxMatches == input.MaxMatches ||
                    this.MaxMatches.Equals(input.MaxMatches)
                ) && 
                (
                    this.Sort == input.Sort ||
                    this.Sort != null &&
                    input.Sort != null &&
                    this.Sort.SequenceEqual(input.Sort)
                ) && 
                (
                    this.SortOld == input.SortOld ||
                    this.SortOld != null &&
                    input.SortOld != null &&
                    this.SortOld.SequenceEqual(input.SortOld)
                ) && 
                (
                    this.Aggs == input.Aggs ||
                    this.Aggs != null &&
                    input.Aggs != null &&
                    this.Aggs.SequenceEqual(input.Aggs)
                ) && 
                (
                    this.Expressions == input.Expressions ||
                    this.Expressions != null &&
                    input.Expressions != null &&
                    this.Expressions.SequenceEqual(input.Expressions)
                ) && 
                (
                    this.Highlight == input.Highlight ||
                    (this.Highlight != null &&
                    this.Highlight.Equals(input.Highlight))
                ) && 
                (
                    this.Source == input.Source ||
                    (this.Source != null &&
                    this.Source.Equals(input.Source))
                ) && 
                (
                    this.Options == input.Options ||
                    this.Options != null &&
                    input.Options != null &&
                    this.Options.SequenceEqual(input.Options)
                ) && 
                (
                    this.Profile == input.Profile ||
                    this.Profile.Equals(input.Profile)
                ) && 
                (
                    this.TrackScores == input.TrackScores ||
                    this.TrackScores.Equals(input.TrackScores)
                );
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.Index != null)
                {
                    hashCode = (hashCode * 59) + this.Index.GetHashCode();
                }
                if (this.Query != null)
                {
                    hashCode = (hashCode * 59) + this.Query.GetHashCode();
                }
                if (this.FulltextFilter != null)
                {
                    hashCode = (hashCode * 59) + this.FulltextFilter.GetHashCode();
                }
                if (this.AttrFilter != null)
                {
                    hashCode = (hashCode * 59) + this.AttrFilter.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Limit.GetHashCode();
                hashCode = (hashCode * 59) + this.Offset.GetHashCode();
                hashCode = (hashCode * 59) + this.MaxMatches.GetHashCode();
                if (this.Sort != null)
                {
                    hashCode = (hashCode * 59) + this.Sort.GetHashCode();
                }
                if (this.SortOld != null)
                {
                    hashCode = (hashCode * 59) + this.SortOld.GetHashCode();
                }
                if (this.Aggs != null)
                {
                    hashCode = (hashCode * 59) + this.Aggs.GetHashCode();
                }
                if (this.Expressions != null)
                {
                    hashCode = (hashCode * 59) + this.Expressions.GetHashCode();
                }
                if (this.Highlight != null)
                {
                    hashCode = (hashCode * 59) + this.Highlight.GetHashCode();
                }
                if (this.Source != null)
                {
                    hashCode = (hashCode * 59) + this.Source.GetHashCode();
                }
                if (this.Options != null)
                {
                    hashCode = (hashCode * 59) + this.Options.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Profile.GetHashCode();
                hashCode = (hashCode * 59) + this.TrackScores.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
