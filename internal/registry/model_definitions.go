// Package registry provides model definitions and lookup helpers for various AI providers.
// Static model metadata is loaded from the embedded models.json file and can be refreshed from network.
package registry

import (
	"strings"
)

const (
	codexBuiltinImage15ModelID         = "gpt-image-1.5"
	codexBuiltinImageModelID           = "gpt-image-2"
	codexBuiltinImage25FlareModelID    = "gpt-image-2.5-flare"
	codexBuiltinImage25SunburstModelID = "gpt-image-2.5-sunburst"
	codexBuiltinImage25ModelID         = "gpt-image-2.5"
	xaiBuiltinImageModelID             = "grok-imagine-image"
	xaiBuiltinImageQualityModelID      = "grok-imagine-image-quality"
	xaiBuiltinImage20ModelID           = "grok-imagine-image-2.0"
	xaiBuiltinVideoModelID             = "grok-imagine-video"
	xaiBuiltinVideo15ModelID           = "grok-imagine-video-1.5"
	xaiBuiltinVideo15PreviewID         = "grok-imagine-video-1.5-preview"
	xaiBuiltinSpeechModelID            = "grok-tts"
	xaiBuiltinSpeechVoiceModelID       = "grok-voice-tts-1.0"
)

// staticModelsJSON mirrors the top-level structure of models.json.
type staticModelsJSON struct {
	Claude      []*ModelInfo `json:"claude"`
	Gemini      []*ModelInfo `json:"gemini"`
	Vertex      []*ModelInfo `json:"vertex"`
	AIStudio    []*ModelInfo `json:"aistudio"`
	CodexFree   []*ModelInfo `json:"codex-free"`
	CodexTeam   []*ModelInfo `json:"codex-team"`
	CodexPlus   []*ModelInfo `json:"codex-plus"`
	CodexPro    []*ModelInfo `json:"codex-pro"`
	Kimi        []*ModelInfo `json:"kimi"`
	Antigravity []*ModelInfo `json:"antigravity"`
	XAI         []*ModelInfo `json:"xai"`
	Devin       []*ModelInfo `json:"devin"`
	Meta        []*ModelInfo `json:"meta"`
}

// GetClaudeModels returns the standard Claude model definitions.
func GetClaudeModels() []*ModelInfo {
	return cloneModelInfos(getModels().Claude)
}

// GetGeminiModels returns the standard Gemini model definitions.
func GetGeminiModels() []*ModelInfo {
	return cloneModelInfos(getModels().Gemini)
}

// GetGeminiVertexModels returns Gemini model definitions for Vertex AI.
func GetGeminiVertexModels() []*ModelInfo {
	return cloneModelInfos(getModels().Vertex)
}

// GetAIStudioModels returns model definitions for AI Studio.
func GetAIStudioModels() []*ModelInfo {
	return cloneModelInfos(getModels().AIStudio)
}

// GetCodexFreeModels returns model definitions for the Codex free plan tier.
func GetCodexFreeModels() []*ModelInfo {
	return WithCodexBuiltins(cloneModelInfos(getModels().CodexFree))
}

// GetCodexTeamModels returns model definitions for the Codex team plan tier.
func GetCodexTeamModels() []*ModelInfo {
	return WithCodexBuiltins(cloneModelInfos(getModels().CodexTeam))
}

// GetCodexPlusModels returns model definitions for the Codex plus plan tier.
func GetCodexPlusModels() []*ModelInfo {
	return WithCodexBuiltins(cloneModelInfos(getModels().CodexPlus))
}

// GetCodexProModels returns model definitions for the Codex pro plan tier.
func GetCodexProModels() []*ModelInfo {
	return WithCodexBuiltins(cloneModelInfos(getModels().CodexPro))
}

// GetKimiModels returns the standard Kimi (Moonshot AI) model definitions.
func GetKimiModels() []*ModelInfo {
	return cloneModelInfos(getModels().Kimi)
}

// GetAntigravityModels returns the standard Antigravity model definitions.
func GetAntigravityModels() []*ModelInfo {
	return cloneModelInfos(getModels().Antigravity)
}

