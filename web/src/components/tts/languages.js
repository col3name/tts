import React from "react";

export default function Languages(props) {
  const detectorEnabled = props.languageDetectorEnabled;
  const allLanguages = [
    {
      "code": "en",
      "name": "en"
    },
    {
      "code": "en-UK",
      "name": "en-UK"
    },
    {
      "code": "en-AU",
      "name": "en-AU"
    },
    {
      "code": "ja",
      "name": "ja"
    },
    {
      "code": "de",
      "name": "de"
    },
    {
      "code": "es",
      "name": "es"
    },
    {
      "code": "ru",
      "name": "ru"
    },
    {
      "code": "ar",
      "name": "ar"
    },
    {
      "code": "bn",
      "name": "bn"
    },
    {
      "code": "cs",
      "name": "cs"
    },
    {
      "code": "da",
      "name": "da"
    },
    {
      "code": "nl",
      "name": "nl"
    },
    {
      "code": "fi",
      "name": "fi"
    },
    {
      "code": "el",
      "name": "el"
    },
    {
      "code": "hi",
      "name": "hi"
    },
    {
      "code": "hu",
      "name": "hu"
    },
    {
      "code": "id",
      "name": "id"
    },
    {
      "code": "km",
      "name": "km"
    },
    {
      "code": "la",
      "name": "la"
    },
    {
      "code": "it",
      "name": "it"
    },
    {
      "code": "no",
      "name": "no"
    },
    {
      "code": "pl",
      "name": "pl"
    },
    {
      "code": "sk",
      "name": "sk"
    },
    {
      "code": "sv",
      "name": "sv"
    },
    {
      "code": "th",
      "name": "th"
    },
    {
      "code": "tr",
      "name": "tr"
    },
    {
      "code": "uk",
      "name": "uk"
    },
    {
      "code": "vi",
      "name": "vi"
    },
    {
      "code": "af",
      "name": "af"
    },
    {
      "code": "bg",
      "name": "bg"
    },
    {
      "code": "ca",
      "name": "ca"
    },
    {
      "code": "cy",
      "name": "cy"
    },
    {
      "code": "et",
      "name": "et"
    },
    {
      "code": "fr",
      "name": "fr"
    },
    {
      "code": "gu",
      "name": "gu"
    },
    {
      "code": "is",
      "name": "is"
    },
    {
      "code": "jv",
      "name": "jv"
    },
    {
      "code": "kn",
      "name": "kn"
    },
    {
      "code": "ko",
      "name": "ko"
    },
    {
      "code": "lv",
      "name": "lv"
    },
    {
      "code": "ml",
      "name": "ml"
    },
    {
      "code": "mr",
      "name": "mr"
    },
    {
      "code": "ms",
      "name": "ms"
    },
    {
      "code": "ne",
      "name": "ne"
    },
    {
      "code": "pt",
      "name": "pt"
    },
    {
      "code": "ro",
      "name": "ro"
    },
    {
      "code": "si",
      "name": "si"
    },
    {
      "code": "sr",
      "name": "sr"
    },
    {
      "code": "su",
      "name": "su"
    },
    {
      "code": "ta",
      "name": "ta"
    },
    {
      "code": "te",
      "name": "te"
    },
    {
      "code": "tl",
      "name": "tl"
    },
    {
      "code": "ur",
      "name": "ur"
    },
    {
      "code": "zh",
      "name": "zh"
    },
    {
      "code": "sw",
      "name": "sw"
    },
    {
      "code": "sq",
      "name": "sq"
    },
    {
      "code": "my",
      "name": "my"
    },
    {
      "code": "mk",
      "name": "mk"
    },
    {
      "code": "hy",
      "name": "hy"
    },
    {
      "code": "hr",
      "name": "hr"
    },
    {
      "code": "eo",
      "name": "eo"
    },
    {
      "code": "bs",
      "name": "bs"
    }
  ]

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
    langSelector(allLanguages);

  return (
    <div>
      <label htmlFor="input">
        Language Detector Enabled
        <input type="checkbox" value={detectorEnabled}
               onChange={(e) => props.onAutoDetectEnabled(e.target.value !== 'true')}/>
      </label>
      {langSelectorHTML}
    </div>
  )

};