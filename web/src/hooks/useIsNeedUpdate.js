import {useCallback, useEffect, useState} from "react"

const useIsNeedUpdate = (callback) => {
  const [isNeedUpdate, setIsNeedUpdate] = useState(false)

  const effect = useCallback(() => {
    if (isNeedUpdate) {
      callback()
      setIsNeedUpdate(false)
    }
  }, [isNeedUpdate, callback])
  useEffect(effect, [effect])

  const needUpdate = (handler) => {
    handler()
    setIsNeedUpdate(true)
  }

  return {
    needUpdate,
  }
}

export default useIsNeedUpdate