# Architecture

## Current milestone

The first milestone delivers a stateless control-plane API with liveness, readiness, Prometheus-compatible metrics, a catalog endpoint, and session creation. State is intentionally in-memory until PostgreSQL and Redis are introduced with migrations and integration tests in Milestone 2.

## Target production architecture

Traffic enters through CloudFront and AWS WAF, reaches an Application Load Balancer, and is served by workloads on Amazon EKS across multiple Availability Zones. PostgreSQL uses Amazon RDS, caching uses ElastiCache, and media objects use Amazon S3. Terraform defines the infrastructure without requiring it to be provisioned during development.

## Reliability principles

- Every service exposes separate liveness and readiness signals.
- Requests use bounded body sizes and server-side timeouts.
- Deployments must support rolling updates and rollback.
- SLIs and SLOs are defined before alert thresholds.
- Failure scenarios produce runbooks and blameless postmortems.

