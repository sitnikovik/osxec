#!/usr/bin/env bash

PERCENT_FILE=${1:-tmp/coverage_total.out}
THRESHOLD=${COVERAGE_THRESHOLD:-90}

# Run coverage generation first even though this is a check script.
if ! make coverage >/dev/null 2>&1; then
    echo "make coverage failed" >&2
    exit 1
fi
if [ ! -f "$PERCENT_FILE" ]; then
    echo "Coverage percent file missing: $PERCENT_FILE" >&2
    exit 1
fi

# validate numeric (allow formats: 1, 1.23, .5, 0.5).
VAL=$(tr -d '\r\n' < "$PERCENT_FILE")
VAL=${VAL%\%}
if ! printf '%s' "$VAL" | awk '/^([0-9]*\.[0-9]+|[0-9]+)$/ { exit 0 } { exit 1 }'; then
    echo "Invalid coverage value: $VAL" >&2
    exit 1
fi

# Check against threshold.
awk -v p="$VAL" -v t="$THRESHOLD" '
BEGIN {
    p0 = p + 0;
    t0 = t + 0;
    if (p0 < t0) {
        printf("❌ Coverage %s%% is below %s%%\n", p, t);
        exit 1
    } else {
        printf("✅ Coverage %s%%\n", p);
    }
}
'
