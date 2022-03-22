function sendResults(event, args) {
  event.sender.send("output", args)
}

module.exports = sendResults
