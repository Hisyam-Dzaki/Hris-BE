# Daun HR product and infrastructure roadmap
Phase 0 — governance, environment distinction, documentation, owner review of infrastructure PR. Existing web development auto-deploy stays unchanged. 
Phase 1 — security hardening after approval: tenant isolation, RBAC, registration, secrets, auth and regression tests.
Phase 2 — infrastructure foundation: minimal Compose, private storage abstraction LOCAL/S3, off-site backup, health/readiness, resource limits, CI validation, rollback. 
Phase 3 — backend/web/mobile core HR integration with API contract.
Phase 4 — workforce MVP: attendance, shift, leave, overtime, approvals, field placement.
Phase 5 — tenant onboarding, quotas, feature flags, subscription and domain strategy.
Phase 6 — production gates, synthetic load test, recovery/security tests, pilot.
Phase 7 — scale only when measurements and business justify it: resize, independent DB, K3s cluster and DB HA; payroll and other modules based on customer demand.
Infrastructure roadmap is [here](docs/infrastructure/README.md); [diagram](docs/infrastructure/architecture.svg). All phases after 0 are proposed, not approval to implement.
