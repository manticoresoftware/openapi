#ifndef success_response_TEST
#define success_response_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define success_response_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/success_response.h"
success_response_t* instantiate_success_response(int include_optional);



success_response_t* instantiate_success_response(int include_optional) {
  success_response_t* success_response = NULL;
  if (include_optional) {
    success_response = success_response_create(
      "0",
      56,
      1,
      "0",
      1
    );
  } else {
    success_response = success_response_create(
      "0",
      56,
      1,
      "0",
      1
    );
  }

  return success_response;
}


#ifdef success_response_MAIN

void test_success_response(int include_optional) {
    success_response_t* success_response_1 = instantiate_success_response(include_optional);

	cJSON* jsonsuccess_response_1 = success_response_convertToJSON(success_response_1);
	printf("success_response :\n%s\n", cJSON_Print(jsonsuccess_response_1));
	success_response_t* success_response_2 = success_response_parseFromJSON(jsonsuccess_response_1);
	cJSON* jsonsuccess_response_2 = success_response_convertToJSON(success_response_2);
	printf("repeating success_response:\n%s\n", cJSON_Print(jsonsuccess_response_2));
}

int main() {
  test_success_response(1);
  test_success_response(0);

  printf("Hello world \n");
  return 0;
}

#endif // success_response_MAIN
#endif // success_response_TEST
