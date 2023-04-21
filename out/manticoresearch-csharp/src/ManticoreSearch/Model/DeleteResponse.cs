/*
 * Manticore Search Client
 *
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 3.3.0
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
    /// Success response
    /// </summary>
    [DataContract(Name = "deleteResponse")]
    public partial class DeleteResponse : IEquatable<DeleteResponse>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="DeleteResponse" /> class.
        /// </summary>
        /// <param name="index">index.</param>
        /// <param name="deleted">deleted.</param>
        /// <param name="id">id.</param>
        /// <param name="result">result.</param>
        public DeleteResponse(string index = default(string), int deleted = default(int), long id = default(long), string result = default(string))
        {
            this.Index = index;
            this.Deleted = deleted;
            this.Id = id;
            this.Result = result;
        }

        /// <summary>
        /// Gets or Sets Index
        /// </summary>
        [DataMember(Name = "_index", EmitDefaultValue = false)]
        public string Index { get; set; }

        /// <summary>
        /// Gets or Sets Deleted
        /// </summary>
        [DataMember(Name = "deleted", EmitDefaultValue = false)]
        public int Deleted { get; set; }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "_id", EmitDefaultValue = false)]
        public long Id { get; set; }

        /// <summary>
        /// Gets or Sets Result
        /// </summary>
        [DataMember(Name = "result", EmitDefaultValue = false)]
        public string Result { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class DeleteResponse {\n");
            sb.Append("  Index: ").Append(Index).Append("\n");
            sb.Append("  Deleted: ").Append(Deleted).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Result: ").Append(Result).Append("\n");
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
            return this.Equals(input as DeleteResponse);
        }

        /// <summary>
        /// Returns true if DeleteResponse instances are equal
        /// </summary>
        /// <param name="input">Instance of DeleteResponse to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(DeleteResponse input)
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
                    this.Deleted == input.Deleted ||
                    this.Deleted.Equals(input.Deleted)
                ) && 
                (
                    this.Id == input.Id ||
                    this.Id.Equals(input.Id)
                ) && 
                (
                    this.Result == input.Result ||
                    (this.Result != null &&
                    this.Result.Equals(input.Result))
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
                hashCode = (hashCode * 59) + this.Deleted.GetHashCode();
                hashCode = (hashCode * 59) + this.Id.GetHashCode();
                if (this.Result != null)
                {
                    hashCode = (hashCode * 59) + this.Result.GetHashCode();
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
