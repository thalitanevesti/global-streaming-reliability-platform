# AGENTS.md

## Mission

Build a production-grade streaming reliability platform demonstrating
Platform Engineering, SRE, DevSecOps, observability and Infrastructure as Code.

## Engineering principles

- Prefer simple, modular and maintainable solutions.
- Avoid overengineering and unjustified infrastructure.
- Apply security by design, least privilege and defense in depth.
- Optimize reliability, performance, cost and resource consumption.
- Preserve existing user changes and repository conventions.
- Do not introduce dependencies without a clear benefit.

## Workflow

Before changing code:

1. Read the relevant files and documentation.
2. Check the current Git branch and working tree.
3. Identify requirements, risks and acceptance criteria.
4. Update `PLAN.md` for substantial work.
5. Never work directly on `main`.

During implementation:

- Make small, cohesive and reviewable changes.
- Validate inputs and handle errors explicitly.
- Use timeouts for external operations.
- Never log secrets, tokens, passwords or sensitive data.
- Never commit `.env`, credentials, private keys or generated artifacts.
- Keep configurations reproducible and version controlled.
- Do not deploy paid cloud resources without explicit authorization.

## Containers and infrastructure

- Use multi-stage container builds and minimal trusted images.
- Run containers as non-root.
- Drop unnecessary Linux capabilities.
- Use read-only filesystems when practical.
- Pin important versions and scan images and dependencies.
- Use Infrastructure as Code.
- Apply private networking and deny-by-default policies.
- Follow least privilege for cloud IAM and CI permissions.

## Observability and reliability

- Use structured logs without sensitive data.
- Provide appropriate health, readiness and liveness checks.
- Collect relevant metrics and distributed traces.
- Define SLIs, SLOs and actionable alerts when applicable.
- Document rollback, recovery and operational procedures.

## Validation

Before concluding:

1. Run formatting and lint checks.
2. Run applicable tests, including race detection.
3. Validate Docker Compose, Dockerfile and IaC.
4. Review security and dependency risks.
5. Inspect the complete Git diff.
6. Correct discovered problems.

Never claim a check passed unless it was executed.
Report any check that could not be performed.

## Git and delivery

- Use branches and pull requests.
- Follow Conventional Commits.
- Never force-push `main`.
- Keep commits focused.
- Update documentation when behavior changes.
- Summarize changed files, validations and remaining risks.
