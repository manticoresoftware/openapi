/**
 * Manticore Search Client
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 4.0.0
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
    instance = new Manticoresearch.BasicSearchRequest();
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

  describe('BasicSearchRequest', function() {
    it('should create an instance of BasicSearchRequest', function() {
      // uncomment below and update the code to test BasicSearchRequest
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be.a(Manticoresearch.BasicSearchRequest);
    });

    it('should have the property aggs (base name: "aggs")', function() {
      // uncomment below and update the code to test the property aggs
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property expressions (base name: "expressions")', function() {
      // uncomment below and update the code to test the property expressions
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property join (base name: "join")', function() {
      // uncomment below and update the code to test the property join
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property highlight (base name: "highlight")', function() {
      // uncomment below and update the code to test the property highlight
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property index (base name: "index")', function() {
      // uncomment below and update the code to test the property index
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property limit (base name: "limit")', function() {
      // uncomment below and update the code to test the property limit
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property maxMatches (base name: "max_matches")', function() {
      // uncomment below and update the code to test the property maxMatches
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property offset (base name: "offset")', function() {
      // uncomment below and update the code to test the property offset
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property options (base name: "options")', function() {
      // uncomment below and update the code to test the property options
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property profile (base name: "profile")', function() {
      // uncomment below and update the code to test the property profile
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property sort (base name: "sort")', function() {
      // uncomment below and update the code to test the property sort
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property source (base name: "_source")', function() {
      // uncomment below and update the code to test the property source
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property trackScores (base name: "track_scores")', function() {
      // uncomment below and update the code to test the property trackScores
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

    it('should have the property query (base name: "query")', function() {
      // uncomment below and update the code to test the property query
      //var instance = new Manticoresearch.BasicSearchRequest();
      //expect(instance).to.be();
    });

  });

}));
