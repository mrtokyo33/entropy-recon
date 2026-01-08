package main

import (
	"context"
	"fmt"

	"entropy-recon/internal/models/primitives"
	"entropy-recon/internal/services"
	"entropy-recon/internal/store"
	storeSql "entropy-recon/internal/store/sql"
	"entropy-recon/internal/tools"
)

func main() {
	domain, err := primitives.NewDomain("example.com")
	if err != nil {
		panic(err)
	}

	assetStore, err := storeSql.NewSQLiteAssetStore("recon.db")
	if err != nil {
		panic(err)
	}

	subfinderTool := &tools.MockSubfinder{}

	service := services.NewDiscoveryService(assetStore, subfinderTool)

	_ = service.Run(context.Background(), domain)

	assets, _ := assetStore.List(store.AssetFilter{})

	for _, a := range assets {
		fmt.Println(a.Type, a.Value, a.Source.Tool)
	}
}
