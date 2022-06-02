import React from 'react';

export default function WordPairForm(props) {
  const onSubmit = (e) => {
    e.preventDefault();
    props.onSubmitWordPair({
      before: e.target[0].value,
      after: e.target[1].value
    })
    e.target[0].value = '';
    e.target[1].value = '';
  };

  return (
    <form onSubmit={onSubmit}>
      <div>
        <label htmlFor="input">
          Before (word to replace)
          <input
            type="text"
            placeholder="before"
          />
        </label>
      </div>
      <div>
        <label htmlFor="input">
          After (word to replace)
          <input
            type="text"
            placeholder="after"
          />
        </label>
      </div>
      <button type="submit">Save</button>
    </form>
  )
};