package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
	"isgod/api"
	"log"
	"time"

	//"log"
	"os"
)

func init() {
	AnnListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "AnnListModel")
}

type AnnListModel struct {
	core.QAbstractListModel

	_         func()          `constructor:"init"`
	_         func()          `signal:"clear,auto"`
	_         func(bool)      `signal:"refresh,auto"`
	_         api.Credentials `property:"Creds"`
	modelData []api.Annoucement
}

func (m *AnnListModel) init() {
	creds, _ := api.ReadEnvCredentials()
	m.SetCreds(creds)
	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)
	m.refresh(true)
	go runDaemon(STOP, m)
}

func runDaemon(stop <-chan struct{}, annList *AnnListModel) {
	ticker := time.NewTicker(time.Duration(CONFIG.RefreshTimeout) * time.Minute)
	for {
		select {
		case <-ticker.C:
			annList.refresh(false)
		case <-stop:
			return
		}
	}
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
func (m *AnnListModel) refresh(reload bool) {
	finger, err := api.FetchFingerprint(CONFIG.Credentials)
	if err == nil {
		if finger.Fingerprint != CONFIG.RecentFingerprint || reload {
			m.removeAll()
			resp, _ := api.FetchAnnoucements(m.Creds(), 0, CONFIG.FetchSize, false)
			m.BeginInsertRows(core.NewQModelIndex(), 0, len(resp.Items)-1)
			m.modelData = append(resp.Items, m.modelData...)
			m.EndInsertRows()
			CONFIG.RecentFingerprint = finger.Fingerprint
			err = CONFIG.Save(CONFIGFILE)
			if err != nil {
				log.Printf("Couldn't save configuration: %v", err)
			}
		}
	}
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
	icon := gui.NewQIcon5(":/tray.png")
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
		STOP <- struct{}{}
		app.Exit(0)
	})
	trayMenu.AddAction("Refresh")
	tray.SetContextMenu(trayMenu)
	return tray
}

func ExecApp() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	APP = widgets.NewQApplication(len(os.Args), os.Args)
	APP.SetQuitOnLastWindowClosed(false)
	quickcontrols2.QQuickStyle_SetStyle("Material")
	view := createView()
	tray := createTray(APP, view)
	tray.Show()
	view.Show()
	APP.Exec()
}
