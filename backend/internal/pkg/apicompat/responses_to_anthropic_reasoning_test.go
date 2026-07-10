package apicompat

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// These tests pin the per-model reasoning mapping to Anthropic's documented
// effort + extended-thinking contract. The bug they guard against: sending
// manual thinking.budget_tokens to Opus 4.7/4.8 (which returns HTTP 400), and
// sending an effort parameter to Haiku 4.5 (which has no effort param). They
// run the real production conversion entrypoint, ResponsesToAnthropicRequest.

func anthReqWithEffort(t *testing.T, model, effort string) *AnthropicRequest {
	t.Helper()
	out, err := ResponsesToAnthropicRequest(&ResponsesRequest{
		Model:     model,
		Reasoning: &ResponsesReasoning{Effort: effort},
		Input:     json.RawMessage(`"hi"`),
	})
	require.NoError(t, err)
	return out
}

// Opus 4.7/4.8: effort ladder incl. xhigh via output_config.effort, and
// NEVER manual thinking.budget_tokens (would 400 upstream).
func TestReasoning_Opus48_EffortNoManualThinking(t *testing.T) {
	for _, model := range []string{"claude-opus-4-8", "claude-opus-4-7", "anthropic/claude-opus-4-8-20260101"} {
		for _, eff := range []string{"low", "medium", "high", "xhigh", "max"} {
			out := anthReqWithEffort(t, model, eff)
			require.NotNil(t, out.OutputConfig, "%s/%s must set output_config", model, eff)
			require.Equal(t, eff, out.OutputConfig.Effort, "%s/%s effort passthrough (xhigh kept)", model, eff)
			require.Nil(t, out.Thinking, "%s/%s must NOT send manual thinking (400 on 4.7/4.8)", model, eff)
		}
	}
}

// Opus 4.6 / Sonnet 4.6: effort low/med/high/max; xhigh clamps to max; no manual thinking.
func TestReasoning_Opus46Sonnet46_ClampXhigh_NoManualThinking(t *testing.T) {
	for _, model := range []string{"claude-opus-4-6", "claude-sonnet-4-6"} {
		out := anthReqWithEffort(t, model, "xhigh")
		require.NotNil(t, out.OutputConfig)
		require.Equal(t, "max", out.OutputConfig.Effort, "%s xhigh must clamp to max (no xhigh tier)", model)
		require.Nil(t, out.Thinking, "%s must not send manual thinking", model)

		for _, eff := range []string{"low", "medium", "high", "max"} {
			o := anthReqWithEffort(t, model, eff)
			require.Equal(t, eff, o.OutputConfig.Effort)
			require.Nil(t, o.Thinking)
		}
	}
}

// Fable 5: effort ladder incl. xhigh; thinking adaptive (no manual budget).
func TestReasoning_Fable5_EffortInclXhigh_NoManualThinking(t *testing.T) {
	out := anthReqWithEffort(t, "claude-fable-5", "xhigh")
	require.NotNil(t, out.OutputConfig)
	require.Equal(t, "xhigh", out.OutputConfig.Effort)
	require.Nil(t, out.Thinking)
}

// Sonnet 5: effort ladder incl. xhigh; thinking adaptive (no manual budget).
func TestReasoning_Sonnet5_EffortInclXhigh_NoManualThinking(t *testing.T) {
	out := anthReqWithEffort(t, "claude-sonnet-5", "xhigh")
	require.NotNil(t, out.OutputConfig)
	require.Equal(t, "xhigh", out.OutputConfig.Effort)
	require.Nil(t, out.Thinking)
}

// Haiku 4.5: NO effort param; manual extended thinking with a budget.
func TestReasoning_Haiku45_ManualThinking_NoEffort(t *testing.T) {
	for _, eff := range []string{"medium", "high", "xhigh", "max"} {
		out := anthReqWithEffort(t, "claude-haiku-4-5", eff)
		require.Nil(t, out.OutputConfig, "haiku does not support the effort parameter")
		require.NotNil(t, out.Thinking, "haiku uses manual extended thinking")
		require.Equal(t, "enabled", out.Thinking.Type)
		require.Greater(t, out.Thinking.BudgetTokens, 0, "haiku thinking needs a budget")
	}
}

