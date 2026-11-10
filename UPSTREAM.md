# Upstream synchronization policy

- Primary protocol upstream: https://github.com/tulir/whatsmeow, resolve default branch via GitHub.
- Public feature upstream: https://github.com/polymorfa/hypermeow, resolve its default branch separately.
- Destination: https://github.com/wppconnect-team/hypermeow, main.
- Daily schedule: 09:30 America/Sao_Paulo through the requested Codex task.

Use a fresh clone or clean worktree. Fetch both upstreams and origin; record exact
SHAs, merge base, ahead/behind and patch-equivalent commits. Preserve local work.
Create a candidate branch for integration. Merge normally, never reset main or
force-push. Resolve conflicts by preserving intended behavior on both sides;
renaming upstream imports must remain consistent with the inherited module path.
Generated protobuf/MEX files require their corresponding sources/generator checks.
Do not lower SQL schema versions or discard fork migrations, namespaces, locks,
performance changes, public API, attribution or protocol corrections.

Run the CI-equivalent tests, vet, build, MEX generation check, JS/WASM build and
pre-commit checks. Test SQL migrations when the store changes, and dependent
server-go builds when public APIs change. Offline passing tests do not waive
account-dependent validation of protocol changes. Never reuse production data
for tests or send messages to unrelated recipients.

Push a candidate and open or update a single upstream-sync PR when appropriate.
Only merge/push main after all applicable checks pass and required branch rules
allow it. Record CI failures or unavailable runners; absence of CI is not success.
Re-fetch main immediately before merge, incorporate intervening changes and rerun
affected checks. Verify the destination SHA after update. No npm publication or
semver tag is implied by this task; Go consumers pin the verified commit.

On conflicts, keep the candidate and a sanitized report. Do not repeatedly open
identical PRs. Reuse the pending candidate where safe. Report source/destination
SHAs, relevant changes, validation, PR/merge link and remaining blockers. If neither
upstream changed and no pending integration needs work, produce no empty commit.

Initial comparison (2026-11-10): Polymorfa f432198 vs tulir c88388f has 37 tulir-only
and 293 fork-only commits; merge-tree reports conflicts in 15 files. This is not
an already synchronized base. Resolve and validate those conflicts before claiming
that the fork includes the current tulir main.
