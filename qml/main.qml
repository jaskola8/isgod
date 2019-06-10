import QtQuick 2.10				//ListView
import QtQuick.Controls 2.5		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomListModel
import QtQuick.Controls.Material 2.12
Rectangle {
    width: 400
	height: 400
	color: "black"

	ColumnLayout {
	    id: list
		anchors.top: parent.top
		anchors.bottom: buttons.top
        width: parent.width
		height: 345
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

	RowLayout {
        id: buttons
        anchors.top: list.bottom
        anchors.bottom: parent.bottom
        width: parent.width

		RoundButton {
		    height: 30
		    width: 20
		    icon.source: "images/refresh.png"
		    anchors.left: parent.left
		    onClicked: listview.model.refresh(false)
	    }

		RoundButton {
		    height: 30
		    width: 20
		    icon.source: "images/reload.png"
	        anchors.right: parent.right
		    onClicked: listview.model.refresh(true)

	    }

    }
}

