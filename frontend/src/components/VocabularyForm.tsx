import React from 'react';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import VocabularyStatus from "@/components/VocabularyStatus";
import VocabularyLastModifiedTime from "@/components/VocabularyLastModifiedTime";
import LanguageSelect from "@/components/LanguageSelect";

interface VocabularyFormProps {
    vocabularyName: string;
    setVocabularyName: (name: string) => void;
    languageCode: string;
    setLanguageCode: (code: string) => void;
    vocabularyState: string;
    lastModifiedTime?: string;
}

const VocabularyForm: React.FC<VocabularyFormProps> = ({
                                                           vocabularyName,
                                                           setVocabularyName,
                                                           languageCode,
                                                           setLanguageCode,
                                                           vocabularyState,
                                                           lastModifiedTime,
                                                       }) => {
    return (
        <Card className="mb-8 shadow-lg border border-gray-200 rounded-lg bg-white dark:bg-gray-800 transition duration-300 ease-in-out hover:shadow-xl">
            <CardHeader className="p-4 bg-gradient-to-r from-purple-500 to-indigo-600 text-white rounded-t-lg">
                <CardTitle className="text-2xl font-semibold">基本情報</CardTitle>
            </CardHeader>
            <CardContent className="p-6 space-y-6">
                {/* 名前入力フィールド */}
                <div className="space-y-2">
                    <Label htmlFor="vocabulary-name" className="text-sm font-medium text-gray-700 dark:text-gray-300">
                        名前
                    </Label>
                    <Input
                        id="vocabulary-name"
                        type="text"
                        value={vocabularyName}
                        onChange={(e) => setVocabularyName(e.target.value)}
                        placeholder="語彙の名前を入力"
                        className="border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring focus:ring-purple-200 focus:ring-opacity-50 focus:border-indigo-500"
                    />
                    <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">
                        語彙名の長さは最大 200 文字です。使用できる文字: a～z、A～Z、0～9、ピリオド (.)、ダッシュ (-)、およびアンダースコア (_)。
                    </p>
                </div>

                {/* 言語選択フィールド */}
                <div className="space-y-2">
                    <LanguageSelect languageCode={languageCode} setLanguageCode={setLanguageCode} />
                </div>

                {/* ステータス表示部分を追加 */}
                <div className="space-y-2">
                    <Label className="text-sm font-medium text-gray-700 dark:text-gray-300">
                        ステータス
                    </Label>
                    <VocabularyStatus status={vocabularyState} />
                </div>

                {/* 変更日時表示フィールド */}
                <div className="space-y-2">
                    <Label className="text-sm font-medium text-gray-700 dark:text-gray-300">変更日時</Label>
                    <VocabularyLastModifiedTime time={lastModifiedTime} />
                </div>
            </CardContent>
        </Card>
    );
};

export default VocabularyForm;
