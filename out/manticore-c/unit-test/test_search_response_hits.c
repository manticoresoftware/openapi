#ifndef search_response_hits_TEST
#define search_response_hits_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define search_response_hits_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/search_response_hits.h"
search_response_hits_t* instantiate_search_response_hits(int include_optional);



search_response_hits_t* instantiate_search_response_hits(int include_optional) {
  search_response_hits_t* search_response_hits = NULL;
  if (include_optional) {
    search_response_hits = search_response_hits_create(
      56,
      list_create()
    );
  } else {
    search_response_hits = search_response_hits_create(
      56,
      list_create()
    );
  }

  return search_response_hits;
}


#ifdef search_response_hits_MAIN

void test_search_response_hits(int include_optional) {
    search_response_hits_t* search_response_hits_1 = instantiate_search_response_hits(include_optional);

	cJSON* jsonsearch_response_hits_1 = search_response_hits_convertToJSON(search_response_hits_1);
	printf("search_response_hits :\n%s\n", cJSON_Print(jsonsearch_response_hits_1));
	search_response_hits_t* search_response_hits_2 = search_response_hits_parseFromJSON(jsonsearch_response_hits_1);
	cJSON* jsonsearch_response_hits_2 = search_response_hits_convertToJSON(search_response_hits_2);
	printf("repeating search_response_hits:\n%s\n", cJSON_Print(jsonsearch_response_hits_2));
}

int main() {
  test_search_response_hits(1);
  test_search_response_hits(0);

  printf("Hello world \n");
  return 0;
}

#endif // search_response_hits_MAIN
#endif // search_response_hits_TEST
