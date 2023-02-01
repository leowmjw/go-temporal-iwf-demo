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
	"net/url"

	ur "github.com/unrolled/render"
)

type iWFStateStartReq struct {
	iwfidl.WorkflowStateStartRequest
}

func (i *iWFStateStartReq) Bind(r *http.Request) error {
	fmt.Println("inside START .. bindingd ...")
	// See if the bind works ..
	//spew.Dump(i.GetWorkflowStateId())
	////TODO implement me
	//panic("implement me")
	//
	// Two below causes crash; last one can manipulate but no effect
	//i.SetWorkflowStateId("BOO")
	//i.SetWorkflowType("unknown")
	//i.Context.WorkflowId = "mleow-1"

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
		// Subscription + Payment page ..
		ri.Get("/subscription", func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("Sub + Payment ...")
			v, err := url.ParseQuery(req.URL.RawQuery)
			if err != nil {
				render.Status(req, http.StatusInternalServerError)
			}
			if v.Has("dood") {
				spew.Dump(v["dood"])
			}
			// render default page ..
			// action=subscribe
			if v.Get("action") == "subscribe" {
				if v.Has("piID") {
					subscription.ConfirmPaymentIntent(v.Get("piID"))
				} else {
					subscription.CreatePaymentIntent("dood")
				}
			} else {
				r := ur.New(ur.Options{
					Directory: "views",
				})
				err := r.HTML(w, http.StatusOK, "home", nil)
				if err != nil {
					render.Status(req, http.StatusInternalServerError)
				}
			}
		})
		// Start workflow ..
		ri.Get("/start", func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("Call start ..")
			// Extract out anything from the req ..
			runID, err := subscription.BasicStartWorkflow(req.Context(),
				&workflow.SubscriptionWorkflow{}, "sub_1MVCpUJJLlTnVKtUCCLFkvMC")
			if err != nil {
				// We have Recoverer so can panic!
				panic(err)
				//w.WriteHeader(http.StatusInternalServerError)
				//w.Write([]byte(err.Error()))
			}
			w.Write([]byte(fmt.Sprintf("Workflow mleow-0 with RunID %s", runID)))
		})
		// Send dummy signal to workflow ..
		ri.Get("/signal", func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("Call signal ...")
		})
		// Receive any webhook ..
		ri.Post("/webhook", func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("Webhook ...")
			r := ur.New()
			r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
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
