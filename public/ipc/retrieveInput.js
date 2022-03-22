const { ipcMain } = require("electron")

function retrieveInput(callback) {
  ipcMain.on("input", callback)
}

module.exports = retrieveInput
