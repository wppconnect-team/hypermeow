# WPPConnect HyperMeow

Public WPPConnect-maintained fork of [Polymorfa HyperMeow](https://github.com/polymorfa/hypermeow),
which derives from [tulir/whatsmeow](https://github.com/tulir/whatsmeow).
This repository is a Go WhatsApp protocol SDK, not the ZuckZapGo application server.

## Provenance and licensing

Initial public source: Polymorfa `f43219816d8600ec963371881a1467d22fa7b3e6`.
The complete reachable history is preserved. Read [LICENSING.md](LICENSING.md),
[LICENSE](LICENSE) and [LICENSE-MIT](LICENSE-MIT) for file-specific licensing.
Existing copyright notices and third-party attributions remain intact.

The inspected ZuckZapGo v2.4.2 executable uses a local `hypermeow-private` replacement.
Its container contains compiled application code and generated documentation, not
that private SDK's compilable source. This public fork does not claim to contain
those private changes or all 303 application API operations.

## Using this fork

The inherited Go module path remains `github.com/polymorfa/hypermeow` so that
internal imports, generated protobuf descriptors and upstream merges remain coherent.
Pin a verified WPPConnect commit using a Go replace directive:

```sh
go mod edit -replace=github.com/polymorfa/hypermeow=github.com/wppconnect-team/hypermeow@<verified-commit>
go mod tidy
```

Applications must consistently import the HyperMeow module. Replacing
`go.mau.fi/whatsmeow` alone does not rewrite its imports or make the APIs identical.
Before migrating WPPConnect Server Go, validate API changes, device stores,
SQL migrations, protocol compatibility, reconnect and backup/restore using a copy
of an authorized test session. Do not point two clients at the same session data.

## Receiving statuses

Personal and group `<status>` publications use encrypted message processing;
channel statuses keep their plaintext processing path. The original transport tag
is preserved for acknowledgments. `MessageInfo.StatusKind()` distinguishes
`personal`, `group` and `channel` publications from ordinary messages, status
mentions and replies. Group-status-v2 content is unwrapped while `RawMessage`
retains the original protobuf.

Pairwise status sessions are stored separately from chat sessions. Explicit
`meta session_scope="status"` is honored by default. Applications that consume
served AB props can install `Client.GetStatusSessionMode` before connecting to
supply `status_infra_1_1_session_split` (0 disabled, 1 receiver enabled, 2 sender
enabled, 3 forced). This receive change does not enable status publishing.

Custom `RemoteSignal` receivers must use `RemoteReceive.StatusSession` to select
the pairwise session and carry `StatusKind` into their application event. Sender
keys remain scoped to the group and sender, and private keys and plaintext stay
in the remote store. Upgrading this library alone does not upgrade a remote
receiver's protocol or storage. Legacy encrypted group publications also need
`RemoteReceiveResult.StatusKind` returned after decrypt and on duplicate delivery,
so the client can select the status receipt class without seeing plaintext.


## Development and checks

See [upstream documentation](README.upstream.md) for SDK features and usage.
The inherited CI supports Go 1.26 and 1.27, checks generated MEX bindings and builds
JS/WASM. The go.mod toolchain preference can be overridden with a supported,
verified Go toolchain on the developer machine.

```sh
go test ./...
go vet ./...
go build ./...
go run ./internal/cmd/genmex -input mex/spec.json -output mex/bindings.go -check
GOOS=js GOARCH=wasm go build ./...
```

Offline tests do not establish successful pairing, WhatsApp message delivery,
private-extension parity or safe migration of a production database.

## Upstream maintenance

Daily maintenance compares against `tulir/whatsmeow` and also records the public
Polymorfa source. [UPSTREAM.md](UPSTREAM.md) defines the merge and validation policy.
No force-push or replacement of WPPConnect/HyperMeow-specific code is permitted.
A conflicting or failing update remains unmerged and is reported with evidence.

The phased Cloud and parity work is tracked in [PLAN.md](PLAN.md).

See the [303-operation contract matrix](docs/parity-contract-matrix.md) for the fixed Zuck/Go/Node baseline, candidate adapters, limitations and acceptance criteria. A matching endpoint is not proof of functional parity.

Custom decrypted-event buffers should implement `store.EventBufferPublicationStore` to retain the group-publication flag while atomically deleting plaintext. The SQL store implements this with migration 21. Old processed rows and custom stores without this extension have unknown legacy classification; no status is inferred from a group chat alone.

## Recovering an interrupted identity reset

The local adapter deletes both chat and status session rows before removing their
shared identity. A failed session deletion therefore leaves the old identity in
place. Use `Device.DeleteIdentityAndSessions` when removing both; custom callers
must preserve that order and serialize other writers to the same backend.

Failed rows stay unavailable on the current Device until deletion succeeds.
Guarded session access retries deletion of its known pending key before using
any stored row. `Device.ReconcileSessionResets(ctx)` can also retry all tracked
failures after storage recovers. A reset that cannot retain all possible failures is refused before it
starts; it does not permanently disable the Device or discard unknown failures.
Admission retries at most two existing failures, while explicit reconciliation
can retry the full bounded set. Both reset history and tracked failures retain at
most 1,024 address entries. Pruning history invalidates older prefetched caches,
which must be loaded again.

These guards are per Device and are not cross-process fencing. A fresh adapter
retains the original identity for surviving rows after a failed library-managed
reset; this relies on backend deletion completion and ordering. Custom-store,
transaction, process-crash and previous-release state acceptance still need their
own evidence. Restarting is not a substitute for repairing a failed backend.

Pending recovery runs before buffered decryption opens its transaction. A previously failed
deletion is not reconciled by a write that can still roll back. SQLStore marks its
transaction callbacks and also detects transactions opened directly on its
underlying database. Custom transactional SessionStores must expose
`SessionTransactionState` or mark their contexts with `WithSessionTransaction`;
callbacks must forward that context. Saved buffered plaintext bypasses new
preparation and decryption. These are source contracts, not database/crash or
previous-release acceptance evidence.

SQL identity cache and transactions: use `SQLStore.DoDecryptionTxn` for identity
reads or writes inside a database transaction. It owns the commit/rollback boundary,
serializes identity-cache access for that device until completion, clears the
transaction's cache entries afterward and invalidates earlier identity-fetch tokens.
Nested calls with its callback context share that boundary. An externally opened
transaction cannot supply its eventual outcome, so identity operations and nested
`DoDecryptionTxn` calls in that context return `ErrUnmanagedIdentityTransaction`
before accessing the cache or running the callback. Phone-number to LID migration
also owns its transaction boundary and refuses to run inside an outer transaction.
Escaped callback contexts are
also rejected. This does not coordinate other SQLStore instances or processes that
write the same device rows; those require separate store ownership and admission.

The SQL rollback/replay regression uses a local SQLite database and authored opaque
session bytes, not real account traffic or decrypted Signal sessions. Run it with
CGO enabled: `CGO_ENABLED=1 go test ./store/sqlstore -run TestSQLIdentityRollbackReplay`.
The SQLite driver is a test-only import; the library still does not register a driver
for applications. Cache serialization and guest-session throughput need acceptance
on the composed native candidate before merge.

## Call-link guests

The `guest` package registers a one-time guest identity for a WhatsApp call
link and logs it in without pairing a phone. The identity lives only in
memory and only for one call:

```go
import "github.com/polymorfa/hypermeow/guest"

device := guest.NewDevice(nil) // fresh keys, in-memory stores
reg, err := guest.Register(ctx, device, guest.Request{
    InviteCode: "https://call.whatsapp.com/voice/TOKEN",
    Name:       "Ada",
})
if err != nil {
    return err
}
client, err := guest.NewClient(device, reg, nil)
if err != nil {
    return err
}
err = client.Connect()
```

`NewClient` applies the guest identity before creating the client, sets the
guest login payload, turns off the post-connect passive IQ and uses the
registered browser's User-Agent. `Request.BaseURL` points registration at
another host, such as a simulator. A guest can only join the
call link it registered for. Messaging, groups and app state don't apply.

Guest registration never follows HTTP redirects, including with a supplied
HTTP client. Point `BaseURL` directly at the intended registration service.
The in-memory guest store retains at most 1024 unconsumed prekeys, including
at most 64 retry keys. Upload attempts are at least one second apart; new
upload batches are generated at most once per minute. A failed upload can
retry its pending keys after the shorter interval. These local resource limits
return an error when reached and never evict an unconsumed key.

Signal identities, sessions and sender keys share a limit of 1024 entries and
16 MiB of stored key/value data, in addition to the prekey limit. Each address
is limited to 512 bytes. Writes beyond those limits return an error without
evicting trusted identities. Replacing existing state or deleting identities
and sessions releases the corresponding capacity. These are local policies
for one-call guests, not WhatsApp server limits.

Registration error strings omit provider response text. Typed `guest.Error`
fields and unwrapped causes remain available for explicit diagnostics and may
contain reflected request data; avoid logging them without redaction.
