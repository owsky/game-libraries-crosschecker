# Game Libraries Crosschecker

I developed this small piece of software with the specific goal to solve a niche problem of mine: being able to tell which games I own across different libraries that are also available on Steam. This allows me to better manage my main library through the Steam client by setting them as owned on other platforms. It's most definitely a first world problem but when you own multiple thousands of games you'll end up buying them multiple times accidentally, believe me.<br>

As a data source I'm using PlayNite and its export to CSV feature so that I can delegate all the web scraping and auth to a reputable app.<br>

## Usage

I'll provide a compiled version as soon as I'm satified with the code, for the time being you can just build it yourself:


First you need to have [Node.js](https://nodejs.org/en/) installed


Then after moving to the project directory with your terminal you need to launch the following commands:

1. `npm install`
2. Pick whichever platform you want to target:
      - `npm run electron:package:win`<br>
      - `npm run electron:package:mac`<br>
      - `npm run electron:package:linux`<br>

The packaged app will be located in \<project-dir\>/dist/.
If you prefer you can skip packaging altogether and launch the development environment with:

`npm start`

## Pivot

Initially I started to develop this project using Go but I quickly discovered that the GUI libraries for Go are kinda terrible so I decided to pivot and rewrite it from scratch using Electron and React. Ultimately I think it was the better choice. The Go module is still available on a separate branch as a CLI tool.
