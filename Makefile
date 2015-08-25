DIRS = ../config
INITDIRS = $(DIRS:%=init-%)
BUILDDIRS = $(DIRS:%=build-%)
TESTDIRS = $(DIRS:%=test-%)
PUBLISHDIRS = $(DIRS:%=publish-%)
CLEANDIRS = $(DIRS:%=clean-%)

build: $(BUILDDIRS)
$(BUILDDIRS):
	$(MAKE) -C $(@:build-%=%) build

init: $(INITDIRS)
$(INITDIRS):
	$(MAKE) -C $(@:init-%=%) init

test: $(TESTDIRS)
$(TESTDIRS):
	$(MAKE) -C $(@:test-%=%) test

publish: $(PUBLISHDIRS)
$(PUBLISHDIRS):
	$(MAKE) -C $(@:publish-%=%) publish

clean: $(CLEANDIRS)
$(CLEANDIRS):
	$(MAKE) -C $(@:clean-%=%) clean


.PHONY: subdirs $(DIRS)
.PHONY: subdirs $(INITDIRS)
.PHONY: subdirs $(BUILDDIRS)
.PHONY: subdirs $(TESTDIRS)
.PHONY: subdirs $(CLEANDIRS)
.PHONY: build init clean test
