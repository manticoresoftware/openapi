# Manticore Search Client
# Copyright (c) 2020-2021, Manticore Software LTD (https://manticoresearch.com)
#
# All rights reserved

defmodule Manticoresearch.Api.IndexTest do
  use ExUnit.Case
  doctest Manticoresearch.Api.Index
  test "perform a search" do
    tesla_client = Manticoresearch.Connection.new
    IO.inspect tesla_client
    sql = ~s[mode=raw&query=CREATE TABLE testrt ( title text, content text, gid integer)]
    doc = ~s[{"table":"testrt","id":1,"doc":{"title":"Hello","content":"world","gid":1}}]
    expected_response = {:ok,
      %Manticoresearch.Model.SearchResponse{
        hits: %Manticoresearch.Model.SearchResponseHits{
          hits: [
            %{
              "_id" => "1",
              "_score" => 2500,
              "_source" => %{"content" => "world", "gid" => 1, "title" => "Hello"}
            }
          ],
          total: 1
        },
        profile: nil,
        timed_out: false,
        took: 0
    }}
    # create table
    create_table = Manticoresearch.Api.Utils.sql(tesla_client, sql)
    IO.inspect create_table
    # insert a document
    Manticoresearch.Api.Index.insert(tesla_client, doc)
    json = ~s[{"table": "testrt",
               "query": {
                  "match": {
                    "*": "hello world"
                  }
                }
              }]
    api_response = Manticoresearch.Api.Search.search(tesla_client, json)
    IO.inspect api_response
    {:ok, search_response} = api_response
    search_response =
      case search_response do
        %Manticoresearch.Model.SearchResponse{} = response -> response
        %{"hits" => _} = response -> struct(Manticoresearch.Model.SearchResponse, response)
        other -> raise "unexpected search response: #{inspect(other)}"
      end
    assert {:ok, search_response} == expected_response
  end
end
