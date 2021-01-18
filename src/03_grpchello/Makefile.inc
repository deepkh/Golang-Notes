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

GRPCHELLO: \
		${PBGRPCHELLO_PHONY} \
		${BINGRPCHELLOCLIENT_PHONY} \
		${BINGRPCHELLOSERVER_PHONY} \
		${GOGRPCHELLOCLIENT_PHONY} \
		${GOGRPCHELLOSERVER_PHONY}
	@echo MAKE GRPCHELLO 

GRPCHELLO_DEV: \
		${PBGRPCHELLO_PHONY_DEV} \
		${BINGRPCHELLOCLIENT_PHONY_DEV} \
		${BINGRPCHELLOSERVER_PHONY_DEV} \
		${GOGRPCHELLOCLIENT_PHONY_DEV} \
		${GOGRPCHELLOSERVER_PHONY_DEV}
	@echo MAKE GRPCHELLO 

GRPCHELLO_CLEAN: \
		${PBGRPCHELLO_PHONY_CLEAN} \
		${BINGRPCHELLOCLIENT_PHONY_CLEAN} \
		${BINGRPCHELLOSERVER_PHONY_CLEAN} \
		${GOGRPCHELLOCLIENT_PHONY_CLEAN} \
		${GOGRPCHELLOSERVER_PHONY_CLEAN}
	@echo MAKE GRPCHELLO_CLEAN 

### grpchello #####################################################################################
BINGRPCHELLOCLIENT: ${BINGRPCHELLOCLIENT_BIN}
	@echo MAKE BINGRPCHELLOCLIENT 

BINGRPCHELLOCLIENT_CLEAN: ${BINGRPCHELLOCLIENT_BIN_CLEAN}
	@echo MAKE BINGRPCHELLOCLIENT_CLEAN 

${BINGRPCHELLOCLIENT_BIN}:
	$(CD) ${BINGRPCHELLOCLIENT} && $(MAKE) -f Makefile.${BINGRPCHELLOCLIENT_NAME} -j8 $@

BINGRPCHELLOCLIENT_DEV: 
	$(CD) ${BINGRPCHELLOCLIENT} && $(MAKE) -f Makefile.${BINGRPCHELLOCLIENT_NAME} -j8 ${BINGRPCHELLOCLIENT_BIN}
	@echo MAKE $@ 

${BINGRPCHELLOCLIENT_BIN_CLEAN}:
	$(CD) ${BINGRPCHELLOCLIENT} && $(MAKE) -f Makefile.${BINGRPCHELLOCLIENT_NAME} $@

### grpchello #####################################################################################
BINGRPCHELLOSERVER: ${BINGRPCHELLOSERVER_BIN}
	@echo MAKE BINGRPCHELLOSERVER 

BINGRPCHELLOSERVER_CLEAN: ${BINGRPCHELLOSERVER_BIN_CLEAN}
	@echo MAKE BINGRPCHELLOSERVER_CLEAN 

${BINGRPCHELLOSERVER_BIN}:
	$(CD) ${BINGRPCHELLOSERVER} && $(MAKE) -f Makefile.${BINGRPCHELLOSERVER_NAME} -j8 $@

BINGRPCHELLOSERVER_DEV: 
	$(CD) ${BINGRPCHELLOSERVER} && $(MAKE) -f Makefile.${BINGRPCHELLOSERVER_NAME} -j8 ${BINGRPCHELLOSERVER_BIN}
	@echo MAKE $@ 

${BINGRPCHELLOSERVER_BIN_CLEAN}:
	$(CD) ${BINGRPCHELLOSERVER} && $(MAKE) -f Makefile.${BINGRPCHELLOSERVER_NAME} $@

