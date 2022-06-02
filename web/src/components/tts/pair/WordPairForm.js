import React, {useRef} from 'react';

export default function WordPairForm(props) {
  const {onSubmitWordPair} = props;
  const beforeInputEl = useRef(null);

  const onSubmit = (e) => {
    e.preventDefault();
    const beforeElement = e.target[0];
    const afterElement = e.target[1];

    const before = beforeElement.value;
    const after = afterElement.value;
    if (before.length < 2) {
      return
    }
    onSubmitWordPair({
      before: before,
      after: after
    })
    beforeElement.value = '';
    afterElement.value = '';
    beforeInputEl.current.focus();
  };

  return (
    <form onSubmit={onSubmit}>
      <div>
        <label className="label">
          <span>Before (word to replace)</span>
          <input
            ref={beforeInputEl}
            type="text"
            placeholder="before"
          />
        </label>
      </div>
      <div>
        <label className="label">
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