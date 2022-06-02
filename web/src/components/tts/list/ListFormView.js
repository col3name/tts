import ListItemForm from "./ListItemForm";
import ListView from "./ListView";
import React from "react";
import {appendToList, deleteByIndex} from "../../../util/util";

export default function ListFormView(props) {
  const onDeleteItem = (index) => deleteByIndex(index, props.list, props.callback);
  const onCreateItem = (item) => {
    if (item.length < props.minLength) {
      return
    }
    if (!props.list.includes(item.toLowerCase())) {
      appendToList(item, props.list, props.callback);
    }
  };

  return (
    <div>
      <p>{props.title}</p>
      <ListItemForm label={props.label}
                    onSubmit={onCreateItem}/>
      <ListView data={props.list}
                onDeletePair={onDeleteItem}
      />
    </div>
  )
};