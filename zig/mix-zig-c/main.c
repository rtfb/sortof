#include <stdio.h>

// implemented in adder.zig
extern int adder(int a, int b);

int main() {
    printf("addered: %d\n", adder(5, 7));
}
