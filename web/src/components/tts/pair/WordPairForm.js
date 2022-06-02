import React, {useRef} from 'react';

export default function WordPairForm(props) {
  const beforeInputEl = useRef(null);

  const onSubmit = (e) => {
    e.preventDefault();
    const before = e.target[0].value;
    const after = e.target[1].value;
    if (before.length < 2) {
      return
    }
    props.onSubmitWordPair({
      before: before,
      after: after
    })
    e.target[0].value = '';
    e.target[1].value = '';
    beforeInputEl.current.focus();
  };

  return (
    <form onSubmit={onSubmit}>
      <div>
        <label className="label" htmlFor="input">
          <span>Before (word to replace)</span>
          <input
            ref={beforeInputEl}
            type="text"
            placeholder="before"
          />
        </label>
      </div>
      <div>
        <label className="label" htmlFor="input">
          <span>After (word to replace)</span>
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