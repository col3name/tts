import {useEffect, useReducer, useRef} from "react"
import useGetSetting from "./useGetSetting"
import {ACTIONS, reducer} from "./reducer";
import {INITIAL_STATE} from "./default";
import {saveSettings} from "../../api";

const useSettingState = () => {
  const [state, dispatch] = useReducer(reducer, INITIAL_STATE)

  const isFirstRef = useRef(true);
  const {isLoading, data, error} = useGetSetting()

  useEffect(() => {
    if (isLoading || error) {
      return
    }
    dispatch({type: ACTIONS.INIT_STATE, payload: data})
  }, [isLoading, data, error])

  useEffect(() => {
    if (isFirstRef.current) {
      isFirstRef.current = false;
      return;
    }
    saveSettings(state)
  }, [state])

  return {
    channelsToListen: {
      value: state.ChannelsToListen,
      onUpdate: (username) => {
        dispatch({type: ACTIONS.UPDATE_USERNAME, payload: username})
      }
    },
    ignoreWords: {
      value: state.IgnoreWords,
      onAddWord: (word) => {
        dispatch({type: ACTIONS.ADD_IGNORE_WORD, payload: word})
      },
      onRemoveWord(word) {
        dispatch({type: ACTIONS.REMOVE_IGNORE_WORD, payload: word})
      }
    },
    language: {
      value: state.Language, onUpdate: (lang) => {
        dispatch({type: ACTIONS.UPDATE_LANGUAGE, payload: lang})
      }
    },
    languageDetectorEnabled: {
      value: state.LanguageDetectorEnabled,
      onUpdate: (isEnabled) => {
        dispatch({type: ACTIONS.UPDATE_LANGUAGE_DETECTOR_ENABLED, payload: isEnabled})
      }
    },
    replacementWordPair: {
      value: state.ReplacementWordPair,
      onAddWordPair: (pair) => {
        dispatch({type: ACTIONS.ADD_WORD_PAIR, payload: pair})
      },
      onRemovePair: (index) => {
        dispatch({type: ACTIONS.REMOVE_WORD_PAIR, payload: index})
      },
    },
    userBanList: {
      value: state.UserBanList,
      onAddItem: (user) => {
        dispatch({type: ACTIONS.ADD_USER_TO_BAN_LIST, payload: user})
      },
      onRemoveItem: (list) => {
        dispatch({type: ACTIONS.REMOVE_USER_TO_BAN_LIST, payload: list})
      },
    },
    volume: {
      value: state.Volume,
      onUpdate: (volume) => {
        dispatch({type: ACTIONS.UPDATE_VOLUME, payload: volume})
      }
    },
    isLoading,
  }
}

export default useSettingState