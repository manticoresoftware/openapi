#!/bin/bash
set -e
do_python() {
  echo "Building Python ..."
  rm out/manticoresearch-python -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  openapitools/openapi-generator-cli generate -i /local/manticore.yml -g python -o /local/out/manticoresearch-python -t /local/templates/python --git-repo-id manticoresearch-python --git-user-id manticoresoftware
  git apply patches/python_bulk.patch
  rm out/manticoresearch-python/test/* -rf
  cp -R test/python/* out/manticoresearch-python/test/ 
  # replace test with our test
  echo "Python done."
}

do_java() {
  echo "Building Java ..."
  rm out/manticore-java -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g java  -o /local/out/manticore-java -t /local/templates/Java --git-repo-id manticoresearch-java --git-user-id manticoresoftware
  git apply patches/java.patch
  echo "Java done."
}

do_javascript() {
  echo "Building Javascript ..."
  rm out/manticore-javascript -rf
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g javascript -o /local/out/manticore-javascript -t /local/templates/Javascript --git-repo-id manticoresearch-javascript --git-user-id manticoresoftware
  rm out/manticore-javascript/test/* -rf
  cp -R test/javascript/* out/manticore-javascript/test/ 
  echo "Javascript done."
}

do_csharp() {
  echo "Building CSharp ..."
  rm out/manticore-csharp -rf
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g csharp  -o /local/out/manticore-csharp -t /local/templates/csharp --git-repo-id manticoresearch-csharp --git-user-id manticoresoftware
  echo "CSharp done."
}

do_ruby() {
  echo "Building Ruby ..."
  rm out/manticore-ruby -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g ruby -o /local/out/manticore-ruby -t /local/templates/ruby-client --git-repo-id manticoresearch-ruby --git-user-id manticoresoftware
  echo "Ruby done."
}

do_swift() {
  echo "Building Swift ..."
  rm out/manticore-swift5 -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"   -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g swift5 -o /local/out/manticore-swift5 -t /local/templates/swift5 --git-repo-id manticoresearch-swift --git-user-id manticoresoftware
  echo "Swift done." 
}

do_perl() {
  echo "Building Perl ..."
  rm out/manticore-perl -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"    -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g perl  -o /local/out/manticore-perl -t /local/templates/perl --git-repo-id manticoresearch-perl --git-user-id manticoresoftware
  echo "Perl done." 
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
