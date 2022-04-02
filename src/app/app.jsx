import { Playnite } from "./components/playnite"
import { Steam } from "./components/steam"
import { Go } from "./components/go"
import { Alert } from "./components/alert"
import "./app.css"

export function App() {
  window.api.listenForError((_, message) => {
    Alert({ title: "Error", text: message })
  })
  return (
    <main>
      <Playnite />
      <Steam />
      <Go />
    </main>
  )
}
