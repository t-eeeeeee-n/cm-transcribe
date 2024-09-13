import { NextResponse } from 'next/server';

// ダミーデータを使ったAPIエンドポイント（実際にはデータベースや他のバックエンドを使用）
const vocabularyData: { [key: string]: any } = {
    'sample-vocab': {
        vocabularyName: 'sample-vocab',
        languageCode: 'ja-JP',
        vocabularies: [
            { phrase: 'サンプル', soundsLike: 'サンプル', ipa: '', displayAs: 'Sample' },
        ],
    },
};

const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

export async function GET(request: Request, { params }: { params: { name: string } }) {
    const { name } = params;

    try {
        // リモートのAPIからデータを取得
        const response = await fetch(`${BACKEND_URL}/api/custom/vocabulary/get?name=${name}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            console.error(`Failed to fetch vocabulary with name '${name}' from backend`);
            return NextResponse.json({ message: `Failed to fetch vocabulary with name '${name}'` }, { status: 404 });
        }

        const data = await response.json();
        // console.log('Response from backend:', data);
        return NextResponse.json(data, { status: 200 });

    } catch (error) {
        console.error('Error fetching data from backend:', error);
        return NextResponse.json({ message: 'サーバーエラーが発生しました。' }, { status: 500 });
    }
}

export async function PUT(request: Request, { params }: { params: { name: string } }) {
    const { name } = params;
    const { vocabularyName, languageCode, vocabularies } = await request.json();

    if (!vocabularyName || !languageCode || !Array.isArray(vocabularies)) {
        console.error('Invalid input data', { vocabularyName, languageCode, vocabularies });
        return NextResponse.json({ message: 'Invalid input data' }, { status: 400 });
    }

    vocabularyData[name] = { vocabularyName, languageCode, vocabularies };
    return NextResponse.json({ message: 'Vocabulary updated successfully' }, { status: 200 });
}