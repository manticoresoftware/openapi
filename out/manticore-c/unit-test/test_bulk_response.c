#ifndef bulk_response_TEST
#define bulk_response_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define bulk_response_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/bulk_response.h"
bulk_response_t* instantiate_bulk_response(int include_optional);



bulk_response_t* instantiate_bulk_response(int include_optional) {
  bulk_response_t* bulk_response = NULL;
  if (include_optional) {
    bulk_response = bulk_response_create(
      0,
      1
    );
  } else {
    bulk_response = bulk_response_create(
      0,
      1
    );
  }

  return bulk_response;
}


#ifdef bulk_response_MAIN

void test_bulk_response(int include_optional) {
    bulk_response_t* bulk_response_1 = instantiate_bulk_response(include_optional);

	cJSON* jsonbulk_response_1 = bulk_response_convertToJSON(bulk_response_1);
	printf("bulk_response :\n%s\n", cJSON_Print(jsonbulk_response_1));
	bulk_response_t* bulk_response_2 = bulk_response_parseFromJSON(jsonbulk_response_1);
	cJSON* jsonbulk_response_2 = bulk_response_convertToJSON(bulk_response_2);
	printf("repeating bulk_response:\n%s\n", cJSON_Print(jsonbulk_response_2));
}

int main() {
  test_bulk_response(1);
  test_bulk_response(0);

  printf("Hello world \n");
  return 0;
}

#endif // bulk_response_MAIN
#endif // bulk_response_TEST
