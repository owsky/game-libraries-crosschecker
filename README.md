# Game Libraries Crosschecker

I developed this small piece of software with the specific goal to solve a niche problem of mine: being able to tell which games I own across different libraries that are also available on Steam. This allows me to better manage my main library through the Steam client by setting them as owned on other platforms. It's most definitely a first world problem but when you own multiple thousands of games you'll end up buying them multiple times accidentally, believe me.<br>

As a data source I'm using Playnite and its export to CSV feature so that I can delegate all the web scraping and auth to a reputable app.<br>

## Usage

For the app to work you need to provide three pieces of information:
1. Your Playnite libraries exported in the CSV format
2. Your Steam user ID
3. Your Steam API key

There are two way to achieve this:
1. Launch the app via the CLI and provide the three arguments in the specified order
2. Create a config.json file with the same information. For ease of use the app automatically creates the config file if no cli arguments are provided. After that you just need to insert the data in the right fields

The app's output will be saved in a output.txt file in the same directory where the app is located

## How to retrieve the information

### Playnite export

[Playnite](https://playnite.link/) is a free and open source app that collects all your libraries' information into a single local database. To export your library you need to navigate to Extentions->Export Library by clicking on Playnite's icon on the top bar.  
In case the export option is not appearing check if the addon is installed at Addons->Installed->Generic->Library Exporter  
If it is not installed look it up under Browse->Generic and install it.

### Steam user ID

There are many ways to retrieve your UID:
If you haven't set a custom URL for your profile you can simply follow the [official instructions](https://help.steampowered.com/en/faqs/view/2816-BE67-5B69-0FEC).  
Otherwise you can use one of the many websites that do this for you. One of them is [steamid.io](https://steamid.io/). Just make sure that you copy the Steam ID 64 format, if it's called differently on other sources you need to look for a numerical string, no other characters should be present.

### Steam API key

Game Libraries Crosschecker needs an API key to retrieve user profiles as per the [Steam API documentation](https://partner.steamgames.com/doc/webapi/IPlayerService#GetOwnedGames). If you don't have one already you can generate one for free [here](https://steamcommunity.com/login/home/?goto=%2Fdev%2Fapikey). The same address will then show you the key afterwards in case you want to access it again.


## Build

If you want to build it yourself:

First you need to have [Go](https://go.dev/) installed.

Then after moving to the project directory with your terminal you need to launch:
`go build .`  

For whatever reason Windows Security doesn't really like Go executables and likes to flag them as viruses, [read this](https://go.dev/doc/faq#virus) for more information.
