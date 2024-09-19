import { NextResponse } from 'next/server';

const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

export async function POST(request: Request) {
    const body = await request.json();

    const { name, language_code, vocabularies } = body;

    // バリデーションなどのサーバーサイド処理を行う
    if (!name || !language_code || !Array.isArray(vocabularies) || vocabularies.length === 0) {
        console.error('Invalid input data:', { name, language_code, vocabularies });
        return NextResponse.json({ message: 'Invalid input data.' }, { status: 400 });
    }

    try {
        // console.log('Sending request to backend:', `${BACKEND_URL}/api/custom/vocabulary/create`);
        // console.log('Request body:', { name, language_code, vocabularies });

        // return NextResponse.json({ message: 'データが保存されました！'}, { status: 200 });

        // 環境変数からバックエンドのURLを使用してリクエストを送信
        const response = await fetch(`${BACKEND_URL}/api/custom/vocabulary/create`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, language_code, vocabularies }),
        });

        if (!response.ok) {
            // バックエンドからエラーレスポンスが返ってきた場合
            let errorMessage = 'Unknown error';
            const contentType = response.headers.get('content-type');

            if (contentType && contentType.includes('application/json')) {
                const errorData = await response.json(); // JSONのパースを試みる
                errorMessage = errorData.message;
                console.error('Error response from backend (JSON):', errorData);
            } else {
                const errorText = await response.text(); // テキストとして取得
                errorMessage = errorText;
                console.error('Error response from backend (not JSON):', errorText);
            }

            return NextResponse.json({ message: `バックエンドエラー: ${errorMessage}` }, { status: response.status });
        }

        // バックエンドのレスポンスを受け取る
        const result = await response.json();
        // console.log('Response from backend:', result);

        // 成功レスポンス
        return NextResponse.json({ message: 'データが保存されました！', data: result }, { status: 200 });
    } catch (error) {
        console.error('Error saving data to backend:', error);
        return NextResponse.json({ message: 'サーバーエラーが発生しました。' }, { status: 500 });
    }
}
