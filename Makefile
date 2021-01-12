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

all: $(GOTUTORIAL_PHONY) $(RUNTIME_BIN)/01_intro$(BINSUFFIX) \
	$(RUNTIME_BIN)/02_protobuffer$(BINSUFFIX) \
	src/02_protobuffer/addressbook/addressbook.pb.cc \
	$(RUNTIME_BIN)/02_protobuffer_cc$(BINSUFFIX)


### 01
$(RUNTIME_BIN)/01_intro$(BINSUFFIX):
	@echo MAKE $@
	cd src/01_intro/hello/ && $(GOBUILD) -o $@ hello.go

### 02_protobuffer for golang
src/02_protobuffer/addressbook/addressbook.pb.go: src/02_protobuffer/addressbook/addressbook.proto
	cd src/02_protobuffer/addressbook/ && protoc -I=. --go_out=. addressbook.proto

$(RUNTIME_BIN)/02_protobuffer$(BINSUFFIX): \
	src/02_protobuffer/addressbook/addressbook.pb.go \
	src/02_protobuffer/main/main.go
	@echo MAKE $@
	#delete addressbook.pb.cc due to "imports addressbook: C++ source files not allowed when not using cgo or SWIG: addressbook.pb.cc"
	$(RM) src/02_protobuffer/addressbook/addressbook.pb.cc
	cd src/02_protobuffer/main/ && $(GOBUILD) -o $@ main.go

### 02_protobuffer for C++
CXXFLAGS=-I${RUNTIME}/include -Isrc/02_protobuffer
src/02_protobuffer/addressbook/%.o: %.cc
	$(CXX) $(CXXFLAGS) $(CFLAGS) -o $@ -c $<

src/02_protobuffer/main/%.o: %.cc
	$(CXX) $(CXXFLAGS) $(CFLAGS) -o $@ -c $<

src/02_protobuffer/addressbook/addressbook.pb.cc: src/02_protobuffer/addressbook/addressbook.proto
	cd src/02_protobuffer/addressbook/ && protoc -I=. --cpp_out=. addressbook.proto

LDFLAGS=-L${RUNTIME}/lib

$(RUNTIME_BIN)/02_protobuffer_cc$(BINSUFFIX): \
	src/02_protobuffer/addressbook/addressbook.pb.o \
	src/02_protobuffer/main/main.o
	@echo MAKE $@
	$(CXX) -o $@ $(filter %.o,$^) $(LDFLAGS) $(LIBPROTOBUF_LDFLAGS)

test: test_greetings

test_greetings:
	cd src/greetings && $(GOTESTV) ./...

clean: $(GOTUTORIAL_PHONY_CLEAN)
	$(RM) $(RUNTIME_BIN)/01_*$(BINSUFFIX)
	$(RM) $(RUNTIME_BIN)/02_*$(BINSUFFIX)
	$(RM) src/02_protobuffer/addressbook/addressbook.pb.o
	$(RM) src/02_protobuffer/main/main.o

