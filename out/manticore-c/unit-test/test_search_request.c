#ifndef search_request_TEST
#define search_request_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define search_request_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/search_request.h"
search_request_t* instantiate_search_request(int include_optional);



search_request_t* instantiate_search_request(int include_optional) {
  search_request_t* search_request = NULL;
  if (include_optional) {
    search_request = search_request_create(
      "test",
      list_create(),
      56,
      56,
      56,
      list_create(),
      0,
      0,
      list_create(),
      1
    );
  } else {
    search_request = search_request_create(
      "test",
      list_create(),
      56,
      56,
      56,
      list_create(),
      0,
      0,
      list_create(),
      1
    );
  }

  return search_request;
}


#ifdef search_request_MAIN

void test_search_request(int include_optional) {
    search_request_t* search_request_1 = instantiate_search_request(include_optional);

	cJSON* jsonsearch_request_1 = search_request_convertToJSON(search_request_1);
	printf("search_request :\n%s\n", cJSON_Print(jsonsearch_request_1));
	search_request_t* search_request_2 = search_request_parseFromJSON(jsonsearch_request_1);
	cJSON* jsonsearch_request_2 = search_request_convertToJSON(search_request_2);
	printf("repeating search_request:\n%s\n", cJSON_Print(jsonsearch_request_2));
}

int main() {
  test_search_request(1);
  test_search_request(0);

  printf("Hello world \n");
  return 0;
}

#endif // search_request_MAIN
#endif // search_request_TEST
