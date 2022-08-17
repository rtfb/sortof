// #include "kernel.h"
#include <stdint.h>

unsigned int get_mhartid() {
    register unsigned int a0 asm ("a0");
    asm volatile (
        "csrr a0, mhartid"
        : "=r"(a0)   // output in a0
    );
    return a0;
}

void* shift_right_addr(void* addr, int bits) {
    unsigned long iaddr = (unsigned long)addr;
    return (void*)(iaddr >> bits);
}

void set_pmpaddr0(void* addr) {
    addr = shift_right_addr(addr, 2);
    asm volatile (
        "csrw   pmpaddr0, %0;"  // set pmpaddr0 to the requested addr
        :                       // no output
        : "r"(addr)             // input in addr
    );
}

void set_pmpaddr1(void* addr) {
    addr = shift_right_addr(addr, 2);
    asm volatile (
        "csrw   pmpaddr1, %0;"  // set pmpaddr1 to the requested addr
        :                       // no output
        : "r"(addr)             // input in addr
    );
}

void set_pmpaddr2(void* addr) {
    addr = shift_right_addr(addr, 2);
    asm volatile (
        "csrw   pmpaddr2, %0;"  // set pmpaddr2 to the requested addr
        :                       // no output
        : "r"(addr)             // input in addr
    );
}

void set_pmpaddr3(void* addr) {
    addr = shift_right_addr(addr, 2);
    asm volatile (
        "csrw   pmpaddr3, %0;"  // set pmpaddr3 to the requested addr
        :                       // no output
        : "r"(addr)             // input in addr
    );
}

void set_pmpcfg0(unsigned long value) {
    asm volatile (
        "csrw   pmpcfg0, %0;"   // set pmpcfg0 to the requested value
        :                       // no output
        : "r"(value)            // input in value
    );
}

unsigned int get_mstatus() {
    register unsigned int a0 asm ("a0");
    asm volatile (
        "csrr a0, mstatus"
        : "=r"(a0)   // output in a0
    );
    return a0;
}

void set_mstatus(unsigned int value) {
    asm volatile (
        "csrw mstatus, %0"
        :            // no output
        : "r"(value) // input in value
    );
}

void* get_mepc() {
    register void* a0 asm ("a0");
    asm volatile (
        "csrr a0, mepc"
        : "=r"(a0)   // output in a0
    );
    return a0;
}

void set_jump_address(void *func) {
    asm volatile (
        "csrw   mepc, %0;"  // set mepc to userland function
        :                   // no output
        : "r"(func)         // input in func
    );
}

void set_mscratch(void* ptr) {
    asm volatile (
        "csrw mscratch, %0"
        :            // no output
        : "r"(ptr)   // input in value
    );
}
