package internal

import (
	"github.com/Bios-Marcel/githuberrordialog"
	"github.com/gotk3/gotk3/gtk"
)

const (
	newTodoText = "Create new To-Do"
)

var (
	counter int
)

func createTodoComponent(parentWindow *gtk.Window) *gtk.Box {
	todoComponent, gtkError := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	githuberrordialog.ShowErrorDialogOnError(gtkError)

	todoList, gtkError := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	githuberrordialog.ShowErrorDialogOnError(gtkError)

	newTodoButton, gtkError := gtk.ButtonNewWithLabel(newTodoText)
	newTodoButton.Connect("clicked", func() {
		todoInputDialog := gtk.MessageDialogNew(
			parentWindow,
			gtk.DIALOG_DESTROY_WITH_PARENT,
			gtk.MESSAGE_QUESTION,
			gtk.BUTTONS_OK_CANCEL,
			"Enter your To-Do")
		defer todoInputDialog.Destroy()

		todoInputDialog.SetTitle(newTodoText)

		contentArea, gtkError := todoInputDialog.GetContentArea()
		githuberrordialog.ShowErrorDialogOnError(gtkError)
		if gtkError != nil {
			return
		}

		entry, gtkError := gtk.EntryNew()
		githuberrordialog.ShowErrorDialogOnError(gtkError)
		if gtkError != nil {
			return
		}
		entry.Connect("activate", func() {
			todoInputDialog.Response(gtk.RESPONSE_OK)
		})

		contentArea.PackStart(entry, true, true, 0)

		todoInputDialog.ShowAll()
		response := todoInputDialog.Run()
		if response != gtk.RESPONSE_OK {
			return
		}

		text, gtkError := entry.GetText()
		githuberrordialog.ShowErrorDialogOnError(gtkError)
		if gtkError != nil {
			return
		}

		todoEntry, gtkError := createTodoEntry(&todoComponent.Container, text)
		githuberrordialog.ShowErrorDialogOnError(gtkError)

		if gtkError == nil {
			todoList.PackStart(todoEntry, false, false, 0)
		}
	})

	githuberrordialog.ShowErrorDialogOnError(gtkError)

	todoComponent.SetMarginStart(5)
	todoComponent.SetMarginEnd(5)
	todoComponent.SetMarginTop(5)
	todoComponent.PackStart(todoList, true, true, 0)
	todoComponent.PackEnd(newTodoButton, false, false, 0)

	return todoComponent
}
