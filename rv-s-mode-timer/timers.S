.option norvc
.section .data
.section .text.init
.global _start

_start:
    li      t0, 0x1800
    csrc    mstatus, t0
    li      t0, 0x8fffffff
    csrw    pmpaddr0, t0
    li      t0, 0xf
    csrw    pmpcfg0, t0
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
    # advance mtimecmp to clear the current MTIP bit in mip:
    li      t1, 0x2004000
    ld      t0, 0(t1)
    li      t2, 500000
    add     t0, t0, t2
    sd      t0, 0(t1)
    # set the STIP bit to cause mret to trap into strap:
    li      t0, 0x20
    csrw    mip, t0
    # set mstatus.SIE=1 again because it gets reset:
    li      t0, 0x2
    csrs    mstatus, t0
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
