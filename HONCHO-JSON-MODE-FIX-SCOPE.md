# OC-router fix: `response_format: json_object` returns fenced JSON (+ tool-args `{}` prefix)

Status: IMPLEMENTED + TESTED locally (go build/vet clean, new regression tests pass, the
`{}`-prefix test proven to fail on old code, full apicompat suite green; the one failing
service test `TestBuildPaymentSubjectAppliesAffixToSubscriptionPlanDefaultName` is
pre-existing and unrelated). NOT yet deployed — the prod router origin is behind the
`oc-router` cloudflared tunnel on a host not reachable from this workspace, so the
build+restart must run on that host (see "Deploy" at bottom).

## What was implemented (commit on branch fix/per-model-effort-thinking)
- `backend/internal/service/gateway_forward_as_responses.go` — `appendRawJSON` now treats
  an empty/whitespace or `{}` accumulator as empty, so a tool_use `content_block_start`
  `{}` seed no longer prefixes the assembled `tool_calls[].function.arguments`
  (fixes `{}{"q":1}` -> `{"q":1}`). Fixes both the CC and Responses buffered paths (shared helper).
- `backend/internal/service/gateway_forward_as_chat_completions.go` — new helpers
  `isJSONObjectResponseFormat(body)` (gjson read of `response_format.type`),
  `stripJSONCodeFence(string)` (unwraps a fully-fenced ```json…``` block, else no-op), and
  `stripJSONFenceFromChoices(resp)`. `handleCCBufferedFromAnthropic` takes a `jsonMode`
  flag; when the client asked for json_object/json_schema, the assembled assistant content
  is unwrapped so strict `json.loads()` callers (Honcho deriver) get bare JSON.
- `backend/internal/service/honcho_json_mode_fix_test.go` — regression tests.

## Original root cause (unchanged)

## Symptom
Honcho's deriver (and any client using OpenAI-style structured output) gets back
markdown-fenced JSON instead of raw JSON, so `json.loads()` fails at
`Expecting value: line 1 column 1 (char 0)` and derivation produces 0 facts.

## Reproduce (from any VM or with a valid OC-router key)
```
curl -s -X POST https://router.tryopencomputer.com/v1/chat/completions \
  -H "Authorization: Bearer $OC_ROUTER_KEY" -H "Content-Type: application/json" \
  --data '{"model":"claude-haiku-4-5","messages":[{"role":"user","content":"Return JSON with key ok true"}],"max_tokens":50,"response_format":{"type":"json_object"}}'
```
Observed `choices[0].message.content` = ```` ```json\n{ "ok": true }\n``` ```` (fenced).
Expected: bare `{"ok": true}`.

Also broken: OpenAI tool/function mode returns malformed
`arguments: "{}{\"ok\": true}"` (double-brace prefix).

## Root cause location
`oc-router/backend/` (Go service). The router proxies OpenAI-compatible
`/v1/chat/completions` to Anthropic-family models. When `response_format.type ==
"json_object"` is requested, the Anthropic model tends to wrap JSON in a ```` ```json ````
fence; the router passes it through verbatim instead of returning bare JSON.

## Fix options (pick during the separate effort)
1. When `response_format.type == "json_object"`, post-process the assistant text:
   strip a leading/trailing ```` ```json ... ``` ```` (or any ``` fence) so
   `content` is bare JSON. Smallest, safest, localized.
2. Inject/strengthen a system instruction for json_object requests ("output only raw
   JSON, no markdown fences") — less reliable than (1).
3. Fix tool/function translation so `tool_calls[].function.arguments` is valid JSON
   (drop the stray `{}` prefix).
Prefer (1)+(3). Add a regression test mirroring the reproduce call: assert
`json.loads(content)` succeeds and no ``` fences remain.

## Verify after fix (fleet-wide benefit; no VM change needed)
On Maya: send a session-scoped main-agent chat, wait ~30s, then
`docker exec honcho-database-1 psql -U postgres -d postgres -tAc "select count(*) from public.documents"`
should grow above 0, and `/api/memory` honcho `fact_count` should rise.

## Why deferred
OC-router is shared prod infra (all customers, all VMs). Change + deploy needs review and
rollback, per the workspace's "confirm before high-stakes actions" rule. Honcho on Maya is
already enabled and persisting messages; this fix unblocks fact-derivation fleet-wide.
