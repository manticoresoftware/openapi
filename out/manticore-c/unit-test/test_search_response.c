#ifndef search_response_TEST
#define search_response_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define search_response_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/search_response.h"
search_response_t* instantiate_search_response(int include_optional);



search_response_t* instantiate_search_response(int include_optional) {
  search_response_t* search_response = NULL;
  if (include_optional) {
    search_response = search_response_create(
      56,
      1,
      {"total":2,"hits":[{"_id":1,"_score":1,"_source":{"gid":11}},{"_id":2,"_score":1,"_source":{"gid":20}}]},
      0
    );
  } else {
    search_response = search_response_create(
      56,
      1,
      {"total":2,"hits":[{"_id":1,"_score":1,"_source":{"gid":11}},{"_id":2,"_score":1,"_source":{"gid":20}}]},
      0
    );
  }

  return search_response;
}


#ifdef search_response_MAIN

void test_search_response(int include_optional) {
    search_response_t* search_response_1 = instantiate_search_response(include_optional);

	cJSON* jsonsearch_response_1 = search_response_convertToJSON(search_response_1);
	printf("search_response :\n%s\n", cJSON_Print(jsonsearch_response_1));
	search_response_t* search_response_2 = search_response_parseFromJSON(jsonsearch_response_1);
	cJSON* jsonsearch_response_2 = search_response_convertToJSON(search_response_2);
	printf("repeating search_response:\n%s\n", cJSON_Print(jsonsearch_response_2));
}

int main() {
  test_search_response(1);
  test_search_response(0);

  printf("Hello world \n");
  return 0;
}

#endif // search_response_MAIN
#endif // search_response_TEST
