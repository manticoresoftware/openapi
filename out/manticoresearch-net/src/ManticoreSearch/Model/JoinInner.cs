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
    /// JoinInner
    /// </summary>
    [DataContract(Name = "join_inner")]
    public partial class JoinInner : IValidatableObject
    {
        /// <summary>
        /// Defines Type
        /// </summary>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum TypeEnum
        {
            /// <summary>
            /// Enum Inner for value: inner
            /// </summary>
            [EnumMember(Value = "inner")]
            Inner = 1,

            /// <summary>
            /// Enum Left for value: left
            /// </summary>
            [EnumMember(Value = "left")]
            Left = 2
        }


        /// <summary>
        /// Gets or Sets Type
        /// </summary>
        [DataMember(Name = "type", IsRequired = true, EmitDefaultValue = true)]
        public TypeEnum Type { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="JoinInner" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected JoinInner() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="JoinInner" /> class.
        /// </summary>
        /// <param name="on">on (required).</param>
        /// <param name="query">query.</param>
        /// <param name="table">table (required).</param>
        /// <param name="type">type (required).</param>
        public JoinInner(List<JoinInnerOnInner> on = default(List<JoinInnerOnInner>), FulltextFilter query = default(FulltextFilter), string table = default(string), TypeEnum type = default(TypeEnum))
        {
            // to ensure "on" is required (not null)
            if (on == null)
            {
                throw new ArgumentNullException("on is a required property for JoinInner and cannot be null");
            }
            this.On = on;
            // to ensure "table" is required (not null)
            if (table == null)
            {
                throw new ArgumentNullException("table is a required property for JoinInner and cannot be null");
            }
            this.Table = table;
            this.Type = type;
            this.Query = query;
        }

        /// <summary>
        /// Gets or Sets On
        /// </summary>
        [DataMember(Name = "on", IsRequired = true, EmitDefaultValue = true)]
        public List<JoinInnerOnInner> On { get; set; }

        /// <summary>
        /// Gets or Sets Query
        /// </summary>
        [DataMember(Name = "query", EmitDefaultValue = false)]
        public FulltextFilter Query { get; set; }

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
            sb.Append("class JoinInner {\n");
            sb.Append("  On: ").Append(On).Append("\n");
            sb.Append("  Query: ").Append(Query).Append("\n");
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