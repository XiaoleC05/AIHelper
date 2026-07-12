# AIHelper

Prompt enhancement and management. Refine requirements into structured prompts, with a searchable library.

## Features

- Analyze prompt weaknesses and suggest improvements
- Built-in templates for coding, writing, translation, and learning
- Tag-based categorization with full-text search
- Recommendations based on usage history

## Architecture

```text
Browser → React Frontend (Oxelia51 unified UI)
  → Go API Layer (prompt processing, template management)
  → PostgreSQL + External API (user-provided key)
```

## Installation

Integrated into the Oxelia51 platform. See [Oxelia51 deployment guide](https://github.com/XiaoleC05/Oxelia51).

## Usage

1. Visit [oxelia51.com](https://oxelia51.com), register and sign in
2. Open AIHelper from the tools menu
3. Enter your API key in settings
4. Describe your requirements or paste a prompt to optimize

## Contributing

1. Fork → 2. Feature branch → 3. Commit → 4. Push → 5. PR

## License

MIT License
