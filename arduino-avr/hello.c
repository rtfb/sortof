
#include <stdio.h>

extern void start();
extern void led(unsigned char onoff);

int main() {
    printf("hello\n");
    start();
    while (1) {
        printf("ledon\n");
        led(1);
        led(0);
    }
}
