"use client"
import React, { useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Typography } from '@/components/ui/typography';

// 単語と信頼度を定義する型
type ConfidenceItem = {
    word: string;
    confidence: string;
};

// ClientComponentのPropsの型を定義
interface ClientComponentProps {
    transcript: string;
    confidenceData: ConfidenceItem[];
    jobName: string; // jobName を追加
}

// 吹き出しスタイルのコンポーネント
const Tooltip = ({ content }: { content: string }) => (
    <div
        className="absolute left-1/2 transform -translate-x-1/2 bg-gray-800 text-white text-xs rounded-md px-2 py-1 z-50">
        {content}
        <div
            className="absolute -top-[5px] left-1/2 transform -translate-x-1/2 border-b-[5px] border-b-gray-800 border-l-[5px] border-l-transparent border-r-[5px] border-r-transparent"></div>
    </div>
);

// 関数コンポーネントを型指定して定義
const ClientComponent: React.FC<ClientComponentProps> = ({ transcript, confidenceData, jobName }) => { // jobName を受け取る
    const [selectedWord, setSelectedWord] = useState<number | null>(null); // 選択された単語を管理

    // スタイル設定関数
    const getWordStyle = (confidence: string) => {
        const confidenceValue = parseFloat(confidence);
        if (confidenceValue >= 0.9) {
            return 'text-black'; // 90%以上の信頼度なら普通に表示
        } else if (confidenceValue >= 0.7) {
            return 'text-blue-500'; // 70%以上90%未満なら青で表示
        } else {
            return 'text-red-500'; // 70%未満なら赤で表示
        }
    };

    // 単語がクリックされた時の処理
    const handleWordClick = (index: number) => {
        if (selectedWord === index) {
            // 同じ単語が再度クリックされたら選択解除
            setSelectedWord(null);
        } else {
            // 違う単語がクリックされたら選択
            setSelectedWord(index);
        }
    };

    return (
        <div className="bg-gray-50 dark:bg-gray-900 min-h-screen py-10 px-4">
            <div className="max-w-2xl mx-auto space-y-6">
                <Card className="shadow-lg border border-gray-200 rounded-lg">
                    <CardHeader
                        className="bg-gradient-to-r from-purple-500 to-indigo-600 text-white p-4 rounded-t-lg mb-5">
                        <CardTitle className="text-2xl font-semibold">ジョブ</CardTitle> {/* CardHeader をジョブに設定 */}
                    </CardHeader>
                    <CardContent>
                        <Typography className="text-gray-900 font-bold text-xl mb-4">{jobName}</Typography> {/* jobName を CardContent に表示 */}
                        <Typography className="text-gray-600 relative">
                            {confidenceData && confidenceData.length > 0 ? (
                                confidenceData.map((item, index) => (
                                    <span
                                        key={index}
                                        className={`relative ${getWordStyle(item.confidence)} hover:underline cursor-pointer`}
                                        onClick={() => handleWordClick(index)}
                                    >
                                        {item.word}
                                        {selectedWord === index && (
                                            <Tooltip
                                                content={`Confidence: ${(parseFloat(item.confidence) * 100).toFixed(1)}%`}
                                            />
                                        )}
                                    </span>
                                ))
                            ) : (
                                'Loading...'
                            )}
                        </Typography>
                    </CardContent>
                </Card>
            </div>
        </div>
    );
};

export default ClientComponent;
