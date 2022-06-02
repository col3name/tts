import WordPairForm from "./WordPairForm";
import WordPairList from "./WordPairList";
import React from "react";

function WordPair(props) {
  const onSubmitWordPair = (pair) => {
    if (pair.before.length < 1) {
      return
    }
    const wordPair = [...props.wordPairList];
    const filter = wordPair.find(item => item.before.toLowerCase() === pair.before);
    if (filter !== undefined) {
      return;
    }
    wordPair.push(pair);
    props.onUpdatePairs(wordPair);
  }

  const onRemovePair = (index) => {
    const pairs = [...props.wordPairList];
    pairs.splice(index, 1);
    props.onUpdatePairs(pairs);
  };

  return (
    <div>
      <p>{props.title}</p>
      <WordPairForm onSubmitWordPair={onSubmitWordPair}/>
      <div>
        <p>{props.subtitle}</p>
        <WordPairList
          pairList={props.wordPairList}
          onDeletePair={onRemovePair}
        />
      </div>
    </div>
  )
}

export default WordPair;