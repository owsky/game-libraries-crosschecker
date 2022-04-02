import Swal from "sweetalert2"
import withReactContent from "sweetalert2-react-content"

import "./alert.css"

const mySwal = withReactContent(Swal)

export function Alert(props) {
  mySwal.fire({
    title: "Error",
    titleText: props.title,
    text: props.text,
    backdrop: true,
    target: "body",
    confirmButtonColor: "orangered",
  })
}
