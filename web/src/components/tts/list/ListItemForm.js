import {useState} from "react"

const ListItemForm = ({label, onSubmit}) => {
  const [value, setValue] = useState('')

  return (
    <form onSubmit={(e) => {
      e.preventDefault()
      onSubmit(value)
      setValue('')
    }}>
      <label>
        {label}
        <input type="text" placeholder={label} value={value} onChange={e => setValue(e.target.value)}/>
      </label>
      <button type="submit">Save</button>
    </form>
  )
}

export default ListItemForm