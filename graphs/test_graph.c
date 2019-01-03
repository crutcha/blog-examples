#include "vendor/unity.h"

void test_the_test(void)
{
    TEST_ASSERT_EQUAL_STRING("HELLO", "HELLO");
}

int main(void)
{
    UnityBegin("test_graph.c");
    RUN_TEST(test_the_test);
    UnityEnd();

    return 0;
}