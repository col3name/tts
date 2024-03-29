# Simple bot for twitch, who can speech message from chat, NO required login
![CI passing](https://github.com/col3name/tts/actions/workflows/prod.yml/badge.svg)

### Feature
- Word replacements
- Word ignore
- Url ignore
- Sequential message reading
- Supported about 60 language
- Volume control
- User ban list
- Automatic language detection

![drawing](docs/preview.png)

### Supported languages
[ISO language code](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes)
- Supported speaking language without automatic detection: en, en-UK, en-AU, ja, de, es, ru, ar, bn, cs, da, nl, fi, el,
  hi, hu, id, km, la, it, no, pl, sk, sv, th, tr, uk, vi, af, bg, ca, cy, et, fr, gu, is, jv, kn, ko, lv, ml, mr, ms,
  ne, pt, ro, si, sr, su, ta, te, tl, ur, zh, sw, sq, my, mk, hy, hr, eo, bs
- Supported automatic language detection: English (en), Spanish (es), French (fr), German (de), Russian (ru), Chinese (
  zh)

## How Works
![drawing](docs/Arch-Diagram.jpg)

## Build

##### Building for windows set this environment variable

GOARCH=amd64 GOOS=windows
```make build && main.exe```

##### Building for linux set this environment variable

GOOS=linux GOARCH=amd64
```go build -o main main.go```
```chmod +x ./main```

##### Building for macos set this environment variable

GOOS=darwin GOARCH=arm64
```go build -o main main.go```
```chmod +x ./main```

## Test

```make test```

