import {useEffect, useState} from "react"
import {getSettings as getSettingsReq} from "../api"

const useGetSettings = () => {
  const [isLoading, setIsLoading] = useState(true)

  const [data, setData] = useState({
    ChannelsToListen: '',
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
      const data = await getSettingsReq();
      const find = data.UserBanList.find((user) => user === data.ChannelsToListen)
      if (find === undefined) {
        data.UserBanList.push(data.ChannelsToListen)
      }
      setData(data);
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