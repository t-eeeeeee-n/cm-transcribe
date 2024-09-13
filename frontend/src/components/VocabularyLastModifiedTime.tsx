import React from 'react';
import { Calendar } from "lucide-react"; // カレンダーアイコンをインポート

interface VocabularyLastModifiedTimeProps {
    time?: string;
}

const VocabularyLastModifiedTime: React.FC<VocabularyLastModifiedTimeProps> = ({ time }) => {
    // 日付をフォーマットする関数
    const formatTime = (timeString: string) => {
        const date = new Date(timeString);
        return new Intl.DateTimeFormat('ja-JP', {
            year: 'numeric',
            month: 'numeric',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            second: 'numeric',
            timeZone: 'UTC', // タイムゾーンをUTCに統一
            hour12: false,
        }).format(date);
    };

    return (
        <div className="flex items-center space-x-2 p-2">
            <Calendar className="h-4 w-4 text-gray-700 dark:text-gray-300" />
            <span className="text-sm text-gray-900 dark:text-gray-100">
                {time ? formatTime(time) : "未設定"}
            </span>
        </div>
    );
};

export default VocabularyLastModifiedTime;
