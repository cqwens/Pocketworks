package pocketworks

import (
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type PocketWorksApp struct {
	pb *pocketbase.PocketBase
}

// InitialSetup initializes the PocketWorksApp
// USAGE
//			if app, err := pocketworks.InitialSetup(); err != nil {
//			  	  log.Fatal(err)
// 			} else if err := app.Start(); err != nil {
// 				   log.Fatal(err)
//		 	}

func InitialSetup() (*PocketWorksApp, error) {
	os.Args = append(os.Args, "serve")
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))
		return se.Next()
	})

	return &PocketWorksApp{pb: app}, nil
}

func (p *PocketWorksApp) Start() error {
	return p.pb.Start()
}
