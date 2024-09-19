# coding: utf-8

"""
    Manticore Search Client

    Сlient for Manticore Search. 

    The version of the OpenAPI document: 4.0.0
    Contact: info@manticoresearch.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from manticoresearch.models.aggregation import Aggregation
from manticoresearch.models.highlight import Highlight
from manticoresearch.models.join_inner import JoinInner
from manticoresearch.models.knn_search_request_all_of_knn import KnnSearchRequestAllOfKnn
from manticoresearch.models.search_request_parameters_sort_inner import SearchRequestParametersSortInner
from manticoresearch.models.search_request_parameters_source import SearchRequestParametersSource
from typing import Optional, Set
from typing_extensions import Self

class KnnSearchRequest(BaseModel):
    """
    Request object for knn search operation
    """ # noqa: E501
    aggs: Optional[Aggregation] = None
    expressions: Optional[Dict[str, StrictStr]] = None
    join: Optional[List[JoinInner]] = None
    highlight: Optional[Highlight] = None
    index: StrictStr
    limit: Optional[StrictInt] = None
    max_matches: Optional[StrictInt] = None
    offset: Optional[StrictInt] = None
    options: Optional[Dict[str, Any]] = None
    profile: Optional[StrictBool] = None
    sort: Optional[List[SearchRequestParametersSortInner]] = None
    source: Optional[SearchRequestParametersSource] = Field(default=None, alias="_source")
    track_scores: Optional[StrictBool] = None
    knn: KnnSearchRequestAllOfKnn
    __properties: ClassVar[List[str]] = ["aggs", "expressions", "join", "highlight", "index", "limit", "max_matches", "offset", "options", "profile", "sort", "_source", "track_scores", "knn"]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of KnnSearchRequest from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of aggs
        if self.aggs:
            _dict['aggs'] = self.aggs.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in join (list)
        _items = []
        if self.join:
            for _item_join in self.join:
                if _item_join:
                    _items.append(_item_join.to_dict())
            _dict['join'] = _items
        # override the default output from pydantic by calling `to_dict()` of highlight
        if self.highlight:
            _dict['highlight'] = self.highlight.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in sort (list)
        _items = []
        if self.sort:
            for _item_sort in self.sort:
                if _item_sort:
                    _items.append(_item_sort.to_dict())
            _dict['sort'] = _items
        # override the default output from pydantic by calling `to_dict()` of source
        if self.source:
            _dict['_source'] = self.source.to_dict()
        # override the default output from pydantic by calling `to_dict()` of knn
        if self.knn:
            _dict['knn'] = self.knn.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of KnnSearchRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "aggs": Aggregation.from_dict(obj["aggs"]) if obj.get("aggs") is not None else None,
            "expressions": obj.get("expressions"),
            "join": [JoinInner.from_dict(_item) for _item in obj["join"]] if obj.get("join") is not None else None,
            "highlight": Highlight.from_dict(obj["highlight"]) if obj.get("highlight") is not None else None,
            "index": obj.get("index"),
            "limit": obj.get("limit"),
            "max_matches": obj.get("max_matches"),
            "offset": obj.get("offset"),
            "options": obj.get("options"),
            "profile": obj.get("profile"),
            "sort": [SearchRequestParametersSortInner.from_dict(_item) for _item in obj["sort"]] if obj.get("sort") is not None else None,
            "_source": SearchRequestParametersSource.from_dict(obj["_source"]) if obj.get("_source") is not None else None,
            "track_scores": obj.get("track_scores"),
            "knn": KnnSearchRequestAllOfKnn.from_dict(obj["knn"]) if obj.get("knn") is not None else None
        })
        return _obj


