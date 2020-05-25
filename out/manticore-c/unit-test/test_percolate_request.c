#ifndef percolate_request_TEST
#define percolate_request_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define percolate_request_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/percolate_request.h"
percolate_request_t* instantiate_percolate_request(int include_optional);

#include "test_percolate_request_query.c"


percolate_request_t* instantiate_percolate_request(int include_optional) {
  percolate_request_t* percolate_request = NULL;
  if (include_optional) {
    percolate_request = percolate_request_create(
       // false, not to have infinite recursion
      instantiate_percolate_request_query(0)
    );
  } else {
    percolate_request = percolate_request_create(
      NULL
    );
  }

  return percolate_request;
}


#ifdef percolate_request_MAIN

void test_percolate_request(int include_optional) {
    percolate_request_t* percolate_request_1 = instantiate_percolate_request(include_optional);

	cJSON* jsonpercolate_request_1 = percolate_request_convertToJSON(percolate_request_1);
	printf("percolate_request :\n%s\n", cJSON_Print(jsonpercolate_request_1));
	percolate_request_t* percolate_request_2 = percolate_request_parseFromJSON(jsonpercolate_request_1);
	cJSON* jsonpercolate_request_2 = percolate_request_convertToJSON(percolate_request_2);
	printf("repeating percolate_request:\n%s\n", cJSON_Print(jsonpercolate_request_2));
}

int main() {
  test_percolate_request(1);
  test_percolate_request(0);

  printf("Hello world \n");
  return 0;
}

#endif // percolate_request_MAIN
#endif // percolate_request_TEST
