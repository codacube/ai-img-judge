package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var (
	envFile string
	client  *genai.Client
	ctx     context.Context
)

var rootCmd = &cobra.Command{
	Use:   "ai-img-tool",
	Short: "A CLI tool to analyze images using the Gemini AI.",
	Long:  `A command-line interface to perform AI-powered analysis on images, including comparison, rating, and suggestions.`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// Load environment variables from the specified .env file (if it exists, else use system env vars)
		_ = godotenv.Load(envFile)

		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("GEMINI_API_KEY not set. Check your .env file (path: %s) or environment variables", envFile)
		}

		ctx = context.Background()
		client, err = genai.NewClient(ctx, &genai.ClientConfig{
			APIKey:  apiKey,
			Backend: genai.BackendGeminiAPI,
		})
		if err != nil {
			return fmt.Errorf("failed to create client: %v", err)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add the persistent --env flag to the root command
	rootCmd.PersistentFlags().StringVarP(&envFile, "env", "e", ".env", "path to .env file")

	rootCmd.AddCommand(compareCmd)
	rootCmd.AddCommand(rateCmd)
	rootCmd.AddCommand(suggestCmd)
}

// --- Helper Functions ---

// loadImage reads a file and prepares it as a genai.Part (image data)
func loadImage(imagePath string) (*genai.Part, error) {
	// Simple MIME type detection from file extension
	var mimeType string
	if strings.HasSuffix(strings.ToLower(imagePath), ".jpg") || strings.HasSuffix(strings.ToLower(imagePath), ".jpeg") {
		mimeType = "image/jpeg"
	} else if strings.HasSuffix(strings.ToLower(imagePath), ".png") {
		mimeType = "image/png"
	} else {
		return nil, fmt.Errorf("unsupported image format for %s: only jpeg and png are supported", imagePath)
	}

	bytes, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read image %s: %v", imagePath, err)
	}

	imgPart := genai.NewPartFromBytes(bytes, mimeType)
	return imgPart, nil
}

// generateAndPrint sends the request to the AI and prints the text response
func generateAndPrint(parts ...*genai.Part) {
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		contents,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
		return
	}

	fmt.Println(result.Text())
}
