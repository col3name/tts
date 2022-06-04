import {useEffect, useState} from "react"
import {getSettings as getSettingsReq} from "../api"
import {stringToArray, stringToListPair} from "../util/util";

const useGetSettings = () => {
  const [isLoading, setIsLoading] = useState(true)

  const [data, setData] = useState( {
    ChannelsToListen: [],
    IgnoreWords: [],
    Language: '',
    LanguageDetectorEnabled: false,
    ReplacementWordPair: [],
    UserBanList: [],
    Volume: 1
  })
  const [error, setError] = useState(undefined)
  useEffect(() => {
    const fetchData = async () => {
      const data = await getSettingsReq()
      setData(data)
      setIsLoading(false)
    }
    try {
      fetchData()
    } catch (e) {
      setError(e)
    }
  }, [])

  return {
    isLoading,
    data,
    error
  }
}

export default useGetSettings