package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequests, secureHeaders)
	dynamicMiddleWare := alice.New(app.session.Enable)

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Post("/snippet/create", dynamicMiddleWare.ThenFunc(app.createSnippet))
	mux.Get("/snippet/create", dynamicMiddleWare.ThenFunc(app.createSnippetForm))
	mux.Get("/snippet/:id", dynamicMiddleWare.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("../../ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// fmt.Println("bitch")

	return standardMiddleware.Then(mux)

}
