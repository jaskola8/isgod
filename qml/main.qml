import QtQuick 2.10				//ListView
import QtQuick.Controls 2.5		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomListModel
import QtQuick.Controls.Material 2.12
Rectangle {
    width: 400
	height: 400
	Material.theme: Material.Dark
	Material.primary: Material.Dark
	Material.foreground: Material.Dark
	Material.background: Material.Dark
	Material.accent: Material.Yellow

    RowLayout {

        id: buttons
        width: parent.width
	    height: 45

		RoundButton {
		    height: 30
		    width: 20
		    icon.source: "images/refresh.png"
	        anchors.top: parent.top
	        anchors.right: parent.right
		    onClicked: listview.model.refresh()
	    }

		RoundButton {
		    height: 30
		    width: 20
		    icon.source: "images/refresh.png"
	        anchors.top: parent.top
	        anchors.left: parent.left
		    onClicked: listview.model.refresh()

	    }

    }

	ColumnLayout {
	    id: list
		anchors.top: buttons.bottom
		anchors.bottom: parent.bottom
        width: parent.width
		ListView {
			id: listview

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: AnnListModel{}
			delegate: Rectangle
			 {
			    width: parent.width
			    height: 50
			    color: "transparent"
                border.color: "black"
                border.width: 1
			        Text {
			        color: "white"
			        leftPadding: 4
			        rightPadding: 4
			        topPadding: 4
			        bottomPadding: 4
			        width: parent.width
			        text: display
			        wrapMode: Text.WordWrap
			        font.pointSize: 9
			        font.bold: true
			        fontSizeMode: Text.Fit
			        }
			}
		}
	}
}

