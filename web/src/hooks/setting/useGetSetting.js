import {useEffect, useState} from "react"
import {getSettings as getSettingsReq} from "../../api"
import {INITIAL_STATE} from "./default";

const useGetSetting = () => {
  const [isLoading, setIsLoading] = useState(true)

  const [data, setData] = useState(INITIAL_STATE)
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
      setError(e);
      setIsLoading(false)
    }
  }, [])

  return {
    isLoading,
    data,
    error
  }
}

export default useGetSetting