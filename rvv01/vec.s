
.globl vec_add_scalar
vec_add_scalar:
	# (float *vec, int nelem, float *scalar);
	# a0: *vec
	# a1: nelem
	# a2: *scalar

	flw f0, 0(a2)    # load *scalar to f0

loop:
	# set vector processor's mode to nelem (in a1) and to 32-bit width (e32).
	# The a3 will be set to vector length (vl CSR's value).
	# If a1 is smaller than the largest possible vector, a3 will be equal to a1
	# (tail handling).
	vsetvli a3, a1, e32

	vlw.v v0, 0(a0)      # load the first batch to v0
	vfadd.vf v2, v0, f0  # add f0 to each element in the vector
	vsw.v v2, 0(a0)      # store the result back

	# advance the pointers: a4 = a3*4 (vector length * 4bytes per elem)
	# a0 += a4
	# a1 -= a3
	# keep looping if a1 non-zero
	slli a4, a3, 2
	add a0, a0, a4
	sub a1, a1, a3
	bne a1, x0, loop

	ret
