import playniteVector from "./img/playnite.svg"
import "./playnite.css"

export function Playnite() {
  return (
    <section id="playnite">
      <img src={playniteVector} alt="playnite logo" id="playnite-logo" />
      <div>
        <h3>Select your Playnite library</h3>
        <input type="file" id="csv-file" accept=".csv" />
      </div>
    </section>
  )
}
