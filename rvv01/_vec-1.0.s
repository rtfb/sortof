
# This should be an equivalent of vec.s compliant to RVV 1.0 spec.
# NOT TESTED.

# a0: *vec
# a1: nelem
# a2: *scalar

.globl vec_add_scalar
vec_add_scalar:   # (float *vec, int nelem, float *scalar);
	vconfig 0x63
	vfld.s v1.s, 0(a2)
loop:
	vsetv1 a3, a1
	vflw v0, 0(a0)
	vfadd.s v2, v1, v0
	vsw v2, 0(a0)

	slli a4, a3, 2
	add a0, a0, a4
	sub a1, a1, a3
	bne a1, x0, loop

	vconfig 0x0
	ret
