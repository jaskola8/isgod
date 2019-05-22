import QtQuick 2.10				//ListView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomListModel

Rectangle {
    width: 400
	height: 400
	color: "#35322F"
    RowLayout {
        id: buttons
        width: parent.width
	    height: 40
		Button {
		    height: parent.height
	        anchors.top: parent.top
	        anchors.right: parent.right
		    Layout.fillWidth: true
		    text: "refresh"
		    onClicked: listview.model.refresh()
	    }
		Button {
		    height: parent.height
	        anchors.top: parent.top
	        anchors.left: parent.left
		    Layout.fillWidth: true
		    text: "refresh"
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
			delegate: Rectangle {
			    width: parent.width
			    height: 50
			    radius: 5
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
			        font.pointSize: 7
			        font.bold: true
			        fontSizeMode: Text.Fit
			        }
			}
		}
	}
}

