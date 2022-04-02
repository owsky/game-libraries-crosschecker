import { Alert } from "../alert"

import "./go.css"

export function Go() {
  return (
    <div id="submit">
      <button className="submit-button" onClick={sendToMain}>
        crosscheck
      </button>
    </div>
  )
}

function validateInput(dialogOutput, uid, apiKey) {
  if (dialogOutput && dialogOutput.type === "text/csv") {
    if (uid.length === 0) {
      Alert({ text: "Please insert your Steam user ID" })
      return false
    } else if (apiKey.length === 0) {
      Alert({ text: "Please insert your Steam API key" })
      return false
    }
  } else {
    Alert({ text: "Please select a CSV file" })
    return false
  }
  return true
}

function sendToMain() {
  const dialogOutput = document.querySelector("#csv-file").files[0]
  const uid = document.querySelector("#steam-userid").value
  const apiKey = document.querySelector("#steam-api-key").value
  if (validateInput(dialogOutput, uid, apiKey)) {
    const csvFilePath = dialogOutput.path
    window.api.sendInput([csvFilePath, uid, apiKey])
  }
}
