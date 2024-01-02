.option norvc
.section .data
.section .text.init
.global _start

_start:
    li      t0, 0x1800
    csrc    mstatus, t0
    li      t0, 0x802
    csrs    mstatus, t0
    li      t0, 0xa0
    csrs    mie, t0
    li      t0, 0x20
    csrs    mideleg, t0
    la      t0, mtrap
    csrw    mtvec, t0
    li      t0, 500000
    li      t1, 0x2004000
    sd      t0, 0(t1)
    la      t1, supervisor
    csrw    mepc, t1
    mret

mtrap:
    li      t0, 0x20
    csrs    mip, t0
    mret

supervisor:
    li      t0, 20
    la      t0, strap
    csrw    stvec, t0
    j       spin

strap:
    j       spin

spin:
    j       spin
