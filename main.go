package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
	"isgod/api"
	"os"
)

func init() {
	AnnListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "AnnListModel")
}

type AnnListModel struct {
	core.QAbstractListModel

	_         func()          `constructor:"init"`
	_         func()          `signal:"clear,auto"`
	_         func()          `signal:"refresh,auto"`
	_         api.Credentials `property:"Creds"`
	modelData []api.Annoucement
}

func (m *AnnListModel) init() {
	creds, _ := api.ReadEnvCredentials()
	m.SetCreds(creds)
	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)
}

func (m *AnnListModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *AnnListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}
	item := m.modelData[index.Row()]
	return core.NewQVariant14(fmt.Sprintf("%v", item.Subject))
}

func (m *AnnListModel) removeAll() {
	m.BeginRemoveRows(core.NewQModelIndex(), 0, len(m.modelData)-1)
	m.modelData = []api.Annoucement{}
	m.EndRemoveRows()
}
func (m *AnnListModel) clear() {
	m.removeAll()
}
func (m *AnnListModel) refresh() {
	m.removeAll()
	resp, _ := api.FetchHeaders(m.Creds(), 0, 10)
	m.BeginInsertRows(core.NewQModelIndex(), 0, len(resp.Items))
	m.modelData = append(resp.Items, m.modelData...)
	m.EndInsertRows()
}

func createView() *quick.QQuickView {
	view := quick.NewQQuickView(nil)
	view.SetFlags(core.Qt__FramelessWindowHint)
	view.ConnectFocusOutEvent(func(event *gui.QFocusEvent) { view.Hide() })
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	return view
}

func createTray(app *widgets.QApplication, view *quick.QQuickView) *widgets.QSystemTrayIcon {
	tray := widgets.NewQSystemTrayIcon(nil)
	icon := gui.NewQIcon5("qml/images/tray.png")
	tray.SetIcon(icon)
	tray.ConnectActivated(func(reason widgets.QSystemTrayIcon__ActivationReason) {
		if reason == widgets.QSystemTrayIcon__Trigger {
			if view.IsVisible() {
				view.Hide()
			} else {
				view.Show()
				view.RequestActivate()
				view.Raise()
			}
		}
	})
	trayMenu := widgets.NewQMenu(nil)
	quit := trayMenu.AddAction("Quit")
	quit.ConnectTriggered(func(bool) {
		app.Exit(0)
	})
	trayMenu.AddAction("Refresh")
	tray.SetContextMenu(trayMenu)
	return tray
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetQuitOnLastWindowClosed(false)
	quickcontrols2.QQuickStyle_SetStyle("Material")
	view := createView()
	tray := createTray(app, view)
	view.Show()
	tray.Show()
	app.Exec()
	/*file, err := os.Open("creds.json")
	if err != nil {
		os.Exit(1)
	}
	creds, err := api.ReadCredentials(file)
	if err != nil {
		os.Exit(2)
	}
	resp, err := api.FetchHeaders(creds, 0, 1)
	if err != nil {
		os.Exit(3)
	}
	ann := resp.Items
	first := ann[0]
	fmt.Printf("Hash: %s \nTemat: %s \nData: %s, Typ: %d", first.Hash, first.Subject, first.ModifiedDate, first.Type)

	*/
}
