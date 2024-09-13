import { Vocabulary } from "@/types/Vocabulary";

export const updateVocabulary = (
    vocabularies: Vocabulary[],
    index: number,
    field: keyof Vocabulary,
    value: string
): Vocabulary[] => {
    const newVocabularies = [...vocabularies];
    newVocabularies[index][field] = value;
    return newVocabularies;
};

export const addVocabulary = (vocabularies: Vocabulary[]): Vocabulary[] => {
    return [...vocabularies, { phrase: '', soundsLike: '', ipa: '', displayAs: '' }];
};

export const removeVocabulary = (vocabularies: Vocabulary[], index: number): Vocabulary[] => {
    return vocabularies.filter((_, i) => i !== index);
};