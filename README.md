# Manticoresearch OpenAPI clients generator

The project allows you to automatically generate Manticoresearch clients for different languages.

For this purpose, *OpenApi Generator* tool is used. More info on it can be found [here](https://github.com/OpenAPITools/openapi-generator).

To generate a client, OpenApi generator needs a config file describing the structure of API. You
can look at it in `manticore.yml`.  To edit and validate the config, you can use an online editor like [Swagger](https://swagger.io/).


## Generated clients

* [Python](https://github.com/manticoresoftware/manticoresearch-python) - it also has a testing suite with incomplete coverage

* [Javascript](https://github.com/manticoresoftware/manticoresearch-javascript) - it also has a testing suite with incomplete coverage

* [Java](https://github.com/manticoresoftware/manticoresearch-java)

* [Elixir](https://github.com/manticoresoftware/manticoresearch-elixir)


## Building

The simplest way to make a build for a specific language, e.g., for Java, is to use our build script:

```
./build.sh java
```

You can also build versions for all supported languages at once:

```
./build.sh all
```

By default, the latest available verison of the OpenApi generator will be used. If you want to apply a specific version, add a corresponding argument to the command call:

```
build.sh Java v5.0.0   
```

where `v5.0.0` is a tag of the generator's docker image you need.


Original versions we used were 'v4.3.1' for Python and 'v5.0.0.' for the rest of the clients.


Note that you can also use Docker and run OpenApi generator docker image directly:

```
docker pull  openapitools/openapi-generator-cli 

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/manticore.yaml -g java -o /local/out/java
```


### Custom overrides

Using the Docker way for generating clients, you should also apply your custom overrides when it's necessary (for example, to change package names).

By default, the generators use standard package names (like *openapi_client*). These names need to be changed at building time (since this is not a responsibility of OpenAPI spec to define how a package for a language must be named).

Package or Project names have their specific variables for each language (there are some general variables for `api/model` package, but some generators don't use them!). There are also other variables that may require a custom value to be set for them. To see these, run

```
docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)" -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli config-help -g XXX
```

where `XXX` is the name of the generator (language).

To add these specific variables, use ` --additional-properties NAME=VALUE` in the `docker run` command (see `build.sh` for examples).

Please note that for Java the library used is `okhttp-gson`, and some of its templates (like the `readme`) must be edited in the `templates/Java/libraries/okhttp-gson/` folder.


## Project folders

- `out` - contains the code of generated clients. Each client is located in a folder named `manticore-X` where X is a client's language.

- `patches` - contains patch files for various clients

- `test` - contains hand-made or modified generated tests (it's desirable to copy them to the corresponding subfolder of `manticore-X`)

- `templates` -  contains mustache templates taken from https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources; These should be updated from time to time


## Specific client issues

### Javascript


* `delete()` is renamed to `callDelete` because `delete()` is a reserved method - the only appropriate fix is to rename doc operations to something else (like `deleteDocuments()`)





