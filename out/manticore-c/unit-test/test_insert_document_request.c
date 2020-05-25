#ifndef insert_document_request_TEST
#define insert_document_request_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define insert_document_request_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/insert_document_request.h"
insert_document_request_t* instantiate_insert_document_request(int include_optional);



insert_document_request_t* instantiate_insert_document_request(int include_optional) {
  insert_document_request_t* insert_document_request = NULL;
  if (include_optional) {
    insert_document_request = insert_document_request_create(
      "0",
      1,
      list_create()
    );
  } else {
    insert_document_request = insert_document_request_create(
      "0",
      1,
      list_create()
    );
  }

  return insert_document_request;
}


#ifdef insert_document_request_MAIN

void test_insert_document_request(int include_optional) {
    insert_document_request_t* insert_document_request_1 = instantiate_insert_document_request(include_optional);

	cJSON* jsoninsert_document_request_1 = insert_document_request_convertToJSON(insert_document_request_1);
	printf("insert_document_request :\n%s\n", cJSON_Print(jsoninsert_document_request_1));
	insert_document_request_t* insert_document_request_2 = insert_document_request_parseFromJSON(jsoninsert_document_request_1);
	cJSON* jsoninsert_document_request_2 = insert_document_request_convertToJSON(insert_document_request_2);
	printf("repeating insert_document_request:\n%s\n", cJSON_Print(jsoninsert_document_request_2));
}

int main() {
  test_insert_document_request(1);
  test_insert_document_request(0);

  printf("Hello world \n");
  return 0;
}

#endif // insert_document_request_MAIN
#endif // insert_document_request_TEST
