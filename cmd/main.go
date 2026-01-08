package main

import (
	"context"
	"fmt"

	"entropy-recon/internal/models/primitives"
	"entropy-recon/internal/services"
	"entropy-recon/internal/store"
	"entropy-recon/internal/tools"
)

func main() {
	domain, err := primitives.NewDomain("example.com")
	if err != nil {
		panic(err)
	}

	assetStore := store.NewMemoryAssetStore()

	subfinderTool := &tools.MockSubfinder{}

	service := services.NewDiscoveryService(assetStore, subfinderTool)

	_ = service.Run(context.Background(), domain)

	assets, _ := assetStore.List(store.AssetFilter{})

	for _, a := range assets {
		fmt.Println(a.Type, a.Value, a.Source.Tool)
	}
}
