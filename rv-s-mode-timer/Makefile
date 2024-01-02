
install:
	sudo apt install gdb-multiarch

timers.o: timers.S
	riscv64-linux-gnu-as $< -o $@

timers: timers.o
	riscv64-linux-gnu-ld -Tvirt.ld timers.o -o $@

run: timers
	docker run -it --name qemu-image --rm \
		--mount type=bind,source=.,target=/host \
		--net=host qemu-image:latest \
		-machine virt -cpu rv64 -smp 1 \
		-s -S -nographic -bios none -kernel timers

gdb:
	gdb-multiarch timers \
		-ex "set architecture riscv:rv64" \
		-ex "target remote :1234" \
		-ex 'display /i $$pc' \
		-ex 'break *_start' \
		-ex 'break *supervisor' \
		-ex 'break *mtrap' \
		-ex 'break *strap' \
		-ex 'display *0x2004000' \
		-ex 'display *0x200BFF8'
