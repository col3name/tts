import React from "react"
import WordPairForm from "./WordPairForm"
import WordPairList from "./WordPairList"

const WordPair = ({title, subtitle, wordPairList, onAddWordPair, onRemovePair}) => {
  const isExist = (pair) => {
    return wordPairList.find(item => item.before.toLowerCase() === pair.before) !== undefined
  }

  return (
    <div>
      <p>{title}</p>
      <WordPairForm onSubmitWordPair={(pair) => {
        if (pair.before.length < 1 || isExist(pair)) {
          return
        }
        onAddWordPair(pair)
      }}/>
      <div>
        <p>{subtitle}: {wordPairList.length}</p>
        <WordPairList
          pairList={wordPairList}
          onDeletePair={index => onRemovePair(index)}
        />
      </div>
    </div>
  )
}

export default WordPair