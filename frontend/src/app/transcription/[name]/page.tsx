import React from 'react';
import Client from './client';

const fetchTranscriptionData = async (jobName: string) => {
    try {
        const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/transcriptions/content?jobName=${jobName}`, {
            headers: {
                'Content-Type': 'application/json',
            },
            cache: 'no-store',
        });

        if (!response.ok) {
            throw new Error('Failed to fetch transcription data');
        }

        return await response.json();
    } catch (error) {
        console.error('Error fetching transcription data:', error);
        throw error; // 呼び出し元にエラーを投げる
    }
};

const Page = async ({ params }: { params: { name: string } }) => {
    const { name } = params;

    try {
        // API関数を呼び出してデータを取得
        const data = await fetchTranscriptionData(name);

        return (
            <div className={"w-full min-h-dvh bg-gray-100"}>
                <Client transcript={data.transcript} confidenceData={data.confidence} jobName={name}/>
            </div>
        );
    } catch (error) {
        console.error(error); // エラーハンドリング
        return <div>Error fetching data. Please try again later.</div>;
    }
};

export default Page;