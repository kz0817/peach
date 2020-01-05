IMG_SIZE := 134217728 # 128MiB

MiB = $(shell expr 1024 '*' 1024)
SECT_SZ = 512

P1_SECTORS := $(shell expr 8 '*' $(MiB) / $(SECT_SZ))
P1_START_SECT := 8192
P1_END_SECT := $(shell expr $(P1_START_SECT) + $(P1_SECTORS) - 1)

P2_SECTORS := $(shell expr 20 '*' $(MiB) / $(SECT_SZ))
P2_START_SECT := $(shell expr $(P1_END_SECT) + 1)
P2_END_SECT := $(shell expr $(P2_START_SECT) + $(P2_SECTORS) - 1)

P3_SECTORS := $(shell expr 10 '*' $(MiB) / $(SECT_SZ))
P3_START_SECT := $(shell expr $(P2_END_SECT) + 1)
P3_END_SECT := $(shell expr $(P3_START_SECT) + $(P3_SECTORS) - 1)



#P4SZ := 

mbr.img: blank.img
	cp $< $@
	parted -s $@ mklabel msdos
	parted -s $@ mkpart primary $(P1_START_SECT)s $(P1_END_SECT)s
	parted -s $@ mkpart primary $(P2_START_SECT)s $(P2_END_SECT)s
	parted -s $@ mkpart primary $(P3_START_SECT)s $(P3_END_SECT)s

# The following command line creates a sparse file.
blank.img:
	dd if=/dev/zero of=$@ bs=$(IMG_SIZE) seek=1 count=0


.PHONY: clean
clean:
	rm -f mbr.img blank.img

show:
	@echo "P1_SECTORS   :" $(P1_SECTORS)
	@echo "P1_START_SECT:" $(P1_START_SECT)
	@echo "P1_END_SECT  :" $(P1_END_SECT)
