import Client from './client';

// Transcription Jobsデータをサーバーサイドで取得する関数
async function fetchTranscriptionJobs() {
    try {
        const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/transcriptions`, {
            headers: {
                'Content-Type': 'application/json',
            },
            cache: 'no-store',
        });

        if (!response.ok) {
            throw new Error('Failed to fetch transcription jobs');
        }

        const data = await response.json();
        return data.jobs;
    } catch (error) {
        console.error('Error fetching transcription jobs:', error);
        throw error; // 呼び出し元にエラーを投げる
    }
}

const Page = async () => {
    try {
        // fetchTranscriptionJobs関数を使ってデータを取得
        const jobs = await fetchTranscriptionJobs();

        // Client Componentにデータを渡して表示
        return (
            <div className={"w-full min-h-dvh bg-gray-100"}>
                <Client jobs={jobs}/>
            </div>
        )
    } catch (error) {
        // エラーハンドリング
        return (
            <div className="text-center text-red-500">
                データの取得に失敗しました。再度お試しください。
            </div>
        );
    }
};

export default Page;