const fs = require("fs")
const { parse } = require("csv-parse")
const { finished } = require("stream/promises")

async function processFile(path) {
  const records = []
  const parser = fs.createReadStream(path).pipe(
    parse({
      delimiter: ",",
      bom: true,
      fromLine: 3,
    })
  )
  parser.on("readable", function () {
    let record
    while ((record = parser.read()) !== null) {
      records.push(record)
    }
  })
  parser.on("error", function (err) {
    console.error(err.message)
  })
  try {
    await finished(parser)
  } catch (error) {
    console.error(error)
  }

  return records
    .filter(el => {
      return el[1] !== "Steam"
    })
    .map(value => value[0])
}

module.exports = processFile
