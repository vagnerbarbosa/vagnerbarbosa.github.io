#!/bin/bash
# check_coverage.sh - Validates that the average project statement coverage meets the minimum.

THRESHOLD=90.0
TOTAL_COV=0
COUNT=0

echo "Checking overall project coverage..."
echo "--------------------------------------------------"

# Run tests and capture output (including stderr)
COVERAGE_OUTPUT=$(go test ./... -cover 2>&1)

# Identify actual test failures (lines containing FAIL, ignoring covdata noise)
ACTUAL_FAILURES=$(echo "$COVERAGE_OUTPUT" | grep "FAIL" | grep -v "covdata")

if [ -n "$ACTUAL_FAILURES" ]; then
    echo "❌ Actual tests failed!"
    echo "$COVERAGE_OUTPUT"
    exit 1
fi

# Extract and aggregate coverage percentages
while read -r line; do
    if [[ $line == *"coverage: "* ]]; then
        pkg=$(echo $line | awk '{print $2}')
        cov=$(echo $line | grep -oP 'coverage: \K[0-9.]+')

        echo "Package $pkg: $cov%"
        TOTAL_COV=$(echo "$TOTAL_COV + $cov" | bc -l)
        COUNT=$((COUNT + 1))
    fi
done <<< "$COVERAGE_OUTPUT"

echo "--------------------------------------------------"
if [ $COUNT -eq 0 ]; then
    echo "❌ No coverage data found!"
    exit 1
fi

AVERAGE=$(echo "$TOTAL_COV / $COUNT" | bc -l)
printf "Project Average Coverage: %.2f%%\n" "$AVERAGE"

if (( $(echo "$AVERAGE < $THRESHOLD" | bc -l) )); then
    echo "❌ Average coverage is below $THRESHOLD%"
    exit 1
else
    echo "🎉 Project average reached the minimum $THRESHOLD% coverage!"
    exit 0
fi
