import {useCallback, useState} from "react";
import useGetSettings from "./useGetSettings";
import useIsNeedUpdate from "./useIsNeedUpdate";
import {saveSettings} from "../api";

function useSettingState() {
  const [channelsToListen, setChannelsToListen] = useState('');
  const [ignoreWords, setIgnoreWords] = useState([]);
  const [language, setLanguage] = useState('en');
  const [languageDetectorEnabled, setLanguageDetectorEnabled] = useState(false);
  const [replacementWordPair, setReplacementWordPair] = useState([]);
  const [userBanList, setUserBanList] = useState([]);
  const [volume, setVolume] = useState(5);

  const updateCallback = useCallback(() => {
    const setting = {
      Id: 1,
      ReplacementWordPair: replacementWordPair,
      IgnoreWords: ignoreWords,
      Language: language,
      LanguageDetectorEnabled: languageDetectorEnabled,
      UserBanList: userBanList,
      ChannelsToListen: channelsToListen,
      Volume: volume,
    };

    saveSettings(setting)
  }, [
    replacementWordPair,
    ignoreWords,
    language,
    languageDetectorEnabled,
    userBanList,
    channelsToListen,
    volume,
  ]);
  const {isNeedUpdateWrapper} = useIsNeedUpdate(updateCallback);

  const {isLoading} = useGetSettings((setting) => {
    setChannelsToListen(setting.ChannelsToListen);
    setIgnoreWords(setting.IgnoreWords);
    setLanguage(setting.Language);
    setLanguageDetectorEnabled(setting.LanguageDetectorEnabled);
    setReplacementWordPair(setting.ReplacementWordPair);
    setUserBanList(setting.UserBanList);
    setVolume(setting.Volume);
  });

  const onUpdateTwitchUsername = (username) => {
    isNeedUpdateWrapper(() => setChannelsToListen(username));
  };
  const onUpdateIgnoreWordList = (list) => {
    isNeedUpdateWrapper(() => setIgnoreWords(list));
  };
  const onUpdateLanguage = (lang) => {
    isNeedUpdateWrapper(() => setLanguage(lang));
  };
  const onUpdateAutoDetectEnabled = (isEnabled) => {
    isNeedUpdateWrapper(() => setLanguageDetectorEnabled(isEnabled));
  };
  const onUpdateReplacementWordPairs = (pairs) => {
    isNeedUpdateWrapper(() => setReplacementWordPair(pairs));
  };
  const onUpdateUserBanList = (list) => {
    isNeedUpdateWrapper(() => setUserBanList(list));
  };
  const onUpdateVolume = volume => {
    isNeedUpdateWrapper(() => setVolume(volume));
  };

  return {
    channelsToListen: {value: channelsToListen, onUpdate: onUpdateTwitchUsername},
    ignoreWords: {value: ignoreWords, onUpdate: onUpdateIgnoreWordList},
    language: {value: language, onUpdate: onUpdateLanguage},
    languageDetectorEnabled: {value: languageDetectorEnabled, onUpdate: onUpdateAutoDetectEnabled},
    replacementWordPair: {value: replacementWordPair, onUpdate: onUpdateReplacementWordPairs},
    userBanList: {value: userBanList, onUpdate: onUpdateUserBanList},
    volume: {value: volume, onUpdate: onUpdateVolume},
    isLoading,
  }
}

export default useSettingState