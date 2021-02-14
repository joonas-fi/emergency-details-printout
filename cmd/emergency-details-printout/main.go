package main

import (
	"bytes"
	"context"
	"embed"
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/function61/gokit/app/aws/lambdautils"
	"github.com/function61/gokit/net/http/httputils"
	"github.com/function61/gokit/os/osutil"
	"github.com/function61/html2pdf/pkg/html2pdfclient"
)

//go:embed assets/*
var assetsFiles embed.FS

//go:embed templates/*
var templates embed.FS

func main() {
	if lambdautils.InLambda() {
		handler, err := newServerHandler()
		osutil.ExitIfError(err)
		lambda.StartHandler(lambdautils.NewLambdaHttpHandlerAdapter(handler))
		return
	}

	osutil.ExitIfError(serve(
		osutil.CancelOnInterruptOrTerminate(nil)))
}

func serve(ctx context.Context) error {
	handler, err := newServerHandler()
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    ":80",
		Handler: handler,
	}

	return httputils.CancelableServer(ctx, srv, func() error { return srv.ListenAndServe() })
}

func newServerHandler() (http.Handler, error) {
	mux := http.NewServeMux()

	toPdfClient, err := html2pdfclient.New(
		html2pdfclient.Function61,
		html2pdfclient.TokenFromEnv)
	if err != nil {
		return nil, err
	}

	printoutTpl, err := parseTemplate("templates/printout.html")
	if err != nil {
		return nil, err
	}

	inputFormTpl, err := parseTemplate("templates/inputform.html")
	if err != nil {
		return nil, err
	}

	// this HTTP app lives under https://joonas.fi/emergency-details-printout/ (thus the fixed prefix)

	// TODO: for safety, don't allow users to see the HTML (image sources etc.)
	mux.HandleFunc("/emergency-details-printout/generate", func(w http.ResponseWriter, r *http.Request) {
		html := printoutHtmlFromRequest(w, r, printoutTpl)
		if html == nil {
			return // error handled already
		}

		if r.PostForm.Get("html") != "" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = w.Write(html)
		} else {
			pdf, err := toPdfClient.Render(r.Context(), string(html), nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/pdf")
			_, _ = io.Copy(w, pdf)
		}
	})

	mux.Handle(
		"/emergency-details-printout/assets/",
		http.StripPrefix(
			"/emergency-details-printout",
			http.FileServer(http.FS(assetsFiles))))

	mux.HandleFunc("/emergency-details-printout/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/emergency-details-printout/" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_ = inputFormTpl.Execute(w, nil)
		} else {
			http.NotFound(w, r)
		}
	})

	return mux, nil
}

// returns nil on error (HTTP error response handled internally)
func printoutHtmlFromRequest(w http.ResponseWriter, r *http.Request, printoutTpl *template.Template) []byte {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	formGet := func(key string) string { // shorthand
		return r.PostForm.Get(key)
	}

	details := struct {
		Translated         // for convenience, embed
		EmergencyNumber    string
		Address            string
		Fire               string
		Bleeding           string
		WaterLeak          string
		MaintenanceCompany string
		Veterinarian       string
		ResuscitationPrint bool
		HeimlichPrint      bool
		TodaysDate         string
	}{
		Translated:         translations[formGet("lang")],
		EmergencyNumber:    formGet("emergencyNumber"),
		Address:            formGet("address"),
		Fire:               formGet("fire"),
		Bleeding:           formGet("bleeding"),
		WaterLeak:          formGet("waterLeak"),
		MaintenanceCompany: formGet("maintenanceCompany"),
		Veterinarian:       formGet("veterinarian"),
		HeimlichPrint:      formGet("heimlich") != "",
		ResuscitationPrint: formGet("resuscitation") != "",
		TodaysDate:         time.Now().UTC().Format("2006-01-02"),
	}

	var buf bytes.Buffer
	if err := printoutTpl.Execute(&buf, details); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return buf.Bytes()
}

func parseTemplate(name string) (*template.Template, error) {
	tpl, err := template.ParseFS(templates, name)
	if err != nil {
		return nil, err
	}

	return tpl, nil
}
