"use client";

import React, { useState } from 'react';
import { Box, Button, Container } from '@mui/material';
import axios from 'axios';
import VocabularyForm from "@/components/VocabularyForm";
import PhraseTable from "@/components/PhraseTable";
import DictionarySettings from "@/components/PhraseTable";

const Client: React.FC = () => {
    // ボキャブラリー名と言語コードの状態管理
    const [vocabularyName, setVocabularyName] = useState('');
    const [languageCode, setLanguageCode] = useState('');

    // フレーズデータの状態管理
    const [phrases, setPhrases] = useState([
        { phrase: '', soundsLike: '', ipa: '', displayAs: '' }
    ]);

    // フレーズの更新
    const handlePhraseChange = (index: number, field: keyof typeof phrases[0], value: string) => {
        const newPhrases = [...phrases];
        newPhrases[index][field] = value;
        setPhrases(newPhrases);
    };

    // フレーズの追加
    const handleAddPhrase = () => {
        setPhrases([...phrases, { phrase: '', soundsLike: '', ipa: '', displayAs: '' }]);
    };

    // フレーズの削除
    const handleRemovePhrase = (index: number) => {
        setPhrases(phrases.filter((_, i) => i !== index));
    };

    // フォームの送信
    const handleSubmit = async () => {
        const body = {
            name: vocabularyName,
            language_code: languageCode,
            vocabularies: phrases
        };

        try {
            const response = await axios.post('/api/saveVocabulary', body);  // APIルートにPOST
            console.log(response.data);
            alert('保存しました！');
        } catch (error) {
            console.error(error);
            alert('保存に失敗しました。');
        }
    };

    return (
        <Box sx={{ backgroundColor: '#f5f5f5', minHeight: '100vh', py: 4 }}>
            <Container maxWidth="md">
                {/* ボキャブラリー設定フォーム */}
                <VocabularyForm
                    vocabularyName={vocabularyName}
                    setVocabularyName={setVocabularyName}
                    languageCode={languageCode}
                    setLanguageCode={setLanguageCode}
                />

                {/* フレーズ設定テーブル */}
                <DictionarySettings
                />

                {/* 保存ボタン */}
                <Box display="flex" justifyContent="center" mt={4}>
                    <Button variant="contained" color="primary" onClick={handleSubmit}>
                        保存
                    </Button>
                </Box>
            </Container>
        </Box>
    );
};

export default Client;
