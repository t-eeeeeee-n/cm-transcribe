import React from 'react';
import { Button } from '@/components/ui/button';
import VocabularyForm from "@/components/VocabularyForm";
import DictionarySettings from "@/components/DictionarySettings";
import { Vocabulary } from "@/types/Vocabulary";
import { Toaster } from 'react-hot-toast';
import VocabularyStatus from "@/components/VocabularyStatus";
import LoadingSpinner from "@/components/LoadingSpinner";  // Toasterのインポート

interface VocabularyLayoutProps {
    vocabularyName: string;
    setVocabularyName: (name: string) => void;
    languageCode: string;
    setLanguageCode: (code: string) => void;
    vocabularies: Vocabulary[];
    onVocabularyChange: (index: number, field: keyof Vocabulary, value: string) => void;
    onAddVocabulary: () => void;
    onRemoveVocabulary: (index: number) => void;
    setVocabulary: (vocabularies: Vocabulary[]) => void;
    vocabularyState: string;
    lastModifiedTime?: string;
    handleSubmit: () => void;
    isLoading: boolean;
    submitButtonText: string;
}

const VocabularyLayout: React.FC<VocabularyLayoutProps> = ({
                                                               vocabularyName,
                                                               setVocabularyName,
                                                               languageCode,
                                                               setLanguageCode,
                                                               vocabularies,
                                                               onVocabularyChange,
                                                               onAddVocabulary,
                                                               onRemoveVocabulary,
                                                               setVocabulary,
                                                               vocabularyState,
                                                               lastModifiedTime,
                                                               handleSubmit,
                                                               isLoading,
                                                               submitButtonText,
                                                           }) => {
    return (
        <div className="bg-gray-50 dark:bg-gray-900 min-h-screen py-10 px-4">
            <Toaster position="top-center" /> {/* Toasterを配置 */}

            <div className="max-w-2xl mx-auto space-y-6">

                {/* Vocabulary Form */}
                <VocabularyForm
                    vocabularyName={vocabularyName}
                    setVocabularyName={setVocabularyName}
                    languageCode={languageCode}
                    setLanguageCode={setLanguageCode}
                    vocabularyState={vocabularyState}
                    lastModifiedTime={lastModifiedTime}
                />

                {/* Dictionary Settings */}
                <DictionarySettings
                    vocabularies={vocabularies}
                    onVocabularyChange={onVocabularyChange}
                    onAddVocabulary={onAddVocabulary}
                    onRemoveVocabulary={onRemoveVocabulary}
                    setVocabulary={setVocabulary}
                />

                {/* Submit Button */}
                <div className="flex justify-center mt-8">
                    <Button
                        className="w-1/3 bg-blue-600 text-white hover:bg-blue-700 transition-colors duration-200 px-6 py-3 rounded-md"
                        onClick={handleSubmit}  // 引数で渡されたhandleSubmitを使う
                    >
                        {isLoading ? <LoadingSpinner /> : submitButtonText}
                    </Button>
                </div>
            </div>
        </div>
    );
};

export default VocabularyLayout;
