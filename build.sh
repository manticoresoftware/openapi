#!/bin/bash
sudo rm out/* -rf
for w in go go-experimental python ruby php csharp c java javascript perl 
do
   echo "doing $w"
  docker run --rm -v ${PWD}:/local  -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.json -g $w -o /local/out/manticore-$w
done
sudo git apply python_bulk.patch
sudo git apply java.patch
