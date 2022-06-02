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

const listToStr = (str, delimiter = ',') => {
  const reduce = str.reduce((result, item) => {
    if (item.length === 0) {
      return result;
    }
    return result + item + delimiter;
  }, '');
  return reduce.substring(0, reduce.length - 1)
}

const delay = ms => new Promise(res => setTimeout(res, ms));

export {appendToList, deleteByIndex, delay, listToStr, deleteLastSymbols}