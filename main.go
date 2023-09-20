package main

import (
	"encoding/base64"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Base64 Encoder / Decoder")
	w.SetContent(makeUI(w))
	w.Resize(fyne.NewSize(w.Canvas().Size().Width, 650))
	w.ShowAndRun()
}

func makeUI(w fyne.Window) fyne.CanvasObject {
	header := canvas.NewText("Base64 Encoder / Decoder", theme.PrimaryColor())
	header.TextSize = 42
	header.Alignment = fyne.TextAlignCenter

	u, _ := url.Parse("https://github.com/able8/base64-encoder-decoder")
	footer := widget.NewHyperlinkWithStyle("github.com/able8/base64-encoder-decoder", u, fyne.TextAlignCenter, fyne.TextStyle{})

	input := widget.NewEntry()
	input.MultiLine = true
	input.Wrapping = fyne.TextWrapBreak
	input.SetPlaceHolder("Input Text Or Read from Clipboard")

	output := widget.NewEntry()
	output.MultiLine = true
	output.Wrapping = fyne.TextWrapBreak
	output.SetPlaceHolder("Output Result")

	encode := widget.NewButtonWithIcon("Encode", theme.MediaSkipNextIcon(), func() {
		if input.Text == "" {
			input.Text = w.Clipboard().Content()
			input.Refresh()
		}
		out := base64.StdEncoding.EncodeToString([]byte(input.Text))
		output.Text = out
		output.Refresh()
	})
	encode.Importance = widget.HighImportance

	clear := widget.NewButtonWithIcon("clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
		input.Text = ""
		input.Refresh()
	})
	clear.Importance = widget.MediumImportance

	decode := widget.NewButtonWithIcon("Decode", theme.MediaSkipPreviousIcon(), func() {
		if input.Text == "" {
			input.Text = w.Clipboard().Content()
			input.Refresh()
		}
		out, err := base64.StdEncoding.DecodeString(input.Text)
		if err == nil {
			output.Text = string(out)
		} else {
			output.Text = err.Error()
		}
		output.Text = string(out)
		output.Refresh()
	})
	decode.Importance = widget.HighImportance

	copy := widget.NewButtonWithIcon("Cut Result", theme.ContentCutIcon(), func() {
		clipboard := w.Clipboard()
		clipboard.SetContent(output.Text)
		output.Text = ""
		output.Refresh()

		input.Text = ""
		input.Refresh()
	})
	copy.Importance = widget.WarningImportance

	return container.NewBorder(header, footer, nil, nil,
		container.NewGridWithRows(2,
			container.NewBorder(nil, container.NewVBox(container.NewGridWithColumns(3, encode, clear, decode), copy), nil, nil, input), output,
		),
	)
}
