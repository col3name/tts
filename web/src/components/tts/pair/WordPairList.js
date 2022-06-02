export default function WordPairList(props) {
  const list = props.pairList;
  return (
    <>
      {list.map((pair, index) => {
        return <div key={index}>{pair.before} => {pair.after}
          <button onClick={() => props.onDeletePair(index)}>X</button>
        </div>;
      })}
    </>
  )
}