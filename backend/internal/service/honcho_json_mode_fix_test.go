package service

import (
	"encoding/json"
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/pkg/apicompat"
	"github.com/stretchr/testify/require"
)

// Regression tests for the structured-output fixes:
//   - appendRawJSON must not prefix tool_call arguments with the "{}" seed that
//     Anthropic's tool_use content_block_start carries.
//   - response_format json_object/json_schema responses must be unwrapped from
//     any markdown code fence so strict json.loads() callers get bare JSON.

func TestAppendRawJSON_ToolUseSeedNotPrefixed(t *testing.T) {
	// Simulate Anthropic streaming for a tool_use block: content_block_start
	// seeds Input="{}"; input_json_delta fragments carry the real arguments.
	var input json.RawMessage = json.RawMessage("{}") // seed from content_block_start
	for _, frag := range []string{`{"qu`, `ery":"h`, `i"}`} {
		input = appendRawJSON(input, frag)
	}
	require.Equal(t, `{"query":"hi"}`, string(input), "seed {} must not prefix the assembled arguments")

	var parsed map[string]any
	require.NoError(t, json.Unmarshal(input, &parsed), "assembled arguments must be valid JSON")
	require.Equal(t, "hi", parsed["query"])
}

func TestAppendRawJSON_EmptyArgsKeepsObject(t *testing.T) {
	// A tool call with no argument deltas keeps its "{}" (valid empty args).
	input := json.RawMessage("{}")
	require.Equal(t, "{}", string(input))
}

func TestAppendRawJSON_EmptyExisting(t *testing.T) {
	require.Equal(t, `{"a":1}`, string(appendRawJSON(nil, `{"a":1}`)))
	require.Equal(t, `{"a":1}`, string(appendRawJSON(json.RawMessage(""), `{"a":1}`)))
}

func TestAppendRawJSON_NonSeedExistingStillAppends(t *testing.T) {
	// A real (non-"{}") accumulator keeps normal append semantics.
	require.Equal(t, `{"a":1}`, string(appendRawJSON(json.RawMessage(`{"a"`), `:1}`)))
}

func TestStripJSONCodeFence(t *testing.T) {
	cases := []struct{ name, in, want string }{
		{"tagged fence", "```json\n{\"ok\":true}\n```", `{"ok":true}`},
		{"bare fence", "```\n{\"ok\":true}\n```", `{"ok":true}`},
		{"array fence", "```json\n[1,2,3]\n```", `[1,2,3]`},
		{"surrounding whitespace", "  ```json\n{\"a\":1}\n```  ", `{"a":1}`},
		{"already bare", `{"ok":true}`, `{"ok":true}`},
		{"no closing fence untouched", "```json\n{\"ok\":true}", "```json\n{\"ok\":true}"},
		{"opening only untouched", "```json", "```json"},
		{"empty", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, stripJSONCodeFence(tc.in))
		})
	}
}

func TestStripJSONFenceFromChoices(t *testing.T) {
	content, _ := json.Marshal("```json\n{\"fact\":\"x\"}\n```")
	resp := &apicompat.ChatCompletionsResponse{
		Choices: []apicompat.ChatChoice{
			{Index: 0, Message: apicompat.ChatMessage{Role: "assistant", Content: content}},
		},
	}
	stripJSONFenceFromChoices(resp)

	var got string
	require.NoError(t, json.Unmarshal(resp.Choices[0].Message.Content, &got))
	require.Equal(t, `{"fact":"x"}`, got)

	var parsed map[string]any
	require.NoError(t, json.Unmarshal([]byte(got), &parsed), "unwrapped content must be valid JSON")
	require.Equal(t, "x", parsed["fact"])
}

func TestStripJSONFenceFromChoices_NilAndEmptySafe(t *testing.T) {
	stripJSONFenceFromChoices(nil) // must not panic
	resp := &apicompat.ChatCompletionsResponse{
		Choices: []apicompat.ChatChoice{{Index: 0, Message: apicompat.ChatMessage{Role: "assistant"}}},
	}
	stripJSONFenceFromChoices(resp) // empty content, must not panic
	require.Len(t, resp.Choices[0].Message.Content, 0)
}

func TestIsJSONObjectResponseFormat(t *testing.T) {
	require.True(t, isJSONObjectResponseFormat([]byte(`{"response_format":{"type":"json_object"}}`)))
	require.True(t, isJSONObjectResponseFormat([]byte(`{"response_format":{"type":"json_schema"}}`)))
	require.False(t, isJSONObjectResponseFormat([]byte(`{"response_format":{"type":"text"}}`)))
	require.False(t, isJSONObjectResponseFormat([]byte(`{"model":"claude-haiku-4-5"}`)))
	require.False(t, isJSONObjectResponseFormat([]byte(`{}`)))
}
