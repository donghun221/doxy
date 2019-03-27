package main

import (
	"flag"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"HelloTencent/api"
	"HelloTencent/examples/ui/data/swagger"
)

var (
	commonEndpoint = flag.String("commonservice_endpoint", "localhost:19270", "endpoint of Common gRPC Service")
)

func main() {
	Run()
}

func Run() (error) {
	// init gateway
	gwmux, err := newGateway()
	if err != nil {
		return err
	}

	httpmux := newHttpServer(gwmux)

	log.Print("Common gRPC Server gateway start at port 1984...")
	http.ListenAndServe(":1984", httpmux)
	return nil
}

func newHttpServer(gwmux http.Handler) (http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	return mux
}

func newGateway() (http.Handler, error) {
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err :=api.RegisterCommonServiceHandlerFromEndpoint(ctx, gwmux, *commonEndpoint, opts); err != nil {
		return nil, err
	}

	if err :=api.RegisterHelloTencentServiceHandlerFromEndpoint(ctx, gwmux, *commonEndpoint, opts); err != nil {
		return nil, err
	}

	return gwmux, nil
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if ! strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("api", p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third-party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}