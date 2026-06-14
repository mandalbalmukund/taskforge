# TESTDESIGN.md

# TaskForge Test Design

## Overview

This document describes the testing strategy for TaskForge.

The primary goals are:

* Verify correctness of workflow execution
* Verify DAG scheduling behavior
* Verify retry and timeout handling
* Verify cancellation behavior
* Verify approval workflows
* Verify shared state propagation
* Verify extensibility requirements
* Verify API contracts

Testing is organized into:

1. Unit Tests
2. Integration Tests
3. End-to-End Tests
4. Failure Injection Tests
5. Concurrency Tests

---

# Test Pyramid

```text
                 E2E Tests
                     ▲
                     │
            Integration Tests
                     ▲
                     │
               Unit Tests
```

Most tests are unit tests.

Critical workflow scenarios are covered through integration and end-to-end tests.

---

# 1. Workflow Validation Tests

## Goal

Verify invalid workflows are rejected before execution.

### Test Cases

### TV-001 Valid DAG

Input:

```text
A -> B -> C
```

Expected:

* Validation succeeds

---

### TV-002 Cycle Detection

Input:

```text
A -> B -> C -> A
```

Expected:

* Validation fails
* Error identifies cycle

---

### TV-003 Missing Dependency

Input:

```text
Task B depends on Task X
```

Expected:

* Validation fails
* Error identifies missing task

---

### TV-004 Unknown Task Type

Input:

```json
{
  "type": "customUnknownTask"
}
```

Expected:

* Validation fails

---

### TV-005 Duplicate Task IDs

Expected:

* Validation fails

---

# 2. DAG Scheduler Tests

## Goal

Verify dependency handling and execution ordering.

---

### TS-001 Linear Workflow

```text
A -> B -> C
```

Expected:

Execution order:

```text
A
B
C
```

---

### TS-002 Parallel Workflow

```text
      A
     / \
    B   C
     \ /
      D
```

Expected:

* B and C run concurrently
* D starts after both complete

---

### TS-003 Multiple Roots

```text
A     B
 \   /
   C
```

Expected:

* A and B start immediately
* C waits

---

### TS-004 Large DAG

100+ tasks

Expected:

* All tasks complete
* No deadlocks

---

# 3. Shared State Tests

## Goal

Verify output propagation.

---

### SS-001 Single Dependency

Task A Output:

```json
{
  "userId": 100
}
```

Task B:

```text
{{tasks.A.userId}}
```

Expected:

Task B receives:

```text
100
```

---

### SS-002 Parallel Producers

Two tasks update shared state simultaneously.

Expected:

* No data corruption
* Both outputs preserved

---

### SS-003 Missing Reference

Reference:

```text
{{tasks.unknown.id}}
```

Expected:

* Validation failure
  or
* Runtime configuration error

---

# 4. Conditional Execution Tests

## Goal

Verify task conditions.

---

### TC-001 Condition True

Expected:

Task executes

Status:

```text
SUCCESS
```

---

### TC-002 Condition False

Expected:

Task skipped

Status:

```text
CONDITION_FALSE
```

---

### TC-003 Downstream Dependency

Expected:

System correctly handles conditionally skipped tasks.

---

# 5. Retry Tests

## Goal

Verify retry behavior.

---

### TR-001 Retryable Failure

Task fails twice.

Task succeeds on third attempt.

Expected:

```text
Attempt 1 -> Fail
Attempt 2 -> Fail
Attempt 3 -> Success
```

Execution succeeds.

---

### TR-002 Retry Exhaustion

Task always fails.

Expected:

```text
FAILED
```

after max retries reached.

---

### TR-003 Non-Retryable Failure

Example:

HTTP 400

Expected:

No retry attempted.

---

# 6. Timeout Tests

## Goal

Verify timeout enforcement.

---

### TT-001 Script Timeout

Task timeout:

```text
5 seconds
```

Script duration:

```text
30 seconds
```

Expected:

```text
TIMED_OUT
```

Process terminated.

---

### TT-002 HTTP Timeout

Remote endpoint never responds.

