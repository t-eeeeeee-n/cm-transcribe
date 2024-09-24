"use client";

import React, { useState } from 'react';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Switch } from '@/components/ui/switch';
import { Button } from '@/components/ui/button';
import { Toaster, toast } from "react-hot-toast";
import axios from 'axios';
import LoadingSpinner from "@/components/LoadingSpinner";
import LanguageSelect from "@/components/LanguageSelect";

const Client: React.FC = () => {
    const [jobName, setJobName] = useState<string>('');
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
        if (!jobName) {
            toast.error("ジョブ名が入力されていません");
            return;
        }
        if (!selectedFile) {
            toast.error("ファイルが選択されていません。");
            return;
        }

        setIsLoading(true);

        try {
            // フロントエンドから直接Golangバックエンドにファイルを送信
            const formData = new FormData();
            formData.append('jobName', jobName);
            formData.append('file', selectedFile);
            formData.append('languageCode', selectedLanguage);
            formData.append('isSpeakerSeparationEnabled', JSON.stringify(isSpeakerSeparationEnabled));
            formData.append('isTimestampEnabled', JSON.stringify(isTimestampEnabled));

            await axios.post('/api/transcription', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });
            toast.success(`文字起こしジョブが正常に開始されました。`);
            // const jobId = response.data?.ID;
            // toast.success(
            //     <span>
            //     "文字起こしジョブが正常に開始されました。"ジョブID:{" "}
            //         <span
            //             style={{ cursor: "pointer", textDecoration: "underline" }}
            //             onClick={() => {
            //                 navigator.clipboard.writeText(jobId);
            //                 toast.success("ジョブIDをコピーしました！");
            //             }}
            //         >
            //         {jobId}
            //     </span>
            // </span>,
            //     { duration: 10000 } // トーストメッセージの表示時間を長くする（10秒）
            // );

        } catch (error) {
            // console.error('エラーが発生しました:', error);// エラーのステータスコードをチェック
            if (axios.isAxiosError(error) && error.response) {
                if (error.response.status === 409) {
                    toast.error("ジョブ名が既に存在します。");
                    return;
                }
            }
            toast.error("文字起こしジョブの開始に失敗しました。再試行してください。");
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="bg-gray-50 dark:bg-gray-900 min-h-screen py-10 px-4">
            <Toaster position="top-center" />

            <div className="max-w-2xl mx-auto space-y-6">
                <Card
                    className="mb-8 shadow-lg border border-gray-200 rounded-lg bg-white dark:bg-gray-800 transition duration-300 ease-in-out hover:shadow-xl">
                    <CardHeader className="p-4 bg-gradient-to-r from-purple-500 to-indigo-600 text-white rounded-t-lg">
                        <CardTitle className="text-2xl font-semibold">文字起こし</CardTitle>
                    </CardHeader>
                    <CardContent className="p-6 space-y-6">
                        {/* ジョブ名入力 */}
                        <div className="space-y-4">
                            <div className="space-y-2">
                                <Label htmlFor="job-name"
                                       className="text-sm font-medium text-gray-700 dark:text-gray-300">
                                    ジョブ名
                                </Label>
                                <Input
                                    type="text"
                                    id="job-name"
                                    value={jobName}
                                    onChange={(e) => setJobName(e.target.value)}
                                    placeholder="ジョブ名を入力してください"
                                    className="border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring focus:ring-purple-200 focus:ring-opacity-50 focus:border-indigo-500"
                                />
                            </div>
                        </div>
                        {/* ファイルのアップロード */}
                        <div className="space-y-4">
                            <div className="space-y-2">
                                <Label htmlFor="file-upload"
                                       className="text-sm font-medium text-gray-700 dark:text-gray-300">
                                    音声ファイルをアップロード
                                </Label>
                                <Input
                                    type="file"
                                    id="file-upload"
                                    onChange={handleFileChange}
                                    className="border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring focus:ring-purple-200 focus:ring-opacity-50 focus:border-indigo-500"
                                />
                            </div>
                        </div>

                        {/* 言語選択とオプション */}
                        <div className="space-y-4">
                            <LanguageSelect languageCode={selectedLanguage} setLanguageCode={setSelectedLanguage} />
                            <div className="space-y-4">
                                <div className="flex items-center space-x-2">
                                    <Switch checked={isSpeakerSeparationEnabled}
                                            onCheckedChange={setIsSpeakerSeparationEnabled}/>
                                    <Label
                                        className="text-sm font-medium text-gray-700 dark:text-gray-300">話者分離を有効にする</Label>
                                </div>
                                <div className="flex items-center space-x-2">
                                    <Switch checked={isTimestampEnabled} onCheckedChange={setIsTimestampEnabled}/>
                                    <Label
                                        className="text-sm font-medium text-gray-700 dark:text-gray-300">タイムスタンプを追加</Label>
                                </div>
                            </div>
                        </div>
                    </CardContent>
                </Card>
                <div className="flex justify-center mt-8">
                    <Button
                        onClick={handleStartTranscription}
                        className="w-1/3 bg-blue-600 text-white hover:bg-blue-700 transition-colors duration-200 px-6 py-3 rounded-md"
                        disabled={!selectedFile || !jobName || isLoading}
                    >
                        {isLoading ? <LoadingSpinner /> : 'Start'}
                    </Button>
                </div>
            </div>
        </div>
    );
};

export default Client;
