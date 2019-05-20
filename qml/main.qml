import QtQuick 2.10				//ListView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomListModel

Item {
	width: 250
	height: 200

	ColumnLayout {
		anchors.fill: parent

		ListView {
			id: listview

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: AnnListModel{}
			delegate: Text {
				text: display
			}
		}

		Button {
			Layout.fillWidth: true

			text: "clear"
			onClicked: listview.model.clear()
		}

		Button {
			Layout.fillWidth: true

			text: "refresh"
			onClicked: listview.model.refresh()
		}

	}
}