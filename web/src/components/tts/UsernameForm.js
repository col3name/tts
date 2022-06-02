import React, {useEffect, useState} from "react";

function UsernameForm(props) {
  const [localUsername, setLocalUsername] = useState('');

  const onSubmit = (e) => {
    e.preventDefault();
    props.onUpdateTwitchUsername(localUsername)
  };

  useEffect(() => {
    let done = false;
    if (!done) {
      setLocalUsername(props.channelsToListen);
    }
    return () => {
      done = false
    }
  }, [props.channelsToListen])

  return <form onSubmit={onSubmit}>
    <label htmlFor="input">
      <p>Twitch username</p>
      <input type="text" value={localUsername} onChange={e => setLocalUsername(e.target.value)}/>
    </label>
    <button type="submit">Watch</button>
    <p>Currently watching {props.channelsToListen}</p>
  </form>
}

export default UsernameForm;