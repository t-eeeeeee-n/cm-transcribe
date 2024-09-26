import {NextResponse} from 'next/server';

const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

// 共通の入力バリデーション関数
function validateInput(body: any) {
    const { name, language_code, vocabularies } = body;
    if (!name || !language_code || !Array.isArray(vocabularies) || vocabularies.length === 0) {
        // console.error('Invalid input data:', { name, language_code, vocabularies });
        return { valid: false, message: 'Invalid input data.' };
    }
    return { valid: true };
}

// バックエンドにリクエストを送る共通関数
async function sendToBackend(method: string, body: any) {
    try {
        const response = await fetch(`${BACKEND_URL}/api/custom/vocabulary`, {
            method,
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(body),
        });

        // custom vocabulary nameが重複している場合、409ステータスコードを処理
        if (response.status === 409) {
            return NextResponse.json({
                message: 'Custom vocabulary name already exists',
            }, { status: 409 });
        }

        if (!response.ok) {
            const contentType = response.headers.get('content-type');
            let errorMessage: string;

            if (contentType && contentType.includes('application/json')) {
                const errorData = await response.json();
                errorMessage = errorData.message;
                // console.error('Error response from backend (JSON):', errorData);
            } else {
                errorMessage = await response.text();
                // console.error('Error response from backend (not JSON):', errorText);
            }

            return NextResponse.json({ message: `バックエンドエラー: ${errorMessage}` }, { status: response.status });
        }

        const result = await response.json();
        return NextResponse.json({ message: 'データが保存されました！', data: result }, { status: 200 });
    } catch (error) {
        // console.error('Error saving data to backend:', error);
        return NextResponse.json({ message: 'サーバーエラーが発生しました。' }, { status: 500 });
    }
}

// POSTリクエストの処理
export async function POST(request: Request) {
    const body = await request.json();
    const validation = validateInput(body);
    if (!validation.valid) {
        return NextResponse.json({ message: validation.message }, { status: 400 });
    }
    return sendToBackend('POST', body);
}

// PUTリクエストの処理
export async function PUT(request: Request) {
    const body = await request.json();
    const validation = validateInput(body);
    if (!validation.valid) {
        return NextResponse.json({ message: validation.message }, { status: 400 });
    }
    return sendToBackend('PUT', body);
}