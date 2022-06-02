import {useCallback, useEffect, useState} from "react";

function useIsNeedUpdate(callback) {
  const [isNeedSave, setIsNeedUpdate] = useState(false)

  const effect = useCallback(() => {
    if (isNeedSave) {
      callback()
      setIsNeedUpdate(false);
    }
  }, [isNeedSave, callback]);
  useEffect(effect, [effect])

  return {
    setIsNeedUpdate,
  }
}

export default useIsNeedUpdate;