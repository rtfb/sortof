#!/bin/bash

./zig-linux-x86_64-0.9.1/zig cc hello.c
./a.out

# ./zig-linux-x86_64-0.9.1/zig cc riscv.c -o riscv.o -target riscv64-linux
# ./zig-linux-x86_64-0.9.1/zig cc -C riscv.c -o riscv.o -target riscv64-linux
# ./zig-linux-x86_64-0.9.1/zig cc targets
# ./zig-linux-x86_64-0.9.1/zig targets
# ./zig-linux-x86_64-0.9.1/zig targets | jq .arch
# ./zig-linux-x86_64-0.9.1/zig cc -C riscv.c -o riscv.o -target riscv64
# ./zig-linux-x86_64-0.9.1/zig cc riscv.c -o riscv.o -target riscv64
# ./zig-linux-x86_64-0.9.1/zig cc riscv.c -o riscv.o -target riscv64-freestanding
