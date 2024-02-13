# Manticore Search Client
# Copyright (c) 2020-2021, Manticore Software LTD (https://manticoresearch.com)
#
# All rights reserved

# Do not edit the class manually.

defmodule Manticoresearch.Model.KnnSearchRequestByVector do
  @moduledoc """
  Request object for knn search operation
  """

  @derive [Poison.Encoder]
  defstruct [
    :"index",
    :"field",
    :"query_vector",
    :"k",
    :"fulltext_filter",
    :"attr_filter",
    :"limit",
    :"offset",
    :"max_matches",
    :"sort",
    :"aggs",
    :"expressions",
    :"highlight",
    :"source",
    :"options",
    :"profile",
    :"track_scores"
  ]

  @type t :: %__MODULE__{
    :"index" => String.t,
    :"field" => String.t,
    :"query_vector" => [float()],
    :"k" => integer(),
    :"fulltext_filter" => Map | nil,
    :"attr_filter" => Map | nil,
    :"limit" => integer() | nil,
    :"offset" => integer() | nil,
    :"max_matches" => integer() | nil,
    :"sort" => [Map] | nil,
    :"aggs" => %{optional(String.t) => Aggregation} | nil,
    :"expressions" => %{optional(String.t) => String.t} | nil,
    :"highlight" => Highlight | nil,
    :"source" => Map | nil,
    :"options" => %{optional(String.t) => Map} | nil,
    :"profile" => boolean() | nil,
    :"track_scores" => boolean() | nil
  }
end

defimpl Poison.Decoder, for: Manticoresearch.Model.KnnSearchRequestByVector do
  import Manticoresearch.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:"aggs", :map, Manticoresearch.Model.Aggregation, options)
    |> deserialize(:"highlight", :struct, Manticoresearch.Model.Highlight, options)
  end
end

