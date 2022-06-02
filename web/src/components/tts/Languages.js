import React from "react";
import ALL_LANGUAGES from "../../api/languages";

export default function Languages(props) {
  const detectorEnabled = props.languageDetectorEnabled;

  const langSelector = (languages) => {
    return <div>
      <label htmlFor="select">
        <select value={props.language} onChange={(e) => props.onSelectLanguage(e.target.value)}>
          {languages.map((lang) => (
            <option key={lang.code}
                    value={lang.code}>{lang.name}</option>
          ))}
        </select>
      </label>
    </div>
  }

  let langSelectorHTML = detectorEnabled ?
    "" :
    langSelector(ALL_LANGUAGES);

  return (
    <div>
      <p>Selected Language {props.language}</p>
      <label htmlFor="input">
        Language Detector Enabled
        <input type="checkbox" value={detectorEnabled}
               onChange={(e) => props.onAutoDetectEnabled(e.target.value !== 'true')}/>
      </label>
      {langSelectorHTML}
    </div>
  )
};