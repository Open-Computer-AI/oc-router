package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPricingServiceGPT56AliasesBypassGenericFallback(t *testing.T) {
	svc := &PricingService{
		pricingData: map[string]*LiteLLMModelPricing{
			"gpt-5.6": openAIGPT54FallbackPricing,
			"gpt-5.4": openAIGPT54FallbackPricing,
		},
	}

	tests := []struct {
		model string
		want  *LiteLLMModelPricing
	}{
		{model: "gpt-5.6", want: openAIGPT56SolFallbackPricing},
		{model: "gpt-5.6-xhigh", want: openAIGPT56SolFallbackPricing},
		{model: "gpt-5.6-max", want: openAIGPT56SolFallbackPricing},
		{model: "gpt-5.6-sol-max", want: openAIGPT56SolFallbackPricing},
		{model: "gpt-5.6-terra-xhigh", want: openAIGPT56TerraFallbackPricing},
		{model: "gpt-5.6-luna-xhigh", want: openAIGPT56LunaFallbackPricing},
		{model: "gpt-5.6-solar", want: nil},
		{model: "gpt-5.6-terrestrial", want: nil},
		{model: "gpt-5.6-lunatic", want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			require.Same(t, tt.want, svc.GetModelPricing(tt.model))
		})
	}
}
