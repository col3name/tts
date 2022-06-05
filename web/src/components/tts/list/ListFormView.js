import ListItemForm from "./ListItemForm"
import ListView from "./ListView"
import React from "react"

const ListFormView = ({title, label, list, minLength, onAddItem, onRemoveItem}) => {
  const onDeleteItem = (index) => onRemoveItem(index)

  const onCreateItem = (item) => {
    if (item.length >= minLength && !list.includes(item.toLowerCase())) {
      onAddItem(item)
    }
  }

  return (
    <div>
      <p>{title}: {list === undefined ? '' : list.length}</p>
      <ListItemForm label={label}
                    onSubmit={onCreateItem}/>
      <ListView data={list}
                onDeletePair={onDeleteItem}
      />
    </div>
  )
}

export default ListFormView