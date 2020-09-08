#!/bin/bash
set -e
do_python() {
  echo "Building Python ..."
  rm out/manticoresearch-python -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  openapitools/openapi-generator-cli generate -i /local/manticore.yml -g python -o /local/out/manticoresearch-python -t /local/templates/python --git-repo-id manticoresearch-python --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch --additional-properties packageName=manticoresearch --additional-properties packageVersion=`cat versions/python`
  rm out/manticoresearch-python/test/* -rf
  cp -R test/python/* out/manticoresearch-python/test/   
  git apply patches/python_bulk.patch
  git apply patches/python_readme.patch
  # replace test with our test
  echo "Python done."
}

do_java() {
  echo "Building Java ..."
  rm out/manticoresearch-java -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)" -e JAVA_OPTS="-Dlog.level=warn"  openapitools/openapi-generator-cli generate -i /local/manticore.yml -g java  -o /local/out/manticoresearch-java -t /local/templates/Java --git-repo-id manticoresearch-java --git-user-id manticoresoftware     --additional-properties apiPackage=com.manticoresearch.client.api --additional-properties modelPackage=com.manticoresearch.client.model  --additional-properties artifactId=manticoresearch  --additional-properties developerName="Manticore Software"  --additional-properties developerEmail="info@manticosearch.com"  --additional-properties  developerOrganization="manticoresearch.com" --additional-properties developerOrganizationUrl="https://github.com/manticoresoftware/manticoresearch-java"  --additional-properties artifactVersionn=`cat versions/java` --additional-properties groupId="com.manticoresearch"  --additional-properties artifactUrl=https://github.com/manticoresoftware/manticoresearch-java --additional-properties licenseName="Apache 2.0" --additional-properties artifactDescription="Client for Manticore Search"
  #git apply patches/java.patch
  echo "Java done."
}

do_javascript() {
  echo "Building Javascript ..."
  rm out/manticoresearch-javascript -rf
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"   openapitools/openapi-generator-cli generate -i /local/manticore.yml -g javascript -o /local/out/manticoresearch-javascript -t /local/templates/Javascript --git-repo-id manticoresearch-javascript --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch  --additional-properties projectVersion=`cat versions/javascript`
  git apply patches/javascript.package.patch
  git apply patches/javascript_readme.patch
  rm out/manticoresearch-javascript/test/* -rf
  cp -R test/javascript/* out/manticoresearch-javascript/test/ 
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

do_go() {
  echo "Building Go ..."
  rm out/manticore-go -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"    -e JAVA_OPTS="-Dlog.level=warn" openapitools/openapi-generator-cli generate -i /local/manticore.yml -g go-experimental  -o /local/out/manticore-go-experimental -t /local/templates/go --git-repo-id manticoresearch-go --git-user-id manticoresoftware
  echo "Perl done." 
}
do_elixir() {
  echo "Building Elixir ..."
  rm out/manticoresearch-elixir -rf
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"   openapitools/openapi-generator-cli generate -i /local/manticore.yml -g elixir  -o /local/out/manticoresearch-elixir -t /local/templates/elixir --git-repo-id manticoresearch-elixir --git-user-id manticoresoftware  --additional-properties packageName="manticoresearch" --additional-properties invokerPackage="Manticoresearch"
  cp test/elixir/api_index_test.exs out/manticoresearch-elixir/test/
  cp patches/test.exs out/manticoresearch-elixir/config/
  cp patches/prod.exs out/manticoresearch-elixir/config/
  echo "Elixir done." 
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
 go)
   do_go
  ;;  
 elixir)
   do_elixir
  ;;    
 all)
   do_python
   do_java
   do_javascript
   do_csharp
   do_ruby
   do_swift
   do_perl
   do_elixir
  ;;
*)
  echo -n "unknown"
  ;;
esac
