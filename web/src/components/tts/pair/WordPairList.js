import './word.css'
import '../scrollbar/style.css'

const WordPairList = ({pairList, onDeletePair}) => {
  return (
    <div className="lists">
      {pairList.map((pair, index) => {
        return <div className="lists__item" key={index}>
          <span>{pair.before} => {pair.after}</span>
          <button onClick={() => onDeletePair(index)}>X</button>
        </div>
      })}
    </div>
  )
}
export default WordPairList