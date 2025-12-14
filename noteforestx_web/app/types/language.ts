export const languageList: LanguageItem[] = [
    { label: "English", flag: "us", code: "en_us" },
    { label: "简体中文", flag: "cn", code: "zh_cn" },
    { label: "繁體中文", flag: "hk", code: "zh_hk" },
    { label: "日本語", flag: "jp", code: "ja_jp" },
    { label: "Español", flag: "es", code: "es_es" },
    { label: "Français", flag: "fr", code: "fr_fr" },
    { label: "Suomi", flag: "fi", code: "fi_fi" },
    { label: "Deutsch", flag: "de", code: "de_de" },
    { label: "Nederlands", flag: "nl", code: "nl_nl" },
];

export type LanguageItem = {
    label: string,
    flag: string,
    code: LanguageCode,
}

export type LanguageCode =
    'en_us' |
    'zh_cn' |
    'zh_hk' |
    'ja_jp' |
    'es_es' |
    'fr_fr' |
    'fi_fi' |
    'de_de' |
    'nl_nl';