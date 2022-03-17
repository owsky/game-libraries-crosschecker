# Game Libraries Crosschecker

I developed this small piece of software with the specific goal to solve a niche problem of mine: being able to tell which games I own across different libraries that are also available on Steam. This allows me to better manage my main library through the Steam client by setting them as owned on other platforms. It's most definitely a first world problem but when you own multiple thousands of games you'll end up buying them multiple times accidentally, believe me<br>

As a data source I'm using PlayNite and its export to CSV feature so that I can delegate all the web scraping and auth to a reputable app.<br>

Please do keep in mind that this is my first Go project, I'll try to fix whatever mistakes I've committed as soon as I learn more about this language.

## Usage

I'll provide a compiled version as soon as I'm satified with the code, for the time being you can just build it yourself by launching:<br>
`go build .`<br>
For whatever reason Windows Security doesn't really like Go executables and likes to flag them as viruses, [read this](https://go.dev/doc/faq#virus) for more information<br>

After placing the CSV file in the same directory of the executable you'll get a .txt file containing the output.
