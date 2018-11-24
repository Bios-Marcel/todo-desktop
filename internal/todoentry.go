package internal

import (
	"github.com/gotk3/gotk3/gtk"
)

func createTodoEntry(parentContainer *gtk.Container, todoText string) (*gtk.Widget, error) {
	todoEntryContainer, gtkError := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	if gtkError != nil {
		return nil, gtkError
	}

	todoTextLabel, gtkError := gtk.LabelNew(todoText)
	if gtkError != nil {
		return nil, gtkError
	}

	deleteButton, gtkError := gtk.ButtonNewWithLabel("X")
	if gtkError != nil {
		return nil, gtkError
	}

	_, gtkError = deleteButton.Connect("clicked", func() {
		todoEntryContainer.Hide()
		todoEntryContainer.Remove(todoTextLabel)
		todoEntryContainer.Remove(deleteButton)
		parentContainer.Remove(todoEntryContainer)
	})
	if gtkError != nil {
		return nil, gtkError
	}

	todoEntryContainer.PackStart(todoTextLabel, false, false, 0)
	todoEntryContainer.PackEnd(deleteButton, false, false, 0)
	todoEntryContainer.ShowAll()

	return &todoEntryContainer.Widget, nil
}
