const { contextBridge, ipcRenderer } = require("electron")

const api = {
  sendInput: args => {
    ipcRenderer.send("input", args)
  },
  listenForError: callback => {
    ipcRenderer.addListener("error", callback)
  },
}

process.once("loaded", () => {
  contextBridge.exposeInMainWorld("api", api)
})
