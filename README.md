# OpenAPI Tryout

The OpenAPI file is `manticore.yaml`. 

## Editing the OpenAPI spec

You can use an online editor line Swagger (https://swagger.io/).



## Building

To generate code for a language (simple form):

`docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/manticore.yaml -g php -o /local/out/php`


The script `build.sh`  is made to perform building for several languages.

You can use it like:

```
./build.sh java

```

I suggest to `docker pull  openapitools/openapi-generator-cli` first.

## Folders

- out - generated code clients. Each is in a folder named `manticore-X`

- patches - contains patch files for various clients

- docker - docker images for CI/testing

- test - hand-made tests (desirable to copy them into `manticore-X` test folder

- templates -  mustache templates taken from https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources; These should be updated from time to time


## Clients

* PHP - works, but are not interested in it, as we have already a client

* CSharp - not tested

* Python - works, has testing suite

* Javascript - works

* Java - works

* Perl - not tested

* Ruby - not tested

* Swift - not tested

* Go -  tried to test it, seems smth broken or need more go knowledge

## Specific client issues

### Python

Solved (with patches):


* problem with `bulk` as the client regex for `json` in `Content-Type` header
and ends up JSON encoding the bulk payload. One fix is to patch regex to be specific
and `bulk()` to accept a string (so user needs to put at input the NDJSON)

* `sql` doesn't work for mode=raw because using `application/x-www-form-urlencoded`
will work for SELECTs, but won't for `mode=raw` which requires the `query` to not
be url encoded. Fix is to use `text/plain` and user needs to provide 
`sql('mode=raw&query=select * from movies')`

### Javascript

Unresolved(can't fix?):

* delete() is renamed to `callDelete` because `delete()` is a reserved method

### Java

Solved (with patches):

* the client seems to throw exception on any content type that is not json, form or multipart 
* another problem is the http client forces addition of charset=utf-8 in `Content-Type`, this breaks
`bulk` even if x-ndjson is set (seems we do an exact check on content-type) (there's a 2017 ticket on the 
http client about this as seems it breaks compat with apis like ones from aws or dropbox since most do exact matching 
on content-type). 

