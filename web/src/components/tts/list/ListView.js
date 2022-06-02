import './lists.css';
import '../scrollbar/style.css';

export default function ListView(props) {
  const {data, onDeletePair} = props;

  return (
    <div className="lists">
      {data.map((item, index) => {
        return <div className="lists__item" key={index}>{item}
          <button onClick={() => onDeletePair(index)}>X</button>
        </div>;
      })}
    </div>
  )
}