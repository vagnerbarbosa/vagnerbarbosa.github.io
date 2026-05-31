# Testing Strategy

This document describes the testing strategy for the professional portfolio website and its import pipeline.

## Overview

The project employs a multi-layered testing approach to ensure data integrity from LinkedIn import to the final rendered HTML.

## Testing Layers

### 1. Unit Tests
- **Focus**: Small, isolated components (parsers, transformers, config managers).
- **Goal**: Validate that individual functions behave correctly under various inputs.
- **Execution**: `go test ./...`

### 2. Integration Tests (Import Pipeline)
- **Focus**: The flow from CSV files $\rightarrow$ YAML configuration.
- **Method**: **Golden Files**.
  - Reference CSVs are parsed.
  - The resulting YAML is compared against a "golden" reference file.
  - Semantic comparison is used (unmarshaling to structs) to ignore insignificant formatting differences.
- **Goal**: Ensure the import pipeline is deterministic and correct.

### 3. End-to-End (E2E) Tests
- **Focus**: Full pipeline from CSV $\rightarrow$ YAML $\rightarrow$ HTML.
- **Method**:
  - A temporary environment is created with a subset of templates and assets.
  - The import pipeline is triggered using test CSV data.
  - The site generator is run.
  - The resulting `index.html` is scanned for specific "markers" (unique strings) that must be present for each section (Experience, Education, Certifications).
- **Goal**: Guarantee that imported data actually renders on the final website.

## Test Data

Test data for integration and E2E tests is located in:
`cmd/import-linkedin/testdata/e2e/`

## Running Tests

### All Tests
```bash
go test -v ./...
```

### E2E Tests Only
```bash
go test -v ./cmd/generator/...
```

## Quality Gates

- **Code Coverage**: The project aims for high test coverage, particularly in the import pipeline (targeting >95%).
- **CI/CD**: All tests are executed automatically on every Pull Request via GitHub Actions.
