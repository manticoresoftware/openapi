#ifndef update_response_TEST
#define update_response_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define update_response_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/update_response.h"
update_response_t* instantiate_update_response(int include_optional);



update_response_t* instantiate_update_response(int include_optional) {
  update_response_t* update_response = NULL;
  if (include_optional) {
    update_response = update_response_create(
      "0",
      56,
      56,
      "0"
    );
  } else {
    update_response = update_response_create(
      "0",
      56,
      56,
      "0"
    );
  }

  return update_response;
}


#ifdef update_response_MAIN

void test_update_response(int include_optional) {
    update_response_t* update_response_1 = instantiate_update_response(include_optional);

	cJSON* jsonupdate_response_1 = update_response_convertToJSON(update_response_1);
	printf("update_response :\n%s\n", cJSON_Print(jsonupdate_response_1));
	update_response_t* update_response_2 = update_response_parseFromJSON(jsonupdate_response_1);
	cJSON* jsonupdate_response_2 = update_response_convertToJSON(update_response_2);
	printf("repeating update_response:\n%s\n", cJSON_Print(jsonupdate_response_2));
}

int main() {
  test_update_response(1);
  test_update_response(0);

  printf("Hello world \n");
  return 0;
}

#endif // update_response_MAIN
#endif // update_response_TEST
