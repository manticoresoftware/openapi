# Manticore Elixir client

Сlient for Manticore Search. 


## Requirements.

Minimum Manticore Search version is >= 4.0.2 with HTTP protocol enabled.


## Building

To install the required dependencies and to build the elixir project, run:

```console
mix local.hex --force
mix do deps.get, compile
```

## Installation

If [available in Hex][], the package can be installed by adding `manticoresearch` to
your list of dependencies in `mix.exs`:

```elixir
def deps do
  [{:manticoresearch, "~> 5.0.0"}]
end
```

Documentation can be generated with [ExDoc][] and published on [HexDocs][]. Once published, the docs can be found at
[https://hexdocs.pm/manticoresearch][docs].

## Configuration

You can override the URL of your server (e.g. if you have a separate development and production server in your
configuration files).

```elixir
config :manticoresearch, base_url: "http://127.0.0.1:9308"
```

Multiple clients for the same API with different URLs can be created passing different `base_url`s when calling
`Manticoresearch.Connection.new/1`:

```elixir
client = Manticoresearch.Connection.new(base_url: "http://127.0.0.1:9308")
```

[exdoc]: https://github.com/elixir-lang/ex_doc
[hexdocs]: https://hexdocs.pm
[available in hex]: https://hex.pm/docs/publish
[docs]: https://hexdocs.pm/manticoresearch
