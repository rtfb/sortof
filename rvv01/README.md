## Experimental code exercising RISC-V Vector extension

Since the Allwinner D1 implements an outdated spec v0.7.1, we need to build a
custom toolchain to compile it. Build the toolchain from here:
https://github.com/brucehoult/riscv-gnu-toolchain/tree/rvv-0.7.1

The official riscv toolchain repo has deleted the `rvv-0.7.1` branch, so thanks
to Bruce Hoult (http://hoult.org/bruce/) for the archaeology.

The spec is still available, though:
https://github.com/riscv/riscv-v-spec/releases/download/0.7.1/riscv-v-spec-0.7.1.pdf.
