import {appendToList, deleteByIndex} from "../../util/util";

const ACTIONS = {
  INIT_STATE: 'initState',
  ADD_WORD_PAIR: 'addWordPair',
  REMOVE_WORD_PAIR: 'removeWordPair',
  UPDATE_USERNAME: 'updateUsername',
  UPDATE_LANGUAGE: 'updateLanguage',
  UPDATE_LANGUAGE_DETECTOR_ENABLED: 'languageDetectorEnabled',
  UPDATE_VOLUME: 'updateVolume',
  UPDATE_IGNORE_WORDS: 'ignoreWords',
  ADD_USER_TO_BAN_LIST: 'addUserToBanList',
  REMOVE_USER_TO_BAN_LIST: 'removeUserFromBanList',
  ADD_IGNORE_WORD: 'addIgnoreWord',
  REMOVE_IGNORE_WORD: 'removeIgnoreWord',
}
const isUniqueValue = (list, value) => {
  return !list.includes(value.toLowerCase())
}
const reducer = (state, action) => {
  const payload = action.payload;
  switch (action.type) {
    case ACTIONS.INIT_STATE: {
      return payload
    }
    case ACTIONS.ADD_WORD_PAIR: {
      const wordPair = [...state.ReplacementWordPair]
      const pair = payload
      const filter = wordPair.find(item => item.before.toLowerCase() === pair.before)
      if (filter !== undefined) {
        return state
      }
      wordPair.push(pair)
      return {
        ...state,
        ReplacementWordPair: wordPair
      }
    }
    case ACTIONS.REMOVE_WORD_PAIR: {
      const pairs = [...state.ReplacementWordPair]

      return {
        ...state,
        ReplacementWordPair: deleteByIndex(pairs, payload),
      }
    }
    case  ACTIONS.UPDATE_USERNAME:
      return {
        ...state,
        ChannelsToListen: payload
      }
    case ACTIONS.UPDATE_IGNORE_WORDS: {
      return {
        ...state,
        IgnoreWords: payload
      }
    }
    case  ACTIONS.UPDATE_LANGUAGE: {
      return {
        ...state,
        Language: payload
      }
    }
    case ACTIONS.UPDATE_LANGUAGE_DETECTOR_ENABLED: {
      return {
        ...state,
        LanguageDetectorEnabled: payload
      }
    }
    case ACTIONS.ADD_USER_TO_BAN_LIST: {
      if (!isUniqueValue(state.UserBanList, payload)) {
        return state
      }
      return {
        ...state,
        UserBanList: appendToList(state.UserBanList, payload)
      };
    }
    case ACTIONS.REMOVE_USER_TO_BAN_LIST: {
      return {
        ...state,
        UserBanList: deleteByIndex(state.UserBanList, payload)
      }
    }
    case ACTIONS.ADD_IGNORE_WORD: {
      if (!isUniqueValue(state.UserBanList, payload)) {
        return
      }
      return {
        ...state,
        IgnoreWords: appendToList(state.IgnoreWords, payload)
      }
    }
    case ACTIONS.REMOVE_IGNORE_WORD: {
      return {
        ...state,
        IgnoreWords: deleteByIndex(state.IgnoreWords, payload)
      }
    }
    case ACTIONS.UPDATE_VOLUME: {
      return {
        ...state,
        Volume: payload
      }
    }
    default: {
      return state
    }
  }
}

export {reducer, ACTIONS}