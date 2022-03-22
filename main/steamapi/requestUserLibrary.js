const axios = require("axios").default

async function requestUserLibrary(uid, apiKey) {
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
    console.log(error)
  }
}

module.exports = requestUserLibrary
