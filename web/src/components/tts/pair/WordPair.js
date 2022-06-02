import WordPairForm from "./WordPairForm";
import WordPairList from "./WordPairList";
import React from "react";

function WordPair(props) {
  const {title, subtitle, wordPairList, onUpdatePairs} = props;

  const onSubmitWordPair = (pair) => {
    if (pair.before.length < 1) {
      return
    }
    const wordPair = [...wordPairList];
    const filter = wordPair.find(item => item.before.toLowerCase() === pair.before);
    if (filter !== undefined) {
      return;
    }
    wordPair.push(pair);
    onUpdatePairs(wordPair);
  }

  const onRemovePair = (index) => {
    const pairs = [...wordPairList];
    pairs.splice(index, 1);
    onUpdatePairs(pairs);
  };

  return (
    <div>
      <p>{title}</p>
      <WordPairForm onSubmitWordPair={onSubmitWordPair}/>
      <div>
        <p>{subtitle}</p>
        <WordPairList
          pairList={wordPairList}
          onDeletePair={onRemovePair}
        />
      </div>
    </div>
  )
}

export default WordPair;