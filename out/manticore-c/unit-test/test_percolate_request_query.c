#ifndef percolate_request_query_TEST
#define percolate_request_query_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define percolate_request_query_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/percolate_request_query.h"
percolate_request_query_t* instantiate_percolate_request_query(int include_optional);



percolate_request_query_t* instantiate_percolate_request_query(int include_optional) {
  percolate_request_query_t* percolate_request_query = NULL;
  if (include_optional) {
    percolate_request_query = percolate_request_query_create(
      0
    );
  } else {
    percolate_request_query = percolate_request_query_create(
      0
    );
  }

  return percolate_request_query;
}


#ifdef percolate_request_query_MAIN

void test_percolate_request_query(int include_optional) {
    percolate_request_query_t* percolate_request_query_1 = instantiate_percolate_request_query(include_optional);

	cJSON* jsonpercolate_request_query_1 = percolate_request_query_convertToJSON(percolate_request_query_1);
	printf("percolate_request_query :\n%s\n", cJSON_Print(jsonpercolate_request_query_1));
	percolate_request_query_t* percolate_request_query_2 = percolate_request_query_parseFromJSON(jsonpercolate_request_query_1);
	cJSON* jsonpercolate_request_query_2 = percolate_request_query_convertToJSON(percolate_request_query_2);
	printf("repeating percolate_request_query:\n%s\n", cJSON_Print(jsonpercolate_request_query_2));
}

int main() {
  test_percolate_request_query(1);
  test_percolate_request_query(0);

  printf("Hello world \n");
  return 0;
}

#endif // percolate_request_query_MAIN
#endif // percolate_request_query_TEST
