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
    /// Payload for updating a document or multiple documents in a table
    /// </summary>
    [DataContract(Name = "updateDocumentRequest")]
    public partial class UpdateDocumentRequest : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="UpdateDocumentRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected UpdateDocumentRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="UpdateDocumentRequest" /> class.
        /// </summary>
        /// <param name="Table">Name of the document table (required).</param>
        /// <param name="Cluster">Name of the document cluster.</param>
        /// <param name="Doc">Object containing the document fields to update (required).</param>
        /// <param name="Id">Document ID.</param>
        /// <param name="Query">Query.</param>
        public UpdateDocumentRequest(string Table = default(string), string Cluster = default(string), Object Doc = default(Object), long Id = default(long), QueryFilter Query = default(QueryFilter))
        {
            // to ensure "Table" is required (not null)
            if (Table == null)
            {
                throw new ArgumentNullException("Table is a required property for UpdateDocumentRequest and cannot be null");
            }
            this.Table = Table;
            // to ensure "Doc" is required (not null)
            if (Doc == null)
            {
                throw new ArgumentNullException("Doc is a required property for UpdateDocumentRequest and cannot be null");
            }
            this.Doc = Doc;
            this.Cluster = Cluster;
            this.Id = Id;
            this.Query = Query;
        }

        /// <summary>
        /// Name of the document table
        /// </summary>
        /// <value>Name of the document table</value>
        [DataMember(Name = "table", IsRequired = true, EmitDefaultValue = true)]
        public string Table { get; set; }

        /// <summary>
        /// Name of the document cluster
        /// </summary>
        /// <value>Name of the document cluster</value>
        [DataMember(Name = "cluster", EmitDefaultValue = false)]
        public string Cluster { get; set; }

        /// <summary>
        /// Object containing the document fields to update
        /// </summary>
        /// <value>Object containing the document fields to update</value>
        /// <example>{gid&#x3D;10}</example>
        [DataMember(Name = "doc", IsRequired = true, EmitDefaultValue = true)]
        public Object Doc { get; set; }

        /// <summary>
        /// Document ID
        /// </summary>
        /// <value>Document ID</value>
        [DataMember(Name = "id", EmitDefaultValue = false)]
        public long Id { get; set; }

        /// <summary>
        /// Gets or Sets Query
        /// </summary>
        [DataMember(Name = "query", EmitDefaultValue = true)]
        public QueryFilter Query { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class UpdateDocumentRequest {\n");
            sb.Append("  Table: ").Append(Table).Append("\n");
            sb.Append("  Cluster: ").Append(Cluster).Append("\n");
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
