/*
 * Manticore Search Client
 * Copyright (c) 2020-2021, Manticore Software LTD (https://manticoresearch.com)
 *
 * All rights reserved
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/Aggregation', 'model/Highlight', 'model/KnnSearchRequestByDocId', 'model/KnnSearchRequestByVector', 'model/SearchRequest'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Aggregation'), require('./Highlight'), require('./KnnSearchRequestByDocId'), require('./KnnSearchRequestByVector'), require('./SearchRequest'));
  } else {
    // Browser globals (root is window)
    if (!root.Manticoresearch) {
      root.Manticoresearch = {};
    }
    root.Manticoresearch.SearchRequest = factory(root.Manticoresearch.ApiClient, root.Manticoresearch.Aggregation, root.Manticoresearch.Highlight, root.Manticoresearch.KnnSearchRequestByDocId, root.Manticoresearch.KnnSearchRequestByVector, root.Manticoresearch.SearchRequest);
  }
}(this, function(ApiClient, Aggregation, Highlight, KnnSearchRequestByDocId, KnnSearchRequestByVector, SearchRequest) {
  'use strict';



  /**
   * The SearchRequest model module.
   * @module model/SearchRequest
   * @version 4.1.0
   */

  /**
   * Constructs a new <code>SearchRequest</code>.
   * @alias module:model/SearchRequest
   * @class
   * @implements module:model/KnnSearchRequestByVector
   * @implements module:model/KnnSearchRequestByDocId
   * @implements module:model/SearchRequest
   * @param index {String} 
   * @param field {String} 
   * @param queryVector {Array.<Number>} 
   * @param k {Number} 
   * @param docId {Number} 
   */
  var exports = function(index, field, queryVector, k, docId) {
    var _this = this;

    KnnSearchRequestByVector.call(_this, index, field, queryVector, k);
    KnnSearchRequestByDocId.call(_this, index, field, docId, k);
    SearchRequest.call(_this, index, field, queryVector, k, docId);
    _this['index'] = index;
    _this['field'] = field;
    _this['query_vector'] = queryVector;
    _this['k'] = k;
    _this['doc_id'] = docId;
  };

  /**
   * Constructs a <code>SearchRequest</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/SearchRequest} obj Optional instance to populate.
   * @return {module:model/SearchRequest} The populated <code>SearchRequest</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      KnnSearchRequestByVector.constructFromObject(data, obj);
      KnnSearchRequestByDocId.constructFromObject(data, obj);
      SearchRequest.constructFromObject(data, obj);
      if (data.hasOwnProperty('index')) {
        obj['index'] = ApiClient.convertToType(data['index'], 'String');
      }
      if (data.hasOwnProperty('field')) {
        obj['field'] = ApiClient.convertToType(data['field'], 'String');
      }
      if (data.hasOwnProperty('query_vector')) {
        obj['query_vector'] = ApiClient.convertToType(data['query_vector'], ['Number']);
      }
      if (data.hasOwnProperty('k')) {
        obj['k'] = ApiClient.convertToType(data['k'], 'Number');
      }
      if (data.hasOwnProperty('fulltext_filter')) {
        obj['fulltext_filter'] = ApiClient.convertToType(data['fulltext_filter'], Object);
      }
      if (data.hasOwnProperty('attr_filter')) {
        obj['attr_filter'] = ApiClient.convertToType(data['attr_filter'], Object);
      }
      if (data.hasOwnProperty('limit')) {
        obj['limit'] = ApiClient.convertToType(data['limit'], 'Number');
      }
      if (data.hasOwnProperty('offset')) {
        obj['offset'] = ApiClient.convertToType(data['offset'], 'Number');
      }
      if (data.hasOwnProperty('max_matches')) {
        obj['max_matches'] = ApiClient.convertToType(data['max_matches'], 'Number');
      }
      if (data.hasOwnProperty('sort')) {
        obj['sort'] = ApiClient.convertToType(data['sort'], [Object]);
      }
      if (data.hasOwnProperty('aggs')) {
        obj['aggs'] = ApiClient.convertToType(data['aggs'], {'String': Aggregation});
      }
      if (data.hasOwnProperty('expressions')) {
        obj['expressions'] = ApiClient.convertToType(data['expressions'], {'String': 'String'});
      }
      if (data.hasOwnProperty('highlight')) {
        obj['highlight'] = Highlight.constructFromObject(data['highlight']);
      }
      if (data.hasOwnProperty('source')) {
		obj['_source'] = ApiClient.convertToType(data['source'], Object);
      }
      if (data.hasOwnProperty('options')) {
        obj['options'] = ApiClient.convertToType(data['options'], {'String': Object});
      }
      if (data.hasOwnProperty('profile')) {
        obj['profile'] = ApiClient.convertToType(data['profile'], 'Boolean');
      }
      if (data.hasOwnProperty('track_scores')) {
        obj['track_scores'] = ApiClient.convertToType(data['track_scores'], 'Boolean');
      }
      if (data.hasOwnProperty('doc_id')) {
        obj['doc_id'] = ApiClient.convertToType(data['doc_id'], 'Number');
      }
      if (data.hasOwnProperty('query')) {
        obj['query'] = ApiClient.convertToType(data['query'], Object);
      }
    }
    return obj;
  }

  /**
   * @member {String} index
   * @default ''
   */
  exports.prototype['index'] = '';
  /**
   * @member {String} field
   * @default ''
   */
  exports.prototype['field'] = '';
  /**
   * @member {Array.<Number>} query_vector
   */
  exports.prototype['query_vector'] = undefined;
  /**
   * @member {Number} k
   */
  exports.prototype['k'] = undefined;
  /**
   * @member {Object} fulltext_filter
   */
  exports.prototype['fulltext_filter'] = undefined;
  /**
   * @member {Object} attr_filter
   */
  exports.prototype['attr_filter'] = undefined;
  /**
   * @member {Number} limit
   */
  exports.prototype['limit'] = undefined;
  /**
   * @member {Number} offset
   */
  exports.prototype['offset'] = undefined;
  /**
   * @member {Number} max_matches
   */
  exports.prototype['max_matches'] = undefined;
  /**
   * @member {Array.<Object>} sort
   */
  exports.prototype['sort'] = undefined;
  /**
   * @member {Object.<String, module:model/Aggregation>} aggs
   */
  exports.prototype['aggs'] = undefined;
  /**
   * @member {Object.<String, String>} expressions
   */
  exports.prototype['expressions'] = undefined;
  /**
   * @member {module:model/Highlight} highlight
   */
  exports.prototype['highlight'] = undefined;
  /**
   * @member {Object} source
   */
  exports.prototype['_source'] = undefined;
  /**
   * @member {Object.<String, Object>} options
   */
  exports.prototype['options'] = undefined;
  /**
   * @member {Boolean} profile
   */
  exports.prototype['profile'] = undefined;
  /**
   * @member {Boolean} track_scores
   */
  exports.prototype['track_scores'] = undefined;
  /**
   * @member {Number} doc_id
   */
  exports.prototype['doc_id'] = undefined;
  /**
   * @member {Object} query
   */
  exports.prototype['query'] = undefined;

  // Implement KnnSearchRequestByVector interface:
  /**
   * @member {String} index
   * @default ''
   */