// Unknown/older models: preserve legacy behavior (effort + manual thinking for
// non-low) so pre-4.6 reasoning models that accept manual budgets keep working.
func TestReasoning_LegacyModel_KeepsManualThinking(t *testing.T) {
	out := anthReqWithEffort(t, "claude-sonnet-4-5", "high")
	require.NotNil(t, out.OutputConfig)
	require.Equal(t, "high", out.OutputConfig.Effort)
	require.NotNil(t, out.Thinking, "legacy models keep manual thinking")
	require.Equal(t, "enabled", out.Thinking.Type)

	// legacy low → effort low, no manual thinking
	low := anthReqWithEffort(t, "claude-sonnet-4-5", "low")
	require.Equal(t, "low", low.OutputConfig.Effort)
	require.Nil(t, low.Thinking)

	// legacy xhigh → clamped to max (older models have no xhigh)
	xh := anthReqWithEffort(t, "claude-sonnet-4-5", "xhigh")
	require.Equal(t, "max", xh.OutputConfig.Effort)
}

// Re-derivation after model mapping (the gateway race-condition fix):
// reasoning was decided for the client model; the upstream model differs.
func TestReapplyAnthropicReasoningForModel(t *testing.T) {
	// Start as if converted for opus-4-8 + xhigh (effort=xhigh, no thinking).
	base := anthReqWithEffort(t, "claude-opus-4-8", "xhigh")
	require.Equal(t, "xhigh", base.OutputConfig.Effort)
	require.Nil(t, base.Thinking)

	// Remapped to opus-4-6 (no xhigh) → must clamp to max, still no manual thinking.
	ReapplyAnthropicReasoningForModel(base, "claude-opus-4-6", "xhigh")
	require.NotNil(t, base.OutputConfig)
	require.Equal(t, "max", base.OutputConfig.Effort, "xhigh must clamp to max for 4.6 after remap")
	require.Nil(t, base.Thinking, "no manual thinking after remap to 4.6")

	// Remapped to haiku-4-5 → effort dropped, manual thinking set.
	h := anthReqWithEffort(t, "claude-opus-4-8", "high")
	ReapplyAnthropicReasoningForModel(h, "claude-haiku-4-5", "high")
	require.Nil(t, h.OutputConfig, "haiku has no effort param after remap")
	require.NotNil(t, h.Thinking)
	require.Equal(t, "enabled", h.Thinking.Type)
	require.Greater(t, h.Thinking.BudgetTokens, 0)

	// Empty effort clears reasoning entirely.
	e := anthReqWithEffort(t, "claude-opus-4-8", "high")
	ReapplyAnthropicReasoningForModel(e, "claude-opus-4-8", "")
	require.Nil(t, e.OutputConfig)
	require.Nil(t, e.Thinking)
}

// The classifier itself, exercised directly for clarity.
func TestAnthropicReasoningPlan(t *testing.T) {
	cases := []struct {
		model                                   string
		useEffort, supportsXhigh, manual, known bool
	}{
		{"claude-opus-4-8", true, true, false, true},
		{"claude-opus-4-7", true, true, false, true},
		{"claude-opus-4-6", true, false, false, true},
		{"claude-sonnet-4-6", true, false, false, true},
		{"claude-fable-5", true, true, false, true},
		{"claude-sonnet-5", true, true, false, true},
		{"claude-haiku-4-5", false, false, true, true},
		{"claude-sonnet-4-5", false, false, false, false},
		{"gpt-5.2", false, false, false, false},
	}
	for _, c := range cases {
		ue, sx, mt, kn := anthropicReasoningPlan(c.model)
		require.Equalf(t, c.useEffort, ue, "%s useEffort", c.model)
		require.Equalf(t, c.supportsXhigh, sx, "%s supportsXhigh", c.model)
		require.Equalf(t, c.manual, mt, "%s manualThinking", c.model)
		require.Equalf(t, c.known, kn, "%s known", c.model)
	}
}
