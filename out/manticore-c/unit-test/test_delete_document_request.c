#ifndef delete_document_request_TEST
#define delete_document_request_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define delete_document_request_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/delete_document_request.h"
delete_document_request_t* instantiate_delete_document_request(int include_optional);



delete_document_request_t* instantiate_delete_document_request(int include_optional) {
  delete_document_request_t* delete_document_request = NULL;
  if (include_optional) {
    delete_document_request = delete_document_request_create(
      "0",
      1,
      0
    );
  } else {
    delete_document_request = delete_document_request_create(
      "0",
      1,
      0
    );
  }

  return delete_document_request;
}


#ifdef delete_document_request_MAIN

void test_delete_document_request(int include_optional) {
    delete_document_request_t* delete_document_request_1 = instantiate_delete_document_request(include_optional);

	cJSON* jsondelete_document_request_1 = delete_document_request_convertToJSON(delete_document_request_1);
	printf("delete_document_request :\n%s\n", cJSON_Print(jsondelete_document_request_1));
	delete_document_request_t* delete_document_request_2 = delete_document_request_parseFromJSON(jsondelete_document_request_1);
	cJSON* jsondelete_document_request_2 = delete_document_request_convertToJSON(delete_document_request_2);
	printf("repeating delete_document_request:\n%s\n", cJSON_Print(jsondelete_document_request_2));
}

int main() {
  test_delete_document_request(1);
  test_delete_document_request(0);

  printf("Hello world \n");
  return 0;
}

#endif // delete_document_request_MAIN
#endif // delete_document_request_TEST
