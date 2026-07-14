# Manticore Search Client
# Copyright (c) 2020-2021, Manticore Software LTD (https://manticoresearch.com)
#
# All rights reserved

from manticoresearch.models.search_response import SearchResponse


def unwrap_search_response(result):
    """Return the SearchResponse from a /search call result."""
    if isinstance(result, SearchResponse):
        return result
    return result.actual_instance