var staticDevinModels = []*ModelInfo{
	devinBuiltinSWE16SlowModelInfo(),
	{
		ID:                  "devin/swe-2",
		Type:                "devin",
		OwnedBy:             "cognition",
		DisplayName:         "SWE-2",
		ContextLength:       262000,
		MaxCompletionTokens: 128000,
		Thinking: &ThinkingSupport{
			Levels: []string{"medium", "high", "max"},
		},
	},
	{
		ID:                  "devin/claude-fable-5-1",
		Type:                "devin",
		OwnedBy:             "anthropic",
		DisplayName:         "Claude Fable 5.1",
		ContextLength:       1000000,
		MaxCompletionTokens: 64000,
		Thinking: &ThinkingSupport{
			Levels: []string{"low", "medium", "high", "xhigh", "max"},
		},
	},
	{
		ID:                  "devin/gpt-6-astra",
		Type:                "devin",
		OwnedBy:             "openai",
		DisplayName:         "GPT-6 Astra",
		ContextLength:       1000000,
		MaxCompletionTokens: 64000,
		Thinking: &ThinkingSupport{
			Levels: []string{"low", "medium", "high", "xhigh", "max"},
		},
	},
	{
		ID:                  "devin/glm-5-2",
		Type:                "devin",
		OwnedBy:             "zhipu",
		DisplayName:         "GLM-5.2",
		ContextLength:       200000,
		MaxCompletionTokens: 64000,
		Thinking: &ThinkingSupport{
			Levels: []string{"none", "high"},
		},
	},
	{
		ID:                  "devin/glm-5-3",
		Type:                "devin",
		OwnedBy:             "zhipu",
		DisplayName:         "GLM-5.3",
		ContextLength:       1048576,
		MaxCompletionTokens: 128000,
		Thinking: &ThinkingSupport{
			Levels: []string{"low", "high", "max"},
		},
	},
	{
		ID:                  "devin/glm-5-3-flash",
		Type:                "devin",
		OwnedBy:             "zhipu",
		DisplayName:         "GLM-5.3 Flash",
		ContextLength:       1000000,
		MaxCompletionTokens: 128000,
		Thinking: &ThinkingSupport{
			Levels: []string{"low", "high", "max"},
		},
	},
	{
		ID:                  "devin/gpt-5-6-sol",
		Type:                "devin",
		OwnedBy:             "openai",
		DisplayName:         "GPT-5.6 Sol",
		ContextLength:       1000000,
		MaxCompletionTokens: 128000,
		Thinking: &ThinkingSupport{
			Levels: []string{"none", "low", "medium", "high", "xhigh", "max"},
		},
	},
	{
		ID:                  "devin/gemini-3-8-flash",
		Type:                "devin",
		OwnedBy:             "google",
		DisplayName:         "Gemini 3.8 Flash",
		ContextLength:       1048576,
		MaxCompletionTokens: 65536,
		Thinking: &ThinkingSupport{
			Levels: []string{"low", "medium", "high"},
		},
	},
	{
		ID:                  "devin/grok-4-6",
		Type:                "devin",
		OwnedBy:             "xai",
		DisplayName:         "Grok 4.6",
		ContextLength:       500000,
		MaxCompletionTokens: 131072,
		Thinking: &ThinkingSupport{
			Levels: []string{"low", "medium", "high", "xhigh"},
		},
	},
	{
		ID:                  "devin/deepseek-v4-flash",
		Type:                "devin",
		OwnedBy:             "deepseek",
		DisplayName:         "DeepSeek V4 Flash",
		ContextLength:       1048576,
		MaxCompletionTokens: 64000,
		Thinking: &ThinkingSupport{
			Levels: []string{"high", "max"},
		},
	},
	{
		ID:                  "devin/deepseek-v4-1-flash",
		Type:                "devin",
		OwnedBy:             "deepseek",
		DisplayName:         "DeepSeek V4.1 Flash",
		ContextLength:       1048576,
		MaxCompletionTokens: 64000,
		Thinking: &ThinkingSupport{
			Levels: []string{"high", "max"},
		},
	},
}

// AntigravityWebSearchModelFor returns the Antigravity model that should run a
// native web search request for modelID.
func AntigravityWebSearchModelFor(modelID string) string {
	modelID = normalizeAntigravityCapabilityModelID(modelID)
	if modelID == "" {
		return ""
	}
	for _, model := range GetGlobalRegistry().GetAvailableModelsByProvider("antigravity") {
		if model == nil {
			continue
		}
		currentModelID := normalizeAntigravityCapabilityModelID(model.ID)
		if currentModelID == "" {
			continue
		}
		if currentModelID == modelID {
			if model.SupportsWebSearch {
				return currentModelID
			}
			return ""
		}
	}
	return ""
}

