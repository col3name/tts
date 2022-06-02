import axios from "axios";
import {listPairToString, stringToArray, stringToListPair} from "../util/util";

const url = 'http://localhost:8000/api/v1/settings';

const getSettings = async () => {
  const resp = await axios.get(url);
  if (resp.status !== 200) {
    return []
  }
  const setting = resp.data;

  return {
    ChannelsToListen: setting.ChannelsToListen,
    IgnoreWords: stringToArray(setting.IgnoreWords),
    Language: setting.Language,
    LanguageDetectorEnabled: setting.LanguageDetectorEnabled,
    ReplacementWordPair: stringToListPair(setting.ReplacementWordPair),
    UserBanList: stringToArray(setting.UserBanList),
    Volume: setting.Volume,
  };
}

const saveSettingsReq = async (data) => {
  const setting = {
    Id: 1,
    ReplacementWordPair: listPairToString(data.ReplacementWordPair),
    IgnoreWords: data.IgnoreWords.join(),
    Language: data.Language,
    LanguageDetectorEnabled: data.LanguageDetectorEnabled,
    UserBanList: data.UserBanList.join(),
    ChannelsToListen: data.ChannelsToListen,
    Volume: data.Volume,
  };
  return await axios.post(url, JSON.stringify(setting));
}

const saveSettings = (setting) => {
  saveSettingsReq(setting).then(resp => {
    if (resp.status === 200) {
      console.log(resp);
    }
  }).catch(err => [
    console.log(err)
  ])
};


export {getSettings, saveSettings}