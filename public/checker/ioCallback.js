const { requestAllGames, requestUserLibrary } = require("../steamapi")
const processFile = require("../io")
const includes = require("./includes")
const { dialog } = require("electron")
const fs = require("fs")
const { returnError } = require("../ipc").returnError

async function ioCallback(event, args) {
  const csvFilePath = args[0]
  const uid = args[1]
  const apiKey = args[2]

  const playniteLibrariesPromise = processFile(csvFilePath, event)
  const wholeSteamLibraryPromise = requestAllGames(event)
  const userLibraryPromise = requestUserLibrary(uid, apiKey, event)

  const [playniteLibraries, wholeSteamLibrary, userLibrary] = await Promise.all(
    [playniteLibrariesPromise, wholeSteamLibraryPromise, userLibraryPromise]
  )

  if (playniteLibraries && wholeSteamLibrary && userLibrary) {
    console.log(typeof playniteLibraries)
    const res = playniteLibraries.filter(el => {
      return includes(wholeSteamLibrary, el) && !includes(userLibrary, el)
    })

    try {
      const { filePath, canceled } = await dialog.showSaveDialog({
        defaultPath: "output.txt",
      })
      if (filePath && !canceled) {
        const stream = fs.createWriteStream(filePath)
        res.forEach(function (v) {
          stream.write(v + "\n")
        })
        stream.end()
      }
    } catch (error) {
      console.error(error)
      returnError(error)
    }
  } else {
    console.error("at least one of the promises failed")
  }
}

module.exports = ioCallback
