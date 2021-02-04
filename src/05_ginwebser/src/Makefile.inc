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

### tcp_proxy #####################################################################################
GINWEBSER: ${GINWEBSER_BIN}
	@echo MAKE GINWEBSER 

GINWEBSER_CLEAN: ${GINWEBSER_BIN_CLEAN}
	@echo MAKE GINWEBSER_CLEAN 

${GINWEBSER_BIN}:
	$(CD) ${GINWEBSER} && $(MAKE) -f Makefile.${GINWEBSER_NAME} -j8 $@

GINWEBSER_DEV: 
	$(CD) ${GINWEBSER} && $(MAKE) -f Makefile.${GINWEBSER_NAME} -j8 ${GINWEBSER_BIN}
	@echo MAKE $@ 

${GINWEBSER_BIN_CLEAN}:
	$(CD) ${GINWEBSER} && $(MAKE) -f Makefile.${GINWEBSER_NAME} $@


### tcp_proxy #####################################################################################
GINWEBSERTEST: ${GINWEBSERTEST_BIN}
	@echo MAKE GINWEBSERTEST 

GINWEBSERTEST_CLEAN: ${GINWEBSERTEST_BIN_CLEAN}
	@echo MAKE GINWEBSERTEST_CLEAN 

${GINWEBSERTEST_BIN}:
	$(CD) ${GINWEBSERTEST} && $(MAKE) -f Makefile.${GINWEBSERTEST_NAME} -j8 $@

GINWEBSERTEST_DEV: 
	$(CD) ${GINWEBSERTEST} && $(MAKE) -f Makefile.${GINWEBSERTEST_NAME} -j8 ${GINWEBSERTEST_BIN}
	@echo MAKE $@ 

${GINWEBSERTEST_BIN_CLEAN}:
	$(CD) ${GINWEBSERTEST} && $(MAKE) -f Makefile.${GINWEBSERTEST_NAME} $@


