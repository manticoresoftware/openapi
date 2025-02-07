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
    /// Response object indicating the success of an operation, such as inserting or updating a document
    /// </summary>
    [DataContract(Name = "successResponse")]
    public partial class SuccessResponse : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="SuccessResponse" /> class.
        /// </summary>
        /// <param name="Index">Name of the document index.</param>
        /// <param name="Table">Name of the document table (alias of index).</param>
        /// <param name="Id">ID of the document affected by the request operation.</param>
        /// <param name="Created">Indicates whether the document was created as a result of the operation.</param>
        /// <param name="Result">Result of the operation, typically &#39;created&#39;, &#39;updated&#39;, or &#39;deleted&#39;.</param>
        /// <param name="Found">Indicates whether the document was found in the index.</param>
        /// <param name="Status">HTTP status code representing the result of the operation.</param>
        public SuccessResponse(string Index = default(string), string Table = default(string), long Id = default(long), bool Created = default(bool), string Result = default(string), bool Found = default(bool), int Status = default(int))
        {
            this.Index = Index;
            this.Table = Table;
            this.Id = Id;
            this.Created = Created;
            this.Result = Result;
            this.Found = Found;
            this.Status = Status;
        }

        /// <summary>
        /// Name of the document index
        /// </summary>
        /// <value>Name of the document index</value>
        [DataMember(Name = "_index", EmitDefaultValue = false)]
        public string Index { get; set; }

        /// <summary>
        /// Name of the document table (alias of index)
        /// </summary>
        /// <value>Name of the document table (alias of index)</value>
        [DataMember(Name = "table", EmitDefaultValue = false)]
        public string Table { get; set; }

        /// <summary>
        /// ID of the document affected by the request operation
        /// </summary>
        /// <value>ID of the document affected by the request operation</value>
        [DataMember(Name = "_id", EmitDefaultValue = false)]
        public long Id { get; set; }

        /// <summary>
        /// Indicates whether the document was created as a result of the operation
        /// </summary>
        /// <value>Indicates whether the document was created as a result of the operation</value>
        [DataMember(Name = "created", EmitDefaultValue = true)]
        public bool Created { get; set; }

        /// <summary>
        /// Result of the operation, typically &#39;created&#39;, &#39;updated&#39;, or &#39;deleted&#39;
        /// </summary>
        /// <value>Result of the operation, typically &#39;created&#39;, &#39;updated&#39;, or &#39;deleted&#39;</value>
        [DataMember(Name = "result", EmitDefaultValue = false)]
        public string Result { get; set; }

        /// <summary>
        /// Indicates whether the document was found in the index
        /// </summary>
        /// <value>Indicates whether the document was found in the index</value>
        [DataMember(Name = "found", EmitDefaultValue = true)]
        public bool Found { get; set; }

        /// <summary>
        /// HTTP status code representing the result of the operation
        /// </summary>
        /// <value>HTTP status code representing the result of the operation</value>
        [DataMember(Name = "status", EmitDefaultValue = false)]
        public int Status { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class SuccessResponse {\n");
            sb.Append("  Index: ").Append(Index).Append("\n");
            sb.Append("  Table: ").Append(Table).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Created: ").Append(Created).Append("\n");
            sb.Append("  Result: ").Append(Result).Append("\n");
            sb.Append("  Found: ").Append(Found).Append("\n");
            sb.Append("  Status: ").Append(Status).Append("\n");
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
