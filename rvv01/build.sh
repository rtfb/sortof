#!/bin/bash

/opt/riscv/bin/riscv64-unknown-linux-gnu-as -march=rv64gv -o vec.o vec.s
/opt/riscv/bin/riscv64-unknown-linux-gnu-gcc -c -o main.o main.c
/opt/riscv/bin/riscv64-unknown-linux-gnu-gcc -o rvv01 main.o vec.o
