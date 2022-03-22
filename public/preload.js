const { contextBridge, ipcRenderer } = require("electron")

const api = {
  sendInput: args => {
    ipcRenderer.send("input", args)
  },
  listenForOutput: callback => {
    ipcRenderer.addListener("output", callback)
  },
}

process.once("loaded", () => {
  contextBridge.exposeInMainWorld("api", api)
})
