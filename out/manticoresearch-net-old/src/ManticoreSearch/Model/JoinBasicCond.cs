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
    /// JoinBasicCond
    /// </summary>
    [DataContract(Name = "joinBasicCond")]
    public partial class JoinBasicCond : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="JoinBasicCond" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected JoinBasicCond() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="JoinBasicCond" /> class.
        /// </summary>
        /// <param name="field">field (required).</param>
        /// <param name="table">table (required).</param>
        public JoinBasicCond(string field = default(string), string table = default(string))
        {
            // to ensure "field" is required (not null)
            if (field == null)
            {
                throw new ArgumentNullException("field is a required property for JoinBasicCond and cannot be null");
            }
            this.Field = field;
            // to ensure "table" is required (not null)
            if (table == null)
            {
                throw new ArgumentNullException("table is a required property for JoinBasicCond and cannot be null");
            }
            this.Table = table;
        }

        /// <summary>
        /// Gets or Sets Field
        /// </summary>
        [DataMember(Name = "field", IsRequired = true, EmitDefaultValue = true)]
        public string Field { get; set; }

        /// <summary>
        /// Gets or Sets Table
        /// </summary>
        [DataMember(Name = "table", IsRequired = true, EmitDefaultValue = true)]
        public string Table { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class JoinBasicCond {\n");
            sb.Append("  Field: ").Append(Field).Append("\n");
            sb.Append("  Table: ").Append(Table).Append("\n");
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
