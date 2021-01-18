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
#!/bin/bash

if [ ! -z "$1" ]; then
	unset PBGRPCHELLO_PHONY
	unset PBGRPCHELLO_PHONY_CLEAN
	if [ "${HAVE_PBGRPCHELLO}" = "1" ]; then
		export PBGRPCHELLO_NAME="pbgrpchello"
		export PBGRPCHELLO="$1"
		export PBGRPCHELLO_OBJS_DIR=
		export PBGRPCHELLO_CCGO="cc/grpchello.pb.cc \
															cc/grpchello.grpc.pb.cc \
															grpchello_grpc.pb.go"
		export PBGRPCHELLO_CCGO_CLEAN="cc/grpchello.pb.cc_clean \
															cc/grpchello.grpc.pb.cc_clean
															grpchello_grpc.pb.go_clean"
		export PBGRPCHELLO_PHONY="PBGRPCHELLO"
		export PBGRPCHELLO_PHONY_DEV="PBGRPCHELLO_DEV"
		export PBGRPCHELLO_PHONY_CLEAN="PBGRPCHELLO_CLEAN"
		export PBGRPCHELLO_CFLAGS="-I${PBGRPCHELLO}/cc"
		export PBGRPCHELLO_LDFLAGS=""
		echo "PBGRPCHELLO=${PBGRPCHELLO}"
	fi
fi

