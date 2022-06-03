import React from "react";
import ALL_SUPPORTED_LANGUAGES from "../../api/languages";

export default function Languages({
                                    language,
                                    languageDetectorEnabled,
                                    onSelectLanguage,
                                    onAutoDetectEnabled
                                  }) {
  const langSelector = (languages) => {
    return <div>
      <label htmlFor="select">
        <select value={language} onChange={(e) => {
          onSelectLanguage(e.target.value)
        }
        }>
          {languages.map((lang) => (
            <option key={lang.code}
                    value={lang.code}>{lang.name}</option>
          ))}
        </select>
      </label>
    </div>
  }

  const langSelectorHTML = languageDetectorEnabled ?
    "" :
    langSelector(ALL_SUPPORTED_LANGUAGES);

  return (
    <div>
      <p>Selected Language {language}</p>
      <label>
        Language Detector Enabled
        <input type="checkbox" value={languageDetectorEnabled}
               onChange={(e) => {
                 onAutoDetectEnabled(e.target.value !== 'true')
               }}/>
      </label>
      {langSelectorHTML}
    </div>
  )
};