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
    /// Response object for successful delete request
    /// </summary>
    [DataContract(Name = "deleteResponse")]
    public partial class DeleteResponse : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="DeleteResponse" /> class.
        /// </summary>
        /// <param name="">The name of the index from which the document was deleted.</param>
        /// <param name="">Number of documents deleted.</param>
        /// <param name="">The ID of the deleted document. If multiple documents are deleted, the ID of the first deleted document is returned.</param>
        /// <param name="">Indicates whether any documents to be deleted were found.</param>
        /// <param name="">Result of the delete operation, typically &#39;deleted&#39;.</param>
        public DeleteResponse(string  = default(string), int  = default(int), long  = default(long), bool  = default(bool), string  = default(string))
        {
            this.Index = ;
            this.Deleted = ;
            this.Id = ;
            this.Found = ;
            this.Result = ;
        }

        /// <summary>
        /// The name of the index from which the document was deleted
        /// </summary>
        /// <value>The name of the index from which the document was deleted</value>
        [DataMember(Name = "_index", EmitDefaultValue = false)]
        public string Index { get; set; }

        /// <summary>
        /// Number of documents deleted
        /// </summary>
        /// <value>Number of documents deleted</value>
        [DataMember(Name = "deleted", EmitDefaultValue = false)]
        public int Deleted { get; set; }

        /// <summary>
        /// The ID of the deleted document. If multiple documents are deleted, the ID of the first deleted document is returned
        /// </summary>
        /// <value>The ID of the deleted document. If multiple documents are deleted, the ID of the first deleted document is returned</value>
        [DataMember(Name = "_id", EmitDefaultValue = false)]
        public long Id { get; set; }

        /// <summary>
        /// Indicates whether any documents to be deleted were found
        /// </summary>
        /// <value>Indicates whether any documents to be deleted were found</value>
        [DataMember(Name = "found", EmitDefaultValue = true)]
        public bool Found { get; set; }

        /// <summary>
        /// Result of the delete operation, typically &#39;deleted&#39;
        /// </summary>
        /// <value>Result of the delete operation, typically &#39;deleted&#39;</value>
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
            sb.Append("  Found: ").Append(Found).Append("\n");
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
