# OpenAPI Tryout

The OpenAPI file is `manticore.yaml`.  You can use an online editor like Swagger (https://swagger.io/) to edit and validate it.

## Building

I suggest to `docker pull  openapitools/openapi-generator-cli` first.

To generate code for a language (simple form):

`docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/manticore.yaml -g php -o /local/out/php`

Depending on language, there could be patches to be applied or some files to overwrite, like tests.

The better is to use the  `build.sh` , which has for every language all the steps required (and also you should update this if you do work for a client)

You can use it like:

```
./build.sh java

```

You can also rebuild all with

```
./build.sh all
```


## Folders

- out - generated code clients. Each is in a folder named `manticore-X`

- patches - contains patch files for various clients

- test - hand-made or modified generated tests (desirable to copy them into `manticore-X` test folder)

- templates -  mustache templates taken from https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources; These should be updated from time to time


## Clients

* PHP - not included - we have hand-made client already

* CSharp - not tested

* Python - works, has testing suite (coverage is not complete)

* Javascript - works, has testing suite (coverage is poor)

* Java - works

* Perl - not tested

* Ruby - not tested

* Swift - not tested

* Go -  tried to test it, seems smth broken or need more go knowledge

## Custom overrides

By default the generators use crap package names (like openapi_client). These needs to be changes at building time (as  this is not part of OpenAPI spec how you name the package for a language).

Package or Project names have specific variables for each language (there are some general variables for api/model package, but some generators don't use these!). There are also other variables that may have an use to set a custom value for them. To see these, run

```
docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)" -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli config-help -g XXX
```

where `XXX` is the name of the generator (language).

Adding one of these specific variables, is made adding ` --additional-properties NAME=VALUE` to the docker run command (see build.sh for examples).

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

Unresolved:

* delete() is renamed to `callDelete` because `delete()` is a reserved method - the only fix is to rename doc operations to something else (like deleteDocuments())

### Java

Solved (with patches):

* the client seems to throw exception on any content type that is not json, form or multipart 
* another problem is the http client forces addition of charset=utf-8 in `Content-Type`, this breaks
`bulk` even if x-ndjson is set (seems we do an exact check on content-type) (there's a 2017 ticket on the 
http client about this as seems it breaks compat with apis like ones from aws or dropbox since most do exact matching 
on content-type). 

## CI testing

Right now only added for javascript and python. There's a manticore service that can be used (use `manticoresearch-manticore` as hostname).

CI jobs run for each client only if a change was made in their folder.

## Pushing to github repositories

The current setup is making use of git subtrees to push the output folders with code to individual github repositories - so it's  a one-way thing.

Adding a new push for a language:

- create the repository on github, add a dummy readme

- back on this repository do `git remote  add manticoresearch-XXX git@github.com:manticoresoftware/manticoresearch-XXX.git`

- the `manticoresearch-XXX` must not exists in `out` folder - if it's the case deleted (and commit)

- `git subtree add --prefix=out/manticoresearch-XXX/ manticoresearch-XXX master`

- you can not use `build.sh` script to write in the `out/manticoresearch-XXX` folder (careful as some early outputs are in `manticore-XXX` format)

Pushing is made during CI using
`git subtree push --prefix=out/manticoresearch-XXX https://$GITHUB_USERNAME:$GITHUB_TOKEN@github.com/manticoresoftware/manticoresearch-XXX.git master`

Pushing is made only if there's a change in `manticoresearch-XXX` folder. 

## Publishing  and versions

In `versions` folder there is a file with the version of the respective package. The version should be increased whenever we want to publish a new version. 
Please note that some (or all) package systems don't allow reuploading a version (at least pypi doesn't), so even for some readme change we need to push a new version.

The Gitlab CI has manual jobs for publishing a package. 
So the workflow is:

- make changes, increase version in version file, commit
- wait until testing and push pipeline finish
- go to Gitlab pipeline and run the manual job (e.g. python_pypi) to publish the package.



