import './lists.css';
import '../scrollbar/style.css';

export default function ListView(props) {
  return (
    <div className="lists">
      {props.data.map((item, index) => {
        return <div className="lists__item" key={index}>{item}
          <button onClick={() => props.onDeletePair(index)}>X</button>
        </div>;
      })}
    </div>
  )
}