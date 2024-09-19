import React from 'react';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Label } from '@/components/ui/label';

interface LanguageSelectProps {
    languageCode: string;
    setLanguageCode: (code: string) => void;
}

// 言語選択肢の配列
const languages = [
    { value: 'ja-JP', label: '日本語 (ja-JP)' },
    { value: 'en-US', label: '英語、米国 (en-US)' },
    { value: 'fr-FR', label: 'フランス語 (fr-FR)' },
    { value: 'es-ES', label: 'スペイン語 (es-ES)' },
];

const LanguageSelect: React.FC<LanguageSelectProps> = ({ languageCode, setLanguageCode }) => {
    return (
        <div className="space-y-2">
            <Label htmlFor="language-select" className="text-sm font-medium text-gray-700 dark:text-gray-300">
                言語
            </Label>
            <Select
                value={languageCode}
                onValueChange={(value) => setLanguageCode(value)}
            >
                <SelectTrigger className="border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring focus:ring-purple-200 focus:ring-opacity-50 focus:border-indigo-500">
                    <SelectValue placeholder="言語を選択" />
                </SelectTrigger>
                <SelectContent>
                    {languages.map((lang) => (
                        <SelectItem key={lang.value} value={lang.value}>
                            {lang.label}
                        </SelectItem>
                    ))}
                </SelectContent>
            </Select>
        </div>
    );
};

export default LanguageSelect;
