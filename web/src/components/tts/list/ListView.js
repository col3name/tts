export default function ListView(props) {
  return (
    <>
      {props.data.map((item, index) => {
        return <div key={index}>{item}
          <button onClick={() => props.onDeletePair(index)}>X</button>
        </div>;
      })}
    </>
  )
}