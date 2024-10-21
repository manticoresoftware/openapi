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
    /// Object containing data for inserting a new document into the index 
    /// </summary>
    [DataContract(Name = "insertDocumentRequest")]
    public partial class InsertDocumentRequest : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="InsertDocumentRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected InsertDocumentRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="InsertDocumentRequest" /> class.
        /// </summary>
        /// <param name="">Name of the index to insert the document into (required).</param>
        /// <param name="">Name of the cluster to insert the document into.</param>
        /// <param name="">Document ID. If not provided, an ID will be auto-generated .</param>
        /// <param name="">Object containing document data  (required).</param>
        public InsertDocumentRequest(string  = default(string), string  = default(string), long  = default(long), Object  = default(Object))
        {
            // to ensure "" is required (not null)
            if ( == null)
            {
                throw new ArgumentNullException(" is a required property for InsertDocumentRequest and cannot be null");
            }
            this.Index = ;
            // to ensure "" is required (not null)
            if ( == null)
            {
                throw new ArgumentNullException(" is a required property for InsertDocumentRequest and cannot be null");
            }
            this.Doc = ;
            this.Cluster = ;
            this.Id = ;
        }

        /// <summary>
        /// Name of the index to insert the document into
        /// </summary>
        /// <value>Name of the index to insert the document into</value>
        [DataMember(Name = "index", IsRequired = true, EmitDefaultValue = true)]
        public string Index { get; set; }

        /// <summary>
        /// Name of the cluster to insert the document into
        /// </summary>
        /// <value>Name of the cluster to insert the document into</value>
        [DataMember(Name = "cluster", EmitDefaultValue = false)]
        public string Cluster { get; set; }

        /// <summary>
        /// Document ID. If not provided, an ID will be auto-generated 
        /// </summary>
        /// <value>Document ID. If not provided, an ID will be auto-generated </value>
        [DataMember(Name = "id", EmitDefaultValue = false)]
        public long Id { get; set; }

        /// <summary>
        /// Object containing document data 
        /// </summary>
        /// <value>Object containing document data </value>
        [DataMember(Name = "doc", IsRequired = true, EmitDefaultValue = true)]
        public Object Doc { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class InsertDocumentRequest {\n");
            sb.Append("  Index: ").Append(Index).Append("\n");
            sb.Append("  Cluster: ").Append(Cluster).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Doc: ").Append(Doc).Append("\n");
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
