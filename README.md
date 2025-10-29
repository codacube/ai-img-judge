# AI-IMG-JUDGE

I have a lot of images that need to be sorted through, and I can't always pick the best one, so I'll let AI decide for me! It can compare two images, rate an image, and suggest improvments

## Notes

- Currently Gemini only (Gemini Developer API)
  - Create a new key here: https://aistudio.google.com/app/api-keys
  - Docs: https://ai.google.dev/gemini-api/docs/image-understanding#go
- Setup an environment variable for the API Key, then in a terminal run `export GEMINI_API_KEY="<your_api_key_here>"`

You can also setup an .env file for the API key if you wish. Inside the .env file, add your key:

```ini
# This is a comment
GEMINI_API_KEY="<your_api_key_here>"
```

- Be sure to add the .env file to .gitignore if you are storing on github

## Installation

```bash
go get google.golang.org/genai
go get github.com/joho/godotenv
go mod tidy
go build
```

## Usage

```bash
ai-img-judge --help
```
