import { Suspense } from 'react';
import Client from './client';

// サーバーサイドでデータを取得する関数
async function fetchVocabularyData(name: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/custom/vocabulary/${name}`, {  // フルURLを指定
        cache: 'no-store', // 最新のデータを取得するためキャッシュを無効化
    });

    if (!response.ok) {
        throw new Error('Failed to fetch vocabulary data');
    }

    return response.json();
}

const Page = async ({ params }: { params: { name: string } }) => {  // async を追加
    const { name } = params;

    // サーバーコンポーネントでデータを取得
    const data = await fetchVocabularyData(name);

    return (
        <Suspense fallback={<div>Loading...</div>}>
            <Client data={data} />
        </Suspense>
    );
}

export default Page;
