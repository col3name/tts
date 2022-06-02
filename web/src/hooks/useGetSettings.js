import {useCallback, useEffect, useState} from "react";
import {getSettings as getSettingsReq} from "../api";

function useGetSettings(callback) {
  const [isLoading, setIsLoading] = useState(true);

  const effect = useCallback(() => {
    const fetchData = async () => {
      const data = await getSettingsReq();
      callback(data);
      setIsLoading(false);
    }
    fetchData();
  }, [callback]);
  useEffect(effect, []);

  return {
    isLoading
  }
}

export default useGetSettings;