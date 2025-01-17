package browser

import (
	"searchall3.5/tuozhan/liulanqi/item"
)

var (
	chromiumList = map[string]struct {
		name        string
		profilePath string
		storage     string
		items       []item.Item
	}{
		"chrome": {
			name:        chromeName,
			profilePath: chromeUserDataPath,
			items:       item.DefaultChromium,
		},
		"edge": {
			name:        edgeName,
			profilePath: edgeProfilePath,
			items:       item.DefaultChromium,
		},
		"chromium": {
			name:        chromiumName,
			profilePath: chromiumUserDataPath,
			items:       item.DefaultChromium,
		},
		"chrome-beta": {
			name:        chromeBetaName,
			profilePath: chromeBetaUserDataPath,
			items:       item.DefaultChromium,
		},

		"360speed": {
			name:        speed360Name,
			profilePath: speed360ProfilePath,
			items:       item.DefaultChromium,
		},
		"qq": {
			name:        qqBrowserName,
			profilePath: qqBrowserProfilePath,
			items:       item.DefaultChromium,
		},

		"sogou": {
			name:        sogouName,
			profilePath: sogouProfilePath,
			items:       item.DefaultChromium,
		},
	}
	firefoxList = map[string]struct {
		name        string
		storage     string
		profilePath string
		items       []item.Item
	}{
		"firefox": {
			name:        firefoxName,
			profilePath: firefoxProfilePath,
			items:       item.DefaultFirefox,
		},
	}
)

var (
	chromeUserDataPath     = homeDir + "/AppData/Local/Google/Chrome/User Data/Default/"
	chromeBetaUserDataPath = homeDir + "/AppData/Local/Google/Chrome Beta/User Data/Default/"
	chromiumUserDataPath   = homeDir + "/AppData/Local/Chromium/User Data/Default/"
	edgeProfilePath        = homeDir + "/AppData/Local/Microsoft/Edge/User Data/Default/"
	speed360ProfilePath    = homeDir + "/AppData/Local/360chrome/Chrome/User Data/Default/"
	qqBrowserProfilePath   = homeDir + "/AppData/Local/Tencent/QQBrowser/User Data/Default/"
	sogouProfilePath       = homeDir + "/AppData/Roaming/SogouExplorer/Webkit/Default/"

	firefoxProfilePath = homeDir + "/AppData/Roaming/Mozilla/Firefox/Profiles/"
)
