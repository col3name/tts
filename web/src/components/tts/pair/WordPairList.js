import './word.css';
import '../scrollbar/style.css';
export default function WordPairList(props) {
  const list = props.pairList;
  return (
    <div className="lists">
      {list.map((pair, index) => {
        return <div className="lists__item" key={index}><span>{pair.before} => {pair.after}</span>
          <button onClick={() => props.onDeletePair(index)}>X</button>
        </div>;
      })}
    </div>
  )
}