package claude

import "testing"

func TestDefaultModels_ContainsCurrentAnthropicClaudeModels(t *testing.T) {
	t.Parallel()

	modelsByID := make(map[string]Model, len(DefaultModels))
	for _, model := range DefaultModels {
		modelsByID[model.ID] = model
	}

	for _, id := range []string{
		"claude-fable-5",
		"claude-opus-4-8",
		"claude-sonnet-5",
		"claude-haiku-4-5-20251001",
	} {
		if _, ok := modelsByID[id]; !ok {
			t.Fatalf("expected current Anthropic model %q to be exposed", id)
		}
	}
}
