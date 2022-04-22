#!/bin/bash
set -e
do_python() {
  echo "Building Python ..."
  rm -rf out/manticoresearch-python 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g python -o /local/out/manticoresearch-python -t /local/templates/python --git-repo-id manticoresearch-python --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch --additional-properties packageName=manticoresearch --additional-properties packageVersion=`cat versions/python`
  rm -rf out/manticoresearch-python/test/* 
  cp LICENSE.txt out/manticoresearch-python/LICENSE.txt
  cp -R test/python/* out/manticoresearch-python/test/   
  git apply patches/python_bulk.patch
  #git apply patches/python_readme.patch
  # replace test with our test
  echo "Python done."
}

do_java() {
  echo "Building Java ...$version"
  rm -rf out/manticoresearch-java 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)" -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g java  -o /local/out/manticoresearch-java -t /local/templates/Java --git-repo-id manticoresearch-java --git-user-id manticoresoftware     --additional-properties apiPackage=com.manticoresearch.client.api --additional-properties modelPackage=com.manticoresearch.client.model  --additional-properties artifactId=manticoresearch  --additional-properties developerName="Manticore Software"  --additional-properties developerEmail="info@manticosearch.com"  --additional-properties  developerOrganization="manticoresearch.com" --additional-properties developerOrganizationUrl="https://github.com/manticoresoftware/manticoresearch-java"  --additional-properties artifactVersion=`cat versions/java` --additional-properties groupId="com.manticoresearch"  --additional-properties artifactUrl=https://github.com/manticoresoftware/manticoresearch-java --additional-properties licenseName="Apache 2.0" --additional-properties artifactDescription="Client for Manticore Search"  --additional-properties  library="jersey2" 
  cp LICENSE.txt out/manticoresearch-java/LICENSE.txt
  #cp docs/java/README.md out/manticoresearch-java/README.md
  cp docs/java/docs/* out/manticoresearch-java/docs/
  #git apply patches/java.patch
  #git apply patches/java/ApiClient.java
  #git apply patches/java/SuccessResponse.java
  rm -rf out/manticoresearch-java/.openapi-generator
  rm -rf out/manticoresearch-java/api
  echo "Java done."
}

do_javascript() {
  echo "Building Javascript ..."
  rm -rf out/manticoresearch-javascript 
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g javascript -o /local/out/manticoresearch-javascript -t /local/templates/Javascript --git-repo-id manticoresearch-javascript --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch  --additional-properties projectVersion=`cat versions/javascript`   --additional-properties  usePromises=true
  git apply patches/javascript.package.patch
  cp LICENSE.txt out/manticoresearch-javascript/LICENSE.txt
  #cp docs/javascript/README.md out/manticoresearch-javascript/README.md
  cp docs/javascript/docs/* out/manticoresearch-javascript/docs/
  rm -rf out/manticoresearch-javascript/test/api/* 
  cp -R test/javascript/api/Manual.spec.js out/manticoresearch-javascript/test/api/Manual.spec.js
  echo "Javascript done."
}

do_csharp() {
  echo "Building CSharp ..."
  rm -rf out/manticore-csharp 
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g csharp  -o /local/out/manticore-csharp -t /local/templates/csharp --git-repo-id manticoresearch-csharp --git-user-id manticoresoftware
  echo "CSharp done."
}

do_ruby() {
  echo "Building Ruby ..."
  rm -rf out/manticore-ruby 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"   -e JAVA_OPTS="-Dlog.level=warn" "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g ruby -o /local/out/manticore-ruby -t /local/templates/ruby-client --git-repo-id manticoresearch-ruby --git-user-id manticoresoftware
  echo "Ruby done."
}

do_swift() {
  echo "Building Swift ..."
  rm -rf out/manticore-swift5 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"   -e JAVA_OPTS="-Dlog.level=warn" "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g swift5 -o /local/out/manticore-swift5 -t /local/templates/swift5 --git-repo-id manticoresearch-swift --git-user-id manticoresoftware
  echo "Swift done." 
}

do_perl() {
  echo "Building Perl ..."
  rm -rf out/manticore-perl 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"    -e JAVA_OPTS="-Dlog.level=warn" "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g perl  -o /local/out/manticore-perl -t /local/templates/perl --git-repo-id manticoresearch-perl --git-user-id manticoresoftware
  echo "Perl done." 
}

do_go() {
  echo "Building Go ..."
  rm -rf out/manticore-go 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"    -e JAVA_OPTS="-Dlog.level=warn" "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g go  -o /local/out/manticore-go-experimental -t /local/templates/go --git-repo-id manticoresearch-go --git-user-id manticoresoftware
  echo "Go done." 
}
do_elixir() {
  echo "Building Elixir ..."
  rm -rf out/manticoresearch-elixir 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g elixir  -o /local/out/manticoresearch-elixir -t /local/templates/elixir --git-repo-id manticoresearch-elixir --git-user-id manticoresoftware  --additional-properties packageName="manticoresearch" --additional-properties invokerPackage="Manticoresearch"
  cp LICENSE.txt out/manticoresearch-elixir/LICENSE.txt
  cp test/elixir/api_index_test.exs out/manticoresearch-elixir/test/
  cp patches/test.exs out/manticoresearch-elixir/config/
  cp patches/dev.exs out/manticoresearch-elixir/config/
  echo "Elixir done." 
}

if [ ! -z $2 ]; then
version=":$2"
else
version=''
fi

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
