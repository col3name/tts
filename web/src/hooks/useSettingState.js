import {useCallback, useEffect, useState} from "react"
import useGetSettings from "./useGetSettings"
import {saveSettings} from "../api"
import useIsNeedUpdate from "./useIsNeedUpdate";

const useSettingState = () => {
  const [channelsToListen, setChannelsToListen] = useState('')
  const [ignoreWords, setIgnoreWords] = useState([])
  const [language, setLanguage] = useState('en')
  const [languageDetectorEnabled, setLanguageDetectorEnabled] = useState(false)
  const [replacementWordPair, setReplacementWordPair] = useState([])
  const [userBanList, setUserBanList] = useState([])
  const [volume, setVolume] = useState(5)

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
    }

    saveSettings(setting).then()
  }, [
    replacementWordPair,
    ignoreWords,
    language,
    languageDetectorEnabled,
    userBanList,
    channelsToListen,
    volume,
  ])

  const {needUpdate} = useIsNeedUpdate(updateCallback)

  const {isLoading, data, error} = useGetSettings()

  useEffect(() => {
    setChannelsToListen(data.ChannelsToListen)
    setIgnoreWords(data.IgnoreWords)
    setLanguage(data.Language)
    setLanguageDetectorEnabled(data.LanguageDetectorEnabled)
    setReplacementWordPair(data.ReplacementWordPair)
    setUserBanList(data.UserBanList)
    setVolume(data.Volume)
  }, [data, error])

  const onUpdateTwitchUsername = (username) => {
    needUpdate(() => setChannelsToListen(username))
  }
  const onUpdateIgnoreWordList = (list) => {
    needUpdate(() => setIgnoreWords(list))
  }
  const onUpdateLanguage = (lang) => {
    needUpdate(() => setLanguage(lang))
  }
  const onUpdateAutoDetectEnabled = (isEnabled) => {
    needUpdate(() => setLanguageDetectorEnabled(isEnabled))
  }
  const onUpdateReplacementWordPairs = (pairs) => {
    needUpdate(() => setReplacementWordPair(pairs))
  }
  const onUpdateUserBanList = (list) => {
    console.log(list)
    needUpdate(() => setUserBanList(list))
  }
  const onUpdateVolume = volume => {
    needUpdate(() => setVolume(volume))
  }

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