# OpenAPI Tryout

The OpenAPI file is `manticore.yaml`.

To generate code for a language:

`docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/manticore.yaml -g php -o /local/out/php`


In php/python there is a test.php/py file using the client.


The yaml was edited with swagger and can be seen at:

https://app.swaggerhub.com/apis/adriannuta/ManticoreSeach/1.0.0
