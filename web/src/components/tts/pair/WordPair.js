import React from "react"
import WordPairForm from "./WordPairForm"
import WordPairList from "./WordPairList"

const WordPair = ({title, subtitle, wordPairList, onUpdatePairs}) => {
  const onSubmitWordPair = (pair) => {
    if (pair.before.length < 1) {
      return
    }
    const wordPair = [...wordPairList]
    const filter = wordPair.find(item => item.before.toLowerCase() === pair.before)
    if (filter !== undefined) {
      return
    }
    wordPair.push(pair)
    onUpdatePairs(wordPair)
  }

  const onRemovePair = (index) => {
    const pairs = [...wordPairList]
    pairs.splice(index, 1)
    onUpdatePairs(pairs)
  }

  return (
    <div>
      <p>{title}</p>
      <WordPairForm onSubmitWordPair={onSubmitWordPair}/>
      <div>
        <p>{subtitle}: {wordPairList.length}</p>
        <WordPairList
          pairList={wordPairList}
          onDeletePair={onRemovePair}
        />
      </div>
    </div>
  )
}

export default WordPair