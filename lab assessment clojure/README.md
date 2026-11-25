# Clojure Practical Assignment

Simple Clojure application covering Sessions 10–11.

## Requirements Covered
1. **Sequence Functions**: `filter`, `map`, and `reduce` in `process-pipeline` (filters CGPA >= 8.5 for Distinction, computes average CGPA).
2. **Ring HTTP Endpoint**: Endpoint `/results` exposed using Ring and Jetty.
3. **Relational Database**: In-memory H2 database with `try/catch` error handling for student records with `cgpa`.
4. **Java Interop**: `(java.time.Instant/now)` for timestamps.

## Run
```bash
lein run
```
