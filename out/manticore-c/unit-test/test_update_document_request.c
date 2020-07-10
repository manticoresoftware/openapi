#ifndef update_document_request_TEST
#define update_document_request_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define update_document_request_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/update_document_request.h"
update_document_request_t* instantiate_update_document_request(int include_optional);



update_document_request_t* instantiate_update_document_request(int include_optional) {
  update_document_request_t* update_document_request = NULL;
  if (include_optional) {
    update_document_request = update_document_request_create(
      "0",
      {"gid":10},
      56,
      {"query":{"match":{"title":"match me"}}}
    );
  } else {
    update_document_request = update_document_request_create(
      "0",
      {"gid":10},
      56,
      {"query":{"match":{"title":"match me"}}}
    );
  }

  return update_document_request;
}


#ifdef update_document_request_MAIN

void test_update_document_request(int include_optional) {
    update_document_request_t* update_document_request_1 = instantiate_update_document_request(include_optional);

	cJSON* jsonupdate_document_request_1 = update_document_request_convertToJSON(update_document_request_1);
	printf("update_document_request :\n%s\n", cJSON_Print(jsonupdate_document_request_1));
	update_document_request_t* update_document_request_2 = update_document_request_parseFromJSON(jsonupdate_document_request_1);
	cJSON* jsonupdate_document_request_2 = update_document_request_convertToJSON(update_document_request_2);
	printf("repeating update_document_request:\n%s\n", cJSON_Print(jsonupdate_document_request_2));
}

int main() {
  test_update_document_request(1);
  test_update_document_request(0);

  printf("Hello world \n");
  return 0;
}

#endif // update_document_request_MAIN
#endif // update_document_request_TEST
