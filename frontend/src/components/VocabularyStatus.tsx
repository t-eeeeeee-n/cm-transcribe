import React from 'react';
import { AiOutlineCheckCircle, AiOutlineClockCircle, AiOutlineCloseCircle, AiOutlineMinusCircle } from 'react-icons/ai'; // アイコンをインポート

interface VocabularyStatusProps {
    status: string;
}

const VocabularyStatus: React.FC<VocabularyStatusProps> = ({ status }) => {
    // ステータスに基づいてアイコンとスタイルを決定
    const getStatusStyle = () => {
        switch (status) {
            case 'READY':
                return { icon: <AiOutlineCheckCircle className="text-green-500" />, text: '登録済み', textColor: 'text-green-500' };
            case 'PENDING':
                return { icon: <AiOutlineClockCircle className="text-yellow-500" />, text: '処理中', textColor: 'text-yellow-500' };
            case 'ERROR':
                return { icon: <AiOutlineCloseCircle className="text-red-500" />, text: 'エラー', textColor: 'text-red-500' };
            case '':
                // 空白が渡された場合、未登録を表示
                return { icon: <AiOutlineMinusCircle className="text-gray-500" />, text: '未登録', textColor: 'text-gray-500' };
            default:
                return { icon: null, text: '-', textColor: 'text-gray-500' };
        }
    };

    const { icon, text, textColor } = getStatusStyle();

    return (
        <div className={`flex items-center space-x-2 p-2 ${textColor}`}>
            {icon}
            <span className="text-sm">{text}</span>
        </div>
    );
};

export default VocabularyStatus;
