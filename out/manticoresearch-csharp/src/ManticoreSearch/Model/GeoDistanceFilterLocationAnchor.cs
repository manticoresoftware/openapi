/*
 * Manticore Search Client
 *
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 3.3.1
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
    /// Geo pin point object
    /// </summary>
    [DataContract(Name = "geoDistanceFilter_location_anchor")]
    public partial class GeoDistanceFilterLocationAnchor : IEquatable<GeoDistanceFilterLocationAnchor>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="GeoDistanceFilterLocationAnchor" /> class.
        /// </summary>
        /// <param name="lat">Geo latitude of pin point in degrees.</param>
        /// <param name="lon">Geo longitude pf pin point in degrees.</param>
        public GeoDistanceFilterLocationAnchor(decimal lat = default(decimal), decimal lon = default(decimal))
        {
            this.Lat = lat;
            this.Lon = lon;
        }

        /// <summary>
        /// Geo latitude of pin point in degrees
        /// </summary>
        /// <value>Geo latitude of pin point in degrees</value>
        [DataMember(Name = "lat", EmitDefaultValue = false)]
        public decimal Lat { get; set; }

        /// <summary>
        /// Geo longitude pf pin point in degrees
        /// </summary>
        /// <value>Geo longitude pf pin point in degrees</value>
        [DataMember(Name = "lon", EmitDefaultValue = false)]
        public decimal Lon { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class GeoDistanceFilterLocationAnchor {\n");
            sb.Append("  Lat: ").Append(Lat).Append("\n");
            sb.Append("  Lon: ").Append(Lon).Append("\n");
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
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as GeoDistanceFilterLocationAnchor);
        }

        /// <summary>
        /// Returns true if GeoDistanceFilterLocationAnchor instances are equal
        /// </summary>
        /// <param name="input">Instance of GeoDistanceFilterLocationAnchor to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(GeoDistanceFilterLocationAnchor input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Lat == input.Lat ||
                    this.Lat.Equals(input.Lat)
                ) && 
                (
                    this.Lon == input.Lon ||
                    this.Lon.Equals(input.Lon)
                );
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                hashCode = (hashCode * 59) + this.Lat.GetHashCode();
                hashCode = (hashCode * 59) + this.Lon.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
