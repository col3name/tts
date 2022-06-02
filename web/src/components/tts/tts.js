import React from "react";
import ListFormView from "./list/ListFormView";
import Languages from "./Languages";
import UsernameForm from "./UsernameForm";
import VolumeInput from "./VolumeInput";
import WordPair from "./pair/WordPair";
import useSettingState from "../../hooks/useSettingState";

export default function TextToSpeech() {
  const {
    channelsToListen,
    ignoreWords,
    language,
    languageDetectorEnabled,
    replacementWordPair,
    userBanList,
    volume,
    isLoading,
  } = useSettingState()

  if (isLoading) {
    return <div>Loading...</div>
  }

  return (
    <>
      <UsernameForm
        channelsToListen={channelsToListen.value}
        onUpdateTwitchUsername={channelsToListen.onUpdate}
      />
      <VolumeInput
        volume={volume.value}
        onUpdateVolume={volume.onUpdate}
      />
      <WordPair
        title="Add word replacement"
        subtitle="Word replacements"
        wordPairList={replacementWordPair.value}

        onUpdatePairs={replacementWordPair.onUpdate}
      />
      <Languages
        language={language.value}
        languageDetectorEnabled={languageDetectorEnabled.value}

        onSelectLanguage={language.onUpdate}
        onAutoDetectEnabled={languageDetectorEnabled.onUpdate}
      />
      <ListFormView title="Banned users"
                    label="User for ban"
                    minLength={4}
                    list={userBanList.value}
                    onUpdate={userBanList.onUpdate}/>
      <ListFormView title="Ignore words"
                    label="ignore word"
                    minLength={2}
                    list={ignoreWords.value}
                    onUpdate={ignoreWords.onUpdate}/>
    </>
  );
};