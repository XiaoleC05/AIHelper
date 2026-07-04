# AIHelper

Transform vague natural language requirements into structured, precise prompts. Build and search a reusable prompt memory library.

## Features

- Analyze prompt weaknesses and fill missing elements automatically
- Convert natural language descriptions into complete prompts with role, context, and output format
- Built-in templates for coding, writing, translation, and learning
- Tag-based categorization of historical prompts with full-text search
- Smart recommendations based on usage history

## Architecture

```text
Browser
  ↓
React Frontend (Oxelia51 unified UI)
  ↓
Go API Layer (prompt processing, template management)
  ↓        ↓
PostgreSQL    LLM API (user-provided key)
```

The online version runs on the Oxelia51 platform. The Go backend handles rule-based preprocessing and template management, while the LLM API is user-provided. The desktop version uses SQLite instead of PostgreSQL and embeds the React frontend within the Go binary.

## Requirements

- Online: Oxelia51 platform (Go, PostgreSQL, React)
- Desktop: standalone executable, no runtime dependencies
- LLM API key (OpenAI, Anthropic, or compatible providers)

## Installation

### Desktop

Download `AIHelper.exe` from [GitHub Releases](https://github.com/XiaoleC05/AIHelper/releases).

### Online

Integrated into the Oxelia51 platform. See [Oxelia51 deployment guide](https://github.com/XiaoleC05/Oxelia51).

## Usage

### Online

1. Visit [oxelia51.com](https://oxelia51.com), register and sign in
2. Open AIHelper from the tools menu
3. Enter your LLM API key in settings
4. Describe your requirements or paste a prompt to optimize

### Desktop

1. Double-click `AIHelper.exe` to start
2. Enter your API key. All data is stored locally.

## Roadmap

- [ ] Prompt optimization and generation
- [ ] Template library with creation and sharing
- [ ] Memory system with smart recommendations

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/xxx`)
3. Commit your changes (`git commit -m 'Add xxx'`)
4. Push the branch (`git push origin feature/xxx`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