Expected:

```text
TIMED_OUT
```

---

# 7. Cancellation Tests

## Goal

Verify execution cancellation.

---

### TCN-001 Cancel Before Start

Expected:

Task never begins.

---

### TCN-002 Cancel Running Workflow

Expected:

* Running tasks receive cancellation
* New tasks not scheduled

Execution status:

```text
CANCELLED
```

---

### TCN-003 Graceful Shutdown

Task handles cancellation signal.

Expected:

Clean termination.

---

# 8. Approval Tests

## Goal

Verify gated workflow execution.

---

### TA-001 Approval Granted

Expected:

Workflow resumes.

---

### TA-002 Approval Rejected

Expected:

Workflow fails.

---

### TA-003 Approval Timeout

Expected:

Configured timeout action triggered.

---

### TA-004 Invalid Approval Token

Expected:

Approval rejected.

---

# 9. HTTP Task Tests

## Goal

Verify HTTP task behavior.

---

### TH-001 HTTP 200

Expected:

SUCCESS

Response stored.

---

### TH-002 HTTP 500

Expected:

Retry triggered.

---

### TH-003 HTTP 400

Expected:

Permanent failure.

---

### TH-004 Network Failure

Expected:

Retry triggered.

---

# 10. Script Task Tests

## Goal

Verify script execution.

---

### TSH-001 Exit Code 0

Expected:

SUCCESS

---

### TSH-002 Exit Code Non-Zero

Expected:

FAILED

---

### TSH-003 STDERR With Exit Code 0

Expected:

SUCCESS

stderr captured.

---

### TSH-004 Script Timeout

Expected:

TIMED_OUT

---

# 11. Kubernetes Task Tests

## Goal

Verify cluster operations.

---

### TK-001 Successful Rollout

Expected:

SUCCESS

---

### TK-002 Deployment Not Found

Expected:

FAILED

---

### TK-003 Cluster Unreachable

Expected:

Retry attempted.

---

# 12. Approval Task Tests

## Goal

Verify approval lifecycle.

---

### TAP-001 Approved

Expected:

APPROVED

---

### TAP-002 Rejected

Expected:

REJECTED

---

### TAP-003 Timed Out

Expected:

TIMED_OUT

---

# 13. Concurrency Tests

## Goal

Verify thread safety.

---

### CON-001 Concurrent State Updates

Expected:

No race conditions.

---

### CON-002 Concurrent Task Completion

Expected:

Dependency counters updated correctly.

---

### CON-003 High Parallelism

100 concurrent tasks.

Expected:

No deadlocks.

No data corruption.

---

# 14. API Tests

## Goal

Verify API contracts.

---

### API-001 Register Workflow

Expected:

201 Created

---

### API-002 Retrieve Workflow

Expected:

200 OK

---

### API-003 Start Execution

Expected:

Execution created.

---

### API-004 Get Execution Status

Expected:

Current task state returned.

---

### API-005 Cancel Execution

Expected:

Execution transitions to CANCELLED.

---

### API-006 Resolve Approval

Expected:

Approval processed.

---

# 15. Extensibility Tests

## Goal

Verify core assignment requirement.

---

### EXT-001 Register New Task Type

Create:

```go
type SlackTask struct {}
```

Register:

```go
registry.Register(&SlackTask{})
```

Expected:

Workflow executes successfully.

No engine source files modified.

---

### EXT-002 Unknown Task Handler

Expected:

Validation failure.

---

# Performance Tests

### PERF-001 Large Workflow

Workflow:

1000 tasks

Expected:

Execution completes.

---

### PERF-002 Parallel Workflow

100 independent tasks.

Expected:

Efficient parallel execution.

---

# Success Criteria

TaskForge is considered production-ready when:

* All validation tests pass
* All scheduler tests pass
* All retry tests pass
* All timeout tests pass
* All approval tests pass
* All cancellation tests pass
* No race conditions detected
* Extensibility tests pass without engine modifications

The extensibility test is considered the most important acceptance criterion because it validates the primary architectural requirement of the given problem statement.