package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var detailed bool

var compareCmd = &cobra.Command{
	Use:   "compare [image1] [image2]",
	Short: "Compares two images and decides which is 'best'",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		imagePath1 := args[0]
		imagePath2 := args[1]

		img1, err := loadImage(imagePath1)
		if err != nil {
			log.Fatal(err)
		}
		img2, err := loadImage(imagePath2)
		if err != nil {
			log.Fatal(err)
		}

		var responseType string
		var prompt *genai.Part
		if detailed {
			prompt = genai.NewPartFromText("You are an expert photography judge. Compare these two photos in detail. Which one is 'best' based on composition, color, lighting, and overall emotional impact Explain your reasoning clearly.")
			responseType = "detailed"
		} else {
			prompt = genai.NewPartFromText("You are an expert photography judge. Compare these two photos and decide which one is 'best' based on composition, color, lighting, and overall emotion impact. Provide a brief explanation for your choice.")
			responseType = "summary"
		}

		fmt.Printf("--- 📸 Comparing Images (%s)... ---\n", responseType)
		generateAndPrint(prompt, img1, img2)
	},
}

func init() {
	// Compare allows for a detailed comparison from the AI
	compareCmd.Flags().BoolVarP(
		&detailed,
		"detailed",
		"d",
		false,
		"Provide a more detailed, in-depth comparison",
	)
}
