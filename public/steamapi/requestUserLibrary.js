const axios = require("axios").default
const returnError = require("../ipc").returnError

async function requestUserLibrary(uid, apiKey, event) {
  try {
    const res = await axios.get(
      "https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/",
      {
        params: {
          steamid: uid,
          key: apiKey,
          include_appinfo: true,
          include_played_free_games: true,
        },
      }
    )
    const games = res.data.response.games
    games.sort((first, second) => first.name.localeCompare(second.name))
    return games
  } catch (error) {
    console.error(error)
    if (error.response.status === 401)
      returnError(event, "The provided API key is invalid")
    else returnError(event, "The provided Steam user ID is invalid")
  }
}

module.exports = requestUserLibrary
