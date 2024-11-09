package main

import (
    "base-gin/config"
    _ "base-gin/docs"
    "base-gin/repository"
    "base-gin/rest"
    "base-gin/server"
    "base-gin/service"
    "base-gin/storage"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    cfg := config.NewConfig()
    storage.InitDB(cfg)
    repository.SetupRepositories()
    service.SetupServices(&cfg)

    app := server.Init(&cfg, repository.GetAccountRepo())

    // Tentukan folder tempat template berada
    app.LoadHTMLGlob("design/templates/*")

    // Menampilkan form.html ketika mengakses route /form
    app.GET("/form", func(c *gin.Context) {
        c.HTML(200, "form.html", nil)
    })

    // Setup Swagger jika dalam mode debug
    if cfg.App.Mode == "debug" {
        app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }

    rest.SetupRestHandlers(app)

    server.Serve(app)
}
