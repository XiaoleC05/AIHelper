package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/XiaoleC05/AIHelper/internal/config"
	"github.com/XiaoleC05/AIHelper/internal/db"
	"github.com/XiaoleC05/AIHelper/internal/handler"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsOrigins() []string {
	if v := os.Getenv("CORS_ALLOWED_ORIGINS"); v != "" {
		return strings.Split(v, ",")
	}
	return []string{"http://localhost:5173"}
}

func main() {
	config.Load()

	if err := db.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	runMigrations()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-User-Id", "X-Username", "X-Role"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	promptHandler := handler.NewPromptHandler()
	templateHandler := handler.NewTemplateHandler()
	enhanceHandler := handler.NewEnhanceHandler()
	settingsHandler := handler.NewSettingsHandler()

	r.GET("/api/health", handler.Health)

	api := r.Group("/api")
	api.Use(handler.AuthMiddleware())
	{
		api.GET("/prompts", promptHandler.List)
		api.POST("/prompts", promptHandler.Create)
		api.GET("/prompts/:id", promptHandler.Get)
		api.PUT("/prompts/:id", promptHandler.Update)
		api.DELETE("/prompts/:id", promptHandler.Delete)
		api.PATCH("/prompts/:id/favorite", promptHandler.ToggleFavorite)

		api.GET("/templates", templateHandler.List)

		api.POST("/enhance", enhanceHandler.Enhance)

		api.GET("/settings", settingsHandler.Get)
		api.PUT("/settings", settingsHandler.Update)
	}

	srv := &http.Server{
		Addr:    ":" + config.Cfg.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func runMigrations() {
	ctx := context.Background()

	initSQL := `
		CREATE SCHEMA IF NOT EXISTS aihelper;

		CREATE TABLE IF NOT EXISTS aihelper.templates (
			id BIGSERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			category TEXT NOT NULL DEFAULT '通用',
			content TEXT NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW()
		);

		CREATE TABLE IF NOT EXISTS aihelper.prompts (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			category TEXT DEFAULT '通用',
			tags TEXT[] DEFAULT '{}',
			variables TEXT[] DEFAULT '{}',
			is_favorite BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);

		CREATE TABLE IF NOT EXISTS aihelper.user_settings (
			user_id BIGINT PRIMARY KEY,
			api_key TEXT DEFAULT '',
			api_base TEXT DEFAULT '',
			model TEXT DEFAULT 'gpt-4o-mini',
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);
	`

	_, err := db.Pool.Exec(ctx, initSQL)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	seedSQL := `
		INSERT INTO aihelper.templates (name, category, content)
		VALUES
			('代码审查助手', '编程', '你是一个资深软件工程师，请对以下代码进行详细审查。关注：1) 代码正确性 2) 性能优化 3) 安全隐患 4) 可维护性。请给出具体的改进建议和代码示例。'),
			('Bug 调试专家', '编程', '你是一个调试专家。我会描述一个 Bug 现象，请按以下步骤帮助我定位问题：1) 分析可能的原因（按概率排序） 2) 提供排查步骤 3) 给出修复方案。Bug 描述：'),
			('API 设计顾问', '编程', '你是一个 API 设计专家。请根据以下需求设计 RESTful API，遵循最佳实践。包含：端点路径、HTTP 方法、请求/响应格式、错误处理、分页策略。'),

			('文章润色助手', '写作', '你是一个专业的文字编辑。请对以下文章进行润色，要求：1) 保持原意不变 2) 提升文字流畅度 3) 修正语法错误 4) 适当增加修辞。请同时输出修改说明。'),
			('大纲生成器', '写作', '你是一个内容策划专家。请为以下主题生成一份详细的写作大纲，包含：1) 标题（含 3 个备选） 2) 核心论点 3) 章节划分 4) 每章节要点 5) 预计字数分配。'),

			('专业术语翻译', '翻译', '你是一个专业翻译，精通中英双语。请翻译以下内容，要求：1) 专业术语准确 2) 语句通顺自然 3) 保留原文格式 4) 对关键术语附上原文。'),
			('多语种翻译', '翻译', '你是多语种翻译专家。请将以下内容翻译为目标语言，要求准确、流畅、符合目标语言的表达习惯。如有歧义，请注释说明。'),

			('数据分析报告', '分析', '你是一个数据分析师。请对以下数据进行深入分析，输出：1) 数据概览 2) 关键发现（至少 5 条） 3) 趋势判断 4) 可视化建议 5) 行动建议。数据如下：'),
			('竞品分析框架', '分析', '你是一个产品经理。请对以下产品进行竞品分析，维度包括：1) 功能对比 2) 用户体验 3) 商业模式 4) 技术架构 5) 市场定位 6) SWOT 分析。'),

			('创意头脑风暴', '创意', '你是一个创意总监。请围绕以下主题进行头脑风暴，输出：1) 10 个创意点子 2) 每个点子的可行性评估 3) 最佳 3 个的执行方案。主题：'),
			('故事创作引擎', '创意', '你是一个资深小说家。请根据以下设定创作故事：1) 包含引人入胜的开头 2) 有明确的故事弧线 3) 人物性格鲜明 4) 结尾出人意料。设定如下：'),

			('万能助手', '通用', '你是一个全能 AI 助手。请根据用户的需求，提供专业、准确、有条理的回答。如果涉及专业领域，请以该领域专家的口吻回答。请保持回答简洁但信息完整。'),
			('学习计划制定', '通用', '你是一个教育专家。请根据我的学习目标制定详细的学习计划，包含：1) 阶段性目标 2) 每日学习任务 3) 推荐资源 4) 练习方法 5) 进度检查点。我的目标：')
		ON CONFLICT DO NOTHING;
	`

	_, err = db.Pool.Exec(ctx, seedSQL)
	if err != nil {
		log.Fatalf("Failed to seed templates: %v", err)
	}

	log.Println("Database migrations completed")
}
