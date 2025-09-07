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