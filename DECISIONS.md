# DECISIONS.md

# TaskForge Decision Log

## Overview

This document records the key architectural and implementation decisions made while building TaskForge, along with alternatives considered and trade-offs accepted.

---

# 1. Technology Choice

## Decision

Used **Go** as the implementation language.

## Why

TaskForge is fundamentally a concurrent workflow orchestration engine.

The language requirements include:
- Efficient concurrency
- Context propagation
- Cancellation support
- Simple deployment model
- Strong standard library
- Low operational overhead

Go's goroutines, channels, and context package make it particularly well suited for workflow execution engines.

---

# 2. Workflow Definition Format

## Decision

Used JSON workflow definitions.

## Why

- Native HTTP API compatibility
- Easy persistence using JSON
- Human-readable
- Easy integration with automation systems

---

# 3. DAG Validation Strategy

## Decision

Validate workflows during registration.

Validation includes:
- Cycle detection
- Unknown task types
- Missing dependencies
- Duplicate task identifiers

---

# 4. Execution Model

## Decision

Used a DAG scheduler based on dependency counting.

Provides O(V + E) scheduling complexity and efficient parallel execution.

---

# 5. Shared State Design

## Decision

Maintain execution state as a shared JSON document.

Allows flexible task outputs and heterogeneous task types.

---

# 6. Variable Resolution

Use:

{{tasks.build.imageId}}
{{tasks.http.response.id}}

---

# 7. Task Extensibility Mechanism

Used a registry-based plugin architecture.

New task types can be added without modifying engine code.

---

# 8. Retry Policy

Retryable:
- Network failures
- HTTP 5xx
- Rate limiting (429)

Non-retryable:
- Validation errors
- HTTP 4xx
- Invalid configuration

---

# 9. Timeout Enforcement

Used context.WithTimeout and CommandContext.

---

# 10. Failure Propagation

Permanent task failures mark dependent tasks as SKIPPED.

---

# 11. Conditional Execution

Introduce CONDITION_FALSE as a distinct state.

---

# 12. Approval Tasks

States:
- WAITING_APPROVAL
- APPROVED
- REJECTED
- TIMED_OUT

---

# 13. Cancellation Strategy

Two-phase cancellation:
1. Stop scheduling new tasks.
2. Graceful shutdown before force termination.

---

# 14. Persistence Strategy

Used PostgreSQL.

---

# 15. Additional Task Types

- Approval Task
- Kubernetes Task

---

# 16. Observability

Track execution status and task status.

---

# 17. Estimated Development Time

Estimated effort: 12–18 hours

Actual effort: 20 hours
