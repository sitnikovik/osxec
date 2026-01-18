#!/bin/sh

# The script calculates coverage percent and outputs to stdout.

go tool cover -func=tmp/coverage.out > tmp/make_coverage_percent.log
tail -n1 tmp/make_coverage_percent.log | awk '{print $NF}'
