CFLAGS = -std=c99
CFLAGS += -lcunit

test: tests.out
	@./tests.out

tests.out: test_graph.c graph.c
	@$(CC) $(CFLAGS) vendor/unity.c graph.c test_graph.c -o tests.out