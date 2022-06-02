import './word.css';
import '../scrollbar/style.css';

export default function WordPairList(props) {
  const {pairList, onDeletePair} = props;

  return (
    <div className="lists">
      {pairList.map((pair, index) => {
        return <div className="lists__item" key={index}>
          <span>{pair.before} => {pair.after}</span>
          <button onClick={() => onDeletePair(index)}>X</button>
        </div>;
      })}
    </div>
  )
}