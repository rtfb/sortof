
Following this guide: <https://ebpf-go.dev/guides/getting-started/>

Needed some software installed:
`sudo apt install clang libbpf-dev llvm`

And a bit of patching up:
`sudo ln -s /usr/include/x86_64-linux-gnu/asm /usr/include/asm` (according to
[this SO answer][1])

[1]: https://stackoverflow.com/a/77465534
