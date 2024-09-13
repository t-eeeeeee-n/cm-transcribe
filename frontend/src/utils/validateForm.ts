import toast from 'react-hot-toast';
import { Vocabulary } from "@/types/Vocabulary";

interface ValidationProps {
    vocabularyName: string;
    languageCode: string;
    vocabularies: Vocabulary[];
}

export const validateForm = ({ vocabularyName, languageCode, vocabularies }: ValidationProps): boolean => {
    if (!vocabularyName) {
        toast.error("名前を入力してください。");
        return false;
    }
    if (!languageCode) {
        toast.error("言語を選択してください。");
        return false;
    }
    if (vocabularies.length === 0 || vocabularies.every(vocabulary => !vocabulary || !vocabulary.phrase || !vocabulary.phrase.trim())) {
        toast.error("少なくとも1つのフレーズを入力するか、CSVをインポートしてください。");
        return false;
    }
    return true;
};
