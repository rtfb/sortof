import cocotb
from cocotb.clock import Clock
from cocotb.triggers import RisingEdge, FallingEdge, Timer, ClockCycles


@cocotb.test()
async def test_my_design(dut):
    dut._log.info("start")
    clock = Clock(dut.clk, 1, units="ms")
    cocotb.start_soon(clock.start())

    dut._log.info("reset")
    dut.rst_n.value = 0 # low to reset
    await ClockCycles(dut.clk, 10)
    dut.rst_n.value = 1 # take out of reset

    # want_hex_digits = [0x4, 0x3, 0x2, 0x1, 0xd, 0xc, 0xb, 0xa]
    want_hex_digits = [4, 3, 2, 1, 13, 12, 11, 10]
    # want_hex_digits = [4, 3, 2, 1, 13, 12, 11]

    await ClockCycles(dut.clk, 1)

    for d in want_hex_digits:
        await ClockCycles(dut.clk, 1)
        assert int(dut.have_output) == 1
        assert int(dut.output_bits) == d
    # await ClockCycles(dut.clk, 10)

    # await ClockCycles(dut.clk, 1)
    # assert int(dut.output_bits.value) == 0x4
    # await ClockCycles(dut.clk, 1)
    # assert int(dut.output_bits.value) == 0x3
    # await ClockCycles(dut.clk, 1)
    # assert int(dut.output_bits.value) == 0x2
    # await ClockCycles(dut.clk, 1)
    # assert int(dut.output_bits.value) == 0x1
    # await ClockCycles(dut.clk, 1)
    # assert int(dut.output_bits.value) == 0xd
    # await ClockCycles(dut.clk, 1)
    # assert int(dut.output_bits.value) == 0xc
    # await ClockCycles(dut.clk, 1)
    # assert int(dut.output_bits.value) == 0xb
    # await ClockCycles(dut.clk, 1)
    # assert int(dut.output_bits.value) == 0xa
