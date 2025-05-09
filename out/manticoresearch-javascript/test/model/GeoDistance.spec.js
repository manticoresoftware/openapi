/**
 * Manticore Search Client
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 5.0.0
 * Contact: info@manticoresearch.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Manticoresearch);
  }
}(this, function(expect, Manticoresearch) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Manticoresearch.GeoDistance();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('GeoDistance', function() {
    it('should create an instance of GeoDistance', function() {
      // uncomment below and update the code to test GeoDistance
      //var instance = new Manticoresearch.GeoDistance();
      //expect(instance).to.be.a(Manticoresearch.GeoDistance);
    });

    it('should have the property locationAnchor (base name: "location_anchor")', function() {
      // uncomment below and update the code to test the property locationAnchor
      //var instance = new Manticoresearch.GeoDistance();
      //expect(instance).to.be();
    });

    it('should have the property locationSource (base name: "location_source")', function() {
      // uncomment below and update the code to test the property locationSource
      //var instance = new Manticoresearch.GeoDistance();
      //expect(instance).to.be();
    });

    it('should have the property distanceType (base name: "distance_type")', function() {
      // uncomment below and update the code to test the property distanceType
      //var instance = new Manticoresearch.GeoDistance();
      //expect(instance).to.be();
    });

    it('should have the property distance (base name: "distance")', function() {
      // uncomment below and update the code to test the property distance
      //var instance = new Manticoresearch.GeoDistance();
      //expect(instance).to.be();
    });

  });

}));
