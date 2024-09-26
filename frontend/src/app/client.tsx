import React from 'react';
import Link from 'next/link';

const Client = () => {

    const linkData: { title: string, description: string, bgColor: string }[] = [
        { title: "ジョブ一覧", description: "作成済みの文字起こしジョブを一覧で確認できます。", bgColor: "purple" },
        { title: "文字起こし開始", description: "音声ファイルをアップロード。文字起こしを開始しましょう。", bgColor: "blue" },
        { title: "カスタム辞書", description: "独自の辞書を作成して文字起こしの精度を向上させましょう。", bgColor: "green" },
    ]

    return (
        <div className="mx-auto max-w-7xl py-10 px-4 text-center">
            <h1 className="text-2xl font-extrabold mb-8 bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 bg-clip-text text-transparent">
                あなたの音声、すぐに文字に。
            </h1>
            <p className="text-lg mb-12 text-gray-600">
                このアプリケーションでは、音声ファイルをアップロードして自動で文字起こしを行ったり、文字起こしジョブを管理することができます。
                また、カスタム辞書を使ってより正確な文字起こしを行うことも可能です。
            </p>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                {linkData.map((item, i) => (
                    <Link href="/transcription/list" key={i}>
                        <div
                            className={`py-5 px-10 bg-${item.bgColor}-500 hover:bg-${item.bgColor}-600 text-white rounded-xl shadow-lg transform transition duration-300 cursor-pointer`}
                        >
                            <h3 className="text-xl font-bold">{item.title}</h3>
                            <p className="mt-4 text-md">{item.description}</p>
                        </div>
                    </Link>
                ))}
            </div>
        </div>
    );
};

export default Client;