package main

import (
	"app/internal/subscription"
	"app/internal/subscription/workflow"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"log"
	"net/http"
)

type iWFStateStartReq struct {
	iwfidl.WorkflowStateStartRequest
}

func (i iWFStateStartReq) Bind(r *http.Request) error {
	fmt.Println("inside START .. bindingd ...")
	// See if the bind works ..
	spew.Dump(i.GetWorkflowStateId())
	////TODO implement me
	//panic("implement me")
	//
	i.SetWorkflowStateId("BOO")
	// All OK after manipulation ..
	return nil
}

type iWFStateDecideReq struct {
	iwfidl.WorkflowStateDecideRequest
}

func (i iWFStateDecideReq) Bind(r *http.Request) error {
	fmt.Println("inside DECIDE .. bindingd ...")
	// See if the bind works ..
	//spew.Dump(i)
	//TODO implement me
	//panic("implement me")
	return nil
}

func setupServer() *http.Server {
	// try with chi instead of others ..
	r := chi.NewRouter()
	// from example; should it be lef tout?
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, req *http.Request) {
		panic("test")
	})

	// RESTy routes for "simulation" resource
	r.Route("/sim", func(ri chi.Router) {
		// Attach common middleware unde rthis route?
		// Start workflow ..
		ri.Get("/start", func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("Call start ..")
			// Extract out anything from the req ..
			runID, err := subscription.BasicStartWorkflow(req.Context(),
				workflow.SubscriptionWorkflow{},
				"InitID", 1)
			if err != nil {
				// We have Recoverer so can panic!
				panic(err)
				//w.WriteHeader(http.StatusInternalServerError)
				//w.Write([]byte(err.Error()))
			}
			w.Write([]byte(fmt.Sprintf("Workflow mleow-o with RunID %s", runID)))
		})
		// Send dummy signal to workflow ..
		ri.Get("/signal", func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("Call signal ...")
		})

	})

	// Attach the iWF callback routes ..
	// StateStart
	r.Post(iwf.WorkflowStateStartApi, func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("iWF Starting ...")
		var v iWFStateStartReq
		err := render.Bind(req, &v)
		if err != nil {
			// We have Recoverer so can panic!
			panic(err)
			//w.WriteHeader(http.StatusInternalServerError)
			//w.Write([]byte(err.Error()))
		}
		// This fails if not using ref method ..
		fmt.Println("AFTER_MANIPULATION: ", v.GetWorkflowStateId())
		resp, ierr := subscription.BasicInvokeStartHandler(req.Context(), v.WorkflowStateStartRequest)
		if ierr != nil {
			panic(ierr)
		}
		//w.Write([]byte("START-OK"))
		// All OK, move on to the next state ..
		render.JSON(w, req, resp)
	})
	// StateDecide
	r.Post(iwf.WorkflowStateDecideApi, func(w http.ResponseWriter, req *http.Request) {
		var v iWFStateDecideReq
		err := render.Bind(req, &v)
		if err != nil {
			// We have Recoverer so can panic!
			panic(err)
			//w.WriteHeader(http.StatusInternalServerError)
			//w.Write([]byte(err.Error()))
		}
		resp, ierr := subscription.BasicInvokeDecideHandler(req.Context(), v.WorkflowStateDecideRequest)
		if ierr != nil {
			panic(ierr)
		}
		//w.Write([]byte("DECIDE-OK"))
		// All OK, move on to the next state ..
		render.JSON(w, req, resp)
	})

	return &http.Server{
		Addr:    ":" + iwf.DefaultWorkerPort,
		Handler: r,
	}
}

func startIWFWorker() (closeFunc func()) {
	// Implement using go-chi instead of gin ..
	wfServer := setupServer()
	// TODO: See the go-chi example for proper shutdown .. timer 30s ..
	go func() {
		if err := wfServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return func() {
		fmt.Println("Closing iWF Worker ...")
		wfServer.Close()
	}

	//return func() {
	//	fmt.Println("TODO: Needs implementation for close of startIWFWorker")
	//}
}
