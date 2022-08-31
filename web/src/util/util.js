const appendToList = (list, value) => {
  const temp = [...list]
  temp.push(value)
  return temp
}

const deleteByIndex = (list, index) => {
  if (index < 0 || index >= list.length) {
    return list
  }
  const tempList = [...list]
  tempList.splice(index, 1)
  return tempList
}

const deleteLastSymbols = (str, count = 1) => {
  return str.substring(0, str.length - count)
}

const isExistInPairList = (pairList, pair) => {
  return pairList.find(item => item.before.toLowerCase() === pair.before) !== undefined
}

const stringToArray = (string) => {
  return string.includes(',') ? string.split(',').filter(item => item.length > 0) : []
}

const listPairToString = (listPair) => {
  let wordPair = listPair.reduce((result, pair, index) => {
    if (pair.before.length === 0) {
      return result
    }
    const ok = index !== listPair.length - 1
    const str = ok ? ',' : ''
    return result + pair.before + ':' + pair.after + str
  }, '')
  while (wordPair[wordPair.length - 1] === ',') {
    wordPair = deleteLastSymbols(wordPair)
  }
  return wordPair
}

const stringToListPair = (string) => {
  while (string[string.length - 1] === ',') {
    string = deleteLastSymbols(string)
  }
  return string.split(',').map(pair => {
    const res = pair.split(':')
    return {
      before: res[0],
      after: res[1],
    }
  })
}

export {
  appendToList,
  isExistInPairList,
  deleteByIndex,
  deleteLastSymbols,
  stringToArray,
  listPairToString,
  stringToListPair,
}