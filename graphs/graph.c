/* 
 * Undirected graph struct represented as a list type where each node will
 * each node can contain 1 or more links to other nodes as pointers. There
 * will be a default max of 3 links allowed from any one node to another node,
 * but we will also allow a way to dynamically resize the node's link array.
 */

#include <stdio.h>

struct graph {
    int vertices;
    struct node **nodes;
};

struct node {
    int visited;
    int node;
};
