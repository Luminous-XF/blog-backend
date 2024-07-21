package main

import (
    "blog-backend/config"
    "blog-backend/global"
    "blog-backend/initialize"
    "blog-backend/pkg/logger"
    "errors"
    "fmt"
    "go.uber.org/zap"
    "net/http"
    "time"
)

func init() {
    initialize.InitProject()
}

func main() {
    addr := fmt.Sprintf(":%d", config.CONFIG.ServerConfig.Addr)
    ReadTimeout := config.CONFIG.ServerConfig.ReadTimeout
    WriteTimeout := config.CONFIG.ServerConfig.WriteTimeout

    s := &http.Server{
        Addr:           addr,
        Handler:        global.Engine,
        ReadTimeout:    ReadTimeout * time.Second,
        WriteTimeout:   WriteTimeout * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
        logger.Error("listen:", zap.String("err", err.Error()))
    }

}
