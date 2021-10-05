package fiber

import (
	"fmt"

	"github.com/go-gluon/gluon"
	"github.com/go-gluon/gluon/log"
	"github.com/gofiber/fiber/v2"
)

//go:generate gluon-generate --target=config

//gluon:config
type FiberConfig struct {
	Enabled               bool   `config:"enabled"`
	Listen                string `config:"listen"`
	DisableStartupMessage bool   `config:"disable-start-message"`
}

type FiberService struct {
	gluon.Annotation `name:"fiber" priority:"100" service:"true"`
	config           *FiberConfig
	app              *fiber.App
}

func (f *FiberService) InitConfig() interface{} {
	f.config = &FiberConfig{
		Enabled:               true,
		Listen:                "localhost:8080",
		DisableStartupMessage: true,
	}
	return f.config
}

func (f *FiberService) Init(info *gluon.GluonInfo, runtime *gluon.Runtime) error {
	f.app = fiber.New(fiber.Config{
		DisableStartupMessage: f.config.DisableStartupMessage,
	})
	return nil
}

func (f *FiberService) Start() {
	if !f.config.Enabled {
		log.Info("fiber service disabled.")
		return
	}
	log.Info(fmt.Sprintf("fiber service listening on http://%s", f.config.Listen))
	log.Error("Error start fiber ", log.Err(f.app.Listen(f.config.Listen)))
}
