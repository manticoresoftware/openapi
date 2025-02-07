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
    /// Object representing the conditions used to perform the join operation
    /// </summary>
    [DataContract(Name = "joinCond")]
    public partial class JoinCond : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="JoinCond" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected JoinCond() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="JoinCond" /> class.
        /// </summary>
        /// <param name="Field">Field to join on (required).</param>
        /// <param name="Table">Joined table (required).</param>
        /// <param name="Type">Type.</param>
        public JoinCond(string Field = default(string), string Table = default(string), Object Type = default(Object))
        {
            // to ensure "Field" is required (not null)
            if (Field == null)
            {
                throw new ArgumentNullException("Field is a required property for JoinCond and cannot be null");
            }
            this.Field = Field;
            // to ensure "Table" is required (not null)
            if (Table == null)
            {
                throw new ArgumentNullException("Table is a required property for JoinCond and cannot be null");
            }
            this.Table = Table;
            this.Type = Type;
        }

        /// <summary>
        /// Field to join on
        /// </summary>
        /// <value>Field to join on</value>
        [DataMember(Name = "field", IsRequired = true, EmitDefaultValue = true)]
        public string Field { get; set; }

        /// <summary>
        /// Joined table
        /// </summary>
        /// <value>Joined table</value>
        [DataMember(Name = "table", IsRequired = true, EmitDefaultValue = true)]
        public string Table { get; set; }

        /// <summary>
        /// Gets or Sets Type
        /// </summary>
        [DataMember(Name = "type", EmitDefaultValue = true)]
        public Object Type { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class JoinCond {\n");
            sb.Append("  Field: ").Append(Field).Append("\n");
            sb.Append("  Table: ").Append(Table).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
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
