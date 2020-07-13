# OpenAPI Tryout

The OpenAPI file is `manticore.yaml`.

To generate code for a language:

`docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/manticore.yaml -g php -o /local/out/php`


In php/python there is a test.php/py file using the client.


The yaml was edited with swagger and can be seen at:

https://app.swaggerhub.com/apis/adriannuta/ManticoreSeach/1.0.0




## Specific client issues

### Python

* problem with `bulk` as the client regex for `json` in `Content-Type` header
and ends up JSON encoding the bulk payload. One fix is to patch regex to be specific
and `bulk()` to accept a string (so user needs to put at input the NDJSON)

* `sql` doesn't work for mode=raw because using `application/x-www-form-urlencoded`
will work for SELECTs, but won't for `mode=raw` which requires the `query` to not
be url encoded. Fix is to use `text/plain` and user needs to provide 
`sql('mode=raw&query=select * from movies')`

### Javascript

* delete() is renamed to `callDelete` because `delete()` is a reserved method


### Java

* the client seems to throw exception on any content type that is not json, form or multipart - it will need patch
to have `bulk` and `sql` work