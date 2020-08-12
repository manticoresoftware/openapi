#!/bin/bash
set -e

do_python() {
  echo "Building Python ..."
  rm out/manticore-python -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g python -o /local/out/manticore-python -t /local/templates/python
  git apply python_bulk.patch
  rm out/manticore-python/test/* -rf
  cp -R test/python/* out/manticore-python/test/ 
  # replace test with our test
  echo "Python done."
}

do_java() {
  echo "Building Java ..."
  rm out/manticore-java -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g java  -o /local/out/manticore-java -t /local/templates/Java
  git apply java.patch
  echo "Java done."
}

do_javascript() {
  echo "Building Javascript ..."
  rm out/manticore-javascript -rf
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g javascript -o /local/out/manticore-javascript -t /local/templates/Javascript
  echo "Javascript done."
}

do_csharp() {
  echo "Building CSharp ..."
  rm out/manticore-csharp -rf
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g csharp  -o /local/out/manticore-csharp -t /local/templates/csharp
  echo "CSharp done."
}

do_ruby() {
  echo "Building Ruby ..."
  rm out/manticore-ruby -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g ruby -o /local/out/manticore-ruby -t /local/templates/ruby-client
  echo "Ruby done."
}

do_swift() {
  echo "Building Swift ..."
  rm out/manticore-swift5 -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g swift5 -o /local/out/manticore-swift5 -t /local/templates/swift5
  echo "Swift done."
}

do_perl() {
  echo "Building Perl ..."
  rm out/manticore-perl -rf
  docker run --rm -v ${PWD}:/local   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g perl  -o /local/out/manticore-perl -t /local/templates/perl
  echo "Perl done." -u "$(id -u):$(id -g)" 
}

case $1 in
 python)
   do_python
  ;;
 java)
   do_java
  ;;
 javascript)
   do_javascript
  ;;
 csharp)
   do_csharp
  ;;
 ruby)
   do_ruby
  ;;
 swift)
   do_swift
  ;;
 perl)
   do_perl
  ;;
 all)
   do_python
   do_java
   do_javascript
   do_csharp
   do_ruby
   do_swift
   do_perl
  ;;
*)
  echo -n "unknown"
  ;;
esac
