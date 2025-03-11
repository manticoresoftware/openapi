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
    /// Object containing the document data for replacing an existing document in a table.
    /// </summary>
    [DataContract(Name = "replaceDocumentRequest")]
    public partial class ReplaceDocumentRequest : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ReplaceDocumentRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ReplaceDocumentRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="ReplaceDocumentRequest" /> class.
        /// </summary>
        /// <param name="Doc">Object containing the new document data to replace the existing one. (required).</param>
        public ReplaceDocumentRequest(Object Doc = default(Object))
        {
            // to ensure "Doc" is required (not null)
            if (Doc == null)
            {
                throw new ArgumentNullException("Doc is a required property for ReplaceDocumentRequest and cannot be null");
            }
            this.Doc = Doc;
        }

        /// <summary>
        /// Object containing the new document data to replace the existing one.
        /// </summary>
        /// <value>Object containing the new document data to replace the existing one.</value>
        [DataMember(Name = "doc", IsRequired = true, EmitDefaultValue = true)]
        public Object Doc { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ReplaceDocumentRequest {\n");
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
