import { NextResponse } from 'next/server';

// POSTリクエストの処理
export async function POST(request: Request) {
    const body = await request.json();

    const { name, language_code, vocabularies } = body;

    // バリデーションなどのサーバーサイド処理を行う
    if (!name || !language_code || !Array.isArray(vocabularies)) {
        return NextResponse.json({ message: 'Invalid input data.' }, { status: 400 });
    }

    // ここでバックエンド処理（例: データベースへの保存など）を行う
    // 仮の処理：コンソールにデータを出力
    console.log('Received data:', { name, language_code, vocabularies });

    // 成功レスポンス
    return NextResponse.json({ message: 'データが保存されました！' }, { status: 200 });
}