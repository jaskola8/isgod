import QtQuick 2.10				//ListView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomListModel

Rectangle {
    width: 300
	height: 300
	color: "#35322F"
	ColumnLayout {
	    id: list
		anchors.top: parent.top
        height: parent.height
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
		Button {
		    height: 20
	        anchors.bottom: parent.bottom
		    Layout.fillWidth: true
		    text: "refresh"
		    onClicked: listview.model.refresh()
	    }
	}
}

