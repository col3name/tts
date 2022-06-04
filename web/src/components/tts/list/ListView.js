import './lists.css'
import '../scrollbar/style.css'

const ListView = ({data, onDeletePair}) => {
  return (
    <div className="lists">
      {data.map((item, index) => {
        return <div className="lists__item" key={index}>{item}
          <button onClick={() => onDeletePair(index)}>X</button>
        </div>
      })}
    </div>
  )
}

export default ListView