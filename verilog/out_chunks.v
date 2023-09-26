module out_chunks (
    input clk,
    input reset,
    output have_output,
    output [CHUNK_SIZE_BITS-1:0]out_bits
);
    localparam CHUNK_SIZE_BITS = 4;
    localparam NUM_CHUNKS = 8;
    localparam ADDR_BITS = $clog2(NUM_CHUNKS);
    localparam DATA_SIZE_BITS = NUM_CHUNKS*CHUNK_SIZE_BITS;

    reg [ADDR_BITS-1:0] bit_addr;
    reg [DATA_SIZE_BITS-1:0]internal_data;
    reg done_output;

    initial begin
        internal_data = 32'habcd1234;
    end

    always @(posedge clk or posedge reset)
    begin
        have_output <= !reset && !done_output;

        if (reset) begin
            bit_addr <= 0;
            done_output <= 0;
            have_output <= 0;
        end else if (!done_output) begin
            {done_output, bit_addr} <= bit_addr + 1;
        end
    end

    assign out_bits = internal_data[bit_addr*CHUNK_SIZE_BITS-1 -: CHUNK_SIZE_BITS];
endmodule
