function returnError(event, args) {
  event.sender.send("error", args)
}

module.exports = returnError
