
ARDUINO := ${HOME}/.arduino15
TOOLCHAIN := ${ARDUINO}/packages/arduino/tools/avr-gcc/7.3.0-atmel3.6.1-arduino7
# DEVICE := /dev/arduino-uno
DEVICE := /dev/ttyACM0

# %.hex: %.asm
# 	avra -fI $<
# 	rm *.eep.hex *.obj *.cof

%.o: %.S
	${TOOLCHAIN}/bin/avr-as -c $< -o $@

%.o: %.c
	${TOOLCHAIN}/bin/avr-gcc -c $< -o $@

all: $(patsubst %.asm,%.hex,$(wildcard *.asm)) \
	$(patsubst %.c,%.o,$(wildcard *.c)) \
	$(patsubst %.S,%.o,$(wildcard *.S))

blinker: blinker.o hello.o
	${TOOLCHAIN}/bin/avr-gcc $< -o $@

blinker.hex: blinker
	${TOOLCHAIN}/bin/avr-objcopy -O ihex $< $@

upload: ${program}.hex
	avrdude -c arduino -p m328p -P ${DEVICE} -b 115200 -U flash:w:$<

monitor:
	picocom --send-cmd "ascii_xfr -s -v -l10" --nolock ${DEVICE}

.PHONY: all upload monitor
