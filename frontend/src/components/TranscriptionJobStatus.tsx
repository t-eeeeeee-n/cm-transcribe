import React from 'react';
import { AiOutlineCheckCircle, AiOutlineClockCircle, AiOutlineCloseCircle, AiOutlineMinusCircle } from 'react-icons/ai';
import { FaSpinner } from 'react-icons/fa'; // IN_PROGRESS 用のアイコン

interface StatusBadgeProps {
    status: string;
}

const getStatusStyle = (status: string) => {
    switch (status) {
        case 'COMPLETED':
            return { icon: <AiOutlineCheckCircle className="text-green-500" />, text: '完了', textColor: 'text-green-500' };
        case 'QUEUED':
            return { icon: <AiOutlineClockCircle className="text-blue-500" />, text: '待機中', textColor: 'text-blue-500' };
        case 'IN_PROGRESS':
            return { icon: <FaSpinner className="text-yellow-500 animate-spin" />, text: '進行中', textColor: 'text-yellow-500' };
        case 'FAILED':
            return { icon: <AiOutlineCloseCircle className="text-red-500" />, text: '失敗', textColor: 'text-red-500' };
        default:
            return { icon: null, text: '不明なステータス', textColor: 'text-gray-500' };
    }
};

const TranscriptionJobStatus: React.FC<StatusBadgeProps> = ({ status }) => {
    const { icon, text, textColor } = getStatusStyle(status);

    return (
        <div className={`flex items-center space-x-2 ${textColor}`}>
            {icon}
            <span>{text}</span>
        </div>
    );
};

export default TranscriptionJobStatus;
