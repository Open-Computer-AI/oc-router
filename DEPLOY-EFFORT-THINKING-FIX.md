# Deploy: per-model effort/thinking fix (router)

This branch (`fix/per-model-effort-thinking`, commit `97f53a9b`) makes the OC
Router honor effort/thinking per Anthropic's spec. It is **verified** (7 new
tests + full apicompat suite + gateway CC tests pass, `go build`/`go vet` clean,
independent audit) but **not yet deployed** — the prod router origin is behind
the Cloudflare Tunnel `oc-router` (`router.tryopencomputer.com → http://localhost:8080`)
on a host that isn't reachable from this workspace, so the final pull/restart
must run on that host.

## What the fix changes
`backend/internal/pkg/apicompat/responses_to_anthropic_request.go`:
- Opus 4.7/4.8 → `output_config.effort` incl. `xhigh`, **no** manual
  `thinking.budget_tokens` (which they 400 on).
- Opus 4.6 / Sonnet 4.6 → effort, `xhigh`→`max` clamp, no manual thinking.
- Fable 5 → effort incl. `xhigh`.
- Haiku 4.5 → **no** effort param, manual `thinking:{enabled,budget_tokens}`.
- Unknown/older models → legacy behavior preserved.
- Plus `ReapplyAnthropicReasoningForModel` re-derives the shape after
  `GetMappedModel`, so a remapped model (e.g. opus-4-8→opus-4-6) can't carry an
  invalid `xhigh`.

## Deploy on the router host (the one running the `oc-router` cloudflared tunnel)

### Option A — rebuild from source (matches the systemd/binary install)
```bash
cd /opt/oc-router            # or wherever the repo/binary lives
git fetch origin && git checkout fix/per-model-effort-thinking   # or: git am < the patch
cd backend && go build -o /opt/oc-router/oc-router ./cmd/server
sudo systemctl restart oc-router
```

### Option B — docker-compose install
The committed compose pulls `weishaw/sub2api:latest` (upstream, WITHOUT this fix).
Build the fixed image and point compose at it:
```bash
git checkout fix/per-model-effort-thinking
docker build -t oc-router:effort-thinking .
# edit deploy/docker-compose.yml: image: oc-router:effort-thinking
docker compose -f deploy/docker-compose.yml up -d --force-recreate oc-router
```
(Or build+push `ghcr.io/open-computer-ai/oc-router` via the Release workflow and
pull that tag.)

## Verify LIVE after deploy (this is the real proof the fix works)
```bash
# Thinking should now come back as reasoning_content on a streamed request:
curl -sN https://router.tryopencomputer.com/v1/chat/completions \
  -H "Authorization: Bearer <router-key>" -H "Content-Type: application/json" \
  -d '{"model":"claude-opus-4-8","stream":true,"reasoning":{"enabled":true,"effort":"xhigh"},
       "max_tokens":800,"messages":[{"role":"user","content":"What is 17*23? Think step by step."}]}' \
  | grep -c reasoning_content        # >0 == thinking now flows (was 0 before)

# And effort should change cost: max should use materially more completion_tokens than off.
```
Before the fix these returned **0** `reasoning_content` and `off`==`max` tokens.
