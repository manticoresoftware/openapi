#ifndef inline_object_TEST
#define inline_object_TEST

// the following is to include only the main from the first c file
#ifndef TEST_MAIN
#define TEST_MAIN
#define inline_object_MAIN
#endif // TEST_MAIN

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include "../external/cJSON.h"

#include "../model/inline_object.h"
inline_object_t* instantiate_inline_object(int include_optional);



inline_object_t* instantiate_inline_object(int include_optional) {
  inline_object_t* inline_object = NULL;
  if (include_optional) {
    inline_object = inline_object_create(
      "0",
      "0"
    );
  } else {
    inline_object = inline_object_create(
      "0",
      "0"
    );
  }

  return inline_object;
}


#ifdef inline_object_MAIN

void test_inline_object(int include_optional) {
    inline_object_t* inline_object_1 = instantiate_inline_object(include_optional);

	cJSON* jsoninline_object_1 = inline_object_convertToJSON(inline_object_1);
	printf("inline_object :\n%s\n", cJSON_Print(jsoninline_object_1));
	inline_object_t* inline_object_2 = inline_object_parseFromJSON(jsoninline_object_1);
	cJSON* jsoninline_object_2 = inline_object_convertToJSON(inline_object_2);
	printf("repeating inline_object:\n%s\n", cJSON_Print(jsoninline_object_2));
}

int main() {
  test_inline_object(1);
  test_inline_object(0);

  printf("Hello world \n");
  return 0;
}

#endif // inline_object_MAIN
#endif // inline_object_TEST