// GetXAIModels returns the standard xAI Grok model definitions.
func GetXAIModels() []*ModelInfo {
	return WithXAIBuiltins(cloneModelInfos(getModels().XAI))
}

// WithCodexBuiltins injects hard-coded Codex-only model definitions that should
// not depend on remote models.json updates. Built-ins replace any matching IDs
// already present in the provided slice.
func WithCodexBuiltins(models []*ModelInfo) []*ModelInfo {
	return upsertModelInfos(models,
		codexBuiltinImage15ModelInfo(),
		codexBuiltinImageModelInfo(),
		codexBuiltinImage25FlareModelInfo(),
		codexBuiltinImage25SunburstModelInfo(),
		codexBuiltinImage25ModelInfo(),
	)
}

// WithXAIBuiltins injects hard-coded xAI image, video, and speech model definitions
// that should not depend on remote models.json updates.
func WithXAIBuiltins(models []*ModelInfo) []*ModelInfo {
	return upsertModelInfos(models, xaiBuiltinImageModelInfo(), xaiBuiltinImageQualityModelInfo(), xaiBuiltinImage20ModelInfo(), xaiBuiltinVideoModelInfo(), xaiBuiltinVideo15ModelInfo(), xaiBuiltinVideo15PreviewModelInfo(), xaiBuiltinSpeechModelInfo(), xaiBuiltinSpeechVoiceModelInfo())
}

func normalizeAntigravityCapabilityModelID(modelID string) string {
	modelID = strings.ToLower(strings.TrimSpace(modelID))
	if open := strings.LastIndex(modelID, "("); open >= 0 && strings.HasSuffix(modelID, ")") {
		modelID = strings.TrimSpace(modelID[:open])
	}
	return modelID
}

func codexBuiltinImage15ModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          codexBuiltinImage15ModelID,
		Object:      "model",
		Created:     1704067200, // 2024-01-01
		OwnedBy:     "openai",
		Type:        "openai",
		DisplayName: "GPT Image 1.5",
		Version:     codexBuiltinImage15ModelID,
	}
}

func codexBuiltinImageModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          codexBuiltinImageModelID,
		Object:      "model",
		Created:     1704067200, // 2024-01-01
		OwnedBy:     "openai",
		Type:        "openai",
		DisplayName: "GPT Image 2",
		Version:     codexBuiltinImageModelID,
	}
}

func codexBuiltinImage25FlareModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          codexBuiltinImage25FlareModelID,
		Object:      "model",
		Created:     1704067200, // 2024-01-01
		OwnedBy:     "openai",
		Type:        "openai",
		DisplayName: "GPT Image 2.5 Flare",
		Version:     codexBuiltinImage25FlareModelID,
	}
}

func codexBuiltinImage25SunburstModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          codexBuiltinImage25SunburstModelID,
		Object:      "model",
		Created:     1704067200, // 2024-01-01
		OwnedBy:     "openai",
		Type:        "openai",
		DisplayName: "GPT Image 2.5 Sunburst",
		Version:     codexBuiltinImage25SunburstModelID,
	}
}

func codexBuiltinImage25ModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          codexBuiltinImage25ModelID,
		Object:      "model",
		Created:     1704067200, // 2024-01-01
		OwnedBy:     "openai",
		Type:        "openai",
		DisplayName: "GPT Image 2.5",
		Version:     codexBuiltinImage25ModelID,
	}
}

func xaiBuiltinImageModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          xaiBuiltinImageModelID,
		Object:      "model",
		Created:     1735689600, // 2025-01-01
		OwnedBy:     "xai",
		Type:        "xai",
		DisplayName: "Grok Imagine Image",
		Name:        xaiBuiltinImageModelID,
		Description: "xAI Grok image generation model.",
	}
}

func xaiBuiltinImageQualityModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          xaiBuiltinImageQualityModelID,
		Object:      "model",
		Created:     1735689600, // 2025-01-01
		OwnedBy:     "xai",
		Type:        "xai",
		DisplayName: "Grok Imagine Image Quality",
		Name:        xaiBuiltinImageQualityModelID,
		Description: "xAI Grok higher-fidelity image generation model.",
	}
}

