import ListItemForm from "./ListItemForm"
import ListView from "./ListView"
import React from "react"
import {appendToList, deleteByIndex} from "../../../util/util"

const ListFormView = ({title, label, list, minLength, onUpdate}) => {
  const onDeleteItem = (index) => onUpdate(deleteByIndex(index, list))
  const onCreateItem = (item) => {
    if (item.length < minLength) {
      return
    }
    if (!list.includes(item.toLowerCase())) {
      onUpdate(appendToList(item, list))
    }
  }

  return (
    <div>
      <p>{title}</p>
      <ListItemForm label={label}
                    onSubmit={onCreateItem}/>
      <ListView data={list}
                onDeletePair={onDeleteItem}
      />
    </div>
  )
}

export default ListFormView