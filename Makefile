mkfile := $(abspath $(lastword $(MAKEFILE_LIST)))
dir := $(dir $(ci_mkfile))

