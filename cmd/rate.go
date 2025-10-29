package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var rateCmd = &cobra.Command{
	Use:   "rate [image]",
	Short: "Rates a single image on a scale of 1-10",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imagePath := args[0]

		img, err := loadImage(imagePath)
		if err != nil {
			log.Fatal(err)
		}

		prompt := genai.NewPartFromText("You are an expert photography critic. Rate this image on a scale of 1 to 10. First, give the numerical score (e.g., 'Score: 8/10'), then provide a brief justification for your rating based on composition, lighting, and focus.")

		fmt.Println("--- ⭐ Rating Image... ---")
		generateAndPrint(prompt, img)
	},
}