func xaiBuiltinImage20ModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          xaiBuiltinImage20ModelID,
		Object:      "model",
		Created:     1786060800, // 2026-08-07
		OwnedBy:     "xai",
		Type:        "xai",
		DisplayName: "Grok Imagine Image 2.0",
		Name:        xaiBuiltinImage20ModelID,
		Description: "xAI Grok image generation model.",
	}
}

func xaiBuiltinVideoModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          xaiBuiltinVideoModelID,
		Object:      "model",
		Created:     1735689600, // 2025-01-01
		OwnedBy:     "xai",
		Type:        "xai",
		DisplayName: "Grok Imagine Video",
		Name:        xaiBuiltinVideoModelID,
		Description: "xAI Grok video generation model.",
	}
}

func xaiBuiltinVideo15ModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          xaiBuiltinVideo15ModelID,
		Object:      "model",
		Created:     1735689600, // 2025-01-01
		OwnedBy:     "xai",
		Type:        "xai",
		DisplayName: "Grok Imagine Video 1.5",
		Name:        xaiBuiltinVideo15ModelID,
		Description: "xAI Grok video generation model.",
	}
}

func xaiBuiltinVideo15PreviewModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          xaiBuiltinVideo15PreviewID,
		Object:      "model",
		Created:     1735689600, // 2025-01-01
		OwnedBy:     "xai",
		Type:        "xai",
		DisplayName: "Grok Imagine Video 1.5 Preview",
		Name:        xaiBuiltinVideo15PreviewID,
		Description: "Compatibility alias for the xAI Grok video generation model.",
	}
}

func xaiBuiltinSpeechModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          xaiBuiltinSpeechModelID,
		Object:      "model",
		Created:     1773619200, // 2026-03-16
		OwnedBy:     "xai",
		Type:        "xai",
		DisplayName: "Grok TTS",
		Name:        xaiBuiltinSpeechModelID,
		Description: "xAI Grok unary text-to-speech model.",
	}
}

func xaiBuiltinSpeechVoiceModelInfo() *ModelInfo {
	return &ModelInfo{
		ID:          xaiBuiltinSpeechVoiceModelID,
		Object:      "model",
		Created:     1773619200, // 2026-03-16
		OwnedBy:     "xai",
		Type:        "xai",
		DisplayName: "Grok Voice TTS 1.0",
		Name:        xaiBuiltinSpeechVoiceModelID,
		Description: "xAI Grok unary text-to-speech model.",
	}
}

func upsertModelInfos(models []*ModelInfo, extras ...*ModelInfo) []*ModelInfo {
	if len(extras) == 0 {
		return models
	}

	extraIDs := make(map[string]struct{}, len(extras))
	extraList := make([]*ModelInfo, 0, len(extras))
	for _, extra := range extras {
		if extra == nil {
			continue
		}
		id := strings.TrimSpace(extra.ID)
		if id == "" {
			continue
		}
		key := strings.ToLower(id)
		if _, exists := extraIDs[key]; exists {
			continue
		}
		extraIDs[key] = struct{}{}
		extraList = append(extraList, cloneModelInfo(extra))
	}

	if len(extraList) == 0 {
		return models
	}

	filtered := make([]*ModelInfo, 0, len(models)+len(extraList))
	for _, model := range models {
		if model == nil {
			continue
		}
		id := strings.TrimSpace(model.ID)
		if id == "" {
			continue
		}
		if _, exists := extraIDs[strings.ToLower(id)]; exists {
			continue
		}
		filtered = append(filtered, model)
	}

	filtered = append(filtered, extraList...)
	return filtered
}

// cloneModelInfos returns a shallow copy of the slice with each element deep-cloned.
func cloneModelInfos(models []*ModelInfo) []*ModelInfo {
	if len(models) == 0 {
		return nil
	}
	out := make([]*ModelInfo, len(models))
	for i, m := range models {
		out[i] = cloneModelInfo(m)
	}
	return out
}

