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
    /// Object with documents to percolate
    /// </summary>
    [DataContract(Name = "percolateRequest")]
    public partial class PercolateRequest : IEquatable<PercolateRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="PercolateRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected PercolateRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="PercolateRequest" /> class.
        /// </summary>
        /// <param name="query">query (required).</param>
        public PercolateRequest(Dictionary<string, Object> query = default(Dictionary<string, Object>))
        {
            // to ensure "query" is required (not null)
            if (query == null)
            {
                throw new ArgumentNullException("query is a required property for PercolateRequest and cannot be null");
            }
            this.Query = query;
        }

        /// <summary>
        /// Gets or Sets Query
        /// </summary>
        [DataMember(Name = "query", IsRequired = true, EmitDefaultValue = false)]
        public Dictionary<string, Object> Query { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class PercolateRequest {\n");
            sb.Append("  Query: ").Append(Query).Append("\n");
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
            return this.Equals(input as PercolateRequest);
        }

        /// <summary>
        /// Returns true if PercolateRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of PercolateRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(PercolateRequest input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Query == input.Query ||
                    this.Query != null &&
                    input.Query != null &&
                    this.Query.SequenceEqual(input.Query)
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
                if (this.Query != null)
                {
                    hashCode = (hashCode * 59) + this.Query.GetHashCode();
                }
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
