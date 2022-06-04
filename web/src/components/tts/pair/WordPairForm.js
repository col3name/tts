import React, {useRef, useState} from 'react'

const WordPairForm = ({onSubmitWordPair}) => {
  const beforeInputEl = useRef(null)

  const [before, setBefore] = useState("")
  const [after, setAfter] = useState("")

  const onSubmit = (e) => {
    e.preventDefault()
    console.log('before', before)
    console.log('after', after)
    if (before.length < 2) {
      return
    }
    onSubmitWordPair({
      before: before,
      after: after
    })
    setBefore('')
    setAfter('')
    beforeInputEl.current.focus()
  }

  return (
    <form onSubmit={onSubmit}>
      <div>
        <label className="label">
          <span>Before (word to replace)</span>
          <input
            ref={beforeInputEl}
            value={before}
            onChange={e => setBefore(e.target.value)}
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
            value={after}
            onChange={e => setAfter(e.target.value)}
          />
        </label>
      </div>
      <button type="submit">Save</button>
    </form>
  )
}

export default WordPairForm