// GetStaticModelDefinitionsByChannel returns static model definitions for a given channel/provider.
// It returns nil when the channel is unknown.
//
// Supported channels:
//   - claude
//   - gemini
//   - gemini-interactions
//   - vertex
//   - aistudio
//   - codex
//   - kimi
//   - antigravity
//   - xai
//   - devin
//   - meta
func GetStaticModelDefinitionsByChannel(channel string) []*ModelInfo {
	key := strings.ToLower(strings.TrimSpace(channel))
	switch key {
	case "claude":
		return GetClaudeModels()
	case "gemini":
		return GetGeminiModels()
	case "gemini-interactions":
		return GetGeminiModels()
	case "vertex":
		return GetGeminiVertexModels()
	case "aistudio":
		return GetAIStudioModels()
	case "codex":
		return GetCodexProModels()
	case "kimi", "kimi-ai", "kimi.ai", "kimi.com":
		return GetKimiModels()
	case "antigravity":
		return GetAntigravityModels()
	case "xai", "x-ai", "grok":
		return GetXAIModels()
	case "devin":
		return GetDevinModels()
	case "meta", "muse":
		return GetMetaModels()
	default:
		return nil
	}
}

// LookupStaticModelInfoByChannel searches one provider-specific static section.
// It does not fall back across providers, so callers can preserve provenance.
func LookupStaticModelInfoByChannel(modelID, channel string) *ModelInfo {
	modelID = strings.TrimSpace(modelID)
	if modelID == "" {
		return nil
	}
	for _, model := range GetStaticModelDefinitionsByChannel(channel) {
		if model != nil && model.ID == modelID {
			return cloneModelInfo(model)
		}
	}
	return nil
}

// GetMetaModels returns the standard Meta Muse model definitions.
func GetMetaModels() []*ModelInfo {
	return cloneModelInfos(getModels().Meta)
}

