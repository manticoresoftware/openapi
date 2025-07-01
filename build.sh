#!/bin/bash

set -e

do_python() {
  echo "Building Python ..."
  rm -rf out/manticoresearch-python 
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g python -o /local/out/manticoresearch-python -t /local/templates/python --git-repo-id manticoresearch-python --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch --additional-properties packageName=manticoresearch --additional-properties packageVersion=`cat versions/python` $build_to_branch
  rm -rf out/manticoresearch-python/test/* 
  cp -r LICENSE.txt out/manticoresearch-python/LICENSE.txt
  cp docs/python/docs/* out/manticoresearch-python/docs/
  # replace test with our tests
  cp -r test/python/* out/manticoresearch-python/test/   
  git apply patches/python_bulk.patch patches/python_sql_api.patch
  echo "Python done."
}

do_python_asyncio() {
  echo "Building Python async..."
  rm -rf out/manticoresearch-python-asyncio
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g python -o /local/out/manticoresearch-python-asyncio -t /local/templates/python --git-repo-id manticoresearch-python-asyncio --git-user-id manticoresoftware  --additional-properties library=asyncio --additional-properties projectName=manticoresearch --additional-properties packageName=manticoresearch-asyncio --additional-properties packageVersion=`cat versions/python` $build_to_branch
  rm -rf out/manticoresearch-python-asyncio/test/* 
  cp -r LICENSE.txt out/manticoresearch-python-asyncio/LICENSE.txt
  #cp docs/python/docs/* out/manticoresearch-python-asyncio/docs/
  # replace test with our tests
  cp -r test/python-asyncio/* out/manticoresearch-python-asyncio/test/   
  git apply patches/python_asyncio_bulk.patch patches/python_asyncio_sql_api.patch
  echo "Python async done."
}

do_rust() {
  echo "Building Rust..."
  rm -rf out/manticoresearch-rust
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g rust -o /local/out/manticoresearch-rust -t /local/templates/rust --git-repo-id manticoresearch-rust --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch --additional-properties packageName=manticoresearch --additional-properties packageVersion=`cat versions/rust` --additional-properties library=hyper $build_to_branch
  cp -r LICENSE.txt out/manticoresearch-rust/LICENSE.txt
  mkdir out/manticoresearch-rust/tests
  cp -r test/rust/* out/manticoresearch-rust/tests/
  git apply patches/rust_request.patch
  echo "Rust done."
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
    --additional-properties prevVersion=$prev_version \
    $build_to_branch
  git apply patches/java.apiclient.patch patches/java.queryfilter.patch patches/java.queryfilter.patch \
  patches/java.searchquery.patch patches/java.sqlresponse.patch 
  cp LICENSE.txt out/manticoresearch-java/LICENSE.txt
  cp docs/java/docs/* out/manticoresearch-java/docs/
  cp -r test/java/api/* out/manticoresearch-java/src/test/java/com/manticoresearch/client/api/
  rm -rf out/manticoresearch-java/.openapi-generator
  rm -rf out/manticoresearch-java/api
  echo "Java done."

}

do_javascript() {
  echo "Building Javascript ..."
  rm -rf out/manticoresearch-javascript 
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn"  "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g javascript -o /local/out/manticoresearch-javascript -t /local/templates/Javascript --git-repo-id manticoresearch-javascript --git-user-id manticoresoftware  --additional-properties projectName=manticoresearch  --additional-properties projectVersion=`cat versions/javascript`   --additional-properties  usePromises=true $build_to_branch
  git apply patches/javascript.sql_api.patch
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
    -g typescript \
    -o /local/out/manticoresearch-typescript \
    -t /local/templates/typescript \
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
    --additional-properties apiDocPath=./ \
    $build_to_branch
  git apply patches/typescript.matchall.patch
  git apply patches/typescript.objectserializer.patch
  git apply patches/typescript.utilsapi.patch
  git apply patches/typescript.indexapi.patch
  cp LICENSE.txt out/manticoresearch-typescript/LICENSE.txt
  mkdir out/manticoresearch-typescript/test/ && cp -R test/typescript/* out/manticoresearch-typescript/test/
 cp -r docs/typescript/docs/* out/manticoresearch-typescript/
  
  # Adding a custom tsup config we use for package build
  #cp misc/typescript/tsup.config.ts out/manticoresearch-typescript/tsup.config.ts
  echo "Typescript done."
}

do_csharp() {
  echo "Building CSharp ..."
  rm -rf out/manticoresearch-net 
  docker run --rm -v ${PWD}:/local   -u "$(id -u):$(id -g)"  -e JAVA_OPTS="-Dlog.level=warn" "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g csharp  -o /local/out/manticoresearch-net -t /local/templates/csharp --library httpclient --git-repo-id manticoresearch-csharp --git-user-id manticoresoftware --additional-properties packageName=ManticoreSearch --additional-properties library=httpclient --additional-properties packageVersion=`cat versions/csharp` $build_to_branch
  git apply patches/net.matchall.patch
  #cp -r gh-actions/csharp/. out/manticoresearch-net
  cp -r docs/csharp/docs/* out/manticoresearch-net/docs/
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
  rm -rf out/manticoresearch-go
  docker run --rm -v ${PWD}:/local  -u "$(id -u):$(id -g)"    -e JAVA_OPTS="-Dlog.level=warn" "openapitools/openapi-generator-cli$version" generate -i /local/manticore.yml -g go  -o /local/out/manticoresearch-go -t /local/templates/go --git-repo-id manticoresearch-go --git-user-id manticoresoftware \
  --additional-properties packageVersion=`cat versions/go` \
  $build_to_branch
  #git apply patches/go.sql_api.patch
  #git apply patches/go.response_error.patch
  #cp patches/go_sql_response.go out/manticoresearch-go/model_sql_response.go
  #Removing redundant files created by the go generator
  cd out/manticoresearch-go &&
  rm -rf model_aggregation_composite_sources_inner_value_terms.go model_aggregation_composite_sources_inner_value.go \
    model_aggregation_composite.go model_aggregation_sort_inner_value.go model_aggregation_terms.go model_attr_filter.go \
    model_basic_search_request.go model_bool_filter_bool.go model_equals_filter_equals.go model_equals_filter.go \
    model_error_response_error_one_of.go model_error_response_error_one_of.go model_geo_filter_geo_distance_location_anchor.go \
    model_geo_filter_geo_distance.go model_geo_filter.go model_highlight_all_of_fields.go model_in_filter.go \
    model_join_basic_cond.go model_join_inner_on_inner_left.go model_join_inner_on_inner.go model_join_inner.go \
    model_knn_doc_id_request.go model_knn_query_vector_request.go model_knn_search_parameters.go model_knn_search_request_all_of_knn.go \
    model_knn_search_request.go model_match_all_filter.go model_match_filter_match.go model_match_filter.go model_match_phrase_filter.go \
    model_query_string_filter.go model_range_filter_range_value.go model_range_filter.go model_search_request_parameters__source.go \
    model_search_request_parameters_sort_inner.go model_search_request_parameters.go model_sort_object.go model_source_by_rules.go
  cd docs &&
  rm -rf AggregationCompositeSourcesInnerValueTerms.md AggregationCompositeSourcesInnerValue.md \
    AggregationComposite.md AggregationSortInnerValue.md AggregationTerms.md AttrFilter.md \
    BasicSearchRequest.md BoolFilterBool.md EqualsFilterEquals.md EqualsFilter.md \
    ErrorResponseErrorOneOf.md ErrorResponseErrorOneOf.md GeoFilterGeoDistanceLocationAnchor.md \
    GeoFilterGeoDistance.md GeoFilter.md HighlightAllOfFields.md InFilter.md \
    JoinBasicCond.md JoinInnerOnInnerLeft.md JoinInnerOnInner.md JoinInner.md \
    KnnSearchRequest.md KnnDocIdRequest.md KnnQueryVectorRequest.md KnnSearchParameters.md KnnSearchRequestAllOfKnn.md \
    MatchAllFilter.md MatchFilterMatch.md MatchFilter.md MatchPhraseFilter.md \
    QueryStringFilter.md RangeFilterRangeValue.md RangeFilter.md SearchRequestParametersSource.md \
    SearchRequestParametersSortInner.md SearchRequestParameters.md SortObject.md SourceByRules.md
  cd ../../../
  cp -r test/go/* out/manticoresearch-go/test/
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
 python-asyncio)
   do_python_asyncio 
  ;;
 rust)
   do_rust
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
