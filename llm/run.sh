#!/bin/bash

./llama.cpp/main -ngl 32 -m nous-capybara-34b.Q4_K_M.gguf \
    --color -c 2048 --temp 0.7 --repeat_penalty 1.1 -n -1 -i -ins
