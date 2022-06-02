import React, {useEffect, useState} from "react";
import WordPairList from "./pair/WordPairList";
import WordPairForm from "./pair/WordPairForm";
import * as PropTypes from "prop-types";
import ListView from "./list/ListView";
import ListFormView from "./list/ListFormView";
import {appendToList, delay, deleteByIndex, deleteLastSymbols, listToStr} from "../../util/util";
import Languages from "./languages";
import {getSettings as getSettingsReq, saveSettings as saveSettingsReq} from "../../api/settings";


ListView.propTypes = {data: PropTypes.arrayOf(PropTypes.string)};

export default function TextToSpeech() {
  const [isLoading, setIsLoading] = useState(false)

  useEffect(() => {
    let mounted = true;

    getSettingsReq()
      .then(setting => {
        if (mounted) {
          console.log(setting);
          const channelsToListen = setting.ChannelsToListen;
          const ignoreWords = setting.IgnoreWords.split(',');
          const banList = setting.UserBanList;
          const userBanList = banList.includes(',') ? banList.split(',').filter(item => item.length > 0) : [];
          setChannelsToListen(channelsToListen);
          setLocalUsername(channelsToListen)
          setIgnoreWords(ignoreWords)
          setLanguage(setting.Language)
          setLanguageDetectorEnabled(setting.LanguageDetectorEnabled)
          let wordPair = setting.ReplacementWordPair;
          while (wordPair[wordPair.length - 1] === ',') {
            wordPair = deleteLastSymbols(wordPair)
          }
          const replacementWordPair = wordPair.split(',').map(pair => {
            const split = pair.split(':');
            return {
              before: split[0],
              after: split[1],
            }
          });
          setReplacementWordPair(replacementWordPair)
          setUserBanList(userBanList)
          setVolume(setting.Volume)
          delay(100)

          setIsLoading(false);
        }
      })
    return () => mounted = false;
  }, []);

  const [channelsToListen, setChannelsToListen] = useState('');
  const [localUsername, setLocalUsername] = useState('');
  const [replacementWordPair, setReplacementWordPair] = useState([]);
  const [userBanList, setUserBanList] = useState([]);
  const [ignoreWords, setIgnoreWords] = useState([]);
  const [language, setLanguage] = useState('en');
  const [languageDetectorEnabled, setLanguageDetectorEnabled] = useState(false);
  const [volume, setVolume] = useState(5);

  const saveSettings = () => {
    let wordPair = replacementWordPair.reduce((result, pair) => {
      if (pair.before.length === 0) {
        return result
      }
      return result + pair.before + ':' + pair.after + ','
    }, '');
    while (wordPair[wordPair.length - 1] === ',') {
      wordPair = deleteLastSymbols(wordPair)
    }
    const setting = {
      Id: 1,
      ReplacementWordPair: wordPair,
      IgnoreWords: listToStr(ignoreWords),
      Language: language,
      LanguageDetectorEnabled: languageDetectorEnabled,
      UserBanList: listToStr(userBanList),
      ChannelsToListen: channelsToListen,
      Volume: volume,
    };
    console.log(setting);
    saveSettingsReq(JSON.stringify(setting)).then(resp => {
      if (resp.status === 200) {
        console.log(resp);
      }
    }).catch(err => [
      console.log(err)
    ])
  };

  const onSubmitTwitchUsername = (e) => {
    e.preventDefault();
    setChannelsToListen(e.target[0].value);
    saveSettings();
  };

  const onSubmitWordPair = (pair) => {
    if (pair.before.length < 1) {
      return
    }
    const filter = replacementWordPair.find(item => item.before.toLowerCase() === pair.before);
    if (filter === undefined) {
      appendToList(pair, replacementWordPair, setReplacementWordPair);
      saveSettings();
    }
  }

  const onChangeVolume = e => {
    setVolume(parseFloat(e.target.value))
    saveSettings();
  };

  if (isLoading) {
    return <div>Loading...</div>
  }

  return (
    <>
      <form onSubmit={onSubmitTwitchUsername}>
        <label htmlFor="input">
          <p>Twitch username</p>
          <input type="text" value={localUsername} onChange={e => setLocalUsername(e.target.value)}/>
        </label>
        <button type="submit">Watch</button>
        <p>Currently watching {channelsToListen}</p>
      </form>

      <div>
        <label htmlFor="input">
          Volume {volume}
          <input type="range" min="0" max="10" value={volume}
                 onChange={(e) => onChangeVolume(e)}/>
        </label>
      </div>
      <div>
        <p>Add word replacement</p>
        <WordPairForm onSubmitWordPair={onSubmitWordPair}/>
        <div>
          <p>Word replacements</p>
          <WordPairList
            pairList={replacementWordPair}
            onDeletePair={(index) => {
              deleteByIndex(index, replacementWordPair, setReplacementWordPair)
              saveSettings();
            }}
          />
        </div>
      </div>
      <p>Selected Language {language}</p>
      <Languages
        language={language}
        languageDetectorEnabled={languageDetectorEnabled}

        onSelectLanguage={(lang) => {
          setLanguage(lang);
          saveSettings();
        }}
        onAutoDetectEnabled={(isEnabled) => {
          console.log(isEnabled);
          setLanguageDetectorEnabled(isEnabled);
          saveSettings();
        }}
      />
      <ListFormView title="Banned users" label="User for ban" minLength={4} list={userBanList}
                    callback={(list) => {
                      setUserBanList(list);
                      saveSettings();
                    }}/>
      <ListFormView title="Ignore words" label="ignore word" minLength={2} list={ignoreWords}
                    callback={(list) => {
                      setIgnoreWords(list);
                      saveSettings();
                    }}/>
    </>
  );
};