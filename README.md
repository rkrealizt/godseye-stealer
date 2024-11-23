<br>

<p align="center">
    <img src="./.github/assets/avatar.png" width=100  >
</p>



<h1 align="center">GodsEye Stealer</h1>

<p align="center">Windows systems, extracting User Data from Discord, Browsers, Crypto Wallets and more, from every user on every disk.</p>

---

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#features">Features</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#remove">Remove</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
    <li><a href="#disclaimer">Disclaimer</a></li>  </ol>
</details>

### Features:

- [antidebug](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/antidebug/antidebug.go): Terminates debugging tools.
- [antivirus](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/antivirus/antivirus.go): Disables Windows Defender and blocks access to antivirus websites.
- [antivm](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/antivm/antivm.go): Detects and exits when running in virtual machines (VMs).
- [browsers](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/browsers/browsers.go):
  - Steals logins, cookies, credit cards, history, and download lists from 37 Chromium-based browsers.
  - Steals logins, cookies, history, and download lists from 10 Gecko browsers.
- [clipper](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/clipper/clipper.go): Replaces the user's clipboard content with a specified crypto address when copying another address.
- [commonfiles](https://github.com/rkrealizt/godseye-stealer/tree/main/modules/commonfiles/commonfiles.go): Steals sensitive files from common locations.
- [discodes](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/discodes/discodes.go): Captures Discord Two-Factor Authentication (2FA) backup codes.
- [discordinjection](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/discordinjection/injection.go):
  - Intercepts login, register, and 2FA login requests.
  - Captures backup codes requests.
  - Monitors email/password change requests.
  - Intercepts credit card/PayPal addition requests.
  - Blocks the use of QR codes for login.
  - Prevents requests to view devices.
- [fakerror](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/fakeerror/fakeerror.go): Trick user into believing the program closed due to an error.
- [games](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/games/games.go): Extracts Epic Games, Uplay, Minecraft (14 launchers) and Riot Games sessions.
- [hideconsole](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/hideconsole/hideconsole.go): Module to hide the console.
- [startup](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/startup/startup.go): Ensures the program runs at system startup.
- [system](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/system/system.go): Gathers CPU, GPU, RAM, IP, location, saved Wi-Fi networks, and more.
- [tokens](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/tokens/tokens.go): Extracts tokens from 4 Discord applications, Chromium-based browsers, and Gecko browsers.
- [uacbypass](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/uacbypass/bypass.go): Grants privileges to steal user data from others users.
- [wallets](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/wallets/wallets.go): Steals data from 10 local wallets and 55 wallet extensions.
- [walletsinjection](https://github.com/rkrealizt/godseye-stealer/blob/main/modules/walletsinjection/walletsinjection.go): Captures mnemonic phrases and passwords from 2 crypto wallets.

## Getting started

### Prerequisites

* [Git](https://git-scm.com/downloads)
* [The Go Programming Language](https://go.dev/dl/)

## Remove

This guide will help you removing godseye-stealer from your system

1. Reset your PC to factory default settings. (I don't even know).

## License
This library is released under the MIT License. See LICENSE file for more informations.

## Disclaimer

### Important Notice: This tool is intended for educational purposes only.

This software, referred to as godseye-stealer, is provided strictly for educational and research purposes. Under no circumstances should this tool be used for any malicious activities, including but not limited to unauthorized access, data theft, or any other harmful actions.

### Usage Responsibility:

By accessing and using this tool, you acknowledge that you are solely responsible for your actions. Any misuse of this software is strictly prohibited, and the creator (nocaob) disclaims any responsibility for how this tool is utilized. You are fully accountable for ensuring that your usage complies with all applicable laws and regulations in your jurisdiction.

### No Liability:

The creator (nocaob) of this tool shall not be held responsible for any damages or legal consequences resulting from the use or misuse of this software. This includes, but is not limited to, direct, indirect, incidental, consequential, or punitive damages arising out of your access, use, or inability to use the tool.

### No Support:

The creator (nocaob) will not provide any support, guidance, or assistance related to the misuse of this tool. Any inquiries regarding malicious activities will be ignored.

### Acceptance of Terms:

By using this tool, you signify your acceptance of this disclaimer. If you do not agree with the terms stated in this disclaimer, do not use the software.
