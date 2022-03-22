const { requestAllGames, requestUserLibrary } = require("../steamapi")
const processFile = require("../io")
const includes = require("./includes")
const { dialog } = require("electron")
const fs = require("fs")

async function ioCallback(event, args) {
  const csvFilePath = args[0]
  const uid = args[1]
  const apiKey = args[2]

  try {
    const playniteLibraries = await processFile(csvFilePath)
    const wholeSteamLibrary = await requestAllGames()
    const userLibrary = await requestUserLibrary(uid, apiKey)

    let res = playniteLibraries.filter(el => {
      return includes(wholeSteamLibrary, el) && !includes(userLibrary, el)
    })

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
  }
}

module.exports = ioCallback
