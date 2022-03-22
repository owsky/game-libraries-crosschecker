import { Playnite } from "./components/playnite"
import { Steam } from "./components/steam"
import { Go } from "./components/go"
import { Attribution } from "./components/attribution"
import "./app.css"

export function App() {
  return (
    <>
      <main>
        <Playnite />
        <Steam />
        <Go />
      </main>
      <Attribution />
    </>
  )
}
