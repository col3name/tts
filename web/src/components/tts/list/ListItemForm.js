export default function ListItemForm(props) {
  return (
    <form onSubmit={(e) => {
      e.preventDefault();
      props.onSubmit(e.target[0].value)
      e.target[0].value = '';
    }}>
      <label>
        {props.label}
        <input type="text" placeholder={props.label}/>
      </label>
      <button type="submit">Save</button>
    </form>
  )
};