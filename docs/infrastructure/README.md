# Daun HR — Infrastructure baseline (2026-11-19)

Status: proposed design, documentation only. Product development, provisioning, production deploy and purchases remain on HOLD until owner approval. Source of truth for [web](https://github.com/Hisyam-Dzaki/hris-web), [backend](https://github.com/Hisyam-Dzaki/Hris-BE), [mobile](https://github.com/Hisyam-Dzaki/Hris-Mobile). See [ADR](ADR-001.md) and [architecture image](architecture.svg).

## Deployment tiers
| Tier | Compute | Scope | Planning budget/month (IDR, not quote) |
|---|---|---|---|
| 0 | local / 1 VPS 1-2 vCPU, 2 GB | demo, synthetic data only, prebuilt images, no full monitoring stack | 0–150k |
| 1a | 1 VPS 2-4 vCPU, 4 GB | small pilot only after load/security/restore tests | obtain quote |
| 1b | 1 VPS 4 vCPU/8 GB or 8 vCPU/16 GB | production sizing based on measured peak; 16 GB conservative not mandatory | 450–650k reference for 16 GB incl basic external storage/backup |
| 2 | app + DB VPS, private LAN | scale bottleneck separately, not HA | 800k–1.2m |
| 3 | 3 K3s server/worker nodes + external DB + LB | app HA, DB still SPOF unless independently replicated | 1.3–2.5m |
| 4 | multi-node app and PostgreSQL HA, DR | contractual availability after tested failover | 3–6m+ |

2 GB must not be treated as a guaranteed production capacity. Build images locally/CI, disable optional workers/Redis until required, set limits and reserve OS memory. No sensitive real users until launch checks pass. Do not share trading infrastructure for production. Stage upgrades 2→4→8→16 GB where provider supports resize; check downtime, disk resize, billing and backups first. Use cloud hourly/monthly billing with configurable sizes, not per-request VPS. A persistent minimum node is required. Do not purchase anything yet. Costs are historical planning placeholders, exclude taxes/egress/notifications/fees and must be verified at purchase.

## Storage and disaster recovery
At first PostgreSQL data and uploads can share one VPS persistent host volumes (/srv/daunhr/postgres, /srv/daunhr/uploads). Do not use writable container layers or public upload directories. Backend object-key interface supports local private storage and S3 later (same API, metadata DB, tenant-scoped authorization, quotas, file type/size checks, signed URLs, malware scan as appropriate). Migration: inventory/checksum, copy, reconcile, switch reads/writes, test rollback; do not delete old copy before validation. Move to S3 before multiple app nodes; use object lifecycle policies per customer contract and applicable law, not arbitrary irreversible deletion. Example 1000 employees × 2 photos/day × 22 days × 200 KB = 8.8 GB/month, excludes documents. Off-site = independent failure domain, separate credentials, encrypted in transit/at rest, versioned/immutable where supported. Backup DB, files, config/recovery metadata; PostgreSQL logical backup plus WAL/PITR for hourly RPO if required, consistent DB/file manifest, checksums. Draft targets production RPO <=1h, RTO <=4h; test restore monthly and before migration, alert on failed backups/WAL gap. Local backups alone do not cover loss of VPS. Retention and region subject to PDP, contract and lawful deletion. No public PostgreSQL/Redis endpoints.

## Dynamic scaling / guards
Use Docker Compose first, K3s learning on local disposable VM separately. Resize vertically on measured pressure; separate DB/worker/storage if bottleneck; then migrate stateless app replicas to K3s. HA embedded etcd requires at least 3 server nodes, odd quorum, low-latency private networking, enough N-1 capacity and a redundant LB; app HA is not DB HA. PostgreSQL HA requires standby/failover/fencing, replication backups; test DR. HPA scales pods (not VPS) with resource requests and metrics-server / queue metrics; cap replicas and cooldown; node autoscaler requires supported provider integration, API quota, spend cap and explicit approval. Start node purchasing manually.
Initial thresholds (tune after benchmark): CPU >65% sustained warning, >85% critical; RAM >70/85%, disk >65/80%, DB pool >70/90%, API p95 >300ms/1s and errors >0.5/2%. Alert only on sustained windows (e.g. 10–15min) with exception for OOM/disk/backup failure; backup alert if RPO target violated (not 12h for 1h RPO). Resource limits, pool/max DB connections, tenant rate limits (429 + Retry-After), request size limits, bounded queues, bounded retries with backoff/jitter and dead-letter, idempotency for attendance/payroll, reserve capacity for authentication/attendance; pause non-critical reports/jobs on overload, recover gradually; do not silently drop accepted work. No automatic purchase without budget approval. Metrics: CPU steal, RAM, disk/IOPS, DB locks/connection, queue age, latency/error rate, backup status, tenant sizes; no raw PII in logs.
Load-test synthetic data only: 50/100/500/1000+ staff, simultaneous attendance, large photo uploads, reports/payroll competing, measure p95/errors, memory and disk, verify duplicates and cross-tenant access, test node/DB failures and restore. Scale when sustained guardrail breach after optimization or availability/SLA requires it, not solely revenue or user count. Track MRR, actual cost/tenant, storage growth and headroom.

## Domain
Examples only; domain availability not verified. Start one parent domain, e.g. daunhr.example, app.daunhr.example and api.daunhr.example, tenant slug in workspace URL or tenant.daunhr.example (wildcard DNS/TLS via DNS-01). Later optional verified hr.customer.example CNAME to ingress, TXT proof / ownership, per-host TLS ACME issuance and renewal, hostname allowlist and tenant mapping, remove dangling DNS risk on offboarding. Custom domains optional per company, not one domain/VPS per employee. Use same API with short-lived authenticated sessions, token tenant membership verification and server-side tenant scoping; host/slug is routing hint only. CSRF/CORS/redirect allowlist, isolated cookies, no trust of arbitrary Host headers. Mobile chooses workspace/login; no separate app build per tenant.
 
## Readiness checklist (not implemented by this document)
- [x] Existing stack inventoried: Go/Gin/GORM/PostgreSQL, Dockerfile/compose, frontend Next.js and mobile Flutter scaffolding.
- [ ] Audit tenant authorization and remove frontend demo hardcoded login before real data.
- [ ] Define immutable tagged images, CI build/test, separate dev/staging/prod secrets, versioned reversible-compatible migrations, deploy health/readiness, graceful shutdown and rollback runbook.
- [ ] Storage interface local/S3, quotas, file access checks, retention/backup/restore.
- [ ] Bounded workers, idempotency, priority queues, resource limits and observability.
- [ ] Domain strategy, TLS, private DB network, secret rotation, least privilege, audit logs, data lifecycle and PDP/PSE applicability review.
- [ ] Synthetic benchmark, provider resize/IOPS/billing verification, capacity and financial gates.
- [ ] K3s sandbox, manifests/Helm only when migration justified; test N-1, rolling update, node/DB failure, backup restore.
No code/config changes authorized in this documentation PR. Review before merge. References: https://docs.k3s.io/installation/ha-embedded/ ; https://kubernetes.io/docs/concepts/workloads/autoscaling/horizontal-pod-autoscale/ ; https://docs.k3s.io/installation/requirements .
