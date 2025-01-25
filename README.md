# Rattus Rex 🐀

Modern AI reasoning chain CLI tool written in Go. Combines DeepSeek and OpenRouter for enhanced conversational capabilities.

## Features
- Two-stage AI processing with DeepSeek reasoning and OpenRouter response
- Real-time TUI interface using Bubble Tea
- Command system for model switching and configuration
- Error handling and logging
- Environment-based configuration

## Installation
```bash
git clone https://github.com/etherealblaade/rattus_rex
cd rattus_rex/cmd
go build
```

## Configuration
To get started, create a .env file in the project root and add the following keys:

```env
DEEPSEEK_API_KEY=your_deepseek_api_key
OPENROUTER_API_KEY=your_openrouter_api_key
```
Replace your_deepseek_api_key and your_openrouter_api_key with your actual API keys.

## Usage
Run the application using the following command:
```bash
./cmd
```

Commands:
* /model <name>: Switch to a different AI model.
* /reasoning: Toggle visibility of the reasoning process.
* /clear: Clear the chat history.
* q or Ctrl+C: Exit the application.



