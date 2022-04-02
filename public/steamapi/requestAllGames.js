const axios = require("axios").default
const returnError = require("../ipc").returnError

async function requestAllGames(event) {
  try {
    const res = await axios.get(
      "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
    )
    const apps = res.data.applist.apps
    apps.sort((first, second) => first.name.localeCompare(second.name))
    return apps
  } catch (error) {
    console.error(error)
    returnError(
      event,
      `Failed to request the entire Steam library: \n${error.toString()}`
    )
  }
}

module.exports = requestAllGames
