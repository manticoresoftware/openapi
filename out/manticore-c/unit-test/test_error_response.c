#ifndef error_response_TEST
#define error_response_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define error_response_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/error_response.h"
error_response_t* instantiate_error_response(int include_optional);



error_response_t* instantiate_error_response(int include_optional) {
  error_response_t* error_response = NULL;
  if (include_optional) {
    error_response = error_response_create(
      list_create(),
      500
    );
  } else {
    error_response = error_response_create(
      list_create(),
      500
    );
  }

  return error_response;
}


#ifdef error_response_MAIN

void test_error_response(int include_optional) {
    error_response_t* error_response_1 = instantiate_error_response(include_optional);

	cJSON* jsonerror_response_1 = error_response_convertToJSON(error_response_1);
	printf("error_response :\n%s\n", cJSON_Print(jsonerror_response_1));
	error_response_t* error_response_2 = error_response_parseFromJSON(jsonerror_response_1);
	cJSON* jsonerror_response_2 = error_response_convertToJSON(error_response_2);
	printf("repeating error_response:\n%s\n", cJSON_Print(jsonerror_response_2));
}

int main() {
  test_error_response(1);
  test_error_response(0);

  printf("Hello world \n");
  return 0;
}

#endif // error_response_MAIN
#endif // error_response_TEST
