package main

import (
	"context"
	"github.com/gabrielmoura/geocoding-go/internal/gtk"
	"github.com/gabrielmoura/geocoding-go/internal/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.InitLogger()

	gtk.Run("com.github.gabrielmoura.geocoding-go", ctx)
}