exports.prototype['index'] = '';

  /**
   * @member {String} field
   * @default ''
   */
exports.prototype['field'] = '';

  /**
   * @member {Array.<Number>} query_vector
   */
exports.prototype['query_vector'] = undefined;

  /**
   * @member {Number} k
   */
exports.prototype['k'] = undefined;

  /**
   * @member {Object} fulltext_filter
   */
exports.prototype['fulltext_filter'] = undefined;

  /**
   * @member {Object} attr_filter
   */
exports.prototype['attr_filter'] = undefined;

  /**
   * @member {Number} limit
   */
exports.prototype['limit'] = undefined;

  /**
   * @member {Number} offset
   */
exports.prototype['offset'] = undefined;

  /**
   * @member {Number} max_matches
   */
exports.prototype['max_matches'] = undefined;

  /**
   * @member {Array.<Object>} sort
   */
exports.prototype['sort'] = undefined;

  /**
   * @member {Object.<String, module:model/Aggregation>} aggs
   */
exports.prototype['aggs'] = undefined;

  /**
   * @member {Object.<String, String>} expressions
   */
exports.prototype['expressions'] = undefined;

  /**
   * @member {module:model/Highlight} highlight
   */
exports.prototype['highlight'] = undefined;

  /**
   * @member {Object} source
   */
