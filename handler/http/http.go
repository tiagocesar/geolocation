package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/tiagocesar/geolocation/clients/grpc_client"
	pb "github.com/tiagocesar/geolocation/handler/grpc/schema"
	"github.com/tiagocesar/geolocation/internal/models"
)

type httpServer struct {
	grpcClient *grpc_client.Client
}

func NewHttpServer(client *grpc_client.Client) *httpServer {
	return &httpServer{
		grpcClient: client,
	}
}

func (h *httpServer) ConfigureAndServe(host, port string) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	router.Get("/locations/{ip}", h.getGeolocationData)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router); err != nil {
		log.Fatalf("Failed to start HTTP server")
	}
}

func (h *httpServer) getGeolocationData(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ip := chi.URLParam(req, "ip")

	if strings.TrimSpace(ip) == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid IP address"))
		return
	}

	result, err := h.grpcClient.GetLocationData(ctx, ip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	location := toLocation(result)

	j, _ := json.Marshal(location)

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, j)
}

func toLocation(response *pb.LocationResponse) models.Geolocation {
	return models.Geolocation{
		IpAddress:   response.GetIp(),
		CountryCode: response.GetCountryCode(),
		Country:     response.GetCountry(),
		City:        response.GetCity(),
		Latitude:    response.GetLatitude(),
		Longitude:   response.GetLongitude(),
	}
}
