/* 
 * Manticore Search API
 *
 * This is the API for Manticore Search HTTP protocol 
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: adrian.nuta@manticoresearch.com
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */

using System;
using System.Linq;
using System.IO;
using System.Text;
using System.Text.RegularExpressions;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Runtime.Serialization;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// UpdateDocumentRequest
    /// </summary>
    [DataContract]
    public partial class UpdateDocumentRequest :  IEquatable<UpdateDocumentRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="UpdateDocumentRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected UpdateDocumentRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="UpdateDocumentRequest" /> class.
        /// </summary>
        /// <param name="index">index (required).</param>
        /// <param name="doc">doc (required).</param>
        /// <param name="id">id.</param>
        /// <param name="query">query.</param>
        public UpdateDocumentRequest(string index = default(string), Object doc = default(Object), long id = default(long), Object query = default(Object))
        {
            // to ensure "index" is required (not null)
            if (index == null)
            {
                throw new InvalidDataException("index is a required property for UpdateDocumentRequest and cannot be null");
            }
            else
            {
                this.Index = index;
            }
            
            // to ensure "doc" is required (not null)
            if (doc == null)
            {
                throw new InvalidDataException("doc is a required property for UpdateDocumentRequest and cannot be null");
            }
            else
            {
                this.Doc = doc;
            }
            
            this.Id = id;
            this.Query = query;
        }
        
        /// <summary>
        /// Gets or Sets Index
        /// </summary>
        [DataMember(Name="index", EmitDefaultValue=true)]
        public string Index { get; set; }

        /// <summary>
        /// Gets or Sets Doc
        /// </summary>
        [DataMember(Name="doc", EmitDefaultValue=true)]
        public Object Doc { get; set; }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name="id", EmitDefaultValue=false)]
        public long Id { get; set; }

        /// <summary>
        /// Gets or Sets Query
        /// </summary>
        [DataMember(Name="query", EmitDefaultValue=false)]
        public Object Query { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class UpdateDocumentRequest {\n");
            sb.Append("  Index: ").Append(Index).Append("\n");
            sb.Append("  Doc: ").Append(Doc).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
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
            return JsonConvert.SerializeObject(this, Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as UpdateDocumentRequest);
        }

        /// <summary>
        /// Returns true if UpdateDocumentRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of UpdateDocumentRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(UpdateDocumentRequest input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Index == input.Index ||
                    (this.Index != null &&
                    this.Index.Equals(input.Index))
                ) && 
                (
                    this.Doc == input.Doc ||
                    (this.Doc != null &&
                    this.Doc.Equals(input.Doc))
                ) && 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
                ) && 
                (
                    this.Query == input.Query ||
                    (this.Query != null &&
                    this.Query.Equals(input.Query))
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
                    hashCode = hashCode * 59 + this.Index.GetHashCode();
                if (this.Doc != null)
                    hashCode = hashCode * 59 + this.Doc.GetHashCode();
                if (this.Id != null)
                    hashCode = hashCode * 59 + this.Id.GetHashCode();
                if (this.Query != null)
                    hashCode = hashCode * 59 + this.Query.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {

            
            // Id (long) minimum
            if(this.Id < (long)1)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for Id, must be a value greater than or equal to 1.", new [] { "Id" });
            }

            yield break;
        }
    }

}