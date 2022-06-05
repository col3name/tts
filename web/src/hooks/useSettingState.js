import {useEffect, useState} from "react"
import useGetSettings from "./useGetSettings"
import {saveSettings} from "../api"

const useSettingState = () => {
  const [channelsToListen, setChannelsToListen] = useState('')
  const [ignoreWords, setIgnoreWords] = useState([])
  const [language, setLanguage] = useState('en')
  const [languageDetectorEnabled, setLanguageDetectorEnabled] = useState(false)
  const [replacementWordPair, setReplacementWordPair] = useState([])
  const [userBanList, setUserBanList] = useState([])
  const [volume, setVolume] = useState(5)

  const [isFirst, setIsFirst] = useState(true);

  const save = () => {
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
  };

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

  useEffect(() => {
    if (!isFirst) {
      save()
    } else {
      setIsFirst(false)
    }
  }, [
    replacementWordPair,
    ignoreWords,
    language,
    languageDetectorEnabled,
    userBanList,
    channelsToListen,
    volume,
  ])

  return {
    channelsToListen: {
      value: channelsToListen, onUpdate: (username) => {
        setChannelsToListen(username);
      }
    },
    ignoreWords: {
      value: ignoreWords, onUpdate: (list) => {
        setIgnoreWords(list)
      }
    },
    language: {
      value: language, onUpdate: (lang) => {
        setLanguage(lang)
      }
    },
    languageDetectorEnabled: {
      value: languageDetectorEnabled, onUpdate: (isEnabled) => {
        setLanguageDetectorEnabled(isEnabled)
      }
    },
    replacementWordPair: {
      value: replacementWordPair, onUpdate: (pairs) => {
        setReplacementWordPair(pairs)
      }
    },
    userBanList: {
      value: userBanList, onUpdate: (list) => {
        setUserBanList(list)
      }
    },
    volume: {
      value: volume, onUpdate: volume => {
        setVolume(volume)
      }
    },
    isLoading,
  }
}

export default useSettingState