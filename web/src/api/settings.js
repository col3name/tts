import axios from "axios";

const url = 'http://localhost:8000/api/v1/settings';

async function getSettings() {
  const resp = await axios.get(url);
  if (resp.status !== 200) {
    return []
  }
  return resp.data;
}

async function saveSettings(data) {
  const resp = await axios.post(url, data);
  console.log(resp)
  return resp
}


export {getSettings, saveSettings}