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
    /// Success response returned after updating one or more documents
    /// </summary>
    [DataContract(Name = "updateResponse")]
    public partial class UpdateResponse : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="UpdateResponse" /> class.
        /// </summary>
        /// <param name="">Name of the document index.</param>
        /// <param name="">Number of documents updated.</param>
        /// <param name="">Document ID.</param>
        /// <param name="">Result of the update operation, typically &#39;updated&#39;.</param>
        public UpdateResponse(string  = default(string), int  = default(int), long  = default(long), string  = default(string))
        {
            this.Index = ;
            this.Updated = ;
            this.Id = ;
            this.Result = ;
        }

        /// <summary>
        /// Name of the document index
        /// </summary>
        /// <value>Name of the document index</value>
        [DataMember(Name = "_index", EmitDefaultValue = false)]
        public string Index { get; set; }

        /// <summary>
        /// Number of documents updated
        /// </summary>
        /// <value>Number of documents updated</value>
        [DataMember(Name = "updated", EmitDefaultValue = false)]
        public int Updated { get; set; }

        /// <summary>
        /// Document ID
        /// </summary>
        /// <value>Document ID</value>
        [DataMember(Name = "_id", EmitDefaultValue = false)]
        public long Id { get; set; }

        /// <summary>
        /// Result of the update operation, typically &#39;updated&#39;
        /// </summary>
        /// <value>Result of the update operation, typically &#39;updated&#39;</value>
        [DataMember(Name = "result", EmitDefaultValue = false)]
        public string Result { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class UpdateResponse {\n");
            sb.Append("  Index: ").Append(Index).Append("\n");
            sb.Append("  Updated: ").Append(Updated).Append("\n");
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