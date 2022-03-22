const { app, BrowserWindow } = require("electron")
const path = require("path")
const url = require("url")

function createWindow() {
  const mainWindow = new BrowserWindow({
    minWidth: 800,
    minHeight: 600,
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
    },
    show: false,
    autoHideMenuBar: true,
  })

  mainWindow.once("ready-to-show", () => {
    mainWindow.show()
  })

  if (app.isPackaged) {
    mainWindow.loadURL(
      url.format({
        pathname: path.join(__dirname, "index.html"),
        protocol: "file:",
        slashes: true,
      })
    )
  } else {
    mainWindow.loadURL("http://localhost:3000")
  }

  if (!app.isPackaged) {
    mainWindow.webContents.openDevTools({ mode: "detach" })
  }
}

app.whenReady().then(() => {
  createWindow()

  app.on("activate", function () {
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow()
    }
  })
})

app.on("window-all-closed", function () {
  if (process.platform !== "darwin") {
    app.quit()
  }
})

require(path.join(__dirname, "checker"))
// require()
