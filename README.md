# OpenAPI Tryout

The OpenAPI file is `manticore.yaml`.

To generate code for a language:

`docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/manticore.yaml -g php -o /local/out/php`


In php/python there is a test.php/py file using the client.


The yaml was edited with swagger and can be seen at:

https://app.swaggerhub.com/apis/adriannuta/ManticoreSeach/1.0.0


## Clients

* PHP - tested, but are not interested in it, as we have already a client

* CSharp - not tested

* Python - tested

* Javascript - tested

* Java - tested

* Perl - not tested

* Ruby - not tested

* Swift - not tested

* Go -  not tested

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

* the client seems to throw exception on any content type that is not json, form or multipart 
* another problem is the http client forces addition of charset=utf-8 in `Content-Type`, this breaks
`bulk` even if x-ndjson is set (seems we do an exact check on content-type) (there's a 2017 ticket on the 
http client about this as seems it breaks compat with apis like ones from aws or dropbox since most do exact matching 
on content-type). 
Both these 2 issues need patches.

