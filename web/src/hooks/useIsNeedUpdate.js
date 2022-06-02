import {useCallback, useEffect, useState} from "react";

function useIsNeedUpdate(callback) {
  const [isNeedUpdate, setIsNeedUpdate] = useState(false)

  const effect = useCallback(() => {
    if (isNeedUpdate) {
      callback()
      setIsNeedUpdate(false);
    }
  }, [isNeedUpdate, callback]);
  useEffect(effect, [effect]);

  const isNeedUpdateWrapper = (handler) => {
    handler();
    setIsNeedUpdate(true);
  };

  return {
    isNeedUpdateWrapper,
  }
}

export default useIsNeedUpdate;