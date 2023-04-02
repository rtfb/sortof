#!/bin/bash

zig build-obj adder.zig
gcc main.c adder.o
./a.out
