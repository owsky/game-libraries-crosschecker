function includes(library, game) {
  let left = 0
  let right = library.length - 1

  while (left <= right) {
    const mid = Math.floor((left + right) / 2)
    const comp = library[mid].name.localeCompare(game)
    if (comp === 0) return true
    else if (comp === -1) left = mid + 1
    else right = mid - 1
  }
  return false
}

module.exports = includes
