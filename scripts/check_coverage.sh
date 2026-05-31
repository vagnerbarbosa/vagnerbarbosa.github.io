#!/bin/bash
# check_coverage.sh - Validates that all project packages meet the minimum statement coverage.

THRESHOLD=95.0
FAILED=0

echo "Checking overall project coverage..."
echo "--------------------------------------------------"

# Run tests and capture output + exit code
COVERAGE_OUTPUT=$(go test ./... -cover)
TEST_EXIT_CODE=$?

# If tests failed, we fail immediately
if [ $TEST_EXIT_CODE -ne 0 ]; then
    echo "❌ Tests failed!"
    echo "$COVERAGE_OUTPUT"
    exit 1
fi

# Extract coverage percentages from output
# Example line: ok  github.com/.../pkg 0.001s coverage: 92.7% of statements
while read -r line; do
    if [[ $line == *"coverage: "* ]]; then
        pkg=$(echo $line | awk '{print $2}')
        cov=$(echo $line | grep -oP 'coverage: \K[0-9.]+')

        if (( $(echo "$cov < $THRESHOLD" | bc -l) )); then
            echo "❌ $pkg: $cov% (FAILED)"
            FAILED=1
        else
            echo "✅ $pkg: $cov% (PASS)"
        fi
    fi
done <<< "$COVERAGE_OUTPUT"

echo "--------------------------------------------------"
if [ $FAILED -eq 0 ]; then
    echo "🎉 All packages reached the minimum $THRESHOLD% coverage!"
    exit 0
else
    echo "❌ Some packages are below $THRESHOLD% coverage."
    exit 1
fi