// LookupStaticModelInfo searches all static model definitions for a model by ID.
// Returns nil if no matching model is found.
func LookupStaticModelInfo(modelID string) *ModelInfo {
	if modelID == "" {
		return nil
	}

	data := getModels()
	allModels := [][]*ModelInfo{
		data.Claude,
		data.Gemini,
		data.Vertex,
		data.AIStudio,
		data.CodexPro,
		data.Kimi,
		data.Antigravity,
		data.XAI,
		data.Devin,
		staticDevinModels,
		data.Meta,
	}
	for _, models := range allModels {
		for _, m := range models {
			if m != nil && m.ID == modelID {
				return cloneModelInfo(m)
			}
		}
	}

	return nil
}
// GetCodeBuddyModels returns the available models for CodeBuddy (Tencent).
// These models are served through the copilot.tencent.com API.
func GetCodeBuddyModels() []*ModelInfo {
	now := int64(1748044800) // 2025-05-24
	return cloneModelInfos([]*ModelInfo{
		{
			ID:                  "auto",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Auto",
			Description:         "Automatic model selection via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-5v-turbo",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-5v Turbo",
			Description:         "GLM-5v Turbo via CodeBuddy",
			ContextLength:       200000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-5.1",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-5.1",
			Description:         "GLM-5.1 via CodeBuddy",
			ContextLength:       200000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-5.0-turbo",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-5.0 Turbo",
			Description:         "GLM-5.0 Turbo via CodeBuddy",
			ContextLength:       200000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-5.0",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-5.0",
			Description:         "GLM-5.0 via CodeBuddy",
			ContextLength:       200000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-4.7",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-4.7",
			Description:         "GLM-4.7 via CodeBuddy",
			ContextLength:       200000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "minimax-m2.7",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "MiniMax M2.7",
			Description:         "MiniMax M2.7 via CodeBuddy",
			ContextLength:       200000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "kimi-k2.5",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Kimi K2.5",
			Description:         "Kimi K2.5 via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "kimi-k2-thinking",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Kimi K2 Thinking",
			Description:         "Kimi K2 Thinking via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
			Thinking:            &ThinkingSupport{ZeroAllowed: true},
		},
		{
			ID:                  "deepseek-v3-2-volc",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "DeepSeek V3.2 (Volc)",
			Description:         "DeepSeek V3.2 via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		// ---- WorkBuddy 侧模型（与 buddy2api workbuddy 通道对齐，2026-09-30）----
		{
			ID:                  "hy3",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Hunyuan HY3",
			Description:         "WorkBuddy HY3 via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "hy3-x",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Hunyuan HY3-X",
			Description:         "WorkBuddy HY3-X via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "hy4-preview",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Hunyuan HY4 Preview",
			Description:         "WorkBuddy HY4 Preview via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "hy4-preview-x",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Hunyuan HY4 Preview X",
			Description:         "WorkBuddy HY4 Preview X via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "default",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Default",
			Description:         "WorkBuddy default routing via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-5.3",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-5.3",
			Description:         "GLM-5.3 via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-5.3-flash",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-5.3 Flash",
			Description:         "GLM-5.3 Flash via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-5.2",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-5.2",
			Description:         "GLM-5.2 via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-4.6",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-4.6",
			Description:         "GLM-4.6 via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "glm-4.6v",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "GLM-4.6V",
			Description:         "GLM-4.6V via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
			SupportedInputModalities: []string{"TEXT", "IMAGE"},
		},
		{
			ID:                  "minimax-m3",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "MiniMax M3",
			Description:         "MiniMax M3 via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "minimax-m2.5",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "MiniMax M2.5",
			Description:         "MiniMax M2.5 via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "kimi-k3-1",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Kimi K3.1",
			Description:         "Kimi K3.1 via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "kimi-k2.8-preview",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Kimi K2.8 Preview",
			Description:         "Kimi K2.8 Preview via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "kimi-k2.7",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Kimi K2.7",
			Description:         "Kimi K2.7 via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "kimi-k2.6",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Kimi K2.6",
			Description:         "Kimi K2.6 via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "deepseek-v4-flash",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "DeepSeek V4 Flash",
			Description:         "DeepSeek V4 Flash via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "deepseek-v4.1-flash",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "DeepSeek V4.1 Flash",
			Description:         "DeepSeek V4.1 Flash via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "deepseek-v4-pro",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "DeepSeek V4 Pro",
			Description:         "DeepSeek V4 Pro via CodeBuddy",
			ContextLength:       128000,
			MaxCompletionTokens: 32768,
		},
		{
			ID:                  "hunyuan-2.0-thinking",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Hunyuan 2.0 Thinking",
			Description:         "Hunyuan 2.0 Thinking via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
			Thinking:            &ThinkingSupport{ZeroAllowed: true},
		},
		{
			ID:                  "hunyuan-chat",
			Object:              "model",
			Created:             now,
			OwnedBy:             "tencent",
			Type:                "codebuddy",
			DisplayName:         "Hunyuan Chat",
			Description:         "Hunyuan Chat via CodeBuddy",
			ContextLength:       256000,
			MaxCompletionTokens: 32768,
		},
	})
}


// GetCodeBuddyIntlModels returns the available models for CodeBuddy International (codebuddy.ai).
func GetCodeBuddyIntlModels() []*ModelInfo {
	now := int64(1748044800)
	textMod := []string{"TEXT"}
	textImageMod := []string{"TEXT", "IMAGE"}
	return cloneModelInfos([]*ModelInfo{
		{
			ID:                       "default-model",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "codebuddy",
			Type:                     "codebuddy-intl",
			DisplayName:              "Default",
			Description:              "Default model selection via CodeBuddy International (x2.00 credits)",
			ContextLength:            176000,
			MaxCompletionTokens:      24000,
			InputTokenLimit:          176000,
			OutputTokenLimit:         24000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "default-model-lite",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "codebuddy",
			Type:                     "codebuddy-intl",
			DisplayName:              "Default Lite",
			Description:              "Default Lite model via CodeBuddy International (x0.67 credits)",
			ContextLength:            176000,
			MaxCompletionTokens:      24000,
			InputTokenLimit:          176000,
			OutputTokenLimit:         24000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gemini-3.1-pro",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "google",
			Type:                     "codebuddy-intl",
			DisplayName:              "Gemini 3.1 Pro",
			Description:              "Gemini 3.1 Pro via CodeBuddy International (x1.32 credits)",
			ContextLength:            400000,
			MaxCompletionTokens:      64000,
			InputTokenLimit:          400000,
			OutputTokenLimit:         64000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gemini-3.0-flash",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "google",
			Type:                     "codebuddy-intl",
			DisplayName:              "Gemini 3.0 Flash",
			Description:              "Gemini 3.0 Flash via CodeBuddy International (x0.33 credits)",
			ContextLength:            400000,
			MaxCompletionTokens:      64000,
			InputTokenLimit:          400000,
			OutputTokenLimit:         64000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gemini-2.5-pro",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "google",
			Type:                     "codebuddy-intl",
			DisplayName:              "Gemini 2.5 Pro",
			Description:              "Gemini 2.5 Pro via CodeBuddy International (x0.90 credits)",
			ContextLength:            400000,
			MaxCompletionTokens:      64000,
			InputTokenLimit:          400000,
			OutputTokenLimit:         64000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gemini-2.5-flash",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "google",
			Type:                     "codebuddy-intl",
			DisplayName:              "Gemini 2.5 Flash",
			Description:              "Gemini 2.5 Flash via CodeBuddy International (x0.22 credits)",
			ContextLength:            400000,
			MaxCompletionTokens:      64000,
			InputTokenLimit:          400000,
			OutputTokenLimit:         64000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gemini-3.1-flash-lite",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "google",
			Type:                     "codebuddy-intl",
			DisplayName:              "Gemini 3.1 Flash Lite",
			Description:              "Gemini 3.1 Flash Lite via CodeBuddy International (x0.17 credits)",
			ContextLength:            200000,
			MaxCompletionTokens:      65536,
			InputTokenLimit:          200000,
			OutputTokenLimit:         65536,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gpt-5.4",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "openai",
			Type:                     "codebuddy-intl",
			DisplayName:              "GPT 5.4",
			Description:              "GPT 5.4 via CodeBuddy International (x1.65 credits)",
			ContextLength:            272000,
			MaxCompletionTokens:      128000,
			InputTokenLimit:          272000,
			OutputTokenLimit:         128000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gpt-5.2",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "openai",
			Type:                     "codebuddy-intl",
			DisplayName:              "GPT 5.2",
			Description:              "GPT 5.2 via CodeBuddy International (x1.25 credits)",
			ContextLength:            272000,
			MaxCompletionTokens:      128000,
			InputTokenLimit:          272000,
			OutputTokenLimit:         128000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gpt-5.3-codex",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "openai",
			Type:                     "codebuddy-intl",
			DisplayName:              "GPT 5.3 Codex",
			Description:              "GPT 5.3 Codex via CodeBuddy International (x1.25 credits)",
			ContextLength:            272000,
			MaxCompletionTokens:      128000,
			InputTokenLimit:          272000,
			OutputTokenLimit:         128000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gpt-5.2-codex",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "openai",
			Type:                     "codebuddy-intl",
			DisplayName:              "GPT 5.2 Codex",
			Description:              "GPT 5.2 Codex via CodeBuddy International (x1.25 credits)",
			ContextLength:            272000,
			MaxCompletionTokens:      128000,
			InputTokenLimit:          272000,
			OutputTokenLimit:         128000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gpt-5.1",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "openai",
			Type:                     "codebuddy-intl",
			DisplayName:              "GPT 5.1",
			Description:              "GPT 5.1 via CodeBuddy International (x0.90 credits)",
			ContextLength:            272000,
			MaxCompletionTokens:      128000,
			InputTokenLimit:          272000,
			OutputTokenLimit:         128000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gpt-5.1-codex",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "openai",
			Type:                     "codebuddy-intl",
			DisplayName:              "GPT 5.1 Codex",
			Description:              "GPT 5.1 Codex via CodeBuddy International (x0.90 credits)",
			ContextLength:            272000,
			MaxCompletionTokens:      128000,
			InputTokenLimit:          272000,
			OutputTokenLimit:         128000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gpt-5.1-codex-max",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "openai",
			Type:                     "codebuddy-intl",
			DisplayName:              "GPT 5.1 Codex Max",
			Description:              "GPT 5.1 Codex Max via CodeBuddy International (x0.90 credits)",
			ContextLength:            200000,
			MaxCompletionTokens:      72000,
			InputTokenLimit:          200000,
			OutputTokenLimit:         72000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "gpt-5.1-codex-mini",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "openai",
			Type:                     "codebuddy-intl",
			DisplayName:              "GPT 5.1 Codex Mini",
			Description:              "GPT 5.1 Codex Mini via CodeBuddy International (x0.18 credits)",
			ContextLength:            272000,
			MaxCompletionTokens:      128000,
			InputTokenLimit:          272000,
			OutputTokenLimit:         128000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                       "deepseek-v3-2-volc",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "deepseek",
			Type:                     "codebuddy-intl",
			DisplayName:              "DeepSeek V3.2",
			Description:              "DeepSeek V3.2 via CodeBuddy International (x0.29 credits)",
			ContextLength:            96000,
			MaxCompletionTokens:      32000,
			InputTokenLimit:          96000,
			OutputTokenLimit:         32000,
			SupportedInputModalities: textMod,
		},
		{
			ID:                       "glm-5.0",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "zhipu",
			Type:                     "codebuddy-intl",
			DisplayName:              "GLM 5.0",
			Description:              "GLM 5.0 via CodeBuddy International (x0.80 credits)",
			ContextLength:            200000,
			MaxCompletionTokens:      48000,
			InputTokenLimit:          200000,
			OutputTokenLimit:         48000,
			SupportedInputModalities: textMod,
		},
		{
			ID:                       "kimi-k2.5",
			Object:                   "model",
			Created:                  now,
			OwnedBy:                  "moonshot",
			Type:                     "codebuddy-intl",
			DisplayName:              "Kimi K2.5",
			Description:              "Kimi K2.5 via CodeBuddy International (x0.45 credits)",
			ContextLength:            164000,
			MaxCompletionTokens:      32000,
			InputTokenLimit:          164000,
			OutputTokenLimit:         32000,
			SupportedInputModalities: textImageMod,
		},
		{
			ID:                        "gemini-3.0-pro-image",
			Object:                    "model",
			Created:                   now,
			OwnedBy:                   "google",
			Type:                      "codebuddy-intl",
			DisplayName:               "Gemini 3.0 Pro Image",
			Description:               "Gemini 3.0 Pro Image via CodeBuddy International (x4.96 credits)",
			ContextLength:             164000,
			MaxCompletionTokens:       4096,
			InputTokenLimit:           164000,
			OutputTokenLimit:          4096,
			SupportedInputModalities:  textImageMod,
			SupportedOutputModalities: []string{"IMAGE"},
		},
		{
			ID:                        "gemini-3.1-flash-image",
			Object:                    "model",
			Created:                   now,
			OwnedBy:                   "google",
			Type:                      "codebuddy-intl",
			DisplayName:               "Gemini 3.1 Flash Image",
			Description:               "Gemini 3.1 Flash Image via CodeBuddy International (x1.78 credits)",
			ContextLength:             164000,
			MaxCompletionTokens:       4096,
			InputTokenLimit:           164000,
			OutputTokenLimit:          4096,
			SupportedInputModalities:  textImageMod,
			SupportedOutputModalities: []string{"IMAGE"},
		},
		{
			ID:                        "gemini-2.5-flash-image",
			Object:                    "model",
			Created:                   now,
			OwnedBy:                   "google",
			Type:                      "codebuddy-intl",
			DisplayName:               "Gemini 2.5 Flash Image",
			Description:               "Gemini 2.5 Flash Image via CodeBuddy International (x1.14 credits)",
			ContextLength:             164000,
			MaxCompletionTokens:       4096,
			InputTokenLimit:           164000,
			OutputTokenLimit:          4096,
			SupportedInputModalities:  textImageMod,
			SupportedOutputModalities: []string{"IMAGE"},
		},
		{
			ID:                        "hunyuan-image-v3.0",
			Object:                    "model",
			Created:                   now,
			OwnedBy:                   "tencent",
			Type:                      "codebuddy-intl",
			DisplayName:               "Hunyuan Image V3",
			Description:               "Hunyuan Image V3 via CodeBuddy International (x5.00 credits)",
			ContextLength:             16384,
			MaxCompletionTokens:       4096,
			InputTokenLimit:           16384,
			OutputTokenLimit:          4096,
			SupportedInputModalities:  textImageMod,
			SupportedOutputModalities: []string{"IMAGE"},
		},
		{
			ID:                        "hunyuan-image-v2.0-general-edit",
			Object:                    "model",
			Created:                   now,
			OwnedBy:                   "tencent",
			Type:                      "codebuddy-intl",
			DisplayName:               "Hunyuan Image Edit",
			Description:               "Hunyuan Image Edit via CodeBuddy International",
			ContextLength:             16384,
			MaxCompletionTokens:       4096,
			InputTokenLimit:           16384,
			OutputTokenLimit:          4096,
			SupportedInputModalities:  textImageMod,
			SupportedOutputModalities: []string{"IMAGE"},
		},
	})
}
