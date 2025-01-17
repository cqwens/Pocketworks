package pocketworks

import (
	"os"
	"time"

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
    errChan := make(chan error, 1)
    go func() {
        errChan <- p.pb.Start()
    }()

    select {
    case err := <-errChan:
        return err
    case <-time.After(100 * time.Millisecond):
        return nil
    }
}
