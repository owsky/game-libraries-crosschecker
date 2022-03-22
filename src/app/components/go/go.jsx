import "./go.css"

export function Go() {
  return (
    <div id="submit">
      <button className="submit-button" onClick={sendToMain}>
        Go
      </button>
    </div>
  )
}

function sendToMain() {
  const csvFilePath = document.querySelector("#csv-file").files[0].path
  const uid = document.querySelector("#steam-userid").value
  const apiKey = document.querySelector("#steam-api-key").value
  window.api.sendInput([csvFilePath, uid, apiKey])
}
