package service

import "testing"

import "github.com/stretchr/testify/require"

func TestNormalizeOpenAIMessagesDispatchModelConfig(t *testing.T) {
	t.Parallel()

	cfg := normalizeOpenAIMessagesDispatchModelConfig(OpenAIMessagesDispatchModelConfig{
		OpusMappedModel:   " gpt-5.5-xhigh ",
		SonnetMappedModel: "gpt-5.5-xhigh",
		HaikuMappedModel:  " gpt-5.4-mini ",
		ExactModelMappings: map[string]string{
			" claude-sonnet-4-5-20250929 ": " gpt-5.5-xhigh ",
			"":                             "gpt-5.4",
			"claude-opus-4-6":              " ",
		},
	})

	require.Equal(t, "gpt-5.5-xhigh", cfg.OpusMappedModel)
	require.Equal(t, "gpt-5.5-xhigh", cfg.SonnetMappedModel)
	require.Equal(t, "gpt-5.4-mini", cfg.HaikuMappedModel)
	require.Equal(t, map[string]string{
		"claude-sonnet-4-5-20250929": "gpt-5.5-xhigh",
	}, cfg.ExactModelMappings)
}

func TestResolveOpenAIMessagesDispatchModelUsesHermesDefaults(t *testing.T) {
	t.Parallel()

	group := &Group{}

	require.Equal(t, "gpt-5.5-xhigh", group.ResolveMessagesDispatchModel("claude-opus-4-8"))
	require.Equal(t, "gpt-5.5-xhigh", group.ResolveMessagesDispatchModel("claude-sonnet-5"))
	require.Equal(t, "gpt-5.4-mini", group.ResolveMessagesDispatchModel("claude-haiku-4-5"))
}
