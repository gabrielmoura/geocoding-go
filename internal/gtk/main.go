package gtk

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/gabrielmoura/geocoding-go/internal/api"
)

//go:embed mui.ui
var gtkTemplate string

func Run(applicationId string, ctx context.Context) {
	app := gtk.NewApplication(applicationId, gio.ApplicationFlagsNone)

	// conecta o sinal de ativação da aplicação passando o contexto
	app.Connect("activate", func() {
		activate(app, ctx)
	})

	// verifica se o contexto foi cancelado e fecha a aplicação
	go func() {
		<-ctx.Done()
		app.Quit()
	}()
	app.Run([]string{})

}

func activate(app *gtk.Application, ctx context.Context) {
	builder := gtk.NewBuilderFromString(gtkTemplate, len(gtkTemplate))

	Window := builder.GetObject("MainWindow").Cast().(*gtk.Window)
	Adress := builder.GetObject("address").Cast().(*gtk.Entry)
	Cep := builder.GetObject("cep").Cast().(*gtk.Entry)
	BtnCep := builder.GetObject("btnCep").Cast().(*gtk.Button)
	BtnAdress := builder.GetObject("btnAddress").Cast().(*gtk.Button)
	Output := builder.GetObject("output").Cast().(*gtk.TextView)

	BtnCep.Connect("clicked", func() {
		cep := Cep.Buffer().Text()
		if cep == "" {
			Output.Buffer().SetText(fmt.Sprintf("Erro: %s", "CEP não pode ser vazio"))
			return
		}
		bind, err := api.GetCep(ctx, cep)

		if err != nil {
			Output.Buffer().SetText(fmt.Sprintf("Erro: %s", err.Error()))
			return
		}

		jsonX, _ := json.MarshalIndent(bind, "", "  ")
		Output.Buffer().SetText(string(jsonX))
	})

	BtnAdress.Connect("clicked", func() {
		adress := Adress.Buffer().Text()
		if adress == "" {
			Adress.SetName("error")
			Output.Buffer().SetText(fmt.Sprintf("Erro: %s", "Endereço não pode ser vazio"))
			return
		}
		bind, err := api.GetGeocoding(ctx, adress)

		if err != nil {
			Adress.SetName("error")
			Output.Buffer().SetText(fmt.Sprintf("Erro: %s", err.Error()))
			return
		}

		jsonX, _ := json.MarshalIndent(bind, "", "  ")
		Output.Buffer().SetText(string(jsonX))
	})

	app.AddWindow(Window)
	Window.Show()
}
