#!/bin/bash
sudo rm out/* -rf
#for w in go-experimental  ruby php csharp c java perl
#do
#   echo "doing $w"
#  docker run --rm -v ${PWD}:/local  -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g $w -o /local/out/manticore-$w
#done
echo "go-experimental"
docker run --rm -v ${PWD}:/local  -e JAVA_OPTS="-Dlog.level=warn"  openapitools/openapi-generator-cli generate -i /local/manticore.yml -g go-experimental  -o /local/out/manticore-go-experimental -t /local/templates/go-experimental

echo "csharp"
docker run --rm -v ${PWD}:/local   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g csharp  -o /local/out/manticore-csharp -t /local/templates/csharp

echo "java"
docker run --rm -v ${PWD}:/local   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g java  -o /local/out/manticore-java -t /local/templates/Java

echo "perl"
docker run --rm -v ${PWD}:/local   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g perl  -o /local/out/manticore-perl -t /local/templates/perl

echo  "python"
docker run --rm -v ${PWD}:/local   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g python  -o /local/out/manticore-python -t /local/templates/python

echo "javascript"
docker run --rm -v ${PWD}:/local   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g javascript -o /local/out/manticore-javascript -t /local/templates/Javascript

echo "ruby"
docker run --rm -v ${PWD}:/local   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g ruby -o /local/out/manticore-ruby -t /local/templates/ruby-client


echo "swift5"
docker run --rm -v ${PWD}:/local   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g swift5 -o /local/out/manticore-swift5 -t /local/templates/swift5
 

sudo git apply python_bulk.patch
sudo git apply java.patch
