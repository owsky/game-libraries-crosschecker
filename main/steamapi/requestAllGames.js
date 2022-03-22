const axios = require("axios").default

async function requestAllGames() {
  try {
    const res = await axios.get(
      "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
    )
    const apps = res.data.applist.apps
    apps.sort((first, second) => first.name.localeCompare(second.name))
    return apps
  } catch (error) {
    console.log(error)
  }
}

module.exports = requestAllGames
