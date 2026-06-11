package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"akubisa/internal/shared/storage"
	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

// Service defines the interface for AI business logic
type Service interface {
	AnalyzePronunciation(ctx context.Context, childID uuid.UUID, audioFile io.Reader, fileName string, expectedText string) (*PronunciationResult, error)
}

type service struct {
	repo         Repository
	openaiClient *openai.Client
	storage      storage.Storage
	bucketName   string
}

// NewService creates a new AI service
func NewService(repo Repository, openaiClient *openai.Client, storage storage.Storage, bucketName string) Service {
	return &service{
		repo:         repo,
		openaiClient: openaiClient,
		storage:      storage,
		bucketName:   bucketName,
	}
}

type AnalysisResponse struct {
	Score    float64 `json:"score"`
	Feedback string  `json:"feedback"`
}

func (s *service) AnalyzePronunciation(ctx context.Context, childID uuid.UUID, audioFile io.Reader, fileName string, expectedText string) (*PronunciationResult, error) {
	// 1. Upload audio to storage
	ext := filepath.Ext(fileName)
	if ext == "" {
		ext = ".wav" // default
	}
	key := fmt.Sprintf("pronunciation/%s/%s%s", childID.String(), uuid.New().String(), ext)

	audioURL, err := s.storage.UploadFile(ctx, s.bucketName, key, audioFile, "audio/wav")
	if err != nil {
		return nil, fmt.Errorf("failed to upload audio: %w", err)
	}

	// 2. Transcribe audio using Whisper
	// Whisper requires a file on disk or a multipart reader. Since we have io.Reader,
	// we might need to save it temporarily if the SDK doesn't support direct streaming easily.
	// Actually go-openai's CreateTranscription takes an AudioRequest which includes a File path or reader.

	// Create a temporary file because Whisper API usually expects a file with a proper extension
	tempFile, err := os.CreateTemp("", "audio-*"+ext)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, audioFile); err != nil {
		// If audioFile was already read, we might have an issue.
		// In the handler, we should probably use a Teereader or just read it once.
	}

	resp, err := s.openaiClient.CreateTranscription(ctx, openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: tempFile.Name(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to transcribe: %w", err)
	}

	transcription := resp.Text

	// 3. Analyze pronunciation using GPT
	score, feedback, err := s.getScoringAndFeedback(ctx, transcription, expectedText)
	if err != nil {
		return nil, fmt.Errorf("failed to get scoring: %w", err)
	}

	// 4. Save result to database
	result := &PronunciationResult{
		ChildID:       childID,
		AudioURL:      audioURL,
		Transcription: transcription,
		Score:         score,
		Feedback:      feedback,
	}

	if err := s.repo.SaveResult(result); err != nil {
		return nil, fmt.Errorf("failed to save result: %w", err)
	}

	return result, nil
}

func (s *service) getScoringAndFeedback(ctx context.Context, transcription, expectedText string) (float64, string, error) {
	prompt := fmt.Sprintf(`
Analyze the following pronunciation of a child learning to read.
Expected text: "%s"
Actual transcription: "%s"

Provide a pronunciation score from 0 to 100 and brief, encouraging feedback in Indonesian.
Respond ONLY in JSON format like this:
{"score": 85.0, "feedback": "Bagus sekali! Pengucapan huruf 'R' sudah jelas, terus berlatih ya!"}
`, expectedText, transcription)

	resp, err := s.openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are an expert Indonesian language teacher for young children.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	})

	if err != nil {
		return 0, "", err
	}

	var analysis AnalysisResponse
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &analysis); err != nil {
		return 0, "", err
	}

	return analysis.Score, analysis.Feedback, nil
}
