DIRS = ../config
INITDIRS = $(DIRS:%=init-%)
BUILDDIRS = $(DIRS:%=build-%)
CLEANDIRS = $(DIRS:%=clean-%)
TESTDIRS = $(DIRS:%=test-%)

build: $(BUILDDIRS)
$(BUILDDIRS):
	$(MAKE) -C $(@:build-%=%) build

init: $(INITDIRS)
$(INITDIRS):
	$(MAKE) -C $(@:init-%=%) init

test: $(TESTDIRS)
$(TESTDIRS):
	$(MAKE) -C $(@:test-%=%) test

clean: $(CLEANDIRS)
$(CLEANDIRS):
	$(MAKE) -C $(@:clean-%=%) clean


.PHONY: subdirs $(DIRS)
.PHONY: subdirs $(INITDIRS)
.PHONY: subdirs $(BUILDDIRS)
.PHONY: subdirs $(TESTDIRS)
.PHONY: subdirs $(CLEANDIRS)
.PHONY: build init clean test
