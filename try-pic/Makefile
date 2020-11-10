
foo.png: foo.pic

bar.png: bar.pic

%.png: %.pic
	cat $< | pic2graph -density 200x200 -units pixelspercentimeter \
		-background white -alpha remove -alpha off -border 10 \
		-bordercolor white > $@
