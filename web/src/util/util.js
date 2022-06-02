const appendToList = (value, list, callback) => {
  const temp = [...list];
  temp.push(value)
  callback(temp);
}

const deleteByIndex = (index, list, callback) => {
  const tempList = [...list];
  tempList.splice(index, 1);
  callback(tempList);
}

const deleteLastSymbols = (str, count = 1) => {
  return str.substring(0, str.length - count)
}

const arrayToString = (arr, delimiter = ',') => {
  const reduce = arr.reduce((result, item) => {
    if (item.length === 0) {
      return result;
    }
    return result + item + delimiter;
  }, '');
  return reduce.substring(0, reduce.length - 1)
}

const stringToArray = (string) => {
  return string.includes(',') ? string.split(',').filter(item => item.length > 0) : []
}

const listPairToString = (listPair) => {
  let wordPair = listPair.reduce((result, pair) => {
    if (pair.before.length === 0) {
      return result
    }
    return result + pair.before + ':' + pair.after + ','
  }, '');
  while (wordPair[wordPair.length - 1] === ',') {
    wordPair = deleteLastSymbols(wordPair)
  }
  return wordPair;
}

const stringToListPair = (string) => {
  while (string[string.length - 1] === ',') {
    string = deleteLastSymbols(string)
  }
  return string.split(',').map(pair => {
    const split = pair.split(':');
    return {
      before: split[0],
      after: split[1],
    }
  });
}

const delay = ms => new Promise(res => setTimeout(res, ms));

export {
  appendToList,
  deleteByIndex,
  delay,
  arrayToString,
  deleteLastSymbols,
  stringToArray,
  listPairToString,
  stringToListPair,
};