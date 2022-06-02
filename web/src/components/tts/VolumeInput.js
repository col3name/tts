import React from "react";

function VolumeInput(props) {
  const {volume, onUpdateVolume} = props;
  const onChange = (e) => {
    onUpdateVolume(e.target.valueAsNumber)
  };

  return <div>
    <label>
      Volume {volume}
      <input type="range" min="0" max="10" value={volume}
             onChange={onChange}/>
    </label>
  </div>
}

export default VolumeInput;