import React from "react";

function VolumeInput(props) {
  const onChange = (e) => {
    const volume = parseFloat(e.target.value);
    props.onUpdateVolume(volume)
  };

  return <div>
    <label htmlFor="input">
      Volume {props.volume}
      <input type="range" min="0" max="10" value={props.volume}
             onChange={onChange}/>
    </label>
  </div>
}

export default VolumeInput;