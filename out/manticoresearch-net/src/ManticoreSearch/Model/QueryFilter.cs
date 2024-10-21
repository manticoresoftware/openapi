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
    /// Object used to apply various conditions, such as full-text matching or attribute filtering, to a search query
    /// </summary>
    [DataContract(Name = "queryFilter")]
    public partial class QueryFilter : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="QueryFilter" /> class.
        /// </summary>
        /// <param name="">Filter object defining a query string.</param>
        /// <param name="">Filter object defining a match keyword.</param>
        /// <param name="">Filter object defining a match phrase.</param>
        /// <param name="">Filter object to select all documents.</param>
        /// <param name="">.</param>
        /// <param name="">.</param>
        /// <param name="">Filter to match a given set of attribute values..</param>
        /// <param name="">Filter to match a given range of attribute values..</param>
        /// <param name="">.</param>
        public QueryFilter(Object  = default(Object), Object  = default(Object), Object  = default(Object), Object  = default(Object), BoolFilter  = default(BoolFilter), Object  = default(Object), Object  = default(Object), Object  = default(Object), GeoDistance  = default(GeoDistance))
        {
            this.QueryString = ;
            this.Match = ;
            this.MatchPhrase = ;
            this.MatchAll = ;
            this.VarBool = ;
            this.PropertyEquals = ;
            this.VarIn = ;
            this.Range = ;
            this.GeoDistance = ;
        }

        /// <summary>
        /// Filter object defining a query string
        /// </summary>
        /// <value>Filter object defining a query string</value>
        [DataMember(Name = "query_string", EmitDefaultValue = true)]
        public Object QueryString { get; set; }

        /// <summary>
        /// Filter object defining a match keyword
        /// </summary>
        /// <value>Filter object defining a match keyword</value>
        [DataMember(Name = "match", EmitDefaultValue = true)]
        public Object Match { get; set; }

        /// <summary>
        /// Filter object defining a match phrase
        /// </summary>
        /// <value>Filter object defining a match phrase</value>
        [DataMember(Name = "match_phrase", EmitDefaultValue = true)]
        public Object MatchPhrase { get; set; }

        /// <summary>
        /// Filter object to select all documents
        /// </summary>
        /// <value>Filter object to select all documents</value>
        [DataMember(Name = "match_all", EmitDefaultValue = true)]
        public Object MatchAll { get; set; }

        /// <summary>
        /// Gets or Sets VarBool
        /// </summary>
        [DataMember(Name = "bool", EmitDefaultValue = false)]
        public BoolFilter VarBool { get; set; }

        /// <summary>
        /// Gets or Sets PropertyEquals
        /// </summary>
        [DataMember(Name = "equals", EmitDefaultValue = true)]
        public Object PropertyEquals { get; set; }

        /// <summary>
        /// Filter to match a given set of attribute values.
        /// </summary>
        /// <value>Filter to match a given set of attribute values.</value>
        [DataMember(Name = "in", EmitDefaultValue = false)]
        public Object VarIn { get; set; }

        /// <summary>
        /// Filter to match a given range of attribute values.
        /// </summary>
        /// <value>Filter to match a given range of attribute values.</value>
        [DataMember(Name = "range", EmitDefaultValue = false)]
        public Object Range { get; set; }

        /// <summary>
        /// Gets or Sets GeoDistance
        /// </summary>
        [DataMember(Name = "geo_distance", EmitDefaultValue = false)]
        public GeoDistance GeoDistance { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class QueryFilter {\n");
            sb.Append("  QueryString: ").Append(QueryString).Append("\n");
            sb.Append("  Match: ").Append(Match).Append("\n");
            sb.Append("  MatchPhrase: ").Append(MatchPhrase).Append("\n");
            sb.Append("  MatchAll: ").Append(MatchAll).Append("\n");
            sb.Append("  VarBool: ").Append(VarBool).Append("\n");
            sb.Append("  PropertyEquals: ").Append(PropertyEquals).Append("\n");
            sb.Append("  VarIn: ").Append(VarIn).Append("\n");
            sb.Append("  Range: ").Append(Range).Append("\n");
            sb.Append("  GeoDistance: ").Append(GeoDistance).Append("\n");
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
