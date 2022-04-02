import Logo18 from "./img/Logo18.svg"
import "./steam.css"

export function Steam() {
  return (
    <section id="steam">
      <ul id="wrapper">
        <li className="steamrow">
          <label htmlFor="steam-userid">Your Steam user ID</label>
          <input type="text" id="steam-userid" />
        </li>
        <li className="steamrow">
          <label htmlFor="steam-api-key">Your Steam API key</label>
          <input type="text" id="steam-api-key" />
        </li>
      </ul>
      <img src={Logo18} alt="random logo" id="random-logo" />
    </section>
  )
}
