#!/bin/bash

set -e

do_python() {
  echo "Building Python ..."
  rm -rf out/manticoresearch-python 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore_new.yml -g python -o /local/out/manticoresearch-python -t /local/templates/python --git-repo-id manticoresearch-python --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch --additional-properties packageName=manticoresearch --additional-properties packageVersion=`cat versions/python` $build_to_branch
  rm -rf out/manticoresearch-python/test/* 
  cp LICENSE.txt out/manticoresearch-python/LICENSE.txt
  # replace test with our test
  cp -R test/python/* out/manticoresearch-python/test/   
  cp -r docs/python/docs/SearchApi.md out/manticoresearch-python/docs/SearchApi.md
  git apply patches/python_bulk.patch
  #git apply patches/python_readme.patch
  echo "Python done."
}

do_java() {
  echo "Building Java ...$version"
  rm -rf out/manticoresearch-java 
  docker run \
    --rm \
    -v ${PWD}:/local \
    -u "$(id -u):$(id -g)" \
    -e JAVA_OPTS="-Dlog.level=warn" \
    "openapitools/openapi-generator-cli$version" generate \
    -i /local/manticore.yml \
    -g java  \
    -o /local/out/manticoresearch-java \
    -t /local/templates/Java \
    --git-repo-id manticoresearch-java \
    --git-user-id manticoresoftware \
    --additional-properties apiPackage=com.manticoresearch.client.api \
    --additional-properties modelPackage=com.manticoresearch.client.model  \
    --additional-properties artifactId=manticoresearch \
    --additional-properties developerName="Manticore Software" \
    --additional-properties developerEmail="info@manticosearch.com" \
    --additional-properties  developerOrganization="manticoresearch.com" \
    --additional-properties developerOrganizationUrl="https://github.com/manticoresoftware/manticoresearch-java" \
    --additional-properties artifactVersion=`cat versions/java` \
    --additional-properties groupId="com.manticoresearch" \
    --additional-properties artifactUrl=https://github.com/manticoresoftware/manticoresearch-java \
    --additional-properties licenseName="Apache 2.0" \
    --additional-properties artifactDescription="Client for Manticore Search" \
    --additional-properties library="jersey3" \
    --additional-properties useJakartaEe=true \
    $build_to_branch
  cp LICENSE.txt out/manticoresearch-java/LICENSE.txt
  cp docs/java/docs/* out/manticoresearch-java/docs/
  rm -rf out/manticoresearch-java/.openapi-generator
  rm -rf out/manticoresearch-java/api
  echo "Java done."

}

do_javascript() {
  echo "Building Javascript ..."
  rm -rf out/manticoresearch-javascript 
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g javascript -o /local/out/manticoresearch-javascript -t /local/templates/Javascript --git-repo-id manticoresearch-javascript --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch  --additional-properties projectVersion=`cat versions/javascript`   --additional-properties  usePromises=true $build_to_branch
  git apply patches/javascript.package.patch
  git apply patches/javascript.jsonbig.patch
  cp LICENSE.txt out/manticoresearch-javascript/LICENSE.txt
  #cp docs/javascript/README.md out/manticoresearch-javascript/README.md
  cp docs/javascript/docs/* out/manticoresearch-javascript/docs/
  rm -rf out/manticoresearch-javascript/test/api/* 
  cp -R test/javascript/api/Manual.spec.js out/manticoresearch-javascript/test/api/Manual.spec.js
  echo "Javascript done."
}

do_typescript() {
  echo "Building Typescript ..."
  rm -rf out/manticoresearch-typescript
  docker run \
    --rm \
    -v ${PWD}:/local \
    -u "$(id -u):$(id -g)" \
    -e JAVA_OPTS="-Dlog.level=warn" \
    "openapitools/openapi-generator-cli$version" generate \
    -i /local/manticore.yml \
    -g typescript-fetch \
    -o /local/out/manticoresearch-typescript \
    -t /local/templates/Typescript \
    --git-repo-id manticoresearch-typescript \
    --git-user-id manticoresoftware \
    --reserved-words-mappings delete=delete \
    --additional-properties npmName=manticoresearch-ts \
    --additional-properties npmVersion=`cat versions/typescript` \
    --additional-properties withoutRuntimeChecks=true \
    --additional-properties stringEnums=true \
    --additional-properties enumPropertyNaming=camelCase \
    --additional-properties modelPropertyNaming=original \
    --additional-properties useSingleRequestParameter=false \
    --additional-properties moduleName=Manticoresearch \
    --additional-properties apiDocPath=docs/ \
    $build_to_branch
  cp LICENSE.txt out/manticoresearch-typescript/LICENSE.txt
  mkdir out/manticoresearch-typescript/test/ && cp -R test/typescript/* out/manticoresearch-typescript/test/
  cp -r docs/typescript/docs out/manticoresearch-typescript/
  # Adding a custom tsup config we use for package build
  cp misc/typescript/tsup.config.ts out/manticoresearch-typescript/tsup.config.ts
  echo "Typescript done."
}

do_csharp() {
  echo "Building CSharp ..."
  rm -rf out/manticoresearch-csharp 
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g csharp-netcore  -o /local/out/manticoresearch-csharp -t /local/templates/csharp-netcore --library httpclient --git-repo-id manticoresearch-csharp --git-user-id manticoresoftware --additional-properties packageName=ManticoreSearch --additional-properties library=httpclient --additional-properties packageVersion=`cat versions/csharp` $build_to_branch
  cp -r gh-actions/csharp/. out/manticoresearch-csharp
  cp -r docs/csharp/docs/SearchApi.md out/manticoresearch-csharp/docs/SearchApi.md
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

if [ ! -z $2 -a $2 != 'v0' ]; then
version=":$2"
else
version=''
fi


if [ ! -z $3 ]; then
	if [ $3 != 'dev' ]; then
		build_to_branch="--additional-properties outBranch=1"
	else
		build_to_branch="--additional-properties devBranch=1"
	fi
else
build_to_branch=''
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
 typescript)
   do_typescript
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
   do_typescript
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
