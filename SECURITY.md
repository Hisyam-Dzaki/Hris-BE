# Security posture and launch gates
Current scaffold is DEVELOPMENT ONLY. No verified production readiness. Findings from static review, not runtime exploit tests.
| Priority | Finding | Required before production |
|---|---|---|
| P0 | Known default JWT signing secret and database password | Separate environment secrets, fail closed, rotate if used outside disposable dev |
| P0 | Employee and company ID routes lack demonstrated tenant scoping; company list unscoped | Server-side authenticated tenant and role checks on all operations, relation ownership, negative regression tests |
| P0 | Public registration accepts company ID; role middleware not wired to CRUD | Reviewed invitation/tenant enrollment and deny-by-default role enforcement |
| P0 | DB port published on host | Private DB network; firewall and restricted dev binding; do not assume actual external exposure |
| P1 | Broad CORS and client mock login | Allowed origins, secure sessions, CSRF and actual backend auth |
| P1 | Startup AutoMigrate and verbose SQL logging | Versioned migrations and redaction |
| P1 | Health endpoint insufficient for readiness | DB readiness, monitoring, resource guards |
Development configuration may remain unchanged for now by owner decision; never use real employee data. Do not run destructive tests against shared dev databases. Future production requires separate environment credentials and DB, backups/restore drills, audited tenant isolation, retention and Indonesian PDP/PSE applicability review. Report vulnerabilities privately to repository owner, not public issues with sensitive details. No code fixes authorized by this documentation PR.
