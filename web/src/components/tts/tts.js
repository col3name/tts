import React from "react";
import * as PropTypes from "prop-types";
import ListView from "./list/ListView";
import ListFormView from "./list/ListFormView";
import Languages from "./Languages";
import UsernameForm from "./UsernameForm";
import VolumeInput from "./VolumeInput";
import WordPair from "./pair/WordPair";
import useSettingState from "../../hooks/useSettingState";

ListView.propTypes = {data: PropTypes.arrayOf(PropTypes.string)};

export default function TextToSpeech() {
  const {
    channelsToListen, onUpdateTwitchUsername,
    ignoreWords, onUpdateIgnoreWordList,
    language, onSelectLanguage,
    languageDetectorEnabled, onAutoDetectEnabled,
    replacementWordPair, onUpdateReplacementWordPairs,
    userBanList, onUpdateUserBanList,
    volume, onUpdateVolume,
    isLoading,
  } = useSettingState()

  if (isLoading) {
    return <div>Loading...</div>
  }

  return (
    <>
      <UsernameForm
        channelsToListen={channelsToListen}
        onUpdateTwitchUsername={onUpdateTwitchUsername}
      />
      <VolumeInput
        volume={volume}
        onUpdateVolume={onUpdateVolume}
      />
      <WordPair
        title="Add word replacement"
        subtitle="Word replacements"
        wordPairList={replacementWordPair}

        onUpdatePairs={onUpdateReplacementWordPairs}
      />

      <Languages
        language={language}
        languageDetectorEnabled={languageDetectorEnabled}

        onSelectLanguage={onSelectLanguage}
        onAutoDetectEnabled={onAutoDetectEnabled}
      />
      <ListFormView title="Banned users"
                    label="User for ban"
                    minLength={4}
                    list={userBanList}
                    callback={onUpdateUserBanList}/>
      <ListFormView title="Ignore words"
                    label="ignore word"
                    minLength={2}
                    list={ignoreWords}
                    callback={onUpdateIgnoreWordList}/>
    </>
  );
};