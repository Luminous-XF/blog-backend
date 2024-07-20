package main

import (
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
	addr := fmt.Sprintf(":%d", global.CONFIG.ServerConfig.Addr)
	ReadTimeout := global.CONFIG.ServerConfig.ReadTimeout
	WriteTimeout := global.CONFIG.ServerConfig.WriteTimeout

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
