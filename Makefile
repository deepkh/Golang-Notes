# Copyright (c) 2018, Gary Huang, deepkh@gmail.com, https://github.com/deepkh
#
# Permission to use, copy, modify, and distribute this software for any
# purpose with or without fee is hereby granted, provided that the above
# copyright notice and this permission notice appear in all copies.
#
# THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
# WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
# MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
# ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
# WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
# ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
# OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

SHELL=/bin/sh

.DEFAULT_GOAL := all

include $(GOTUTORIAL)/${MAKEFILE_DEP}

define create_test_file
	printf HI; \
	head -c $(2) </dev/urandom > $(1); \
	md5sum $(1) >> $(3).md5
endef

all: $(GOTUTORIAL_PHONY) test_file/10M

test_file:
	mkdir $@

test_file/10M: test_file
	$(call create_test_file,test_file/10M,10000000,test_file/md5sum)
	$(call create_test_file,test_file/20M,20000000,test_file/md5sum)
	$(call create_test_file,test_file/30M,30000000,test_file/md5sum)

clean: $(GOTUTORIAL_PHONY_CLEAN) 

