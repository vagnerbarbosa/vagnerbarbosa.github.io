#!/bin/bash
# check_coverage.sh - Validates that all project packages have 100% statement coverage.

THRESHOLD=100.0
FAILED=0

echo "Checking overall project coverage..."
echo "--------------------------------------------------"

# Get coverage for all packages
COVERAGE_OUTPUT=$(go test ./... -cover)

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
    echo "🎉 All packages reached 100% coverage!"
    exit 0
else
    echo "❌ Some packages are below $THRESHOLD% coverage."
    exit 1
fi
