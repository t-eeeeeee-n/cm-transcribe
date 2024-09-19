"use client";

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import axios from 'axios';
import { Vocabulary } from "@/types/Vocabulary";
import { addVocabulary, removeVocabulary, updateVocabulary } from "@/utils/vocabularyUtils";
import VocabularyLayout from "@/components/VocabularyLayout";  // 共通レイアウトをインポート
import { validateForm } from "@/utils/validateForm";  // 共通バリデーション関数をインポート
import toast from 'react-hot-toast';  // react-hot-toastのインポート

interface ClientProps {
    data: {
        vocabularyName: string;
        languageCode: string;
        vocabularies: Vocabulary[];
        vocabularyState: string;
        lastModifiedTime?: string;
    };
}

const Client: React.FC<ClientProps> = ({ data }) => {
    const router = useRouter();

    const [vocabularyName, setVocabularyName] = useState<string>(data.vocabularyName);
    const [languageCode, setLanguageCode] = useState<string>(data.languageCode);
    const [vocabularies, setVocabularies] = useState<Vocabulary[]>(data.vocabularies);

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

        const body = {
            name: vocabularyName,
            language_code: languageCode,
            vocabularies: vocabularies
        };

        try {
            await axios.put(`/api/vocabulary/${data.vocabularyName}`, body);
            toast.success('更新しました！');  // 成功メッセージ
            router.push('/custom/vocabulary');  // 更新後にリダイレクト
        } catch (error) {
            console.error('更新に失敗しました。', error);
            toast.error('更新に失敗しました。');  // エラーメッセージ
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
            vocabularyState={data.vocabularyState}
            lastModifiedTime={data.lastModifiedTime}
            handleSubmit={handleSubmit}  // handleSubmit関数を渡す
            submitButtonText="更新"
        />
    );
};

export default Client;
