import React, { useRef, useState } from 'react';
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { X, Upload, AlertTriangle } from "lucide-react";
import { read, utils } from 'xlsx';
import { Vocabulary } from "@/types/Vocabulary";

interface UploadCSVProps {
    onUpload: (vocabularies: Vocabulary[]) => void;
    fileName: string | null;
    setFileName: React.Dispatch<React.SetStateAction<string | null>>;
}

const UploadCSV = ({ onUpload, fileName, setFileName }: UploadCSVProps) => {
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const fileInputRef = useRef<HTMLInputElement | null>(null);

    const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (!file) return;

        setFileName(file.name);

        try {
            const reader = new FileReader();
            reader.onload = (e) => {
                const arrayBuffer = e.target?.result as ArrayBuffer;
                const uint8Array = new Uint8Array(arrayBuffer);
                const workbook = read(uint8Array, { type: 'array' });
                const sheetName = workbook.SheetNames[0];
                const worksheet = workbook.Sheets[sheetName];
                const data = utils.sheet_to_json<any>(worksheet);

                const normalizedVocabularies = data.map((vocabulary: any) => ({
                    phrase: vocabulary.Phrase || '',
                    soundsLike: vocabulary.SoundsLike || '',
                    ipa: vocabulary.ipa || '',
                    displayAs: vocabulary.DisplayAs || '',
                }));

                onUpload(normalizedVocabularies);
            };
            reader.readAsArrayBuffer(file);
        } catch (error) {
            setErrorMessage('ファイルの読み込みに失敗しました。');
            // console.error('Error reading file:', error);
        }
    };

    const handleRemoveFile = () => {
        setFileName(null);
        setErrorMessage(null);
        onUpload([]);
    };

    const handleButtonClick = () => {
        fileInputRef.current?.click();
    };

    return (
        <div className="mt-6 space-y-6">
            <div className="flex justify-center">
                {/* ファイル入力 */}
                <Input
                    type="file"
                    accept=".csv, .xlsx"
                    onChange={handleFileChange}
                    className="hidden"
                    ref={fileInputRef}
                />
                <Button variant="outline" className="cursor-pointer hover:bg-blue-100 transition-all duration-300" onClick={handleButtonClick}>
                    <Upload className="mr-2 h-5 w-5 text-blue-600" />
                    CSVファイルを選択
                </Button>
            </div>

            {fileName && (
                <Card className="shadow-lg rounded-lg border border-gray-200">
                    <CardContent className="flex items-center justify-between p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-all duration-300">
                        <div>
                            <p className="text-sm text-gray-500">アップロードしたファイル:</p>
                            <p className="font-medium text-gray-800">{fileName}</p>
                        </div>
                        <Button variant="ghost" size="icon" onClick={handleRemoveFile} className="hover:bg-red-100 transition-all duration-300">
                            <X className="h-5 w-5 text-red-500" />
                        </Button>
                    </CardContent>
                </Card>
            )}

            {errorMessage && (
                <Alert variant="destructive" className="flex items-center gap-3 p-4 rounded-lg border-l-4 border-red-500 bg-red-50 text-red-800 shadow-lg transition-transform transform hover:scale-105">
                    <AlertTriangle className="h-5 w-5 text-red-600" />
                    <div>
                        <AlertTitle className="font-bold">エラー</AlertTitle>
                        <AlertDescription className="text-sm">{errorMessage}</AlertDescription>
                    </div>
                    <Button variant="outline" size="sm" onClick={() => setErrorMessage(null)} className="ml-auto">
                        閉じる
                    </Button>
                </Alert>
            )}
        </div>
    );
};

export default UploadCSV;
