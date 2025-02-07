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
    /// JoinOn
    /// </summary>
    [DataContract(Name = "joinOn")]
    public partial class JoinOn : IValidatableObject
    {
        /// <summary>
        /// Defines VarOperator
        /// </summary>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum OperatorEnum
        {
            /// <summary>
            /// Enum Eq for value: eq
            /// </summary>
            [EnumMember(Value = "eq")]
            Eq = 1
        }


        /// <summary>
        /// Gets or Sets VarOperator
        /// </summary>
        [DataMember(Name = "operator", EmitDefaultValue = false)]
        public OperatorEnum? VarOperator { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="JoinOn" /> class.
        /// </summary>
        /// <param name="Right">Right.</param>
        /// <param name="Left">Left.</param>
        /// <param name="VarOperator">VarOperator.</param>
        public JoinOn(JoinCond Right = default(JoinCond), JoinCond Left = default(JoinCond), OperatorEnum? VarOperator = default(OperatorEnum?))
        {
            this.Right = Right;
            this.Left = Left;
            this.VarOperator = VarOperator;
        }

        /// <summary>
        /// Gets or Sets Right
        /// </summary>
        [DataMember(Name = "right", EmitDefaultValue = false)]
        public JoinCond Right { get; set; }

        /// <summary>
        /// Gets or Sets Left
        /// </summary>
        [DataMember(Name = "left", EmitDefaultValue = false)]
        public JoinCond Left { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class JoinOn {\n");
            sb.Append("  Right: ").Append(Right).Append("\n");
            sb.Append("  Left: ").Append(Left).Append("\n");
            sb.Append("  VarOperator: ").Append(VarOperator).Append("\n");
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
