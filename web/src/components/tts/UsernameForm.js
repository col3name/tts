import React, {useEffect, useState} from "react"

const UsernameForm = ({channelsToListen, onUpdateTwitchUsername}) => {
  const [localUsername, setLocalUsername] = useState(channelsToListen)

  const onSubmit = (e) => {
    e.preventDefault()
    onUpdateTwitchUsername(localUsername)
  }

  useEffect(() => {
    setLocalUsername(channelsToListen)
  }, [channelsToListen])

  return <form onSubmit={onSubmit}>
    <label>
      <p>Twitch username</p>
      <input type="text" value={localUsername} onChange={e => setLocalUsername(e.target.value)}/>
    </label>
    <button type="submit">Watch</button>
    <p>Currently watching: {channelsToListen}</p>
  </form>
}

export default UsernameForm;