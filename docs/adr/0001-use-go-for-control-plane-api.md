# ADR 0001: Use Go for the control-plane API

- Status: Accepted
- Date: 2026-09-07

## Context

The platform needs predictable resource usage, fast startup, straightforward concurrency, and a small deployable artifact.

## Decision

Use Go for the first API service and prefer the standard library until an external dependency provides measurable value.

## Consequences

The service compiles to a static binary, runs in a minimal non-root container, and has a small initial dependency surface. Engineers must keep explicit HTTP timeouts and graceful shutdown behavior.

