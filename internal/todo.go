package internal

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/Bios-Marcel/githuberrordialog"
)

//Startup sets up all components and shows the application
func Startup() {
	gtk.Init(nil)

	githuberrordialog.Repository = "github.com/Bios-Marcel/todo-desktop"
	application, gtkError := gtk.ApplicationNew("me.biosmarcel.tododesktop", glib.APPLICATION_FLAGS_NONE)
	githuberrordialog.ShowErrorDialogOnError(gtkError)

	mainWindow, gtkError := gtk.ApplicationWindowNew(application)
	githuberrordialog.ShowErrorDialogOnError(gtkError)

	mainWindow.SetTypeHint(gdk.WINDOW_TYPE_HINT_DESKTOP)
	mainWindow.SetDecorated(false)
	mainWindow.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)
	mainWindow.SetSizeRequest(600, 600)

	mainWindow.Add(createTodoComponent(&mainWindow.Window))
	mainWindow.ShowAll()

	gtk.Main()
}
