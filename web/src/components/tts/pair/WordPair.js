import React from "react"
import WordPairForm from "./WordPairForm"
import WordPairList from "./WordPairList"
import {isExistInPairList} from "../../../util/util";

const WordPair = ({title, subtitle, wordPairList, onAddWordPair, onRemovePair}) => {
  return (
    <div>
      <p>{title}</p>
      <WordPairForm onSubmitWordPair={(pair) => {
        if (pair.before.length < 1 || isExistInPairList(wordPairList, pair)) {
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