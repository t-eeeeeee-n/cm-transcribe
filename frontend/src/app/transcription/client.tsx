"use client";

import React, { useState } from 'react';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Switch } from '@/components/ui/switch';
import { Button } from '@/components/ui/button';
import { Toaster, toast } from "react-hot-toast";
import axios from 'axios';

const Client: React.FC = () => {
    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [selectedLanguage, setSelectedLanguage] = useState<string>('ja-JP');
    const [isSpeakerSeparationEnabled, setIsSpeakerSeparationEnabled] = useState<boolean>(false);
    const [isTimestampEnabled, setIsTimestampEnabled] = useState<boolean>(false);
    const [isLoading, setIsLoading] = useState<boolean>(false);

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            const file = e.target.files[0];
            const allowedExtensions = ['audio/mpeg', 'audio/wav', 'audio/mp3', 'audio/m4a'];

            if (!allowedExtensions.includes(file.type)) {
                toast.error('無効なファイル形式です。対応している形式: MP3, WAV, M4A');
                setSelectedFile(null);
                return;
            }

            setSelectedFile(file);
        }
    };

    const handleStartTranscription = async () => {
        if (!selectedFile) {
            toast.error("ファイルが選択されていません。");
            return;
        }

        setIsLoading(true);

        try {
            // フロントエンドから直接Golangバックエンドにファイルを送信
            const formData = new FormData();
            formData.append('file', selectedFile);
            formData.append('languageCode', selectedLanguage);
            formData.append('isSpeakerSeparationEnabled', JSON.stringify(isSpeakerSeparationEnabled));
            formData.append('isTimestampEnabled', JSON.stringify(isTimestampEnabled));

            const response = await axios.post('/api/transcribe', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });

            if (response.status === 200) {
                const jobId = response.data?.ID;
                // toast.success(`文字起こしジョブが正常に開始されました。ジョブID: ${jobId}`);
                toast.success(
                    <span>
                    文字起こしジョブが正常に開始されました。ジョブID:{" "}
                        <span
                            style={{ cursor: "pointer", textDecoration: "underline" }}
                            onClick={() => {
                                navigator.clipboard.writeText(jobId);
                                toast.success("ジョブIDをコピーしました！");
                            }}
                        >
                        {jobId}
                    </span>
                </span>,
                    { duration: 10000 } // トーストメッセージの表示時間を長くする（10秒）
                );
            } else {
                throw new Error('文字起こしジョブの開始に失敗しました');
            }

        } catch (error) {
            console.error('エラーが発生しました:', error);
            toast.error("エラーが発生しました。再試行してください。");
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="bg-gray-50 dark:bg-gray-900 min-h-screen py-10 px-4">
            <Toaster position="top-center" />

            <div className="max-w-2xl mx-auto space-y-6">
                <Card className="mb-8 shadow-lg border border-gray-200 rounded-lg bg-white dark:bg-gray-800 transition duration-300 ease-in-out hover:shadow-xl">
                    <CardHeader className="p-4 bg-gradient-to-r from-purple-500 to-indigo-600 text-white rounded-t-lg">
                        <CardTitle className="text-2xl font-semibold">文字起こし</CardTitle>
                    </CardHeader>
                    <CardContent className="p-6 space-y-6">
                        {/* ファイルのアップロード */}
                        <div className="space-y-4">
                            <Label htmlFor="file-upload" className="text-sm font-medium text-gray-700 dark:text-gray-300">
                                音声ファイルをアップロード
                            </Label>
                            <Input
                                type="file"
                                id="file-upload"
                                onChange={handleFileChange}
                                className="border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring focus:ring-purple-200 focus:ring-opacity-50 focus:border-indigo-500"
                            />
                        </div>

                        {/* 言語選択とオプション */}
                        <div className="space-y-4">
                            <div className="space-y-2">
                                <Label htmlFor="language-select" className="text-sm font-medium text-gray-700 dark:text-gray-300">
                                    言語を選択
                                </Label>
                                <Select value={selectedLanguage} onValueChange={setSelectedLanguage}>
                                    <SelectTrigger className="w-full border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring focus:ring-purple-200 focus:ring-opacity-50 focus:border-indigo-500">
                                        <SelectValue placeholder="言語を選択" />
                                    </SelectTrigger>
                                    <SelectContent>
                                        <SelectItem value="ja-JP">日本語 (ja-JP)</SelectItem>
                                        <SelectItem value="en-US">英語 (en-US)</SelectItem>
                                    </SelectContent>
                                </Select>
                            </div>
                            <div className="space-y-4">
                                <div className="flex items-center space-x-2">
                                    <Switch checked={isSpeakerSeparationEnabled} onCheckedChange={setIsSpeakerSeparationEnabled} />
                                    <Label className="text-sm font-medium text-gray-700 dark:text-gray-300">話者分離を有効にする</Label>
                                </div>
                                <div className="flex items-center space-x-2">
                                    <Switch checked={isTimestampEnabled} onCheckedChange={setIsTimestampEnabled} />
                                    <Label className="text-sm font-medium text-gray-700 dark:text-gray-300">タイムスタンプを追加</Label>
                                </div>
                            </div>
                            <div className="flex justify-end mt-6">
                                <Button
                                    onClick={handleStartTranscription}
                                    className="w-1/3 bg-green-600 text-white hover:bg-green-700 transition-all duration-300 rounded-lg"
                                    disabled={!selectedFile || isLoading}
                                >
                                    {isLoading ? '処理中...' : 'Start'}
                                </Button>
                            </div>
                        </div>
                    </CardContent>
                </Card>
            </div>
        </div>
    );
};

export default Client;