exports.prototype['source'] = undefined;

  /**
   * @member {Object.<String, Object>} options
   */
exports.prototype['options'] = undefined;

  /**
   * @member {Boolean} profile
   */
exports.prototype['profile'] = undefined;

  /**
   * @member {Boolean} track_scores
   */
exports.prototype['track_scores'] = undefined;

  // Implement KnnSearchRequestByDocId interface:
  /**
   * @member {String} index
   * @default ''
   */
exports.prototype['index'] = '';

  /**
   * @member {String} field
   * @default ''
   */
exports.prototype['field'] = '';

  /**
   * @member {Number} doc_id
   */
exports.prototype['doc_id'] = undefined;

  /**
   * @member {Number} k
   */
exports.prototype['k'] = undefined;

  /**
   * @member {Object} fulltext_filter
   */
exports.prototype['fulltext_filter'] = undefined;

  /**
   * @member {Object} attr_filter
   */
exports.prototype['attr_filter'] = undefined;

  /**
   * @member {Number} limit
   */
exports.prototype['limit'] = undefined;

  /**
   * @member {Number} offset
   */
exports.prototype['offset'] = undefined;

  /**
   * @member {Number} max_matches
   */
exports.prototype['max_matches'] = undefined;

  /**
   * @member {Array.<Object>} sort
   */
exports.prototype['sort'] = undefined;

  /**
   * @member {Object.<String, module:model/Aggregation>} aggs
   */
exports.prototype['aggs'] = undefined;

  /**
   * @member {Object.<String, String>} expressions
   */
exports.prototype['expressions'] = undefined;

  /**
   * @member {module:model/Highlight} highlight
   */
exports.prototype['highlight'] = undefined;

  /**
   * @member {Object} source
   */
exports.prototype['source'] = undefined;

  /**
   * @member {Object.<String, Object>} options
   */
exports.prototype['options'] = undefined;

  /**
   * @member {Boolean} profile
   */
exports.prototype['profile'] = undefined;

  /**
   * @member {Boolean} track_scores
   */
exports.prototype['track_scores'] = undefined;

  // Implement SearchRequest interface:
  /**
   * @member {String} index
   * @default ''
   */
exports.prototype['index'] = '';

  /**
   * @member {String} field
   * @default ''
   */
exports.prototype['field'] = '';

  /**
   * @member {Array.<Number>} query_vector
   */
exports.prototype['query_vector'] = undefined;

  /**
   * @member {Number} k
   */
exports.prototype['k'] = undefined;

  /**
   * @member {Object} fulltext_filter
   */
exports.prototype['fulltext_filter'] = undefined;

  /**
   * @member {Object} attr_filter
   */
exports.prototype['attr_filter'] = undefined;

  /**
   * @member {Number} limit
   */
exports.prototype['limit'] = undefined;

  /**
   * @member {Number} offset
   */
exports.prototype['offset'] = undefined;

  /**
   * @member {Number} max_matches
   */
exports.prototype['max_matches'] = undefined;

  /**
   * @member {Array.<Object>} sort
   */
exports.prototype['sort'] = undefined;

  /**
   * @member {Object.<String, module:model/Aggregation>} aggs
   */
exports.prototype['aggs'] = undefined;

  /**
   * @member {Object.<String, String>} expressions
   */
exports.prototype['expressions'] = undefined;

  /**
   * @member {module:model/Highlight} highlight
   */
exports.prototype['highlight'] = undefined;

  /**
   * @member {Object} source
   */
exports.prototype['source'] = undefined;

  /**
   * @member {Object.<String, Object>} options
   */
exports.prototype['options'] = undefined;

  /**
   * @member {Boolean} profile
   */
exports.prototype['profile'] = undefined;

  /**
   * @member {Boolean} track_scores
   */
exports.prototype['track_scores'] = undefined;

  /**
   * @member {Number} doc_id
   */
exports.prototype['doc_id'] = undefined;

  /**
   * @member {Object} query
   */
exports.prototype['query'] = undefined;



  return exports;
}));


