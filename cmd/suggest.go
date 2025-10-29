package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var suggestCmd = &cobra.Command{
	Use:   "suggest [image]",
	Short: "Suggests improvements for a single image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imagePath := args[0]

		img, err := loadImage(imagePath)
		if err != nil {
			log.Fatal(err)
		}

		prompt := genai.NewPartFromText("You are a helpful photo editor. Analyze this image and provide 3-5 specific, actionable suggestions on how it could be improved (e.g., 'Crop the image to place the subject on the right third,' 'Increase the contrast slightly,' 'Warm the color temperature').")

		fmt.Println("--- 💡 Suggesting Improvements... ---")
		generateAndPrint(prompt, img)
	},
}
