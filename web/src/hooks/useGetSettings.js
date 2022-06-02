import {useEffect, useState} from "react";
import {getSettings as getSettingsReq} from "../api";
import {delay} from "../util/util";

function useGetSettings(callback) {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    let mounted = true;

    getSettingsReq()
      .then(data => {
        if (mounted) {
          callback(data);

          delay(2000).then();

          setIsLoading(false);
        }
      })
    return () => mounted = false;
  }, []);

  return {
    isLoading
  }
}

export default useGetSettings;