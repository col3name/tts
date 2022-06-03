export default function ListItemForm({label, onSubmit}) {
  return (
    <form onSubmit={(e) => {
      e.preventDefault();
      onSubmit(e.target[0].value)
      e.target[0].value = '';
    }}>
      <label>
        {label}
        <input type="text" placeholder={label}/>
      </label>
      <button type="submit">Save</button>
    </form>
  )
};