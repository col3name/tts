CREATE TABLE settings
(
    replacement_word_pair TEXT         NOT NULL DEFAULT '',
    ignore_words          TEXT         NOT NULL DEFAULT '',
    language              VARCHAR(255) NOT NULL DEFAULT 'en',
    lang_detector_enabled BOOLEAN               DEFAULT false,
    userBanList           TEXT         NOT NULL DEFAULT '',
    channelsToListen      VARCHAR(255) NOT NULL,
    volume                SMALLINT     NOT NULL DEFAULT 1
);