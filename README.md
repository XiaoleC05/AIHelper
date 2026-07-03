# AIHelper — Prompt Generator, Optimizer & Memory Manager

> Stop writing prompts from scratch. Let AIHelper craft, refine, and remember your best prompts.

## Why AIHelper?

When working with AI (ChatGPT, Claude, etc.), the quality of your prompt directly determines the quality of the output. But most people face the same problems:

- Describing what you want in vague natural language leads to mediocre results
- Tweaking prompts manually is tedious, and you lose track of what worked
- Great prompts get buried in chat history, never to be found again

**AIHelper** solves all three. Turn rough ideas into structured prompts, optimize existing ones with one click, and keep a searchable memory of everything that worked.

## Features

| Feature | What It Does |
|---------|-------------|
| **Prompt Optimizer** | Paste your prompt → AI analyzes weaknesses → outputs an improved version |
| **Prompt Generator** | Describe what you need in plain language → get a structured, production-ready prompt |
| **Template Library** | Ready-to-use templates for coding, writing, translation, learning, and more |
| **Prompt Memory** | Save all your prompts with tags, full-text search, and smart recommendations |

## How It Works

AIHelper uses a two-stage optimization engine:

1. **Rule-based preprocessing** — detects missing elements (role, context, format, constraints) and fills gaps
2. **LLM refinement** — sends the preprocessed prompt to an AI model for final polishing

The result is a prompt that includes proper role definition, clear task description, output format, and quality constraints — things most people forget to specify.

## Tech Stack

| Environment | Backend | Database | Frontend | Special |
|-------------|---------|----------|----------|---------|
| Online (Oxelia51) | Go | PostgreSQL | React | LLM API |
| Desktop (exe) | Go | SQLite | Embedded React | LLM API |

- Online version shares the Oxelia51 platform's PostgreSQL + Redis instances
- Desktop version compiles to a single `.exe` with zero dependencies

## API Key

AIHelper does **not** provide its own AI API. You must bring your own API key from any supported provider (OpenAI, Anthropic, etc.). Your key is stored locally and never leaves your machine.

## Getting Started

### Online (via Oxelia51)

1. Visit [oxelia51.com](https://oxelia51.com) and sign in
2. Go to AIHelper in the tools menu
3. Enter your API key in settings
4. Start optimizing prompts

### Desktop (exe)

1. Download `AIHelper.exe` from [GitHub Releases](https://github.com/XiaoleC05/AIHelper/releases)
2. Run the executable — it starts a local web interface
3. Enter your API key
4. Everything runs locally, nothing is sent to any server

## Roadmap

- [ ] Prompt Optimizer & Generator (Priority 1-2)
- [ ] Template Library with sharing (Priority 3)
- [ ] Memory system with smart recommendations (Priority 4)

## Status

Concept phase. Development not yet started.
