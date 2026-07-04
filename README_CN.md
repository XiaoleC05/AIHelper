# AIHelper

将模糊的自然语言需求转化为结构化的精准提示词，并建立可检索的提示词记忆库。

## Features

- 分析原始提示词的缺失要素并自动补全
- 将自然语言描述转化为包含角色、上下文、输出格式的完整提示词
- 内置编程、写作、翻译、学习等分类模板
- 按标签归类历史提示词，支持全文检索
- 基于历史使用自动推荐相关提示词

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

在线版运行于 Oxelia51 平台。Go 后端负责规则补全和模板管理，大模型 API 由用户自行提供。桌面版使用 SQLite 替代 PostgreSQL，通过 Go 二进制内嵌 React 前端运行。

## Requirements

- 在线版：Oxelia51 平台（Go + PostgreSQL + React）
- 桌面版：独立可执行文件，无需运行时依赖
- 大模型 API Key（OpenAI、Anthropic 等）

## Installation

### 桌面版

从 [GitHub Releases](https://github.com/XiaoleC05/AIHelper/releases) 下载 `AIHelper.exe`。

### 在线版

在线版集成于 Oxelia51 平台，参见 [Oxelia51 部署指南](https://github.com/XiaoleC05/Oxelia51)。

## Usage

### 在线

1. 访问 [oxelia51.com](https://oxelia51.com) 注册并登录
2. 进入 AIHelper 工具页
3. 在设置中填入大模型 API Key
4. 输入需求或粘贴待优化的提示词

### 桌面

1. 双击 `AIHelper.exe` 启动
2. 填入 API Key，所有数据存储在本地

## Roadmap

- [ ] 提示词优化与生成功能
- [ ] 模板库及创建、分享功能
- [ ] 记忆系统与智能推荐

## Contributing

1. Fork 本仓库
2. 创建功能分支 (`git checkout -b feature/xxx`)
3. 提交变更 (`git commit -m 'Add xxx'`)
4. 推送分支 (`git push origin feature/xxx`)
5. 提交 Pull Request

## License

This project is licensed under the MIT License.
