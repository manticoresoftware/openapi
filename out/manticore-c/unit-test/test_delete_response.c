#ifndef delete_response_TEST
#define delete_response_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define delete_response_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/delete_response.h"
delete_response_t* instantiate_delete_response(int include_optional);



delete_response_t* instantiate_delete_response(int include_optional) {
  delete_response_t* delete_response = NULL;
  if (include_optional) {
    delete_response = delete_response_create(
      "0",
      56,
      56,
      "0"
    );
  } else {
    delete_response = delete_response_create(
      "0",
      56,
      56,
      "0"
    );
  }

  return delete_response;
}


#ifdef delete_response_MAIN

void test_delete_response(int include_optional) {
    delete_response_t* delete_response_1 = instantiate_delete_response(include_optional);

	cJSON* jsondelete_response_1 = delete_response_convertToJSON(delete_response_1);
	printf("delete_response :\n%s\n", cJSON_Print(jsondelete_response_1));
	delete_response_t* delete_response_2 = delete_response_parseFromJSON(jsondelete_response_1);
	cJSON* jsondelete_response_2 = delete_response_convertToJSON(delete_response_2);
	printf("repeating delete_response:\n%s\n", cJSON_Print(jsondelete_response_2));
}

int main() {
  test_delete_response(1);
  test_delete_response(0);

  printf("Hello world \n");
  return 0;
}

#endif // delete_response_MAIN
#endif // delete_response_TEST
