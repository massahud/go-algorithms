benchtime=1s
.PHONY: help bench-sort bench-biglist test

help: ## Shows this help.
# Target copied from https://gist.github.com/prwhite/8168133
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
##
bench-sort: ## Runs benchmarks of the sorting methods, saving sort-memory.out and sort-cpu.out pprof files
##               - with default benchtime: 1s:
##                 make bench-sort
##               - with specific benchtime
##                 make bench-sort benchtime=2s
##               - with number of iteration per cpu
##                 make bench-sort benchtime=10x
bench-sort:
	go test -bench='Sort' -benchtime=$(benchtime) -benchmem -memprofile sort-memory.out -cpuprofile sort-cpu.out

##
bench-biglist: ## Runs benchmarks of the biglist creation methods, saving biglist-memory.out and biglist-cpu.out pprof files
##                  - with default benchtime: 1s:
##                    make bench-biglist
##                  - with specific benchtime
##                    make bench-biglist benchtime=2s
##                  - with number of iteration per cpu
##                    make bench-biglist benchtime=10x	
bench-biglist:
	go test -bench='BigList' -benchtime=$(benchtime) -benchmem -memprofile sort-memory.out -cpuprofile sort-cpu.out
##

test: ## Runs tests
	go test -v
