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

if [ -z "${HAVE_GOCOMPILER}" ];then
	export HAVE_GOCOMPILER=1
fi

#GPRC already included protobuf 3.13.0
#if [ -z "${HAVE_PROTOBUF}" ];then
#	export HAVE_PROTOBUF=1
#fi

#if [ -z "${HAVE_LIB_PROTOBUF}" ];then
#	export HAVE_LIB_PROTOBUF=1
#fi

if [ -z "${HAVE_GRPC}" ];then
	export HAVE_GRPC=1
fi

if [ -z "${HAVE_PROTOCGENGO}" ];then
	export HAVE_PROTOCGENGO=1
fi
