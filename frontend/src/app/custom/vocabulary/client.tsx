"use client";

import React, { useState } from 'react';
import axios from 'axios';
import { Vocabulary } from "@/types/Vocabulary";
import { addVocabulary, removeVocabulary, updateVocabulary } from "@/utils/vocabularyUtils";
import VocabularyLayout from "@/components/VocabularyLayout";
import { validateForm } from "@/utils/validateForm";  // 共通バリデーション関数をインポート
import toast from 'react-hot-toast';  // react-hot-toastのインポート

const Client: React.FC = () => {
    const [vocabularyName, setVocabularyName] = useState('');
    const [languageCode, setLanguageCode] = useState<string>("ja-JP");
    const [vocabularies, setVocabularies] = useState<Vocabulary[]>([]);
    const [isLoading, setIsLoading] = useState<boolean>(false);

    const handleVocabularyChange = (index: number, field: keyof Vocabulary, value: string) => {
        setVocabularies(updateVocabulary(vocabularies, index, field, value));
    };

    const handleAddVocabulary = () => {
        setVocabularies(addVocabulary(vocabularies));
    };

    const handleRemoveVocabulary = (index: number) => {
        setVocabularies(removeVocabulary(vocabularies, index));
    };

    const handleSubmit = async () => {
        // バリデーションチェック
        if (!validateForm({ vocabularyName, languageCode, vocabularies })) return;

        setIsLoading(true);

        const body = {
            name: vocabularyName,
            language_code: languageCode,
            vocabularies: vocabularies
        };

        try {
            await axios.post('/api/vocabulary', body);
            toast.success('保存しました！');
        } catch (error) {
            // console.error('保存に失敗しました。', error);
            if (axios.isAxiosError(error) && error.response) {
                if (error.response.status === 409) {
                    toast.error("カスタムボキャブラリー名が既に存在します。");
                    return;
                }
            }
            toast.error('保存に失敗しました。');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <VocabularyLayout
            vocabularyName={vocabularyName}
            setVocabularyName={setVocabularyName}
            languageCode={languageCode}
            setLanguageCode={setLanguageCode}
            vocabularies={vocabularies}
            onVocabularyChange={handleVocabularyChange}
            onAddVocabulary={handleAddVocabulary}
            onRemoveVocabulary={handleRemoveVocabulary}
            setVocabulary={setVocabularies}
            vocabularyState={""}
            handleSubmit={handleSubmit}
            isLoading={isLoading}
            submitButtonText="保存"
        />
    );
};

export default Client;
