/*
 * Manticore Search Client
 * Copyright (c) 2020-2021, Manticore Software LTD (https://manticoresearch.com)
 *
 * All rights reserved
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
    // create a new instance
    //instance = new Manticoresearch.AggregationComposite();
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

  describe('AggregationComposite', function() {
    it('should create an instance of AggregationComposite', function() {
      // uncomment below and update the code to test AggregationComposite
      //var instance = new Manticoresearch.AggregationComposite();
      //expect(instance).to.be.a(Manticoresearch.AggregationComposite);
    });

    it('should have the property size (base name: "size")', function() {
      // uncomment below and update the code to test the property size
      //var instance = new Manticoresearch.AggregationComposite();
      //expect(instance).to.be();
    });

    it('should have the property sources (base name: "sources")', function() {
      // uncomment below and update the code to test the property sources
      //var instance = new Manticoresearch.AggregationComposite();
      //expect(instance).to.be();
    });

  });

}));
