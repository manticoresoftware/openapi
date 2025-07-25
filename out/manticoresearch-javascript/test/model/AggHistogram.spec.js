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
    instance = new Manticoresearch.AggHistogram();
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

  describe('AggHistogram', function() {
    it('should create an instance of AggHistogram', function() {
      // uncomment below and update the code to test AggHistogram
      //var instance = new Manticoresearch.AggHistogram();
      //expect(instance).to.be.a(Manticoresearch.AggHistogram);
    });

    it('should have the property field (base name: "field")', function() {
      // uncomment below and update the code to test the property field
      //var instance = new Manticoresearch.AggHistogram();
      //expect(instance).to.be();
    });

    it('should have the property interval (base name: "interval")', function() {
      // uncomment below and update the code to test the property interval
      //var instance = new Manticoresearch.AggHistogram();
      //expect(instance).to.be();
    });

    it('should have the property offset (base name: "offset")', function() {
      // uncomment below and update the code to test the property offset
      //var instance = new Manticoresearch.AggHistogram();
      //expect(instance).to.be();
    });

    it('should have the property keyed (base name: "keyed")', function() {
      // uncomment below and update the code to test the property keyed
      //var instance = new Manticoresearch.AggHistogram();
      //expect(instance).to.be();
    });

  });

}));
