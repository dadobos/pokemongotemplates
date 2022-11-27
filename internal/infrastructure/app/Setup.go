package app

import (
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/dadobos/pokemongotemplates/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

type ApplicationState struct {
	HTTPServer *http.Server
	Handler    *gin.Engine
	Config     *infrastructure.ApplicationEnvironment
}

type ApplicationServer struct {
	State ApplicationState
}

var templateFunctionMap = template.FuncMap{

	"GetCurrentYear":      infrastructure.GetCurrentYear,
	"GetPrevPagePokemons": infrastructure.GetPrevPagePokemons,
	"GetNextPagePokemons": infrastructure.GetNextPagePokemons,
}

func (s *ApplicationServer) registerHandlers() {
	var files []string

	templateLocation := fmt.Sprintf("%s/web/templates", s.State.Config.TemplateLocation)

	err := filepath.Walk(templateLocation, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".gohtml") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}

	s.State.Handler.SetFuncMap(templateFunctionMap)

	s.State.Handler.LoadHTMLFiles(files...)
	s.State.Handler.GET("/", s.pageHandler())
	s.State.Handler.GET("/pagination", s.paginationPageHandler())

	s.State.Handler.GET("/form/send-pagination", s.slugSendPaginationRequestHandler())
}

func NewApplicationServer(userOptions *ApplicationState) *ApplicationServer {
	state := userOptions
	if state == nil {
		state = &ApplicationState{}
	}

	if state.Handler == nil {
		gin.SetMode(gin.ReleaseMode)
		state.Handler = gin.Default()
	}

	if state.HTTPServer == nil {
		state.HTTPServer = &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  100 * time.Second,
			Addr:         ":8080",
			Handler:      state.Handler,
		}
	}

	if state.Config == nil {
		config := infrastructure.GetConfig()
		state.Config = &config
	}

	s := ApplicationServer{
		State: ApplicationState{
			HTTPServer: state.HTTPServer,
			Handler:    state.Handler,
			Config:     state.Config,
		},
	}
	s.registerHandlers()
	return &s

}
