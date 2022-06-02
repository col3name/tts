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
  const {setIsNeedUpdate} = useIsNeedUpdate(updateCallback);

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
    setChannelsToListen(username);
    setIsNeedUpdate(true);
  };
  const onUpdateIgnoreWordList = (list) => {
    setIgnoreWords(list);
    setIsNeedUpdate(true);
  };

  const onSelectLanguage = (lang) => {
    setLanguage(lang);
    setIsNeedUpdate(true);
  };
  const onAutoDetectEnabled = (isEnabled) => {
    setLanguageDetectorEnabled(isEnabled);
    setIsNeedUpdate(true);
  };
  const onUpdateReplacementWordPairs = (pairs) => {
    setReplacementWordPair(pairs)
    setIsNeedUpdate(true);
  };
  const onUpdateUserBanList = (list) => {
    setUserBanList(list);
    setIsNeedUpdate(true);
  };
  const onUpdateVolume = volume => {
    setVolume(volume)
    setIsNeedUpdate(true);
  };

  return {
    replacementWordPair,
    ignoreWords,
    language,
    languageDetectorEnabled,
    userBanList,
    channelsToListen,
    volume,
    isLoading,
    onUpdateTwitchUsername,
    onUpdateIgnoreWordList,
    onSelectLanguage,
    onAutoDetectEnabled,
    onUpdateReplacementWordPairs,
    onUpdateUserBanList,
    onUpdateVolume,
  }
}

export default useSettingState