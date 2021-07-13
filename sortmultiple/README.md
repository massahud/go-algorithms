# sortmultiple

This package was created to benchmark possible solutions to the problem of sorting n lists 
without repetition into a single list with all numbers greater than 0.

The main reason for this is because I was thinking with a friend that since a map does allocations, maybe a single sort and a new pass removing repeated values of a big list will be faster than adding the values as keys of a map and then sorting its keys.

Use `make` to run the benchmarks.